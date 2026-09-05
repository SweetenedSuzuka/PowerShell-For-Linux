"""
获取官方页面 markdown 并提取字段。
供 snapshot.py 与 check_official_changes.py 导入。
"""
import re
import time
import urllib.request

USER_AGENT = 'Mozilla/5.0 (Windows NT 10.0; Win64; x64)'
HEADERS = {'User-Agent': USER_AGENT, 'Accept': 'text/markdown'}


def fetch_page_markdown(url, timeout=40, retries=4):
    """返回 (status, text)。
    status 取值：OK / HTTPxxx / FAILxxx。"""
    last_status = 'FAIL:unknown'
    for attempt in range(retries):
        try:
            time.sleep(0.4)
            request = urllib.request.Request(url, headers=HEADERS)
            with urllib.request.urlopen(request, timeout=timeout) as response:
                if response.status == 429:
                    time.sleep(10)
                    last_status = 'HTTP429'
                    continue
                body = response.read().decode('utf-8', 'replace')
            body = body.replace('\r\n', '\n').replace('\r', '\n')
            return 'OK', body
        except Exception as error:
            last_status = 'FAIL:' + repr(error)[:80]
            time.sleep(3)
    return last_status, ''


def extract_frontmatter(text):
    match = re.search(r'\A---\n(.*?)\n---\n', text, re.S)
    return match.group(1) if match else ''


def frontmatter_field(frontmatter_text, key):
    match = re.search(r'(?m)^' + re.escape(key) + r':\s*(.+?)\s*$', frontmatter_text)
    return match.group(1).strip() if match else ''


def normalize_whitespace(text):
    return re.sub(r'\s+', ' ', text or '').strip()


def extract_title(text):
    match = re.search(r'(?m)^#\s+(.+?)\s*$', text)
    return (match.group(1).strip() if match else '')


def extract_synopsis(text):
    """标题行后第一段正文。"""
    match = re.search(r'(?m)^#\s+\S+\s*\n(.*?)(?=\n##\s|\Z)', text, re.S)
    if not match:
        return ''
    synopsis_block = match.group(1)
    synopsis_block = re.sub(r'(?m)^\s*-\s*模块:.*$', '', synopsis_block)
    synopsis_block = re.sub(r'\[([^\]]*)\]\([^)]*\)', r'\1', synopsis_block)
    for line in synopsis_block.splitlines():
        stripped = normalize_whitespace(line)
        if stripped and not stripped.startswith(('-', '|', '!', '<')):
            return stripped
    return ''


def extract_syntax(text):
    """## 语法条目内全部 Syntax 代码块原文拼接（中英文标题皆可）。"""
    match = re.search(r'(?m)^##\s+(语法|syntax)\s*$\n(.*?)(?=^##\s|\Z)', text, re.S | re.I)
    if not match:
        return ''
    code_blocks = re.findall(r'```Syntax\n(.*?)```', match.group(2), re.S)
    return '\n'.join(normalize_whitespace(code_block) for code_block in code_blocks)


def extract_example1(text):
    """第一个示例条目：标题和首个 powershell 代码块。"""
    match = re.search(r'(?m)^##\s+(示例|examples)\s*$\n(.*?)(?=^##\s|\Z)', text, re.S | re.I)
    if not match:
        return ''
    body = match.group(2)
    title_match = re.search(r'(?m)^###\s+(.+?)\s*$', body)
    if not title_match:
        return ''
    title = re.sub(r'^\s*(示例|Example)\s*\d+\s*[:：\-–—.]?\s*', '', title_match.group(1)).strip()
    example_body = body[title_match.end():]
    example_body = re.split(r'(?m)^###\s+', example_body, maxsplit=1)[0]
    code_match = re.search(r'```powershell\n(.*?)```', example_body, re.S)
    code = normalize_whitespace(code_match.group(1)) if code_match else ''
    return title + '\n' + code


def snapshot_page(url):
    """获取页面内容并提取全部字段，返回 dict（含 _status）。"""
    status, text = fetch_page_markdown(url)
    if status != 'OK':
        return {'_status': status}
    frontmatter_text = extract_frontmatter(text)
    return {
        '_status': 'OK',
        'title': extract_title(text),
        'moniker': frontmatter_field(frontmatter_text, 'defaultMoniker') or frontmatter_field(frontmatter_text, 'default_moniker'),
        'git_commit': frontmatter_field(frontmatter_text, 'git_commit_id'),
        'ms_date': frontmatter_field(frontmatter_text, 'ms.date'),
        'synopsis': extract_synopsis(text),
        'syntax': extract_syntax(text),
        'example1': extract_example1(text),
    }
