package api

import (
	"net/http"
	"strings"
)

func RequestToExternalAPI(requestBody []byte) (*http.Response, error) {
	resp, err := http.Post("https://sycret.ru/service/apigendoc/apigendoc", "application/json", strings.NewReader(string(requestBody)))
	if err != nil {
		return nil, err
	}
	return resp, nil
}
