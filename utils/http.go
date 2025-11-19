package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
	"time"
)

var client *http.Client

const loginUrl = "/api/method/login"

type LongRes struct {
	Message  string `json:"message"`
	HomePage string `json:"home_page"`
	FullName string `json:"full_name"`
}

func Init() (err error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return
	}
	client = &http.Client{
		Jar: jar,
	}

	baseUrl := os.Getenv("ERP_BASE_URL")
	if baseUrl == "" {
		return errors.New("ERP_BASE_URL is not set")
	}
	username := os.Getenv("ERP_USERNAME")
	if username == "" {
		return errors.New("ERP_USERNAME is not set")
	}
	password := os.Getenv("ERP_PASSWORD")
	if password == "" {
		return errors.New("ERP_PASSWORD is not set")
	}

	url, err := url.JoinPath(baseUrl, loginUrl)
	if err != nil {
		return
	}
	req, err := http.NewRequest("POST", url, strings.NewReader(fmt.Sprintf("{\"usr\":\"%s\",\"pwd\":\"%s\"}", username, password)))
	if err != nil {
		return
	}
	defer req.Body.Close()

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	resp, err := DoRequest(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("login failed")
	}

	lr := LongRes{}
	err = json.NewDecoder(resp.Body).Decode(&lr)
	if err != nil {
		return
	}

	if lr.Message != "Logged In" {
		return errors.New("login failed")
	}

	fmt.Printf("logged in as %s\n", lr.FullName)

	return
}

func SaveHttpResponse(resp http.Response) {

	contentType := resp.Header.Get("Content-Type")

	extension := "txt"

	switch contentType {
	case "application/json":
		extension = "json"
	case "application/xml":
		extension = "xml"
	case "text/html":
		extension = "html"
	case "text/plain":
		extension = "txt"
	}

	os.Mkdir("logs", 0755)
	filePath := fmt.Sprintf("logs/http-%d-%s.%s", resp.StatusCode, time.Now().Format("2006-01-02"), extension)
	f, err := os.Create(filePath)
	HandelErr(err)
	_, err = io.Copy(f, resp.Body)
	HandelErr(err)

	fmt.Printf("response saved to %s\n", filePath)
	if resp.StatusCode >= 300 && resp.StatusCode < 200 {
		HandelErr(fmt.Errorf("http error: %d", resp.StatusCode))
	}

	fmt.Println(resp.Status)
}

func DoRequest(req *http.Request) (resp *http.Response, err error) {
	resp, err = client.Do(req)
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		SaveHttpResponse(*resp)
		err = errors.Join(fmt.Errorf("http error: %d", resp.StatusCode), errors.New("failed to get response"))
		return
	}

	return
}
