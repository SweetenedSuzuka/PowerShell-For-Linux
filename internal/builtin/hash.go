package builtin

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
	"strings"

	"powershell/internal/lang"
	"powershell/internal/object"
)

// hash.go 实现文件哈希类 cmdlet。

func cmdGetFileHash(c *Context) ([]*object.PSObject, error) {
	path := firstPathArg(c)
	if path == "" {
		return nil, nil
	}
	algorithm := "SHA256"
	if a, ok := c.Args.Str("Algorithm"); ok {
		algorithm = a
	}
	var out []*object.PSObject
	paths, derr := expandWildcard(c, path)
	if derr != nil {
		return errf(c, "%v", derr)
	}
	for _, p := range paths {
		data, err := os.ReadFile(p)
		if err != nil {
			continue
		}
		hash, err := computeHash(algorithm, data)
		if err != nil {
			return errf(c, "%s", lang.T(lang.MsgHashAlgoUnsupported, algorithm))
		}
		o := object.Object("Microsoft.PowerShell.Utility.FileHash", nil)
		o.AddProp("Algorithm", strings.ToUpper(algorithm))
		o.AddProp("Hash", strings.ToLower(hash))
		o.AddProp("Path", p)
		out = append(out, o)
	}
	return out, nil
}

func computeHash(algorithm string, data []byte) (string, error) {
	switch strings.ToUpper(algorithm) {
	case "SHA256", "SHA2_256":
		sum := sha256.Sum256(data)
		return fmt.Sprintf("%x", sum), nil
	case "SHA1":
		sum := sha1.Sum(data)
		return fmt.Sprintf("%x", sum), nil
	case "MD5":
		sum := md5.Sum(data)
		return fmt.Sprintf("%x", sum), nil
	case "SHA512", "SHA2_512":
		sum := sha512.Sum512(data)
		return fmt.Sprintf("%x", sum), nil
	}
	return "", fmt.Errorf("%s", lang.T(lang.MsgUnsupported))
}

// ---- 注册 ----

func init() {
	Register("Get-FileHash", []ParamSpec{
		{Name: "Path", Position: 0, PositionSet: true, Type: "path"},
		{Name: "Algorithm", Type: "string"},
	}, cmdGetFileHash)
}
