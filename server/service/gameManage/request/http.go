package request

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// DataFormat 定义数据格式类型
type DataFormat string

const (
	// JSONFormat JSON格式
	JSONFormat DataFormat = "json"
	// FormFormat 表单格式
	FormFormat DataFormat = "form"
)

func HttpRequest(
	urlStr string,
	method string,
	headers map[string]string,
	params map[string]string,
	data any,
	format DataFormat) (*http.Response, error) {
	// 创建URL
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	// 添加查询参数
	// 保留原始URL中的参数
	query := u.Query()
	for k, v := range params {
		// 直接设置参数值，避免重复编码
		query.Set(k, v)
	}
	// 使用Encode()方法确保参数正确编码
	u.RawQuery = query.Encode()

	// 确保URL编码符合预期
	fmt.Printf("最终请求URL: %s\n", u.String())

	// 处理请求体数据
	var body io.Reader = nil
	if data != nil {
		buf := new(bytes.Buffer)
		switch format {
		case FormFormat:
			// 处理表单数据
			if formData, ok := data.(map[string]string); ok {
				form := url.Values{}
				for k, v := range formData {
					form.Add(k, v)
				}
				buf.WriteString(form.Encode())
			} else {
				return nil, errors.New("表单数据必须是map[string]string类型")
			}
		case JSONFormat:
			// 处理JSON数据
			b, err := json.Marshal(data)
			if err != nil {
				return nil, err
			}
			buf.Write(b)
		default:
			return nil, fmt.Errorf("不支持的数据格式: %s", format)
		}
		body = buf
	}

	// 创建请求
	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	// 设置请求头
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// 设置Content - Type
	if data != nil {
		switch format {
		case FormFormat:
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
		case JSONFormat:
			req.Header.Set("Content-Type", "application/json")
		}
	}

	// 发送请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	// 处理gzip压缩的响应
	if resp.Header.Get("Content-Encoding") == "gzip" {
		gz, err := gzip.NewReader(resp.Body)
		if err != nil {
			return resp, err // 这里返回原始响应，由调用者决定如何处理错误
		}
		// 替换响应体为解压后的Reader
		resp.Body = io.NopCloser(gz)
		// 清除Content-Encoding头，避免调用者重复处理
		resp.Header.Del("Content-Encoding")
	}

	// 返回响应，让调用者处理
	return resp, nil
}
