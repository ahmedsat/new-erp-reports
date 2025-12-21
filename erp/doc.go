package erp

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/ahmedsat/erp-reports-cli/utils"
)

func GetDoc(path string) (result []byte, err error) {
	url, err := url.JoinPath(os.Getenv("ERP_BASE_URL"), path)
	if err != nil {
		return
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	req.Header.Add("Authorization", fmt.Sprintf("token %s", os.Getenv("ERP_AUTH_TOKEN")))

	resp, err := utils.DoRequest(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		utils.SaveHttpResponse(*resp)
		err = errors.Join(fmt.Errorf("http error: %d", resp.StatusCode), errors.New("failed to get response"))
		return
	}

	result, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}

func DeleteDoc(path string) (result []byte, err error) {
	url, err := url.JoinPath(os.Getenv("ERP_BASE_URL"), path)
	if err != nil {
		return
	}

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return
	}

	req.Header.Add("Authorization", fmt.Sprintf("token %s", os.Getenv("ERP_AUTH_TOKEN")))

	resp, err := utils.DoRequest(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != 202 {
		utils.SaveHttpResponse(*resp)
		err = errors.Join(fmt.Errorf("http error: %d", resp.StatusCode), errors.New("failed to get response"))
		return
	}

	result, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}
