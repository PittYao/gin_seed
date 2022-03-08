// Package statics embed 静态文件
package statics

import "embed"

// Files 静态文件资源
//go:embed  html/*
var Files embed.FS
