package utils

import (
	"strings"
	"path/filepath"
)

// CheckFileExt 检查文件扩展名是否在允许的列表中
// @param filename 文件名
// @param allowedExts 允许的扩展名列表(如 [".xlsx", ".xls"])
// @return bool 是否允许
func CheckFileExt(filename string, allowedExts []string) bool {
	if filename == "" {
		return false
	}
	
	// 获取文件扩展名
	ext := strings.ToLower(filepath.Ext(filename))
	if ext == "" {
		return false
	}
	
	// 检查扩展名是否在允许列表中
	for _, allowed := range allowedExts {
		if strings.ToLower(allowed) == ext {
			return true
		}
	}
	
	return false
}