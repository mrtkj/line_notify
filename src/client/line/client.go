package line

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type Config struct {
	Token string
}

type Client interface {
	SendMessage(msg string) error
}

func NewClient(config Config) Client {
	return clientImpl{
		config:     config,
		httpClient: &http.Client{},
	}
}

type clientImpl struct {
	config     Config
	httpClient *http.Client
}

func (c clientImpl) SendMessage(msg string) error {
	token := "Bearer " + c.config.Token
	apiURL := "https://notify-api.line.me/api/notify"

	values := url.Values{"message": {msg}}
	req, _ := http.NewRequest("POST", apiURL, strings.NewReader(values.Encode()))
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}

	dumpResp, _ := httputil.DumpResponse(resp, true)
	log.Printf("%s", dumpResp)
	return nil
}
