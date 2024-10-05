package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

var (
	token   string
	channel string
	port    int
	webhook string
	api     *slack.Client
)

func main() {
	flag.StringVar(&token, "token", "", "Slack Bot Token")
	flag.StringVar(&channel, "channel", "", "Slack Channel ID")
	flag.IntVar(&port, "port", 8080, "Port to listen on")
	flag.StringVar(&webhook, "webhook", "", "Webhook URL (optional)")
	flag.Parse()

	if token == "" || channel == "" {
		log.Fatal("Both token and channel must be provided")
	}

	api = slack.New(token)

	http.HandleFunc("/slack/events", handleSlackEvents)
	http.HandleFunc("/send", handleIncomingMessage)

	log.Printf("slack-bot-server is running on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func handleSlackEvents(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ev, err := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionNoVerifyToken())
	if err != nil {
		log.Printf("Error parsing Slack event: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if ev.Type == slackevents.URLVerification {
		var r *slackevents.ChallengeResponse
		err := json.Unmarshal([]byte(body), &r)
		if err != nil {
			log.Printf("Error unmarshalling challenge response: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text")
		w.Write([]byte(r.Challenge))
	}

	if ev.Type == slackevents.CallbackEvent {
		innerEvent := ev.InnerEvent
		switch ev := innerEvent.Data.(type) {
		case *slackevents.MessageEvent:
			if ev.ChannelType == "channel" && ev.Channel == channel {
				if webhook != "" {
					sendToWebhook(ev.Text)
				}
			}
		}
	}
}

func handleIncomingMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	_, _, err = api.PostMessage(channel, slack.MsgOptionText(string(body), false))
	if err != nil {
		log.Printf("Error sending message to Slack: %v", err)
		http.Error(w, fmt.Sprintf("Error sending message to Slack: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func sendToWebhook(message string) {
	resp, err := http.Post(webhook, "application/json", strings.NewReader(message))
	if err != nil {
		log.Printf("Error sending message to webhook: %v", err)
		return
	}
	defer resp.Body.Close()
}
