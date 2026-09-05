"""
对比快照与官方现状，列出有改动的条目。
在脚本目录产生 official_changes.md，有变化时退出码为 1。
用法示例：
  python tools/docs-maint/check_official_changes.py --only Get-ChildItem
  python tools/docs-maint/check_official_changes.py --limit 20
更多选项见 --help。
"""
import argparse
import concurrent.futures
import datetime
import difflib
import json
import os
import sys

SCRIPT_DIR = os.path.dirname(os.path.abspath(__file__))
sys.path.insert(0, SCRIPT_DIR)
from official_pages import snapshot_page

COMPARED_FIELDS = ['synopsis', 'syntax', 'example1']
FIELD_LABELS = {'synopsis': '简介', 'syntax': '语法', 'example1': '示例1'}


def short_diff(old_text, new_text, context=6):
    diff_lines = list(difflib.unified_diff((old_text or '').splitlines(), (new_text or '').splitlines(),
                                           lineterm='', n=context))
    return '\n'.join(diff_lines[:2 * context + 6])


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument('--only', default='', help='只处理指定指令，逗号分隔')
    parser.add_argument('--limit', type=int, default=0, help='只处理前 N 条，0 表示全部')
    parser.add_argument('--workers', type=int, default=3, help='并发抓取数')
    parser.add_argument('--reference', default=os.path.join(SCRIPT_DIR, 'snapshots.json'), help='参考快照路径')
    parser.add_argument('--report', default=os.path.join(SCRIPT_DIR, 'official_changes.md'), help='结果输出目录')
    args = parser.parse_args()

    sources = json.load(open(os.path.join(SCRIPT_DIR, 'sources.json'), encoding='utf-8'))
    reference = json.load(open(args.reference, encoding='utf-8'))
    command_names = sorted(reference)
    if args.only:
        only_names = set(raw_name.strip() for raw_name in args.only.split(','))
        command_names = [command_name for command_name in command_names if command_name in only_names]
    if args.limit:
        command_names = command_names[:args.limit]

    page_targets = [(command_name, url) for command_name in command_names for url in reference[command_name]['pages']]
    print('names:', len(command_names), 'urls:', len(page_targets), flush=True)

    def snapshot_one(target):
        command_name, url = target
        return command_name, url, snapshot_page(url)

    current_snapshots = {}
    completed_count = 0
    with concurrent.futures.ThreadPoolExecutor(max_workers=args.workers) as executor:
        for command_name, url, page_snapshot in executor.map(snapshot_one, page_targets):
            current_snapshots.setdefault(command_name, {})[url] = page_snapshot
            completed_count += 1
            if completed_count % 20 == 0:
                print('...', completed_count, '/', len(page_targets), flush=True)

    report_lines = ['# 微软官方文档产生的变化汇总', '',
                    '参考快照：`%s`，检查时间：%s。' % (
                        args.reference, datetime.datetime.now().isoformat(timespec='seconds')), '']
    changed_command_count = 0

    for command_name in command_names:
        doc_file = sources.get(command_name, {}).get('file', reference[command_name].get('file', '?'))
        change_blocks = []
        for url, reference_page in reference[command_name]['pages'].items():
            current_page = current_snapshots[command_name].get(url, {})
            view_label = url.rsplit('view=', 1)[-1] if 'view=' in url else url.rsplit('/', 1)[-1][:40]
            if current_page.get('_status') != 'OK':
                change_blocks.append('## %s（%s）\n\n- 页面获取失败：%s（参考快照 %s）。\n'
                                     % (command_name, view_label, current_page.get('_status'), reference_page.get('_status')))
                continue
            if reference_page.get('_status') != 'OK':
                change_blocks.append('## %s（%s）\n\n- 页面恢复正常（参考快照 %s）。\n'
                                     % (command_name, view_label, reference_page.get('_status')))
                continue
            if current_page.get('title') != reference_page.get('title'):
                change_blocks.append('## %s（%s）\n\n- 页面标题变化：%s → %s。\n'
                                     % (command_name, view_label, reference_page.get('title'), current_page.get('title')))
            if current_page.get('moniker') != reference_page.get('moniker'):
                change_blocks.append('## %s（%s）\n\n- moniker 变化：%s → %s。\n'
                                     % (command_name, view_label, reference_page.get('moniker'), current_page.get('moniker')))
            for field_name in COMPARED_FIELDS:
                if (current_page.get(field_name) or '') != (reference_page.get(field_name) or ''):
                    change_blocks.append('## %s（%s）%s变化\n\n- 文件：%s，条目 `### %s`，出处：%s\n\n```diff\n%s\n```\n'
                                         % (command_name, view_label, FIELD_LABELS[field_name], doc_file, command_name, url,
                                            short_diff(reference_page.get(field_name), current_page.get(field_name))))
            if current_page.get('git_commit') != reference_page.get('git_commit') and not any(
                    (current_page.get(field_name) or '') != (reference_page.get(field_name) or '') for field_name in COMPARED_FIELDS):
                change_blocks.append('## %s（%s）\n\n- 官方提交变化（%s → %s），但简介/语法/示例1没有变化。\n'
                                     % (command_name, view_label, (reference_page.get('git_commit') or '')[:8],
                                        (current_page.get('git_commit') or '')[:8]))
        if change_blocks:
            changed_command_count += 1
            report_lines += change_blocks

    report_lines.append('共 %d 条有变化（检查 %d 条）。' % (changed_command_count, len(command_names)))
    report_lines.append('')
    with open(args.report, 'w', encoding='utf-8', newline='\n') as report_file:
        report_file.write('\n'.join(report_lines))
    print('changed:', changed_command_count, '/', len(command_names), flush=True)
    sys.exit(1 if changed_command_count else 0)


if __name__ == '__main__':
    main()
