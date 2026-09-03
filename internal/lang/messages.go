package lang

// Msg 是提示文本的编号。文本本体按语言放在各语言表文件里（zh.go、en.go）。
type Msg int

const (
	// ---- shell：横幅与帮助 ----
	MsgBannerDesktop Msg = iota
	MsgBannerCore
	MsgUsage
	MsgReadonlyVar
	MsgDriveUnsupported

	// ---- parser：解析错误与 token 描述 ----
	MsgParseExpectText
	MsgParseExpectWhat
	MsgParseUnexpectedAfter
	MsgParseCatchAfterTry
	MsgParseTryMissingHandler
	MsgParseFinallyAfterTry
	MsgParseAssignOp
	MsgParseExpectBrace
	MsgParseForeachVar
	MsgParseForeachIn
	MsgParseDoWhile
	MsgParseFuncParam
	MsgParseFuncName
	MsgParseParamList
	MsgParsePipeRightCmd
	MsgParseCmpOp
	MsgParseTernaryColon
	MsgParseIncDecVar
	MsgParseOpMissingLeft
	MsgParseUnexpectedArg
	MsgParseInputRedirect
	MsgParseRedirectAfterCmd
	MsgParseUnexpectedOp
	MsgParseTypeLiteralName
	MsgParseTypeLiteralRbracket
	MsgParseParamTypeName
	MsgParseParamTypeRbracket
	MsgParseNamedBlockPosition
	MsgParseNamedBlockDuplicate
	MsgConvertFail
	MsgTypeUnknown
	MsgStaticMemberNotFound
	MsgBindConvertFail
	MsgInvokeTargetMissing
	MsgParseUnexpectedToken
	MsgParseHashEntry
	MsgTokEOF
	MsgTokNewline
	MsgTokWord
	MsgTokNumber
	MsgTokString
	MsgTokOp

	// ---- eval：求值与脚本 ----
	MsgDivideByZero
	MsgFormatIndexOut
	MsgIndexOutOfRange
	MsgParamOnlyInFunc
	MsgRedirectWrite
	MsgCommandNotFound
	MsgScriptReadFail
	MsgScriptParseFail
	MsgIncompleteInput

	// ---- builtin：cmdlet 报错与提示 ----
	MsgBindNoParam
	MsgBindSwitchNoValue
	MsgBindErrorActionInvalid
	MsgErrorActionPreferenceInvalid
	MsgWriteErrorPrefix
	MsgWriteVerbosePrefix
	MsgWriteWarningPrefix
	MsgWriteDebugPrefix
	MsgDateParseFail
	MsgPathNotFoundFmt
	MsgPathNotFoundForSet
	MsgCannotWrite
	MsgCannotOpen
	MsgCannotCreate
	MsgCannotDelete
	MsgCopyDestNotDir
	MsgCopyNeedsRecurse
	MsgSkipNegative
	MsgCannotRename
	MsgRenameDestExists
	MsgWhatIfPerform
	MsgConfirmPrompt
	MsgCannotClear
	MsgPropNotFound
	MsgVarExists
	MsgAliasExists
	MsgServiceNeedSystemd
	MsgServiceListFail
	MsgPingNotFound
	MsgFormatHexNotFound
	MsgNewObjectBadType
	MsgHelpAliasTo
	MsgHelpNotFound
	MsgHelpNameHeader
	MsgHelpSyntaxHeader
	MsgHelpAliasHeader
	MsgProcIdNotFound
	MsgProcNameNotFound
	MsgPSVersionBad
	MsgPSVersionSet
	MsgHashAlgoUnsupported

	// ---- object 与内部提示 ----
	MsgUnsupported

	// ---- external 与 main ----
	MsgExternalExecFail
	MsgFlagParseFail
)
