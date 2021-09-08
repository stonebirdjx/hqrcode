// Author stone-bird created on 2021/8/30 23:03.
// Email 1245863260@qq.com or g1245863260@gmail.com.
// Use of fetch func
package qrcore

import (
	"fmt"
	dq "github.com/tuotoo/qrcode"
	"hqrcode/qrcore/encode"
	"hqrcode/qrflag"
	"image/color"
	"log"
	"os"
	"strconv"
	"strings"
)

func DecodeMain(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	mx ,err := dq.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mx.Content)
}

// encode qr
func EncodeMain() {
	qr := new(encode.QrInfo)

	// background color deal
	bgcolor := strings.TrimSpace(*qrflag.BgColor)
	bgs := strings.Split(bgcolor, ",")
	if len(bgs) != 4 {
		log.Fatal("background type err")
	}

	bgr, err := strconv.Atoi(strings.TrimSpace(bgs[0]))
	if err != nil {
		log.Fatal("background rgba r is not int type")
	}

	bgg, err := strconv.Atoi(strings.TrimSpace(bgs[1]))
	if err != nil {
		log.Fatal("background rgba g is not int type")
	}

	bgb, err := strconv.Atoi(strings.TrimSpace(bgs[2]))
	if err != nil {
		log.Fatal("background rgba b is not int type")
	}

	bga, err := strconv.Atoi(strings.TrimSpace(bgs[3]))
	if err != nil {
		log.Fatal("background rgba a is not int type")
	}

	qr.BgColor = color.RGBA{R: uint8(bgr), G: uint8(bgg), B: uint8(bgb), A: uint8(bga)}

	// qr color deal
	qrcolor := strings.TrimSpace(*qrflag.QrColor)
	qcs := strings.Split(qrcolor, ",")
	if len(qcs) != 4 {
		log.Fatal("qrcolor type err")
	}

	qcr, err := strconv.Atoi(strings.TrimSpace(qcs[0]))
	if err != nil {
		log.Fatal("qrcolor rgba r is not int type")
	}

	qcg, err := strconv.Atoi(strings.TrimSpace(qcs[1]))
	if err != nil {
		log.Fatal("qrcolor rgba g is not int type")
	}

	qcb, err := strconv.Atoi(strings.TrimSpace(qcs[2]))
	if err != nil {
		log.Fatal("qrcolor rgba b is not int type")
	}

	qca, err := strconv.Atoi(strings.TrimSpace(qcs[3]))
	if err != nil {
		log.Fatal("qrcolor rgba a is not int type")
	}

	qr.QrColor = color.RGBA{R: uint8(qcr), G: uint8(qcg), B: uint8(qcb), A: uint8(qca)}
	qr.Content = strings.TrimSpace(*qrflag.Content)
	qr.Size = *qrflag.Size
	qr.Level = strings.TrimSpace(*qrflag.Level)
	qr.Logo = strings.TrimSpace(*qrflag.Logo)
	qr.LogoSize = *qrflag.LogoSize
	qr.FileName = strings.TrimSpace(*qrflag.FileName)

	if qr.LogoSize > qr.Size {
		log.Fatal("logo size must le qr size")
	}
	qr.Encode()
}
