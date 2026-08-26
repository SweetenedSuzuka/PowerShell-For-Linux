package lang

// 中文提示
var zh = map[Msg]string{
	// ---- shell ----
	MsgBannerDesktop: "Linux PowerShell by SweetenedSuzuka\n版权没有(C) MacroHard Corporation。无法保留所有权利。\n\n尝试新的跨平台 PowerShell https://aka.ms/pscore6\n别想了，你试不了……等等，PowerShell 7好像还真支持Linux。",
	MsgBannerCore:    "PowerShell For Linux by SweetenedSuzuka\n版权没有(C) MacroHard Corporation。无法保留所有权利。\n\n输入 help 查看帮助。",
	MsgUsage: `用法:
  powershell [-Version 5.1|7] [-NoLogo] [-NoProfile] [-Command <命令>] [-File <脚本>]

参数:
  -Version <5.1|7>  选择 PowerShell 风格（5.X 或 7.X）
  -NoLogo           不显示启动横幅
  -NoProfile        不加载启动脚本
  -Command <命令>   执行命令后退出（- 表示从 stdin 读取）
  -File <脚本>      执行 .ps1 脚本后退出
  -Help, -?         显示本帮助

运行时切换风格:
  Set-PSVersion 5.1   切到 Windows PowerShell 5.X 风格
  Set-PSVersion 7     切到 PowerShell 7.X 风格
`,
	MsgReadonlyVar:      "无法对只读变量 $%s 赋值",
	MsgDriveUnsupported: "不支持的盘符 %s:。Linux 上只用 C 盘表示系统盘，没有其它盘的概念。",

	// ---- parser ----
	MsgParseExpectText:          "期望 '%s'，实际遇到 %s",
	MsgParseExpectWhat:          "期望 %s，实际遇到 %s",
	MsgParseUnexpectedAfter:     "语句后出现意外的 token：%s",
	MsgParseCatchAfterTry:       "catch 必须跟在 try 之后",
	MsgParseTryMissingHandler:   "try 语句缺少 catch 或 finally 块。",
	MsgParseFinallyAfterTry:     "finally 必须跟在 try 之后",
	MsgParseAssignOp:            "期望赋值运算符",
	MsgParseExpectBrace:         "期望 '{'，实际遇到 %s",
	MsgParseForeachVar:          "foreach 需要循环变量",
	MsgParseForeachIn:           "foreach 需要 'in'",
	MsgParseDoWhile:             "do 需要 while (cond)",
	MsgParseFuncParam:           "函数参数应为 $变量名",
	MsgParseFuncName:            "函数名",
	MsgParseParamList:           "param 需要参数列表 (...) ",
	MsgParsePipeRightCmd:        "管道右侧必须是命令",
	MsgParseCmpOp:               "意外的比较运算符 -%s",
	MsgParseTernaryColon:        "三元运算符缺少 ':'",
	MsgParseIncDecVar:           "增量/减量运算符只能用于变量",
	MsgParseOpMissingLeft:       "运算符 -%s 缺少左操作数",
	MsgParseUnexpectedArg:       "意外的参数 -%s",
	MsgParseInputRedirect:       "PowerShell 不支持 '<' 输入重定向",
	MsgParseRedirectAfterCmd:    "重定向 '>' 只能用在命令之后",
	MsgParseUnexpectedOp:        "意外的运算符 %s",
	MsgParseTypeLiteralName:     "类型字面量需要类型名",
	MsgParseTypeLiteralRbracket: "类型字面量需要 ']'",
	MsgParseParamTypeName:       "形参类型标注需要类型名",
	MsgParseParamTypeRbracket:   "形参类型标注需要 ']'",
	MsgConvertFail:              "无法将值“%s”转换为类型“%s”。",
	MsgTypeUnknown:              "无法找到类型 %s。",
	MsgStaticMemberNotFound:     "类型 %s 中不存在成员 %s。",
	MsgBindConvertFail:          "无法把实参“%s”转换成形参 %s 声明的类型“%s”。",
	MsgInvokeTargetMissing:      "调用运算符 & 需要调用目标。",
	MsgParseUnexpectedToken:     "意外的 token：%s",
	MsgParseHashEntry:           "哈希表项需要 '=' 或 ':'",
	MsgTokEOF:                   "输入结束",
	MsgTokNewline:               "换行",
	MsgTokWord:                  "“%s”",
	MsgTokNumber:                "数字 %s",
	MsgTokString:                "字符串",
	MsgTokOp:                    "运算符 %s",

	// ---- eval ----
	MsgDivideByZero:    "尝试除以零",
	MsgFormatIndexOut:  "格式串占位符 %q 越界（参数个数 %d）",
	MsgParamOnlyInFunc: "param 语句只能在函数或脚本中使用",
	MsgRedirectWrite:   "无法写入重定向目标 %s：%v",
	MsgCommandNotFound: "无法将“%s”项识别为 cmdlet、函数、脚本文件或可运行程序的名称。",
	MsgScriptReadFail:  "无法读取脚本 %s：%v",
	MsgScriptParseFail: "脚本 %s 解析错误：%v",
	MsgIncompleteInput: "输入不完整：缺少块结尾或续行内容。",

	// ---- builtin ----
	MsgBindNoParam:         "找不到与参数名称 \"-%s\" 匹配的参数。",
	MsgBindSwitchNoValue:   "参数 \"-%s\" 缺少值。",
	MsgWriteErrorPrefix:    "错误:",
	MsgWriteVerbosePrefix:  "详细:",
	MsgWriteWarningPrefix:  "警告:",
	MsgWriteDebugPrefix:    "调试:",
	MsgDateParseFail:       "无法解析日期时间：%s",
	MsgPathNotFoundFmt:     "找不到路径 %s。",
	MsgPathNotFoundForSet:  "找不到路径 %s，因为该路径不存在。",
	MsgCannotWrite:         "无法写入 %s。",
	MsgCannotOpen:          "无法打开 %s。",
	MsgCannotCreate:        "无法创建 %s：%v",
	MsgCannotDelete:        "无法删除 %s：%v",
	MsgCopyDestNotDir:      "复制多个源时目标 %s 必须是已存在的目录。",
	MsgCopyNeedsRecurse:    "目录 %s 需要 -Recurse。",
	MsgCannotRename:        "无法重命名 %s：%v",
	MsgRenameDestExists:    "无法重命名为 %s，因为该项已存在。",
	MsgWhatIfPerform:       "What if: 对目标“%s”执行 %s。",
	MsgConfirmPrompt:       "确认\n操作：%s\n目标：%s\n[Y] 是 [A] 全部是 [N] 否 [L] 全部否 [?] 帮助（默认为“Y”）: ",
	MsgCannotClear:         "无法清空 %s。",
	MsgPropNotFound:        "路径 %s 不存在属性 %s。",
	MsgVarExists:           "变量 $%s 已存在。",
	MsgAliasExists:         "别名 %s 已存在。",
	MsgServiceNeedSystemd:  "Get-Service : 需要 systemd（systemctl）。",
	MsgServiceListFail:     "Get-Service : 无法读取服务列表。",
	MsgPingNotFound:        "Test-Connection : 找不到 ping 命令。",
	MsgFormatHexNotFound:   "Format-Hex : 找不到路径 %s。",
	MsgNewObjectBadType:    "New-Object : 不支持的类型 %s。",
	MsgHelpAliasTo:         "别名 %s → %s",
	MsgHelpNotFound:        "找不到名为 %s 的帮助。",
	MsgHelpNameHeader:      "名称",
	MsgHelpSyntaxHeader:    "语法",
	MsgHelpAliasHeader:     "别名",
	MsgProcIdNotFound:      "找不到进程标识符为 %s 的进程。",
	MsgProcNameNotFound:    "找不到名为 \"%s\" 的进程。",
	MsgPSVersionBad:        "Set-PSVersion : 不支持的版本 %q（请用 5.1 或 7.x）",
	MsgPSVersionSet:        "已切换到 %s 风格。",
	MsgHashAlgoUnsupported: "Get-FileHash : 不支持的算法 %s",

	// ---- object 与内部提示 ----
	MsgUnsupported: "不支持",

	// ---- external 与 main ----
	MsgExternalExecFail: "无法执行命令 %s：%v",
	MsgFlagParseFail:    "参数解析失败。",
}
