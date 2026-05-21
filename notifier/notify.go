package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/SHAIK14/pulsemon/checker"
)

type Body struct {
	ChatId string `json:"chat_id"`
	Text   string `json:"text"`
}

func Notify(result checker.Result) error {
	token := os.Getenv("TELEGRAM_TOKEN")
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)
	parts := strings.Split(result.Error.Error(), ": ")
	msgError := parts[len(parts)-1]
	msg := fmt.Sprintf("🚨 Down:%s\nURL:%s\nError:%s\n", result.Name, result.URL, msgError)
	body := Body{
		ChatId: os.Getenv("TELEGRAM_CHAT_ID"),
		Text:   msg,
	}

	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(jsonBytes)

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Post(url, "application/json", reader)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil

}
