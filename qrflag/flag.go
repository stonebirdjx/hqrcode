// Author stone-bird created on 2021/8/30 22:48.
// Email 1245863260@qq.com or g1245863260@gmail.com.
// Use of qr flag
package qrflag

import (
	"flag"
	"fmt"
	"os"
)

var (
	Decode   = flag.String("decode", "", "enter file name decode qr")
	Encode   = flag.Bool("encode", true, "whether encode qr")
	Content  = flag.String("content", "", "enter qr content")
	BgColor  = flag.String("bgcolor", "255,255,255,255", "qr background color")
	QrColor  = flag.String("qrcolor", "0,0,0,255", "qr background color")
	Level    = flag.String("level", "medium", "enter qr level must in (low, medium, high, highest)")
	Logo     = flag.String("logo", "", "enter qr logo")
	FileName = flag.String("o", "hqrcode.png", "enter qr output name.png")
	Size     = flag.Int("size", 256, "enter qr size")
	LogoSize = flag.Int("logosize", 64, "enter qr logo size")
	v        = flag.Bool("v", false, "show hqrcode tool version")
	version  = flag.Bool("version", false, "show hqrcode tool help text")
	h        = flag.Bool("h", false, "show hqrcode tool help text")
	help     = flag.Bool("help", false, "show hqrcodel tool help text")
)

func init() {
	flag.Parse()

	if *h || *help {
		flag.Usage()
		os.Exit(0)
	}

	if *v || *version {
		fmt.Println("hqrcode/@jx-v1.0.0")
		os.Exit(0)
	}
}
