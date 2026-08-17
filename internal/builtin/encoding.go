package builtin

import (
	"strings"
	"unicode/utf16"
)

// encodeText 按 PowerShell 编码名把文本编码为字节；不认识的编码名按 UTF-8 处理。
// bom 为 true 时写 BOM（追加模式传 false 避免重复写 BOM）。
func encodeText(enc, text string, bom bool) []byte {
	switch strings.ToLower(strings.ReplaceAll(enc, "-", "")) {
	case "ascii":
		// 非 ASCII 字符替换为 '?'（对齐 .NET ASCII 编码）
		var out []byte
		for _, r := range text {
			if r < 128 {
				out = append(out, byte(r))
			} else {
				out = append(out, '?')
			}
		}
		return out
	case "utf8bom":
		if bom {
			return append([]byte{0xEF, 0xBB, 0xBF}, []byte(text)...)
		}
		return []byte(text)
	case "utf8nobom":
		return []byte(text)
	case "unicode":
		return utf16Bytes(text, false, bom)
	case "bigendianunicode":
		return utf16Bytes(text, true, bom)
	case "utf32":
		return utf32Bytes(text, false, bom)
	case "bigendianutf32":
		return utf32Bytes(text, true, bom)
	default:
		return []byte(text)
	}
}

// utf16Bytes 编码为 UTF-16（LE/BE），可选 BOM（utf16.Encode 负责代理对）。
func utf16Bytes(text string, be, bom bool) []byte {
	var out []byte
	if bom {
		if be {
			out = append(out, 0xFE, 0xFF)
		} else {
			out = append(out, 0xFF, 0xFE)
		}
	}
	for _, u := range utf16.Encode([]rune(text)) {
		if be {
			out = append(out, byte(u>>8), byte(u))
		} else {
			out = append(out, byte(u), byte(u>>8))
		}
	}
	return out
}

// utf32Bytes 编码为 UTF-32（LE/BE），可选 BOM。
func utf32Bytes(text string, be, bom bool) []byte {
	var out []byte
	if bom {
		if be {
			out = append(out, 0x00, 0x00, 0xFE, 0xFF)
		} else {
			out = append(out, 0xFF, 0xFE, 0x00, 0x00)
		}
	}
	for _, r := range text {
		if be {
			out = append(out, byte(r>>24), byte(r>>16), byte(r>>8), byte(r))
		} else {
			out = append(out, byte(r), byte(r>>8), byte(r>>16), byte(r>>24))
		}
	}
	return out
}
