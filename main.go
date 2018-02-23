package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

const (
	TargetWinbat = "winbat"
	TargetBash   = "bash"

	compileURL = "http://batsh.org/compile"
)

type CodeResponse struct {
	Code string `json:"code"`
}
type ErrResponse struct {
	Err string `json:"err"`
}

func compile(target string, code string) (string, error) {
	form := url.Values{
		"target": {target},
		"code":   {code},
	}
	body := bytes.NewBufferString(form.Encode())
	resp, err := http.Post(compileURL, "application/x-www-form-urlencoded", body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var respSucc CodeResponse
	err = json.Unmarshal(data, &respSucc)
	var respErr ErrResponse
	err = json.Unmarshal(data, &respErr)

	if respErr.Err != "" {
		return "", errors.New("[compile error] " + respErr.Err)
	}

	return respSucc.Code, nil
}

var targetLang string
var sourceFilePath string

func init() {
	flag.StringVar(&targetLang, "target", "", "target langage: bash/winbat")
	flag.StringVar(&sourceFilePath, "src", "", "source file path")
}

func main() {
	flag.Parse()

	if targetLang == "" {
		fmt.Fprintln(os.Stderr, "target not defined")
		os.Exit(-1)
	}
	if sourceFilePath == "" {
		fmt.Fprintln(os.Stderr, "source file path not defined")
		os.Exit(-1)
	}

	data, err := ioutil.ReadFile(sourceFilePath)
	if err != nil {
		panic(err)
	}

	source := string(data)
	code, err := compile(targetLang, source)
	if err != nil {
		panic(err)
	}
	fmt.Println(code)
}
