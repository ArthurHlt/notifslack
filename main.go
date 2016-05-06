package main

import (
	"os"
	"github.com/codegangsta/cli"
	"fmt"
	"strings"
	"encoding/json"
	"net/http"
	"bytes"
	"errors"
	"strconv"
)

var channel string
var url string
var username string
var iconUrl string
var iconEmoji string

func main() {
	app := cli.NewApp()
	app.Name = "notifslack"
	app.Usage = "Send notification to slack"
	app.Version = "1.0.0"
	app.UsageText = "notifslack --url https://hooks.slack.com/services/XXXX [global options] \"my text to send\""
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "channel, c",
			Destination: &channel,
			Value: "",
			Usage: "Optional. Override channel to send message to. #channel and @user forms are allowed.",
		},
		cli.StringFlag{
			Name: "url",
			Destination: &url,
			Value: "",
			Usage: "Required. The webhook URL as provided by Slack. Usually in the form: https://hooks.slack.com/services/XXXX",
		},
		cli.StringFlag{
			Name: "username, u",
			Destination: &username,
			Value: "",
			Usage: "Optional. Override name of the sender of the message.",
		},
		cli.StringFlag{
			Name: "icon-url, i",
			Destination: &iconUrl,
			Value: "",
			Usage: "Optional. Override icon by providing URL to the image.",
		},
		cli.StringFlag{
			Name: "icon-emoji, e",
			Destination: &iconEmoji,
			Value: "",
			Usage: "Optional. Override icon by providing emoji code (e.g. :ghost:).",
		},
	}
	app.Action = sendIt
	app.Run(os.Args)
}
func sendIt(c *cli.Context) error {
	if url == "" {
		fmt.Println("You must set an url to slack, e.g.: notifslack --url https://hooks.slack.com/services/XXXX \"my text to send\"")
		os.Exit(1)
	}
	if len(c.Args()) == 0 {
		fmt.Println("You must set an url to slack, e.g.: notifslack --url https://hooks.slack.com/services/XXXX \"my text to send\"")
		os.Exit(1)
	}
	notif := &Notification{
		Text: strings.Join(c.Args(), " "),
	}
	if username != "" {
		notif.Username = username
	}
	if channel != "" {
		notif.Channel = channel
	}
	if iconUrl != "" {
		notif.IconURL = iconUrl
	}
	if iconEmoji != "" {
		notif.IconEmoji = iconEmoji
	}
	jsonMessage, err := json.MarshalIndent(notif, "", "\t")
	fatalIf("Error during mashalling", err)

	client := http.DefaultClient
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonMessage))
	fatalIf("Error during request creation", err)

	req.Header.Set("Content-type", "application/json")
	resp, err := client.Do(req)
	fatalIf("Error during sending notification", err)
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fatalIf("Something goes wrong", errors.New(strconv.Itoa(resp.StatusCode) + " " + resp.Status))
	}
	fmt.Println("Message sent.")
	return nil
}
func fatalIf(doing string, err error) {
	if err != nil {
		fatal(doing + ": " + err.Error())
	}
}
func fatal(message string) {
	fmt.Fprintln(os.Stdout, message)
	os.Exit(1)
}