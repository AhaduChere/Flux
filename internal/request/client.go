package request

import (
	"io"
	"net/http"
	"strings"
	"time"
)

type Request struct {
	URL    string
	Method string
	Body   *string
}

type Response struct {
	StatusCode int
	Duration   time.Duration
	Body       string
}

// NOTE: Sends request and returns response
func SendRequest(Request Request) (Response, error) {
	startime := time.Now()
	var bodyReader io.Reader
	if Request.Body != nil {
		bodyReader = strings.NewReader(*Request.Body)
	}
	req, err := http.NewRequest(Request.Method, Request.URL, bodyReader)
	if err != nil {
		return Response{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()
	StatusCode := resp.StatusCode
	duration := time.Since(startime)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Response{}, err
	}

	return Response{
		Body:       string(body),
		StatusCode: StatusCode,
		Duration:   duration,
	}, nil

}
