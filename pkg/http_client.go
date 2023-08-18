package pkg

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"noneland/backend/interview/pkg/errors"
)

func HttpDoReturnString(client *http.Client, req *http.Request) (body string, resp *http.Response, err error) {
	resp, err = client.Do(req)
	if err != nil {
		return "", nil, errors.Join3rdParty(errors.ErrSystem, err)
	}
	defer resp.Body.Close()

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", nil, errors.Join3rdParty(errors.ErrSystem, err)
	}

	return string(buf), resp, nil
}

func HttpDoReturnType[T any](client *http.Client, req *http.Request) (T, error) {
	data := new(T)

	resp, err := client.Do(req)
	if err != nil {
		return *data, errors.Join3rdParty(errors.ErrSystem, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, data)
	if err != nil {
		return *data, errors.Join3rdParty(errors.ErrSystem, err)
	}

	return *data, nil
}

func NewHttpClient() *http.Client {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}

	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	return &http.Client{
		Timeout:   time.Second * 10,
		Transport: transport,
	}
}
