package main

type Notification struct {
	Text      string      `json:"text"`
	Username  string      `json:"username"`
	IconURL   interface{} `json:"icon_url"`
	IconEmoji interface{} `json:"icon_emoji"`
	Channel   interface{} `json:"channel"`
}
