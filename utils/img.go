package utils

import (
	"bytes"
	"encoding/base64"
	"errors"
	"image"
	"image/draw"
	_"image/gif"
	"image/jpeg"
	_"image/jpeg"
	"image/png"
	_"image/png"
	"io"
	"os"
)

//加载图片
//返回 图片、图片类型、错误
func LoadImg(r io.Reader) (image.Image,string, error) {
	src, aa, err := image.Decode(r)
	if err != nil {
		return nil,aa, errors.New("image decode err :" + err.Error())
	}
	return src,aa, nil
}
func ImgToReader(img image.Image) (*bytes.Reader,int64) {
	b := EncodeImg(img)
	return bytes.NewReader(b), int64(len(b))
}
func EncodeImg(img image.Image) []byte {
	i := bytes.NewBuffer(nil)
	jpeg.Encode(i,img,nil)
	return i.Bytes()
}
func ImgToBase64(img image.Image) string {
	emptyBuff := bytes.NewBuffer(nil)
	jpeg.Encode(emptyBuff, img, nil)
	return base64.StdEncoding.EncodeToString(emptyBuff.Bytes())
}

func TrimmingImage(src image.Image, x, y, w, h int) (image.Image, error) {
	var subImg image.Image
	if rgbImg, ok := src.(*image.YCbCr); ok {
		subImg = rgbImg.SubImage(image.Rect(x, y, x+w, y+h)).(*image.YCbCr) //图片裁剪x0 y0 x1 y1
	} else if rgbImg, ok := src.(*image.RGBA); ok {
		subImg = rgbImg.SubImage(image.Rect(x, y, x+w, y+h)).(*image.RGBA) //图片裁剪x0 y0 x1 y1
	} else if rgbImg, ok := src.(*image.NRGBA); ok {
		subImg = rgbImg.SubImage(image.Rect(x, y, x+w, y+h)).(*image.NRGBA) //图片裁剪x0 y0 x1 y1
	} else {
		return subImg, errors.New("图片解码失败")
	}
	return subImg, nil
}

func Draw(img image.Image,r image.Image,x,y,w,h int) image.Image {
	r1 := image.Rectangle{
		image.Point{x,y},
		image.Point{x+w,y+10},
	}
	r2 := image.Rectangle{
		image.Point{x,y},
		image.Point{x+10,y+h},
	}
	r3 := image.Rectangle{
		image.Point{x+w,y},
		image.Point{x+w+10,y+h},
	}
	r4 := image.Rectangle{
		image.Point{x,y+h},
		image.Point{x+w+10,y+h+10},
	}
	if _,ok := img.(*image.YCbCr);ok {
		b := img.Bounds()
		m := image.NewRGBA(image.Rect(0,0,b.Dx(),b.Dy()))
		draw.Draw(m,img.Bounds(),img,b.Min ,draw.Src)
		img = m
	}
	draw.Draw(img.(draw.Image),r1,r,r1.Min,draw.Src)
	draw.Draw(img.(draw.Image),r2,r,r2.Min,draw.Src)
	draw.Draw(img.(draw.Image),r3,r,r3.Min,draw.Src)
	draw.Draw(img.(draw.Image),r4,r,r4.Min,draw.Src)
	return img
}

func CopyImg(img image.Image) image.Image {
	b := img.Bounds()
	c_img := image.NewRGBA(image.Rect(0, 0, b.Max.X, b.Max.Y))
	draw.Draw(c_img,b,img,b.Min,draw.Src)
	return c_img
}
func SaveImg(img image.Image,path string,mod string) error {
	//解码
	i := bytes.NewBuffer(nil)
	switch mod {
	case "jpeg":
		jpeg.Encode(i,img,nil)
	case "png":
		png.Encode(i,img)
	default:
		jpeg.Encode(i,img,nil)
	}
	//创建文件
	f,err := os.Create(path)
	if err != nil {
		return err
	}
	f.Write(i.Bytes())
	f.Close()
	return nil
}