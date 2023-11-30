package captcha

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dchest/captcha"
	"net/http"
	"path"
	"strings"
	"time"
)

type Mgr struct {
	*http.ServeMux
	length int
	prefix string
	width  int
	height int
}

type Options func(mgr *Mgr)

func WithLength(length int) Options {
	return func(mgr *Mgr) {
		mgr.length = length
	}
}

func WithPrefix(prefix string) Options {
	return func(mgr *Mgr) {
		prefix = strings.Trim(prefix, "/")
		if prefix != "" {
			mgr.prefix = "/" + prefix
		} else {
			mgr.prefix = ""
		}
	}
}

func WithWidth(width int) Options {
	return func(mgr *Mgr) {
		mgr.width = width
	}
}

func WitHeight(height int) Options {
	return func(mgr *Mgr) {
		mgr.height = height
	}
}

func NewMgr(opts ...Options) *Mgr {
	res := &Mgr{
		ServeMux: http.NewServeMux(),
		prefix:   "/captcha",
		length:   4,
		width:    captcha.StdWidth,
		height:   captcha.StdHeight,
	}

	for _, opt := range opts {
		opt(res)
	}

	res.handler()

	return res
}

func (mgr *Mgr) handler() {
	mgr.HandleFunc(mgr.prefix+"/getinfo", func(writer http.ResponseWriter, request *http.Request) {
		captchaId := captcha.NewLen(mgr.length)
		prefixURL := strings.TrimSuffix(request.RequestURI, "/getinfo")
		res, err := json.Marshal(map[string]any{
			"id":  captchaId,
			"url": fmt.Sprintf("%s/image?file=%s.png", prefixURL, captchaId),
		})
		if err != nil {
			fmt.Println("err:", err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(http.StatusOK)
		writer.Write(res)
	})
	mgr.HandleFunc(mgr.prefix+"/image", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		file := r.Form.Get("file")
		ext := path.Ext(file)
		id := file[:len(file)-len(ext)]
		reload := r.Form.Get("reload")
		download := r.Form.Get("download")
		lang := r.Form.Get("lang")

		if ext == "" || id == "" {
			http.NotFound(w, r)
			return
		}
		if reload != "" {
			captcha.Reload(id)
		}
		if Serve(w, r, id, ext, lang, download == "1", captcha.StdWidth, captcha.StdHeight) == captcha.ErrNotFound {
			http.NotFound(w, r)
		}
	})
	mgr.HandleFunc(mgr.prefix+"/verify", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		id := r.Form.Get("id")
		value := r.Form.Get("value")
		if id == "" || value == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "param error")
		}
		if captcha.VerifyString(id, value) {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "ok")
		} else {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "failed")
		}
	})
}

func Serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	var content bytes.Buffer
	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		err := captcha.WriteImage(&content, id, width, height)
		if err != nil {
			fmt.Println("err:", err)
		}
	case ".wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		err := captcha.WriteAudio(&content, id, lang)
		if err != nil {
			fmt.Println("err:", err)
		}
	default:
		return captcha.ErrNotFound
	}

	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}
