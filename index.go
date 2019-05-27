package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

var (
	callbackURL = "https://api.telegram.org/bot" + os.Getenv("BOT_TOKEN") + "/sendMessage"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	var v telegram
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if err := v.Validate(); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if err := handle(&v); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

type telegram struct {
	UpdateID int `json:"update_id"`
	Message  struct {
		MessageID int `json:"message_id"`
		From      struct {
			ID           int    `json:"id"`
			IsBot        bool   `json:"is_bot"`
			FirstName    string `json:"first_name"`
			LastName     string `json:"last_name"`
			LanguageCode string `json:"language_code"`
		} `json:"from"`
		Chat struct {
			ID        int    `json:"id"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Type      string `json:"type"`
		} `json:"chat"`
		Date int    `json:"date"`
		Text string `json:"text"`
	} `json:"message"`
}

func (t telegram) Validate() error {
	fmt.Fprintf(os.Stderr, "%+v", t)
	return nil
}

func handle(t *telegram) error {
	value, err := lookup(t.Message.Text)
	if err != nil {
		return err
	}

	var response = struct {
		ChatID int    `json:"chat_id"`
		Text   string `json:"text"`
	}{
		ChatID: t.Message.Chat.ID,
		Text:   value,
	}
	b, err := json.Marshal(response)
	if err != nil {
		return err
	}
	res, err := http.Post(callbackURL, "application/json", bytes.NewReader(b))
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode >= 400 {
		io.Copy(os.Stderr, res.Body)
		return fmt.Errorf("bad status: %d", res.StatusCode)
	}
	return nil
}

func lookup(msg string) (string, error) {
	return os.Getenv("BLIMPY_" + strings.ReplaceAll(strings.ToUpper(msg), " ", "_")), nil
}
