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
	Solution string `json:"solution"`
}

func getCaptcha(w http.ResponseWriter, r *http.Request) {
	myCaptcha := captcha.New()

	image := captcha.NewImage(myCaptcha, []byte("what the fuck"), captcha.StdWidth, captcha.StdHeight)
	captchaResponse := CaptchaResponse{*image.Paletted, "qwertz"}

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