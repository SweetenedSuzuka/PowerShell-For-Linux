# 文档维护工具

这个目录包含了一系列脚本，可以用于比对文档中对原版 PowerShell 的描述与微软的官方文档是否有出入。
作用是获取官方页面，与之前保存的参考快照进行对比，把所有改变的条目列出。

## 脚本说明

- `collect_source_urls.py` 从两份中文详解的 `出处：` 行收集 URL，写入 `sources.json`。
- `snapshot.py` 逐页获取官方 markdown，提取简介、语法、示例1、标题、moniker、ms_date、git_commit。
- 生成 `snapshots.json`（参考快照）与 `meta.json`。
- `check_official_changes.py` 重新获取页面，与参考快照比对，生成 `official_changes.md`（这个文件会被 .gitignore 排除），有变化时脚本退出码为 1。
- `official_pages.py` 负责页面获取与字段提取，供 `snapshot.py` 与 `check_official_changes.py` 导入。

## 使用方法

在仓库根目录执行：

```bash
python tools/docs-maint/collect_source_urls.py   # 文档出处变了才需运行。
python tools/docs-maint/snapshot.py              # 全部条目的检查需要较长时间。
python tools/docs-maint/snapshot.py --only Get-ChildItem
python tools/docs-maint/check_official_changes.py
python tools/docs-maint/check_official_changes.py --only Get-ChildItem,Get-Service
```

`snapshot.py` 与 `check_official_changes.py` 另有 `--limit`（限制获取的条目数量）与 `--workers`（并发数）。
`check_official_changes.py` 另有 `--reference` 与 `--report`（指定参考快照与报告路径）。
`collect_source_urls.py` 附带产生 `uncovered.txt`（无出处条目清单）。

报告每条列出文件、条目锚点、官方出处链接、前后 diff。
moniker 变化、标题变化、页面失效与恢复单独列出。

## 备忘

- 版本以页内 `defaultMoniker` 为准，而不是 URL 里的 view 参数。
- 回落有 301 跳转与 200 直接返回两种。
- 只检查 URL，发现不了后一种。
- Dism 等 server 视图页只有英文标题（`## Syntax`）。
