package image

import (
	"bytes"
	"context"
	"embed"
	"errors"
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	"io"

	"github.com/aslrousta/rand"
	"github.com/minio/minio-go/v7"
	"golang.org/x/image/draw"
)

//go:embed *.png
var icons embed.FS

var (
	ErrUnsupportedFormat = errors.New("image: unsupported format")
	ErrNotFound          = errors.New("image: not found")

	mc     *minio.Client
	bucket string
	broken image.Image
)

func init() {
	brokenIcon, _ := icons.Open("broken.png")
	broken, _, _ = image.Decode(brokenIcon)
	brokenIcon.Close()
}

// Init initializes the image service.
func Init(c *minio.Client, bucketName string) {
	mc = c
	bucket = bucketName
}

// Store stores the image in the storage and returns its name.
func Store(ctx context.Context, r io.Reader) (string, error) {
	data, err := reencode(r)
	if err != nil {
		return "", ErrUnsupportedFormat
	}
	name := rand.MustString(16, rand.All)
	_, err = mc.PutObject(
		ctx, bucket, name, bytes.NewReader(data), int64(len(data)),
		minio.PutObjectOptions{ContentType: "image/png"},
	)
	if err != nil {
		return "", err
	}
	return name, nil
}

// Serve serves the image with the given name and desired width and height.
func Serve(ctx context.Context, w io.Writer, name string, width, height int) (err error) {
	var o *minio.Object
	var img image.Image
	if width < 0 || height < 0 || height > 1024 {
		// goto broken
		return ErrNotFound
	}
	o, err = mc.GetObject(ctx, bucket, name, minio.GetObjectOptions{})
	if err != nil {
		// goto broken
		return ErrNotFound
	}
	defer o.Close()
	img, _, err = image.Decode(o)
	if err != nil {
		// goto broken
		return ErrNotFound
	}

	if width == 0 && height == 0 {
		return png.Encode(w, img)
	}

	img = resize(img, width, height)
	return png.Encode(w, img)
// broken:
// 	img = genIcon(broken, width, height)
// 	if err := png.Encode(w, img); err != nil {
// 		return err
// 	}
// 	return ErrNotFound
}

func resize(src image.Image, w, h int) image.Image {
	ws, hs := src.Bounds().Max.X, src.Bounds().Max.Y
	ox, oy := 0, 0
	switch {
	case w == 0 && h == 0:
		return src
	case w == 0:
		w = h * ws / hs
	case h == 0:
		h = w * hs / ws
	case ws*h > w*hs:
		wi := w * hs / h
		ox = (ws - wi) / 2
		ws = wi
	case ws*h < w*hs:
		hi := h * ws / w
		oy = (hs - hi) / 2
		hs = hi
	}
	dst := image.NewRGBA(image.Rect(0, 0, w, h))
	crop := image.Rect(0, 0, ws, hs).Add(image.Pt(ox, oy))
	draw.CatmullRom.Scale(dst, dst.Bounds(), src, crop, draw.Src, nil)
	return dst
}

func genIcon(icon image.Image, width, height int) image.Image {
	switch {
	case width == 0 && height == 0:
		width, height = 400, 300
	case width == 0:
		width = height * 4 / 3
	case height == 0:
		height = width * 3 / 4
	}
	canvas := image.NewNRGBA(image.Rect(0, 0, width, height))
	gray := image.NewUniform(color.Gray{0xee})
	draw.Draw(canvas, canvas.Bounds(), gray, image.Point{}, draw.Src)
	offset := image.Pt(width/2-icon.Bounds().Max.X/2, height/2-icon.Bounds().Max.Y/2)
	draw.Draw(canvas, icon.Bounds().Add(offset), icon, image.Point{}, draw.Over)
	return canvas
}

func reencode(r io.Reader) ([]byte, error) {
	img, _, err := image.Decode(r)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
