package webdav

import (
	"github.com/studio-b12/gowebdav"
)

// NewClient 创建一个webdav的客户端
func NewClient(url, user, password string) *gowebdav.Client {
	return gowebdav.NewClient(url, user, password)
}
