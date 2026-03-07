package request

import (
	"io"
	"net/http"
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

func SendRequest(Request Request) (Response, error) {
	startime := time.Now()
	req, err := http.NewRequest(Request.Method, Request.URL, nil)
	if err != nil {
		return Response{}, err
	}
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
