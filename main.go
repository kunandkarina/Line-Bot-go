package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

var (
	bot            *messaging_api.MessagingApiAPI
	channelSecret  string
	err            error
	notifyToken    string
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	channelSecret = os.Getenv("ChannelSecret")
	bot, err = messaging_api.NewMessagingApiAPI(
		os.Getenv("ChannelAccessToken"),
	)
	if err != nil {
		log.Fatal(err)
	}


	notifyToken = os.Getenv("LineNotifyToken")
	if notifyToken == "" {
		log.Fatal("NotifyToken is not set")
	}
	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/callback", callbackHandler)

	// This is just sample code.
	// For actual use, you must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	fmt.Println("http://localhost:" + port + "/")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func callbackHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("/callback called...")

	cb, err := webhook.ParseRequest(channelSecret, req)
	if err != nil {
		log.Printf("Cannot parse request: %+v\n", err)
		if errors.Is(err, webhook.ErrInvalidSignature) {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	log.Println("Handling events...")
	for _, event := range cb.Events {
		log.Printf("/callback called%+v...\n", event)

		switch e := event.(type) {
		case webhook.MessageEvent:
			switch message := e.Message.(type) {
			case webhook.TextMessageContent:
				usermessage := message.Text
				// Send notification via Line Notify
				if err := lineNotifyMessage(notifyToken, usermessage); err != nil {
					log.Printf("Failed to send Line Notify: %v\n", err)
				} else {
					log.Println("Sent Line Notify.")
				}
				// reply
				if _, err = bot.ReplyMessage(
					&messaging_api.ReplyMessageRequest{
						ReplyToken: e.ReplyToken,
						Messages: []messaging_api.MessageInterface{
							messaging_api.TextMessage{
								Text: "收到" + usermessage + "的需求了",
							},
						},
					},
				); err != nil {
					log.Print(err)
				} else {
					log.Println("Sent text reply.")
				}
			}
		default:
			log.Printf("Unsupported message: %T\n", event)
		}
	}
}

func lineNotifyMessage(token, msg string) error {
	apiUrl := "https://notify-api.line.me/api/notify"
	data := url.Values{}
	data.Set("message", msg)
	req, err := http.NewRequest("POST", apiUrl, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return errors.New("status code is not 200")
	}
	return nil
}