package gobui

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"io"
	"net/http"
	"time"
)

type Display struct {
	enc io.ReadSeeker
	mod time.Time
}

func NewDisplay(port int, path string) *Display {
	disp := &Display{}

	// index.html must reside at path
	if path == "" {
		path = "."
	}
	http.Handle("/", http.FileServer(http.Dir(path)))
	http.HandleFunc("/image.png", disp.servImg)

	go http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	return disp
}

func (d *Display) Load(img image.Image) {
	var buf bytes.Buffer
	png.Encode(&buf, img)
	d.enc = bytes.NewReader(buf.Bytes())
	d.mod = time.Now()
}

func (d *Display) servImg(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Cache-Control", "no-store")
	http.ServeContent(w, r, "", d.mod, d.enc)
}
