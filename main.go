package main

import (
	"fmt"
	"github.com/dchest/captcha"
	"log"
	"net/http"
	"encoding/json"
)

type CaptchaResponse struct {
	ImageBase64 captcha.Image `json:"image"`
	Solution []byte `json:"solution"`
}

func getCaptcha(w http.ResponseWriter, r *http.Request) {
	myCaptcha := captcha.New()

	foo := captcha.RandomDigits(6)
	image := captcha.NewImage(myCaptcha, foo, captcha.StdWidth, captcha.StdHeight)
	captchaResponse := CaptchaResponse{*image, foo}

	b, _ := json.Marshal(captchaResponse)

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func main() {
	http.HandleFunc("/", getCaptcha)
	fmt.Println("Server is at localhost:8666")
	if err := http.ListenAndServe("localhost:8666", nil); err != nil {
		log.Fatal(err)
	}
}