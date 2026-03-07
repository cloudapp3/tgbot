#!/usr/bin/env python3
import argparse
import html as htmlmod
import re
import subprocess
import sys
import urllib.request
from pathlib import Path

OFFICIAL_BOT_API_URL = "https://core.telegram.org/bots/api"


def fetch_html(url: str) -> str:
    curl = subprocess.run(
        ["curl", "-fsSL", url],
        check=False,
        capture_output=True,
        text=True,
    )
    if curl.returncode == 0 and curl.stdout:
        return curl.stdout

    with urllib.request.urlopen(url) as response:
        return response.read().decode("utf-8", errors="ignore")


HEADING_RE = re.compile(
    r'<h([34])><a class="anchor" name="([^"]+)" href="#[^"]+">'
    r'<i class="anchor-icon"></i></a>(.*?)</h\1>',
    re.S,
)


def clean_text(text: str) -> str:
    text = re.sub(r"<code>(.*?)</code>", r"\1", text, flags=re.S)
    text = re.sub(r"<[^>]+>", "", text)
    return " ".join(htmlmod.unescape(text).split())


def parse_sections(html: str):
    headings = list(HEADING_RE.finditer(html))
    sections = []
    for index, match in enumerate(headings):
        if match.group(1) != "4":
            continue
        title = clean_text(match.group(3))
        body_start = match.end()
        body_end = headings[index + 1].start() if index + 1 < len(headings) else len(html)
        body = html[body_start:body_end]
        sections.append((title, body))
    return sections


def extract_first_table(body: str):
    table_match = re.search(r'<table class="table">(.*?)</table>', body, re.S)
    if not table_match:
        return [], []
    table_html = table_match.group(1)
    headers = [clean_text(item) for item in re.findall(r"<th>(.*?)</th>", table_html, re.S)]
    rows = []
    for row_html in re.findall(r"<tr>(.*?)</tr>", table_html, re.S):
        cells = [clean_text(item) for item in re.findall(r"<t[dh]>(.*?)</t[dh]>", row_html, re.S)]
        if cells and cells != headers:
            rows.append(cells)
    return headers, rows


def extract_latest_version(html: str):
    match = re.search(
        r'<h4><a class="anchor" name="[^"]+" href="#[^"]+">'
        r'<i class="anchor-icon"></i></a>([^<]+)</h4>\s*<p><strong>(Bot API [^<]+)</strong></p>',
        html,
        re.S,
    )
    if not match:
        return "unknown", "unknown"
    return clean_text(match.group(2)), clean_text(match.group(1))


def parse_official_schema(html: str):
    official_types = {}
    official_methods = {}
    for title, body in parse_sections(html):
        headers, rows = extract_first_table(body)
        body_text = clean_text(body[:500])

        if re.match(r"^[a-z][A-Za-z0-9]+$", title):
            if (
                headers[:1] == ["Parameter"]
                or "Requires no parameters." in body_text
                or "Returns " in body_text
                or "returns " in body_text
            ):
                official_methods[title] = [row[0] for row in rows if row]
            continue

        if re.match(r"^[A-Z][A-Za-z0-9]+$", title):
            if headers or body_text:
                official_types[title] = [row[0] for row in rows if row]

    return official_types, official_methods


STRUCT_RE = re.compile(r"type\s+([A-Z][A-Za-z0-9]+)\s+struct\s*\{(.*?)\n\}", re.S)
INTERFACE_RE = re.compile(r"type\s+([A-Z][A-Za-z0-9]+)\s+interface\s*\{", re.S)
JSON_TAG_RE = re.compile(r'`json:"([^",]+)')
PARAMS_RE = re.compile(
    r'//\s+([A-Z][A-Za-z0-9]+Params) contains params for Telegram method "([^"]+)"\.\s*'
    r'\ntype\s+\1\s+struct\s*\{(.*?)\n\}',
    re.S,
)


def parse_repo_types(repo_root: Path):
    types = {}
    for relative_path in ["sdk_types.go", "sdk_core.go", "sdk_files.go"]:
        text = (repo_root / relative_path).read_text(encoding="utf-8")
        for match in STRUCT_RE.finditer(text):
            types[match.group(1)] = JSON_TAG_RE.findall(match.group(2))
    return types


def parse_repo_type_interfaces(repo_root: Path):
    text = (repo_root / "sdk_types.go").read_text(encoding="utf-8")
    return set(INTERFACE_RE.findall(text))


def parse_repo_methods(repo_root: Path):
    text = (repo_root / "sdk_methods.go").read_text(encoding="utf-8")
    methods = {}
    for match in PARAMS_RE.finditer(text):
        methods[match.group(2)] = JSON_TAG_RE.findall(match.group(3))
    return methods


def compare(repo_root: Path, official_types, official_methods):
    repo_types = parse_repo_types(repo_root)
    repo_interfaces = parse_repo_type_interfaces(repo_root)
    repo_methods = parse_repo_methods(repo_root)

    missing_types = []
    missing_methods = []
    missing_type_fields = {}
    missing_method_params = {}

    for name, fields in official_types.items():
        if name not in repo_types and name not in repo_interfaces:
            missing_types.append(name)
            continue
        if name in repo_types:
            missing = [field for field in fields if field not in repo_types[name]]
            if missing:
                missing_type_fields[name] = missing

    for name, params in official_methods.items():
        if name not in repo_methods:
            missing_methods.append(name)
            continue
        missing = [param for param in params if param not in repo_methods[name]]
        if missing:
            missing_method_params[name] = missing

    return {
        "missing_types": missing_types,
        "missing_methods": missing_methods,
        "missing_type_fields": missing_type_fields,
        "missing_method_params": missing_method_params,
    }


def main() -> int:
    parser = argparse.ArgumentParser(description="Verify this SDK against the official Telegram Bot API docs.")
    parser.add_argument("--html", help="Use a previously downloaded Telegram Bot API HTML file.")
    parser.add_argument("--url", default=OFFICIAL_BOT_API_URL, help="Official Telegram Bot API documentation URL.")
    parser.add_argument("--repo-root", default=".", help="Repository root to inspect.")
    args = parser.parse_args()

    if args.html:
        html = Path(args.html).read_text(encoding="utf-8", errors="ignore")
    else:
        html = fetch_html(args.url)

    version, version_date = extract_latest_version(html)
    official_types, official_methods = parse_official_schema(html)
    result = compare(Path(args.repo_root), official_types, official_methods)

    print(f"Official source: {args.url}")
    print(f"Latest official version: {version} ({version_date})")
    print(f"Official named types: {len(official_types)}")
    print(f"Official methods: {len(official_methods)}")

    if (
        not result["missing_types"]
        and not result["missing_methods"]
        and not result["missing_type_fields"]
        and not result["missing_method_params"]
    ):
        print("Coverage OK: repository matches official Telegram Bot API names and field/parameter sets.")
        print("Note: InputFile and ResponseParameters are runtime support types that live outside sdk_types.go.")
        return 0

    if result["missing_types"]:
        print("Missing types:", ", ".join(result["missing_types"]))
    if result["missing_methods"]:
        print("Missing methods:", ", ".join(result["missing_methods"]))
    if result["missing_type_fields"]:
        print("Types with missing fields:")
        for name, fields in sorted(result["missing_type_fields"].items()):
            print(f"  - {name}: {', '.join(fields)}")
    if result["missing_method_params"]:
        print("Methods with missing parameters:")
        for name, params in sorted(result["missing_method_params"].items()):
            print(f"  - {name}: {', '.join(params)}")
    return 1


if __name__ == "__main__":
    sys.exit(main())
