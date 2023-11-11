package wr840n

import (
	"net/http"
	"io"
	"strings"
)

func MakeRequest(url string, payload string, auth string) (string, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(payload))
	if err != nil {
		return "", err
	}
	req.Header.Add("Referer", "http://192.168.0.1/mainFrame.htm")
	req.Header.Add("Content-Type", "text/plain")
	req.Header.Add("charset", "UTF-8")
	req.Header.Add("Cookie", "Authorization="+auth)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}