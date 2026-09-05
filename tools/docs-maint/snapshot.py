"""
逐页获取 sources.json 中的出处页，提取官方字段存快照。
默认获取全部，一般需要很长时间。
在脚本目录生成 snapshots.json 和 meta.json。
用法示例：
  python tools/docs-maint/snapshot.py --only Get-ChildItem,Get-Item
  python tools/docs-maint/snapshot.py --limit 10
更多选项见 --help。
"""
import argparse
import concurrent.futures
import datetime
import json
import os
import sys

SCRIPT_DIR = os.path.dirname(os.path.abspath(__file__))
sys.path.insert(0, SCRIPT_DIR)
from official_pages import snapshot_page


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument('--only', default='', help='只处理指定指令，逗号分隔')
    parser.add_argument('--limit', type=int, default=0, help='只处理前 N 条，0 表示全部')
    parser.add_argument('--workers', type=int, default=3, help='并发抓取数')
    args = parser.parse_args()

    sources = json.load(open(os.path.join(SCRIPT_DIR, 'sources.json'), encoding='utf-8'))
    command_names = sorted(sources)
    if args.only:
        only_names = set(raw_name.strip() for raw_name in args.only.split(','))
        command_names = [command_name for command_name in command_names if command_name in only_names]
    if args.limit:
        command_names = command_names[:args.limit]

    page_targets = [(command_name, url) for command_name in command_names for url in sources[command_name]['urls']]
    print('names:', len(command_names), 'urls:', len(page_targets), flush=True)

    try:
        reference = json.load(open(os.path.join(SCRIPT_DIR, 'snapshots.json'), encoding='utf-8'))
    except FileNotFoundError:
        reference = {}
    snapshots = dict(reference)

    def snapshot_one(target):
        command_name, url = target
        return command_name, url, snapshot_page(url)

    completed_count = 0
    with concurrent.futures.ThreadPoolExecutor(max_workers=args.workers) as executor:
        for command_name, url, page_snapshot in executor.map(snapshot_one, page_targets):
            record = snapshots.setdefault(command_name, {'file': sources[command_name]['file'], 'pages': {}})
            record['pages'][url] = page_snapshot
            completed_count += 1
            if completed_count % 20 == 0:
                print('...', completed_count, '/', len(page_targets), flush=True)
    with open(os.path.join(SCRIPT_DIR, 'snapshots.json'), 'w', encoding='utf-8') as snapshot_file:
        json.dump(snapshots, snapshot_file, ensure_ascii=False, indent=1)
    with open(os.path.join(SCRIPT_DIR, 'meta.json'), 'w', encoding='utf-8') as meta_file:
        json.dump({'taken_at': datetime.datetime.now().isoformat(timespec='seconds'),
                   'names': len(snapshots)}, meta_file, ensure_ascii=False, indent=1)
    failed_count = sum(1 for command_name in command_names for url, page in snapshots[command_name]['pages'].items()
                       if page.get('_status') != 'OK')
    print('done, failed pages:', failed_count, flush=True)


if __name__ == '__main__':
    main()
