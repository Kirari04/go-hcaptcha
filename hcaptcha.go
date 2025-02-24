// Package hcaptcha handles hCaptcha (http://hcaptcha.com) form submissions
//
// This package is designed to be called from within an HTTP server or web framework
// which offers hCaptcha form inputs and requires them to be evaluated for correctness
//
// Edit the hcaptchaPrivateKey constant before building and using
package hcaptcha

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type RecaptchaResponse struct {
	Success     bool      `json:"success"`
	Score       float64   `json:"score"`
	Action      string    `json:"action"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}

const hcaptchaServerName = "https://api.hcaptcha.com/siteverify"

var hcaptchaPrivateKey string

// check uses the client ip address, the challenge code from the hCaptcha form,
// and the client's response input to that challenge to determine whether or not
// the client answered the hCaptcha input question correctly.
// It returns a boolean value indicating whether or not the client answered correctly.
func check(remoteip, response string) (r RecaptchaResponse, err error) {
	resp, err := http.PostForm(hcaptchaServerName,
		url.Values{"secret": {hcaptchaPrivateKey}, "remoteip": {remoteip}, "response": {response}})
	if err != nil {
		log.Printf("Post error: %s\n", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Read error: could not read body: %s", err)
		return
	}
	err = json.Unmarshal(body, &r)
	if err != nil {
		log.Println("Read error: got invalid JSON: %s", err)
		return
	}
	return
}

// Confirm is the public interface function.
// It calls check, which the client ip address, the challenge code from the hCaptcha form,
// and the client's response input to that challenge to determine whether or not
// the client answered the hCaptcha input question correctly.
// It returns a boolean value indicating whether or not the client answered correctly.
func Confirm(remoteip, response string) (result bool, err error) {
	resp, err := check(remoteip, response)
	result = resp.Success
	return
}

// Init allows the webserver or code evaluating the hCaptcha form input to set the
// hCaptcha private key (string) value, which will be different for every domain.
func Init(key string) {
	hcaptchaPrivateKey = key
}
