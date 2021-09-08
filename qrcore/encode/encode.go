// Author stone-bird created on 2021/9/8 7:46.
// Email 1245863260@qq.com or g1245863260@gmail.com.
// Use of encode qr
package encode

import (
	"bytes"
	"github.com/nfnt/resize"
	"github.com/skip2/go-qrcode"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"
)

type QrInfo struct {
	FileName string
	Content  string
	Size     int
	Logo     string
	LogoSize int
	BgColor  color.Color
	QrColor  color.Color
	Level    string //low Level L: 7% error recovery.
	//medium Level M: 15% error recovery. Good default choice.
	//high Level Q: 25% error recovery.
	//highest Level H: 30% error recovery.
}

func (qr *QrInfo) Encode() {
	level := qrcode.Medium
	switch qr.Level {
	case "low":
		level = qrcode.Low
	case "high":
		level = qrcode.High
	case "highest":
		level = qrcode.Highest
	}
	if qr.Logo != "" {
		qr.encodeLogo(level)
	} else {
		qr.encodeBase(level)
	}

}

func (qr *QrInfo) encodeLogo(level qrcode.RecoveryLevel) {
	var err error
	lf, err := os.Open(qr.Logo)
	if err != nil {
		log.Fatal(err)
	}
	defer lf.Close()

	var logoImg image.Image
	switch {
	case strings.HasSuffix(qr.Logo, ".png"):
		logoImg, err = png.Decode(lf)
	case strings.HasSuffix(qr.Logo, ".jpg"), strings.HasSuffix(qr.Logo, ".jpeg"):
		logoImg, err = jpeg.Decode(lf)
	default:
		log.Fatal("loge must jpeg, jpg, png picture")
	}
	if err != nil {
		log.Fatal(err)
	}

	// logoImg resize
	logoImg = resize.Resize(uint(qr.LogoSize), uint(qr.LogoSize), logoImg, resize.Lanczos3)

	var buf bytes.Buffer

	qc, err := qrcode.New(qr.Content, level)
	if err != nil {
		log.Fatal(err)
	}
	qc.BackgroundColor = qr.BgColor
	qc.ForegroundColor = qr.QrColor
	err = qc.Write(qr.Size, &buf)
	if err != nil {
		log.Fatal(err)
	}

	//png
	fp, err := png.Decode(&buf)
	if err != nil {
		log.Fatal(err)
	}
	// fp bounds
	b := fp.Bounds()

	//dst iamge
	rgba := image.NewRGBA(b)

	// draw qrcode
	draw.Draw(rgba, b, fp, image.Point{X: 0, Y: 0}, draw.Src)

	//draw logo
	// position center
	offset := image.Pt((b.Max.X-logoImg.Bounds().Max.X)/2, (b.Max.Y-logoImg.Bounds().Max.Y)/2)
	draw.Draw(rgba, logoImg.Bounds().Add(offset), logoImg, image.Point{X: 0, Y: 0}, draw.Over)

	//save png
	f, err := os.OpenFile(qr.FileName, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = png.Encode(f, rgba)
	if err != nil {
		log.Fatal(err)
	}
}

func (qr *QrInfo) encodeBase(level qrcode.RecoveryLevel) {
	qc, err := qrcode.New(qr.Content, level)
	if err != nil {
		log.Fatal(err)
	}
	qc.BackgroundColor = qr.BgColor
	qc.ForegroundColor = qr.QrColor
	f, err := os.OpenFile(qr.FileName, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	err = qc.Write(qr.Size, f)
	if err != nil {
		log.Fatal(err)
	}
}
