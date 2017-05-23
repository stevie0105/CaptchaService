package main

import (
	"fmt"
	"github.com/dchest/captcha"
	"log"
	"net/http"
	"encoding/json"
	"encoding/base64"
	"bytes"
	"image/jpeg"
	"strconv"
)

type CaptchaResponse struct {
	Image string `json:"image"`
	Solution string `json:"solution"`
}

func getCaptcha() ([]byte, error) {
	myCaptcha := captcha.New()

	digits := captcha.RandomDigits(6)
	image := captcha.NewImage(myCaptcha, digits, captcha.StdWidth, captcha.StdHeight)

	solution := convert(digits)

	buffer := new(bytes.Buffer)
	jpeg.Encode(buffer, image, nil)
	img := base64.StdEncoding.EncodeToString(buffer.Bytes())

	captchaResponse := CaptchaResponse{img, solution}

	b, err := json.Marshal(captchaResponse)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func handleCaptcha(w http.ResponseWriter, r *http.Request) {
	b, err := getCaptcha()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func convert(b []byte) string {
	var buffer bytes.Buffer
	for i := range b {
		buffer.WriteString(strconv.Itoa(int(b[i])))
	}
	return buffer.String()
}

func init() {
	http.HandleFunc("/", handleCaptcha)
}

func main() {
	fmt.Println("Server is at localhost:8666")
	if err := http.ListenAndServe(":8666", nil); err != nil {
		log.Fatal(err)
	}
}
