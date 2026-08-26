package lang

// 英文提示
var en = map[Msg]string{
	// ---- shell ----
	MsgBannerDesktop: "Linux PowerShell by SweetenedSuzuka\nCopyright (c) MacroHard Corporation. No rights reserved.\n\nTry the new cross-platform PowerShell https://aka.ms/pscore6\nDon't even try — you can't... wait, PowerShell 7 actually supports Linux.",
	MsgBannerCore:    "PowerShell For Linux by SweetenedSuzuka\nCopyright (c) MacroHard Corporation. No rights reserved.\n\nType 'help' to get help.",
	MsgUsage: `Usage:
  powershell [-Version 5.1|7] [-NoLogo] [-NoProfile] [-Command <command>] [-File <script>]

Options:
  -Version <5.1|7>  Pick the PowerShell style (5.X or 7.X)
  -NoLogo           Do not show the startup banner
  -NoProfile        Do not load the profile script
  -Command <cmd>    Run the command and exit (- reads from stdin)
  -File <script>    Run a .ps1 script and exit
  -Help, -?         Show this help

Switch style at runtime:
  Set-PSVersion 5.1   Switch to the Windows PowerShell 5.X style
  Set-PSVersion 7     Switch to the PowerShell 7.X style
`,
	MsgReadonlyVar:      "Cannot assign to read-only variable $%s.",
	MsgDriveUnsupported: "Unsupported drive %s:. Linux only uses C for the system drive; no other drives exist.",

	// ---- parser ----
	MsgParseExpectText:          "Expected '%s', got %s",
	MsgParseExpectWhat:          "Expected %s, got %s",
	MsgParseUnexpectedAfter:     "Unexpected token after statement: %s",
	MsgParseCatchAfterTry:       "catch must follow try",
	MsgParseTryMissingHandler:   "The Try statement is missing its Catch or Finally block.",
	MsgParseFinallyAfterTry:     "finally must follow try",
	MsgParseAssignOp:            "Expected an assignment operator",
	MsgParseExpectBrace:         "Expected '{', got %s",
	MsgParseForeachVar:          "foreach requires a loop variable",
	MsgParseForeachIn:           "foreach requires 'in'",
	MsgParseDoWhile:             "do requires while (cond)",
	MsgParseFuncParam:           "Function parameter must be a $variable",
	MsgParseFuncName:            "function name",
	MsgParseParamList:           "param requires a parameter list (...) ",
	MsgParsePipeRightCmd:        "Right side of pipeline must be a command",
	MsgParseCmpOp:               "Unexpected comparison operator -%s",
	MsgParseTernaryColon:        "Ternary operator is missing ':'",
	MsgParseIncDecVar:           "Increment/decrement operators apply to variables only",
	MsgParseOpMissingLeft:       "Operator -%s is missing its left operand",
	MsgParseUnexpectedArg:       "Unexpected argument -%s",
	MsgParseInputRedirect:       "PowerShell does not support '<' input redirection",
	MsgParseRedirectAfterCmd:    "Redirect '>' can only follow a command",
	MsgParseUnexpectedOp:        "Unexpected operator %s",
	MsgParseTypeLiteralName:     "Type literal requires a type name",
	MsgParseTypeLiteralRbracket: "Type literal requires ']'",
	MsgParseParamTypeName:       "Parameter type annotation requires a type name",
	MsgParseParamTypeRbracket:   "Parameter type annotation requires ']'",
	MsgParseNamedBlockPosition:  "begin/process/end blocks must appear together at the start of a function body",
	MsgParseNamedBlockDuplicate: "Duplicate %s block in function body",
	MsgConvertFail:              "Cannot convert value \"%s\" to type \"%s\".",
	MsgTypeUnknown:              "Unable to find type %s.",
	MsgStaticMemberNotFound:     "Type %s does not have a member %s.",
	MsgBindConvertFail:          "Cannot convert argument \"%s\" for parameter %s to the declared type \"%s\".",
	MsgInvokeTargetMissing:      "The call operator & requires a call target.",
	MsgParseUnexpectedToken:     "Unexpected token: %s",
	MsgParseHashEntry:           "Hashtable entry requires '=' or ':'",
	MsgTokEOF:                   "end of input",
	MsgTokNewline:               "newline",
	MsgTokWord:                  "\"%s\"",
	MsgTokNumber:                "number %s",
	MsgTokString:                "string",
	MsgTokOp:                    "operator %s",

	// ---- eval ----
	MsgDivideByZero:    "Attempted to divide by zero.",
	MsgFormatIndexOut:  "Format placeholder %q is out of range (%d arguments)",
	MsgIndexOutOfRange: "Index out of range: the array subscript is invalid.",
	MsgParamOnlyInFunc: "The param statement is only allowed in a function or script",
	MsgRedirectWrite:   "Cannot write to redirection target %s: %v",
	MsgCommandNotFound: "The term '%s' is not recognized as the name of a cmdlet, function, script file, or runnable program.",
	MsgScriptReadFail:  "Cannot read script %s: %v",
	MsgScriptParseFail: "Parse error in script %s: %v",
	MsgIncompleteInput: "Incomplete input: missing a block terminator or continuation.",

	// ---- builtin ----
	MsgBindNoParam:         "A parameter cannot be found that matches parameter name '-%s'.",
	MsgBindSwitchNoValue:   "Missing an argument for parameter '-%s'.",
	MsgWriteErrorPrefix:    "ERROR:",
	MsgWriteVerbosePrefix:  "VERBOSE:",
	MsgWriteWarningPrefix:  "WARNING:",
	MsgWriteDebugPrefix:    "DEBUG:",
	MsgDateParseFail:       "Cannot parse date and time: %s",
	MsgPathNotFoundFmt:     "Cannot find path '%s'.",
	MsgPathNotFoundForSet:  "Cannot find path '%s' because it does not exist.",
	MsgCannotWrite:         "Cannot write to %s.",
	MsgCannotOpen:          "Cannot open %s.",
	MsgCannotCreate:        "Cannot create %s: %v",
	MsgCannotDelete:        "Cannot delete %s: %v",
	MsgCopyDestNotDir:      "Destination %s must be an existing directory when copying multiple sources.",
	MsgCopyNeedsRecurse:    "Directory %s requires -Recurse.",
	MsgCannotRename:        "Cannot rename %s: %v",
	MsgRenameDestExists:    "Cannot rename to %s because it already exists.",
	MsgWhatIfPerform:       "What if: On target '%s', run %s.",
	MsgConfirmPrompt:       "Confirm\nOperation: %s\nTarget: %s\n[Y] Yes [A] Yes to All [N] No [L] No to All [?] Help (default is \"Y\"): ",
	MsgCannotClear:         "Cannot clear %s.",
	MsgPropNotFound:        "Path %s does not have property '%s'.",
	MsgVarExists:           "A variable with name $%s already exists.",
	MsgAliasExists:         "An alias with name %s already exists.",
	MsgServiceNeedSystemd:  "Get-Service : systemd (systemctl) is required.",
	MsgServiceListFail:     "Get-Service : Cannot read the service list.",
	MsgPingNotFound:        "Test-Connection : The ping command cannot be found.",
	MsgFormatHexNotFound:   "Format-Hex : Cannot find path '%s'.",
	MsgNewObjectBadType:    "New-Object : Unsupported type %s.",
	MsgHelpAliasTo:         "Alias %s → %s",
	MsgHelpNotFound:        "Cannot find help for %s.",
	MsgHelpNameHeader:      "NAME",
	MsgHelpSyntaxHeader:    "SYNTAX",
	MsgHelpAliasHeader:     "ALIASES",
	MsgProcIdNotFound:      "Cannot find a process with the process identifier %s.",
	MsgProcNameNotFound:    "Cannot find a process with the name \"%s\".",
	MsgPSVersionBad:        "Set-PSVersion : Unsupported version %q (use 5.1 or 7.x)",
	MsgPSVersionSet:        "Style switched to %s.",
	MsgHashAlgoUnsupported: "Get-FileHash : Unsupported algorithm %s",

	// ---- object 与内部提示 ----
	MsgUnsupported: "not supported",

	// ---- external 与 main ----
	MsgExternalExecFail: "Cannot execute command %s: %v",
	MsgFlagParseFail:    "Failed to parse command-line options.",
}
