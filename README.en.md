# PowerShell For Linux

[中文](README.md) | English

> As everyone knows, PowerShell is one of the most widely used command-line shells.  
> Yet while it runs on Windows with no dependencies at all, on Linux it demands a .NET installation.  
> That is hardly in the spirit of Linux.
> So to fix this, we built PowerShell For Linux.  
> ~~With such a fine project joining the Linux ecosystem, even cats would switch to HarmonyOS overnight.~~

**A PowerShell-style interpreter that runs on Linux.**

It brings PowerShell's commands, object pipeline, and scripting onto Linux, working against the real Linux filesystem.  
Supports both the Windows PowerShell 5.X and 7.X command formats.

**Depends only on the Go standard library — zero third-party dependencies.**

> Some commands are implemented in Go; others are mapped to their Linux counterparts, so you can control Linux with PowerShell-style commands.  
> ~~Colloquially known as seeking out suffering for no reason.~~

> [!NOTE]
> This project is for fun only. It is not affiliated with Microsoft's PowerShell, and you really shouldn't use it on Linux.

## Features

- **Object pipeline** — objects flow between commands as they should; `Where-Object`, `Sort-Object`, `Select-Object`, `ForEach-Object` and friends all work normally.
- **Dual command formats** — banner, `$PSVersionTable`, aliases, and automatic variables all follow the 5.X / 7.X switch together.
- **Multi-language support** — the banner, help text, and error messages match your interface language against the system variables `LANGUAGE`/`LC_ALL`/`LC_MESSAGES`/`LANG` in order; Chinese and English are supported for now, and unsupported languages fall back to Chinese. Date display and locale info follow the language setting.
- **External command passthrough** — anything not built in is executed straight from the system PATH, with its exit code recorded in `$LASTEXITCODE`.
- **115 cmdlets** — 115 built-in cmdlets plus 77 of their aliases, covering files, paths, processes, services, system information, JSON/CSV conversion, formatting, and more.
- **Scripts and syntax** — `.ps1` scripts run directly; methods on objects, strings, and dates can be called directly; language features include `try/catch/finally`, `throw`, script blocks, `param()` declarations, scope modifiers (`$script:`/`$global:`/`$local:`), `switch` branches, and more.
- **Line editing** — history, Tab completion, multi-line continuation.
- **Zero external dependencies** — a static binary; the target machine needs nothing else installed.

## How it works

Commands fall into two groups: some are mapped to their corresponding Linux commands, others are implemented in Go, so a system can be controlled with PowerShell-style commands.

- **Native Go implementation** — the object pipeline itself, along with file, path, formatting, JSON/CSV, and object-manipulation commands (such as `Get-ChildItem`, `Where-Object`, `ConvertTo-Json`), produce and consume real objects inside Go.
- **Mapped system commands** — system-level operations call the native Linux tools directly:
  - `Get-Service` → `systemctl`
  - `Test-Connection` → `ping`
  - `Get/Set-Clipboard` → `xclip`/`xsel`
  - `Set-TimeZone` → `timedatectl`
  - `Set-Date` → `date`
  - and so on; where root is needed, `sudo` is added automatically.
- **Everything else** — passed through untouched to whatever is on the system PATH (`sudo`, `mkdir`, and `grep` work just as you'd expect).

## Building

Build by hand; no installation step involved:

```sh
go build -o powershell .        # build directly on the current system

# cross-compile a static Linux binary from another OS
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o powershell .
```

## Installing

Install into PATH with the script:

```sh
./install.sh              # installs to /usr/local/bin (use sudo ./install.sh if permissions fall short)
./install.sh ~/.local     # installs to ~/.local/bin (provided it's already on PATH)
```

Once installed, type `powershell` to enter the shell. The prompt is the Windows-style `PS C:\>`; type `exit` (or press Ctrl+D) to leave.

## Usage

```text
powershell [-Version 5.1|7] [-NoLogo] [-NoProfile] [-Command <command>] [-File <script>] [-NoExit] [-NonInteractive] [-ExecutionPolicy <policy>] [-WorkingDirectory <dir>]
```

| Option | Description |
| :--- | :--- |
| `-Version <5.1\|7>` | Start with the 5.X or 7.X command format (default: 7.X) |
| `-Command <command>` | Execute the command, then exit (`-` reads from standard input) |
| `-File <script>` | Execute a `.ps1` script, then exit |
| `-NoLogo` | Skip the startup banner |
| `-NoProfile` | Do not load the profile script |
| `-NoExit` | Stay interactive after running, do not exit |
| `-NonInteractive` | Run non-interactively: confirmation prompts are declined, input reads fail |
| `-ExecutionPolicy <policy>` | Execution policy (value checked only, nothing restricted) |
| `-WorkingDirectory <dir>` | Start in this directory |
| `-?` / `-Help` | Show help |

Switch format at runtime: `Set-PSVersion 5.1` (to 5.X), `Set-PSVersion 7` (back to 7.X).

`C:` stands for the system drive — which is simply the root directory — so `cd C:\tmp` is the same as `cd /tmp`. Other drive letters (`D:` and so on) don't exist on Linux and produce an error.

## Built-in commands

115 cmdlets are built in, plus 77 aliases. For each command's usage, version and distro requirements, and its bash equivalent, see [docs/CommandReference.en.md](docs/CommandReference.en.md).  
Details live in [docs/OriginalCrossPlatformDetails.en.md](docs/OriginalCrossPlatformDetails.en.md) and [docs/OriginalWindowsDetails.en.md](docs/OriginalWindowsDetails.en.md), showing original commands' purposes, version differences and official examples alongside details of this program's implementations.

## Scripts

`.ps1` scripts run directly:

```sh
powershell -File script.ps1
```

Or, once inside the shell, run `.\script.ps1`.

See the documents under `docs/` for the full coverage picture.  
Broadly speaking, the commands making up a script behave the same as genuine PowerShell or close enough to it (commands mapped straight to bash instructions may differ in output), so most scripts will run.

## Pipeline and redirection

```text
|           Passes the left command's output to the right command (object pipeline)
> file      Writes output to the file (overwrite)
>> file     Writes output to the file (append)
2> file     Writes the error stream to the file
2>$null     Discards errors
&&          Runs the right side only if the left succeeded ($? is true)
||          Runs the right side only if the left failed
```

## Files

| Path | Purpose |
| :--- | :--- |
| `~/.powershell_history` | Command history (plain text, one entry per line) |

## Documentation

- [docs/CommandReference.en.md](docs/CommandReference.en.md) — For users: mapping, versions, distro requirements, and examples for every command.
- [docs/OriginalCrossPlatformDetails.en.md](docs/OriginalCrossPlatformDetails.en.md) — Original cross-platform commands: purposes, version differences, usage, official examples.
- [docs/OriginalWindowsDetails.en.md](docs/OriginalWindowsDetails.en.md) — Original Windows-only commands: purposes, version differences, usage, official examples.

## Testing

```sh
go test ./...                                           # unit tests
./powershell -NoLogo -NoProfile -File test/核验.ps1     # runs every example from the reference document
./powershell -NoLogo -NoProfile -File test/内容回归.ps1  # Regression test script
```

## Project layout

```text
PowerShell-For-Linux/
├── main.go              # Entry point: argument parsing, starts the REPL / -Command / -File
├── install.sh           # Install script: builds and installs into PATH
├── go.mod
├── internal/
│   ├── lexer/           # Lexical analysis
│   ├── parser/          # Parsing
│   ├── ast/             # Syntax tree
│   ├── eval/            # Evaluator: expressions, statements, pipelines, command dispatch
│   ├── builtin/         # Built-in cmdlets (Go implementations + mapped system commands)
│   ├── object/          # Object model (PSObject) and formatted output
│   ├── lang/            # Multi-language strings
│   ├── shell/           # Session state: variables, aliases, 5.X/7.X style, banner, prompt
│   ├── external/        # External command passthrough
│   └── repl/            # Interactive REPL (line editing, history, completion)
├── docs/
│   ├── 指令参考.md                  # Categories, behavior, and examples for every command in PowerShell For Linux
│   ├── 指令详解-原版跨平台指令.md   # Original PowerShell cross-platform commands: purposes, version differences, usage, official examples, compared against this program's implementations
│   └── 指令详解-原版Windows指令.md  # Original PowerShell Windows-only commands: purposes, version differences, usage, official examples, compared against this program's implementations
├── test/
│   ├── 核验.ps1          # Verifies every example from the reference document
│   ├── 内容回归.ps1      # Regression test script
│   ├── 功能演示.ps1      # Language feature showcase
│   ├── 格式测试5.ps1     # Tests for the 5.X command format
│   └── 格式测试7.ps1     # Tests for the 7.X command format
└── tools/
    └── docs-maint/       # Document maintenance tool
```
