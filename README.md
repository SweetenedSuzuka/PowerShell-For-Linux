# PowerShell For Linux

> 众所周知，PowerShell是使用最广泛的命令行Shell之一。  
> 但是它却不能在Linux上使用。  
> 凭什么只有Windows能用PowerShell？  
> 因此，为了解决这个问题，我们制作了这款PowerShell For Linux。  
> ~~有这么好的项目进入Linux生态，猫听了连夜换成HarmonyOS。~~

**在 Linux 上运行的 PowerShell 风格解释器。**

把 PowerShell 的命令、对象管道和脚本搬到 Linux 上，操作的是真实的 Linux 文件系统。  
支持 Windows PowerShell 5.X 与 7.X 两套命令格式。

**仅依赖 Go 标准库，零第三方依赖。**

> 部分命令通过Go实现，部分命令会映射为 Linux 的相应命令，可以做到用 PowerShell 风格的命令控制 Linux。  
> ~~俗称没苦硬吃。~~

> [!NOTE]
> 本项目仅供娱乐，与 Microsoft 出品的 PowerShell 无关，亦不建议真的在 Linux 上使用。

## 特性

- **对象管道** —— 命令之间正常传递对象，`Where-Object`、`Sort-Object`、`Select-Object`、`ForEach-Object` 等照常可用。  
- **双命令格式** —— 5.X 和 7.X 的横幅、`$PSVersionTable`、别名、自动变量随切换联动。  
- **多语言支持** —— 横幅、帮助与报错信息会按系统变量 `LANGUAGE`/`LC_ALL`/`LC_MESSAGES`/`LANG` 依次匹配界面语言；目前支持中文和英语，未被支持的语言会使用默认的中文，日期显示与区域信息跟随语言设置。  
- **外部命令透传** —— 没内置的命令直接在系统 PATH 里执行，退出码写进 `$LASTEXITCODE`。  
- **支持 115 个 cmdlet 命令** —— 支持 115 个内置 cmdlet 以及它们的 77 个别名，覆盖文件、路径、进程、服务、系统信息、JSON/CSV 转换、格式化等。  
- **脚本与语法** —— `.ps1` 脚本可直接执行；对象、字符串与日期的方法可直接调用；支持 `try/catch/finally`、`throw`、脚本块、`param()` 参数声明、作用域修饰符（`$script:`/`$global:`/`$local:`）、`switch` 分支等语言特性。  
- **行编辑** —— 历史、Tab 补全、多行续行。  
- **零外部依赖** —— 静态二进制，目标机不用装任何额外的东西。  

## 实现方式

将 PowerShell 的命令分为两类，其中一部分映射为对应的 Linux 命令，另一部分用 Go 实现，实现在 Linux 上使用 PowerShell 风格的命令控制系统。

- **Go 原生复现** —— 对象管道本身，以及文件、路径、格式化、JSON/CSV、对象操作这些命令（如 `Get-ChildItem`、`Where-Object`、`ConvertTo-Json` 等），在 Go 里直接产生和使用对象。  
- **映射系统命令** —— 系统层面的操作直接调用 Linux 上的原生工具：  
  - `Get-Service` → `systemctl`  
  - `Test-Connection` → `ping`  
  - `Get/Set-Clipboard` → `xclip`/`xsel`  
  - `Set-TimeZone` → `timedatectl`  
  - `Set-Date` → `date`  
  - 诸如此类，需要 root 的自动加 `sudo`。  
- **其它命令** —— 直接透传执行系统 PATH 里的命令（`sudo`、`mkdir`、`grep` 照常能用）。

## 构建

手动构建，不涉及安装：

```sh
go build -o powershell .        # 在当前系统上直接构建

# 在其它系统上交叉编译出 Linux 静态二进制
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o powershell .
```

## 安装

用脚本装进 PATH：

```sh
./install.sh              # 装到 /usr/local/bin（权限不足时用 sudo ./install.sh）
./install.sh ~/.local     # 装到 ~/.local/bin（前提是它已在 PATH 里）
```

装好后直接敲 `powershell` 进入命令行，提示符是 Windows 风格的 `PS C:\>`；敲 `exit`（或按 Ctrl+D）退出。

## 用法

```text
powershell [-Version 5.1|7] [-NoLogo] [-NoProfile] [-Command <命令>] [-File <脚本>]
```

| 选项 | 说明 |
| :--- | :--- |
| `-Version <5.1\|7>` | 用 5.X 还是 7.X 的命令格式启动（默认 7.X） |
| `-Command <命令>` | 执行命令后退出（`-` 表示从标准输入读） |
| `-File <脚本>` | 执行 `.ps1` 脚本后退出 |
| `-NoLogo` | 不显示启动横幅 |
| `-NoProfile` | 不加载启动脚本 |
| `-?` / `-Help` | 显示帮助 |

运行时切命令格式：`Set-PSVersion 5.1`（切到 5.X）、`Set-PSVersion 7`（切回 7.X）。

`C:` 表示系统盘（就是根目录），`cd C:\tmp` 等同 `cd /tmp`；其它盘符（`D:` 等）在 Linux 上不存在，输入会报错。

## 内置命令

内置 115 个 cmdlet 命令，外加 77 个别名。每条命令的用法、版本和发行版要求，以及它对应 bash 的什么写法，见 [docs/指令参考.md](docs/指令参考.md)。  
逐参数、逐字段的映射规则见 [docs/指令详解.md](docs/指令详解.md)。

## 脚本

`.ps1` 脚本可以直接执行：

```sh
powershell -File 脚本.ps1
```

或进入命令行后运行 `.\脚本.ps1`。

支持范围详见 `docs/` 下的文档。  
总体来说，组成脚本的指令行为和原版 PowerShell 一致，或者出入不是很大（如一些直接映射为 bash 指令的命令行输出可能存在差异），通常都可以执行。

## 管道与重定向

```text
|           把左侧命令的输出传给右侧命令（对象管道）
> file      把输出写进文件（覆盖）
>> file     把输出写进文件（追加）
2> file     把错误流写进文件
2>$null     丢弃错误
&&          左边成功了才执行右边（$? 为真）
||          左边失败了才执行右边
```

## 文件

| 路径 | 用途 |
| :--- | :--- |
| `~/.powershell_history` | 命令历史（纯文本，每行一条） |

## 文档

- [docs/指令参考.md](docs/指令参考.md) —— 面向使用者：每条指令的映射、版本、发行版与例子。
- [docs/指令详解.md](docs/指令详解.md) —— 开发文档：每条指令逐参数、逐字段的映射规则。
- [docs/未实现指令.md](docs/未实现指令.md) —— 未实现指令的完整清单及原因。
- [docs/指令分类.md](docs/指令分类.md) —— 三个核心模块指令的全量归类。

## 测试

```sh
go test ./...                                           # 单元测试
./powershell -NoLogo -NoProfile -File test/核验.ps1     # 逐条核验参考文档里的示例
./powershell -NoLogo -NoProfile -File test/内容回归.ps1  # 内容级回归：检查输出的实际内容
```

## 项目结构

```text
powershellForLinux/
├── main.go              # 入口：参数解析，启动 REPL / -Command / -File
├── install.sh           # 安装脚本：编译并装进 PATH
├── go.mod
├── internal/
│   ├── lexer/           # 词法分析
│   ├── parser/          # 语法分析
│   ├── ast/             # 语法树
│   ├── eval/            # 求值器：表达式、语句、管道、命令调度
│   ├── builtin/         # 内置 cmdlet（Go 实现 + 映射系统命令）
│   ├── object/          # 对象模型（PSObject）与格式化输出
│   ├── lang/            # 多语言文本
│   ├── shell/           # 会话状态：变量、别名、5.X/7.X 风格、横幅、提示符
│   ├── external/        # 外部命令透传
│   └── repl/            # 交互式 REPL（行编辑、历史、补全）
├── docs/
│   ├── 指令参考.md       # 用户文档：每条指令的分类、行为、两个案例
│   ├── 指令详解.md       # 开发文档：逐参数、逐字段
│   ├── 指令分类.md       # 三个核心模块指令的归类
│   └── 未实现指令.md     # 未实现指令清单及原因
└── test/
    ├── 核验.ps1          # 逐条核验参考文档里的示例
    ├── 内容回归.ps1      # 回归测试脚本
    ├── 功能演示.ps1      # 语言特性演示
    ├── 格式测试5.ps1     # 5.X 命令格式测试
    └── 格式测试7.ps1     # 7.X 命令格式测试
```