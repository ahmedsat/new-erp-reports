package utils

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

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
	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		err = errors.Join(err, fmt.Errorf("%s : failed to send request", WhereAmI()))
		return
	}

	if resp.StatusCode != 200 {
		SaveHttpResponse(*resp)
		err = errors.Join(fmt.Errorf("http error: %d", resp.StatusCode), fmt.Errorf("%s : failed to get response", WhereAmI()))
		return
	}

	return
}
