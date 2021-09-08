// Author stone-bird created on 2021/8/30 22:52.
// Email 1245863260@qq.com or g1245863260@gmail.com.
// Use of encode or decode qr
package main

import (
	"hqrcode/qrcore"
	"hqrcode/qrflag"
	"strings"
)

func main() {
	decode := strings.TrimSpace(*qrflag.Decode)
	switch {
	case decode != "":
		qrcore.DecodeMain(decode)
	case *qrflag.Encode:
		qrcore.EncodeMain()
	default:
		return
	}
}
