package tgbot

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"time"
)

const (
	defaultBotAPIURL           = "https://api.telegram.org"
	defaultBotHTTPTimeout      = 8 * time.Second
	defaultBotResponseBodySize = int64(4 << 20)
)

// BotOption configures a Bot client.
type BotOption func(*Bot)

// Bot is a generic Telegram Bot API client used by API wrappers.
type Bot struct {
	token  string
	apiURL string
	client *http.Client
	debug  bool
}

// APIResponse is the Telegram Bot API response envelope.
type APIResponse struct {
	OK          bool                `json:"ok"`
	Result      json.RawMessage     `json:"result,omitempty"`
	ErrorCode   int                 `json:"error_code,omitempty"`
	Description string              `json:"description,omitempty"`
	Parameters  *ResponseParameters `json:"parameters,omitempty"`
}

// ResponseParameters contains extra Telegram error hints.
type ResponseParameters struct {
	MigrateToChatID int64 `json:"migrate_to_chat_id,omitempty"`
	RetryAfter      int   `json:"retry_after,omitempty"`
}

// APIError wraps Telegram API failures.
type APIError struct {
	StatusCode int
	Code       int
	Message    string
	RetryAfter int
}

func (value *APIError) Error() string {
	if value == nil {
		return "telegram api error"
	}
	if strings.TrimSpace(value.Message) == "" {
		if value.Code > 0 {
			return fmt.Sprintf("telegram api error: code=%d", value.Code)
		}
		if value.StatusCode > 0 {
			return fmt.Sprintf("telegram api error: status=%d", value.StatusCode)
		}
		return "telegram api error"
	}
	if value.Code > 0 {
		return fmt.Sprintf("telegram api error: code=%d message=%s", value.Code, value.Message)
	}
	if value.StatusCode > 0 {
		return fmt.Sprintf("telegram api error: status=%d message=%s", value.StatusCode, value.Message)
	}
	return fmt.Sprintf("telegram api error: %s", value.Message)
}

// IsTooManyRequests reports whether an error is a Telegram 429 rate-limit error.
func IsTooManyRequests(err error) bool {
	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		return false
	}
	return apiErr.StatusCode == http.StatusTooManyRequests || apiErr.Code == http.StatusTooManyRequests
}

// WithHTTPClient overrides the HTTP client used by the Bot.
func WithHTTPClient(client *http.Client) BotOption {
	return func(bot *Bot) {
		if client != nil {
			bot.client = client
		}
	}
}

// WithAPIURL overrides Telegram API base URL. Useful for local Bot API servers.
func WithAPIURL(apiURL string) BotOption {
	return func(bot *Bot) {
		trimmed := strings.TrimRight(strings.TrimSpace(apiURL), "/")
		if trimmed != "" {
			bot.apiURL = trimmed
		}
	}
}

// WithDebug toggles lightweight debug logging for requests and errors.
func WithDebug(debug bool) BotOption {
	return func(bot *Bot) {
		bot.debug = debug
	}
}

// NewBot creates a new Telegram Bot API client.
func NewBot(token string, opts ...BotOption) (*Bot, error) {
	trimmedToken := strings.TrimSpace(token)
	if trimmedToken == "" {
		return nil, fmt.Errorf("telegram bot token is empty")
	}

	bot := &Bot{
		token:  trimmedToken,
		apiURL: defaultBotAPIURL,
		client: &http.Client{Timeout: defaultBotHTTPTimeout},
	}
	for _, opt := range opts {
		if opt != nil {
			opt(bot)
		}
	}
	if bot.client == nil {
		bot.client = &http.Client{Timeout: defaultBotHTTPTimeout}
	}
	return bot, nil
}

// Do sends a raw Telegram method call and returns the raw result payload.
func (bot *Bot) Do(ctx context.Context, method string, params any) (json.RawMessage, error) {
	var result json.RawMessage
	if err := bot.call(ctx, method, params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// Call sends a Telegram method call and decodes the result into result.
func (bot *Bot) Call(ctx context.Context, method string, params any, result any) error {
	return bot.call(ctx, method, params, result)
}

func (bot *Bot) call(ctx context.Context, method string, params any, result any) error {
	if bot == nil {
		return fmt.Errorf("telegram bot client is nil")
	}
	if strings.TrimSpace(bot.token) == "" {
		return fmt.Errorf("telegram bot token is empty")
	}
	methodName := strings.TrimSpace(method)
	if methodName == "" {
		return fmt.Errorf("telegram method is empty")
	}
	if ctx == nil {
		ctx = context.Background()
	}

	httpMethod := http.MethodPost
	var (
		bodyReader  io.Reader
		contentType string
	)
	if isNilValue(params) {
		httpMethod = http.MethodGet
		bodyReader = nil
	} else {
		encoded, encodedContentType, err := encodeBotRequestBody(params)
		if err != nil {
			return err
		}
		contentType = encodedContentType
		bodyReader = bytes.NewReader(encoded)
	}

	requestURL := fmt.Sprintf("%s/bot%s/%s", bot.apiURL, bot.token, methodName)
	request, err := http.NewRequestWithContext(ctx, httpMethod, requestURL, bodyReader)
	if err != nil {
		return fmt.Errorf("build telegram request: %w", err)
	}
	if httpMethod == http.MethodPost {
		if strings.TrimSpace(contentType) == "" {
			contentType = "application/json"
		}
		request.Header.Set("Content-Type", contentType)
	}

	response, err := bot.client.Do(request)
	if err != nil {
		return fmt.Errorf("send telegram request: %w", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(io.LimitReader(response.Body, defaultBotResponseBodySize))
	if err != nil {
		return fmt.Errorf("read telegram response: %w", err)
	}

	var envelope APIResponse
	if len(body) > 0 {
		_ = json.Unmarshal(body, &envelope)
	}

	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
		return buildBotAPIError(response.StatusCode, envelope, body)
	}
	if len(body) == 0 {
		if bot.debug {
			fmt.Printf("tg debug: %s returned empty body\n", methodName)
		}
		return nil
	}

	if err := json.Unmarshal(body, &envelope); err != nil {
		return fmt.Errorf("decode telegram envelope: %w", err)
	}
	if !envelope.OK {
		return buildBotAPIError(response.StatusCode, envelope, body)
	}
	if result == nil || len(envelope.Result) == 0 {
		return nil
	}
	if handled, err := decodeIntoKnownUnionResult(result, envelope.Result); handled {
		if err != nil {
			return fmt.Errorf("decode telegram union result for %s: %w", methodName, err)
		}
		return nil
	}
	if err := json.Unmarshal(envelope.Result, result); err != nil {
		return fmt.Errorf("decode telegram result for %s: %w", methodName, err)
	}
	return nil
}

type botUploadPart struct {
	formField string
	fileName  string
	reader    io.Reader
	closer    io.Closer
}

func encodeBotRequestBody(params any) ([]byte, string, error) {
	payload, uploads, ok, err := buildPayloadFromParams(params)
	if err != nil {
		return nil, "", err
	}

	if !ok {
		encoded, marshalErr := json.Marshal(params)
		if marshalErr != nil {
			return nil, "", fmt.Errorf("marshal telegram params: %w", marshalErr)
		}
		if len(encoded) == 0 {
			encoded = []byte("{}")
		}
		return encoded, "application/json", nil
	}

	if len(uploads) == 0 {
		encoded, marshalErr := json.Marshal(payload)
		if marshalErr != nil {
			return nil, "", fmt.Errorf("marshal telegram params: %w", marshalErr)
		}
		if len(encoded) == 0 {
			encoded = []byte("{}")
		}
		return encoded, "application/json", nil
	}

	formBody, contentType, err := buildMultipartBody(payload, uploads)
	if err != nil {
		return nil, "", err
	}
	return formBody, contentType, nil
}

func buildPayloadFromParams(params any) (map[string]any, []botUploadPart, bool, error) {
	if isNilValue(params) {
		return nil, nil, false, nil
	}

	value := reflect.ValueOf(params)
	for value.Kind() == reflect.Ptr {
		if value.IsNil() {
			return nil, nil, false, nil
		}
		value = value.Elem()
	}
	if value.Kind() != reflect.Struct {
		return nil, nil, false, nil
	}

	typ := value.Type()
	payload := make(map[string]any)
	uploads := make([]botUploadPart, 0)
	attachCounter := 0

	for i := 0; i < typ.NumField(); i++ {
		fieldType := typ.Field(i)
		if fieldType.PkgPath != "" {
			continue
		}
		jsonName, omitEmpty := parseJSONTag(fieldType.Tag.Get("json"), fieldType.Name)
		if jsonName == "" {
			continue
		}

		fieldValue := value.Field(i)
		if omitEmpty && isZeroFieldValue(fieldValue) {
			continue
		}

		raw := fieldValue.Interface()
		normalized, err := normalizeParamValue(raw, jsonName, &uploads, &attachCounter)
		if err != nil {
			return nil, nil, true, fmt.Errorf("encode %s: %w", jsonName, err)
		}
		payload[jsonName] = normalized
	}

	return payload, uploads, true, nil
}

func normalizeParamValue(value any, fieldPath string, uploads *[]botUploadPart, attachCounter *int) (any, error) {
	if inputFile, isInputFile := toInputFile(value); isInputFile {
		formValue, uploadPart, err := inputFileAsFormValue(fieldPath, inputFile, *attachCounter)
		if err != nil {
			return nil, err
		}
		if uploadPart != nil {
			*uploads = append(*uploads, *uploadPart)
			*attachCounter = *attachCounter + 1
		}
		return formValue, nil
	}
	return normalizeReflectValue(reflect.ValueOf(value), fieldPath, uploads, attachCounter)
}

func normalizeReflectValue(value reflect.Value, fieldPath string, uploads *[]botUploadPart, attachCounter *int) (any, error) {
	if !value.IsValid() {
		return nil, nil
	}

	switch value.Kind() {
	case reflect.Interface, reflect.Ptr:
		if value.IsNil() {
			return nil, nil
		}
		return normalizeReflectValue(value.Elem(), fieldPath, uploads, attachCounter)
	case reflect.Struct:
		if inputFile, isInputFile := toInputFile(value.Interface()); isInputFile {
			formValue, uploadPart, err := inputFileAsFormValue(fieldPath, inputFile, *attachCounter)
			if err != nil {
				return nil, err
			}
			if uploadPart != nil {
				*uploads = append(*uploads, *uploadPart)
				*attachCounter = *attachCounter + 1
			}
			return formValue, nil
		}

		typ := value.Type()
		result := make(map[string]any)
		for i := 0; i < typ.NumField(); i++ {
			fieldType := typ.Field(i)
			if fieldType.PkgPath != "" {
				continue
			}

			jsonName, omitEmpty := parseJSONTag(fieldType.Tag.Get("json"), fieldType.Name)
			if jsonName == "" {
				continue
			}

			fieldValue := value.Field(i)
			if omitEmpty && isZeroFieldValue(fieldValue) {
				continue
			}

			normalized, err := normalizeReflectValue(fieldValue, joinFieldPath(fieldPath, jsonName), uploads, attachCounter)
			if err != nil {
				return nil, err
			}
			result[jsonName] = normalized
		}
		return result, nil
	case reflect.Map:
		if value.Type().Key().Kind() != reflect.String {
			return value.Interface(), nil
		}
		keys := value.MapKeys()
		sort.Slice(keys, func(i, j int) bool {
			return keys[i].String() < keys[j].String()
		})

		result := make(map[string]any, len(keys))
		for _, key := range keys {
			keyName := key.String()
			normalized, err := normalizeReflectValue(value.MapIndex(key), joinFieldPath(fieldPath, keyName), uploads, attachCounter)
			if err != nil {
				return nil, err
			}
			result[keyName] = normalized
		}
		return result, nil
	case reflect.Slice, reflect.Array:
		if value.Type().Elem().Kind() == reflect.Uint8 {
			return value.Interface(), nil
		}
		length := value.Len()
		result := make([]any, 0, length)
		for idx := 0; idx < length; idx++ {
			normalized, err := normalizeReflectValue(value.Index(idx), fmt.Sprintf("%s.%d", fieldPath, idx), uploads, attachCounter)
			if err != nil {
				return nil, err
			}
			result = append(result, normalized)
		}
		return result, nil
	default:
		return value.Interface(), nil
	}
}

func joinFieldPath(base, child string) string {
	if strings.TrimSpace(base) == "" {
		return child
	}
	if strings.TrimSpace(child) == "" {
		return base
	}
	return base + "." + child
}

func buildMultipartBody(payload map[string]any, uploads []botUploadPart) ([]byte, string, error) {
	for _, upload := range uploads {
		if upload.closer != nil {
			defer upload.closer.Close()
		}
	}

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	keys := make([]string, 0, len(payload))
	for key := range payload {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		valueText, err := stringifyMultipartValue(payload[key])
		if err != nil {
			return nil, "", fmt.Errorf("encode multipart field %s: %w", key, err)
		}
		if err := writer.WriteField(key, valueText); err != nil {
			return nil, "", fmt.Errorf("write multipart field %s: %w", key, err)
		}
	}

	for _, upload := range uploads {
		part, err := writer.CreateFormFile(upload.formField, upload.fileName)
		if err != nil {
			return nil, "", fmt.Errorf("create multipart file %s: %w", upload.formField, err)
		}
		if _, err := io.Copy(part, upload.reader); err != nil {
			return nil, "", fmt.Errorf("write multipart file %s: %w", upload.formField, err)
		}
	}

	if err := writer.Close(); err != nil {
		return nil, "", fmt.Errorf("finalize multipart body: %w", err)
	}

	return body.Bytes(), writer.FormDataContentType(), nil
}

func inputFileAsFormValue(fieldName string, file InputFile, attachIndex int) (string, *botUploadPart, error) {
	sources := file.sourceCount()
	if sources == 0 {
		return "", nil, fmt.Errorf("input file source is empty")
	}
	if sources > 1 {
		return "", nil, fmt.Errorf("input file source is ambiguous")
	}

	if strings.TrimSpace(file.FileID) != "" {
		return strings.TrimSpace(file.FileID), nil, nil
	}
	if strings.TrimSpace(file.URL) != "" {
		return strings.TrimSpace(file.URL), nil, nil
	}

	attachKey := fmt.Sprintf("attach_%s_%d", sanitizeAttachKey(fieldName), attachIndex)
	resultValue := "attach://" + attachKey

	if strings.TrimSpace(file.FilePath) != "" {
		path := strings.TrimSpace(file.FilePath)
		fd, err := os.Open(path)
		if err != nil {
			return "", nil, fmt.Errorf("open file path %s: %w", path, err)
		}
		return resultValue, &botUploadPart{
			formField: attachKey,
			fileName:  file.normalizedFileName(filepath.Base(path)),
			reader:    fd,
			closer:    fd,
		}, nil
	}

	if file.Reader != nil {
		return resultValue, &botUploadPart{
			formField: attachKey,
			fileName:  file.normalizedFileName(attachKey + ".bin"),
			reader:    file.Reader,
		}, nil
	}

	return "", nil, fmt.Errorf("input file source is empty")
}

func sanitizeAttachKey(value string) string {
	raw := strings.TrimSpace(value)
	if raw == "" {
		return "file"
	}
	var b strings.Builder
	for _, r := range raw {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '_' {
			b.WriteRune(r)
			continue
		}
		b.WriteByte('_')
	}
	result := strings.Trim(strings.ToLower(b.String()), "_")
	if result == "" {
		return "file"
	}
	return result
}

func stringifyMultipartValue(value any) (string, error) {
	switch casted := value.(type) {
	case string:
		return casted, nil
	case json.RawMessage:
		return string(casted), nil
	case bool:
		if casted {
			return "true", nil
		}
		return "false", nil
	}

	encoded, err := json.Marshal(value)
	if err != nil {
		return "", err
	}
	return string(encoded), nil
}

func parseJSONTag(tagValue, fallback string) (string, bool) {
	if tagValue == "-" {
		return "", false
	}
	if strings.TrimSpace(tagValue) == "" {
		return fallback, false
	}
	parts := strings.Split(tagValue, ",")
	name := strings.TrimSpace(parts[0])
	if name == "" {
		name = fallback
	}
	omitEmpty := false
	for _, opt := range parts[1:] {
		if strings.TrimSpace(opt) == "omitempty" {
			omitEmpty = true
			break
		}
	}
	return name, omitEmpty
}

func isZeroFieldValue(value reflect.Value) bool {
	if value.Kind() == reflect.Interface && !value.IsNil() {
		inner := value.Elem()
		if inner.IsValid() {
			if inputFile, ok := toInputFile(inner.Interface()); ok {
				return inputFile.isEmpty()
			}
		}
	}
	if value.CanInterface() {
		if inputFile, ok := toInputFile(value.Interface()); ok {
			return inputFile.isEmpty()
		}
	}
	return value.IsZero()
}

func toInputFile(value any) (InputFile, bool) {
	switch casted := value.(type) {
	case InputFile:
		return casted, true
	case *InputFile:
		if casted == nil {
			return InputFile{}, true
		}
		return *casted, true
	default:
		return InputFile{}, false
	}
}

func buildBotAPIError(statusCode int, envelope APIResponse, rawBody []byte) error {
	message := strings.TrimSpace(envelope.Description)
	if message == "" {
		message = strings.TrimSpace(string(rawBody))
	}
	if message == "" {
		message = "telegram api request failed"
	}

	apiErr := &APIError{
		StatusCode: statusCode,
		Code:       envelope.ErrorCode,
		Message:    message,
	}
	if envelope.Parameters != nil {
		apiErr.RetryAfter = envelope.Parameters.RetryAfter
	}
	return apiErr
}

func isNilValue(value any) bool {
	if value == nil {
		return true
	}
	rv := reflect.ValueOf(value)
	switch rv.Kind() {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Interface, reflect.Func, reflect.Chan:
		return rv.IsNil()
	default:
		return false
	}
}
