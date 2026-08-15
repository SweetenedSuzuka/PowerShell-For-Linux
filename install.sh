#!/bin/sh
# 安装脚本：编译 powershell 并装进 PREFIX/bin。
#
# 用法：
#   ./install.sh              装到 /usr/local/bin（权限不足时用 sudo ./install.sh）
#   ./install.sh ~/.local     装到 ~/.local/bin（前提是它已在 PATH 里）
set -e

PREFIX="${1:-/usr/local}"
BINDIR="$PREFIX/bin"

echo "==> 编译 powershell"
go build -o powershell .

echo "==> 安装到 $BINDIR/powershell"
install -d "$BINDIR"
install -m 0755 powershell "$BINDIR/powershell"

echo "==> 完成。敲 powershell 进入命令行（exit 或 Ctrl+D 退出）"
