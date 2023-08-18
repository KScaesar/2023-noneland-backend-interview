package pkg

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"noneland/backend/interview/pkg/errors"
)

func HttpGetJsonBody[T any](client *http.Client, url string) (T, error) {
	data := new(T)

	resp, err := client.Get(url)
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
