package rstarhelper

import "github.com/go-resty/resty/v2"

// BotManager
type BotManager struct {
	Url string
}

// NewBotManager
func NewBotManager(url string) *BotManager {
	return &BotManager{
		Url: url,
	}
}

// Bot
type Bot struct {
	Url        string
	WxID       string
	HttpClient *resty.Client
}

// NewBot
func (bm *BotManager) NewBot(wxID string) *Bot {
	return &Bot{
		Url:        bm.Url,
		WxID:       wxID,
		HttpClient: resty.New(),
	}
}
