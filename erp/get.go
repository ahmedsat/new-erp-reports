package erp

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/ahmedsat/erp-reports-cli/utils"
)

func Get[T any](doctype string, filters utils.Filters, fields utils.List) (result []T, err error) {
	url, err := url.JoinPath(os.Getenv("ERP_BASE_URL"), "/api/resource", doctype)
	if err != nil {
		return
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	q := req.URL.Query()
	q.Add("limit", "0")

	if len(filters) != 0 {
		q.Add("filters", filters.String())
	}

	if len(fields) == 0 {
		fields = utils.List{"*"}
	}
	q.Add("fields", fields.String())

	req.URL.RawQuery = q.Encode()

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

	decoder := json.NewDecoder(resp.Body)

	var response struct {
		Data []T `json:"data"`
	}

	err = decoder.Decode(&response)
	if err != nil {
		return
	}

	result = response.Data

	return
}

func Get1[T any](doctype, id string) (result T, err error) {
	url, err := url.JoinPath(os.Getenv("ERP_BASE_URL"), "/api/resource", doctype, id)
	if err != nil {
		return
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

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

	decoder := json.NewDecoder(resp.Body)

	var response = struct {
		Data T `json:"data"`
	}{
		Data: result,
	}

	err = decoder.Decode(&response)
	if err != nil {
		return
	}

	return

}
