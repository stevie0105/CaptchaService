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
	"strings"
)

type CaptchaResponse struct {
	Image string `json:"image"`
	Solution string `json:"solution"`
}

func getCaptcha(w http.ResponseWriter, r *http.Request) {
	myCaptcha := captcha.New()

	digits := captcha.RandomDigits(6)
	image := captcha.NewImage(myCaptcha, digits, captcha.StdWidth, captcha.StdHeight)

	solution := convert(digits)

	buffer := new(bytes.Buffer)
	jpeg.Encode(buffer, image, nil)
	img := base64.StdEncoding.EncodeToString(buffer.Bytes())

	captchaResponse := CaptchaResponse{img, solution}

	b, _ := json.Marshal(captchaResponse)

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func convert( b []byte ) string {
	s := make([]string,len(b))
	for i := range b {
		s[i] = strconv.Itoa(int(b[i]))
	}
	return strings.Join(s,"")
}

func main() {
	http.HandleFunc("/", getCaptcha)
	fmt.Println("Server is at localhost:8666")
	if err := http.ListenAndServe("localhost:8666", nil); err != nil {
		log.Fatal(err)
	}
}