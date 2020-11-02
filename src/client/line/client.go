package line

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

func SendMessage(msg string) {
	token := "Bearer " + os.Getenv("LINE_NOTIFY_TOKEN")
	apiURL := "https://notify-api.line.me/api/notify"

	values := url.Values{"message": {msg}}
	req, _ := http.NewRequest("POST", apiURL, strings.NewReader(values.Encode()))
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	dumpResp, _ := httputil.DumpResponse(resp, true)
	log.Printf("%s", dumpResp)
}
