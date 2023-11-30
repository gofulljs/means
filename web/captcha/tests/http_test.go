package tests

import (
	"fmt"
	"github.com/gofulljs/means/web/captcha"
	"net/http"
	"testing"
)

func TestCaptchaHttp(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello world!")
	})

	mux.Handle("/captcha/", http.StripPrefix("/captcha", captcha.NewMgr(captcha.WithPrefix("/"))))
	fmt.Println("handle at :8088")
	http.ListenAndServe(":8088", mux)
}
