package rstarhelper

import (
	"encoding/json"
	"errors"

	"github.com/go-resty/resty/v2"
)

// Bot
type Bot struct {
	Url        string
	WxID       string
	HttpClient *resty.Client
	Option     map[string]string // Bot Http Client Option
}

// NewBot
func (bm *BotManager) NewBot(wxID string) *Bot {
	return &Bot{
		Url:        bm.Url,
		WxID:       wxID,
		HttpClient: resty.New(),
		Option: map[string]string{
			"wxid":    wxID,
			"timeout": "10",
		},
	}
}

// InitWxids
func (bot *Bot) InitWxids(body InitWxidsReq) (rsp InitWxidsRsp, err error) {

	var errRsp ErrorRsp

	resp, err := bot.HttpClient.R().
		SetBody(body).
		SetQueryParams(bot.Option).
		SetQueryParam("funcname", "InitWxids").
		Post(bot.Url + LuaApiCallerPATH)

	if err != nil {
		return
	}

	// json unmarshal
	_ = json.Unmarshal(resp.Body(), &errRsp)

	if errRsp.Ret != 0 {
		err = errors.New(errRsp.ErrMsg)
		return
	}

	if err = json.Unmarshal(resp.Body(), &rsp); err != nil {
		return
	}

	return
}
