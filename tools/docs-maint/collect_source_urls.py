"""从两份中文详解文档收集每个条目的出处 URL。
在脚本目录输出 sources.json，结构为：
{name: {file, version, module, urls: [...]}}
无出处的条目会另存至 uncovered.txt。"""
import json
import os
import re

SCRIPT_DIR = os.path.dirname(os.path.abspath(__file__))
REPO_ROOT = os.path.dirname(os.path.dirname(SCRIPT_DIR))

DOC_PATHS = [os.path.join(REPO_ROOT, 'docs', '指令详解-原版跨平台指令.md'),
             os.path.join(REPO_ROOT, 'docs', '指令详解-原版Windows指令.md')]


def main():
    sections_by_name, uncovered = {}, []
    for doc_path in DOC_PATHS:
        doc_text = open(doc_path, encoding='utf-8').read()
        for section in re.split(r'(?m)^### ', doc_text)[1:]:
            section_name = section.splitlines()[0].strip()
            match_version = re.search(r'(?m)^版本：(.+?)\s*$', section)
            match_module = re.search(r'(?m)^模块：(.+?)\s*$', section)
            match_source = re.search(r'(?m)^出处：(.+?)\s*$', section)
            urls = re.findall(r'https://learn\.microsoft\.com/[^\s)]+', match_source.group(1)) if match_source else []
            urls = [url.rstrip(').,') for url in urls]
            if not urls:
                uncovered.append(section_name)
                continue
            sections_by_name[section_name] = {
                'file': os.path.basename(doc_path),
                'version': match_version.group(1).strip() if match_version else '',
                'module': match_module.group(1).strip() if match_module else '',
                'urls': urls,
            }
    with open(os.path.join(SCRIPT_DIR, 'sources.json'), 'w', encoding='utf-8') as sources_file:
        json.dump(sections_by_name, sources_file, ensure_ascii=False, indent=1)
    with open(os.path.join(SCRIPT_DIR, 'uncovered.txt'), 'w', encoding='utf-8') as uncovered_file:
        uncovered_file.write('\n'.join(sorted(set(uncovered))) + '\n')
    view_counts = {}
    for section_info in sections_by_name.values():
        for url in section_info['urls']:
            view_match = re.search(r'view=([A-Za-z0-9.\-]+)', url)
            view_name = view_match.group(1) if view_match else 'NOVIEW'
            view_counts[view_name] = view_counts.get(view_name, 0) + 1
    print('sections:', len(sections_by_name), 'uncovered:', len(set(uncovered)))
    print('view_counts:', view_counts)


if __name__ == '__main__':
    main()
