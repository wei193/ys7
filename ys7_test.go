package ys7

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"
)

var (
	AppKey       = "you AppKey"
	Secret       = "you Secret"
	DeviceSerial = "you new DeviceSerial"
	ValidateCode = "you new Device's ValidateCode"
	ChannelNo    = 1

	ys = &Ys7{
		AppKey: AppKey,
		Secret: Secret,
	}
)

func tlog(data ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		fmt.Print(time.Now().Format("2006/01/02 15:04:05.999 "), filepath.Base(file), ":", line, " ")
	}
	fmt.Println(data...)
}
