package api

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func DoGET(url string, queryParams [][]string) ([]byte, error) {
	if queryParams != nil {
		var params []string
		for _, param := range queryParams {
			params = append(params, param[0]+"="+param[1])
		}
		url = url + "?" + strings.Join(params, "&")
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://"+url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
