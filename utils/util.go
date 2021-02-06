package utils

import (
	"path/filepath"
)

// 获取文件名称
func GetFileName(path string) string {
	return filepath.Base(path)
}
