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

func Get[T any](path string, filters utils.Filters, fields utils.List) (result []T, err error) {
	url, err := url.JoinPath(os.Getenv("BASE_URL"), path)
	if err != nil {
		err = errors.Join(err, fmt.Errorf("%s : failed to join url", utils.WhereAmI()))
		return
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		err = errors.Join(err, fmt.Errorf("%s : failed to create request", utils.WhereAmI()))
		return
	}

	req.Header.Add("Authorization", fmt.Sprintf("token %s", os.Getenv("AUTH_TOKEN")))

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
		err = errors.Join(err, fmt.Errorf("%s : failed to send request", utils.WhereAmI()))
		return
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	var response struct {
		Data []T `json:"data"`
	}

	err = decoder.Decode(&response)
	if err != nil {
		err = errors.Join(err, fmt.Errorf("%s : failed to decode response", utils.WhereAmI()))
		return
	}

	result = response.Data

	return
}
