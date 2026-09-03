package builtin

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"powershell/internal/object"
)

// web.go 实现网络请求类 cmdlet。

func cmdInvokeWebRequest(c *Context) ([]*object.PSObject, error) {
	uri, _ := c.Args.Str("Uri")
	if uri == "" {
		return nil, nil
	}
	body, method := httpRequestArgs(c)
	resp, err := doHTTPRequest(method, uri, body)
	if err != nil {
		return errf(c, "Invoke-WebRequest : %v", err)
	}
	text := strings.TrimSuffix(resp, "\n")
	var lines []*object.PSObject
	if text != "" {
		for _, ln := range strings.Split(text, "\n") {
			lines = append(lines, object.Str(ln))
		}
	}
	return lines, nil
}

func cmdInvokeRestMethod(c *Context) ([]*object.PSObject, error) {
	uri, _ := c.Args.Str("Uri")
	if uri == "" {
		return nil, nil
	}
	body, method := httpRequestArgs(c)
	resp, err := doHTTPRequest(method, uri, body)
	if err != nil {
		return errf(c, "Invoke-RestMethod : %v", err)
	}
	// 尝试解析 JSON
	var v any
	if err := json.Unmarshal([]byte(resp), &v); err == nil {
		return []*object.PSObject{jsonToObject(v)}, nil
	}
	return []*object.PSObject{object.Str(strings.TrimSuffix(resp, "\n"))}, nil
}

// httpRequestArgs 取请求的 HTTP 方法与请求体（默认 GET）。
func httpRequestArgs(c *Context) (string, string) {
	method := "GET"
	if m, ok := c.Args.Str("Method"); ok && m != "" {
		method = m
	}
	body := ""
	if b := c.Args.Get("Body"); b != nil {
		body = b.String()
	}
	return body, method
}

// doHTTPRequest 发一次 HTTP 请求，返回响应文本；30 秒超时。
func doHTTPRequest(method, uri, body string) (string, error) {
	client := &http.Client{Timeout: 30 * time.Second}
	req, err := http.NewRequest(method, uri, nil)
	if err != nil {
		return "", err
	}
	if body != "" && (method == "POST" || method == "PUT" || method == "PATCH") {
		req.Body = io.NopCloser(strings.NewReader(body))
		req.ContentLength = int64(len(body))
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	return string(data), err
}

// ---- 注册 ----

func init() {
	Register("Invoke-WebRequest", []ParamSpec{
		{Name: "Uri", Position: 0, PositionSet: true, Type: "string"},
		{Name: "Method", Type: "string"},
		{Name: "Body", Type: "object"},
	}, cmdInvokeWebRequest)
	Register("Invoke-RestMethod", []ParamSpec{
		{Name: "Uri", Position: 0, PositionSet: true, Type: "string"},
		{Name: "Method", Type: "string"},
		{Name: "Body", Type: "object"},
	}, cmdInvokeRestMethod)
}
