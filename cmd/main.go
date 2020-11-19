package main

import (
	"github.com/dezhishen/webdav-sync/internal/configutil"
	"github.com/dezhishen/webdav-sync/internal/webdav"
)

func main() {
	err := configutil.LoadConfigs("../configs", []string{"webdav", "sync"})
	if err != nil {
		panic(err)
	}
	fileList := configutil.GetConfigStringMapString("sync", "fileList")
	if fileList != nil {
		url := configutil.GetConfigString("webdav", "url")
		user := configutil.GetConfigString("webdav", "user")
		password := configutil.GetConfigString("webdav", "password")
		//load webdavclient
		client := webdav.NewClient(url, user, password)
		//开始同步
		for k, v := range fileList {
			client.Read()
		}
	}
}
