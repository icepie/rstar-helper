package rstarhelper

const (
	LuaApiCallerPATH = "/v1/LuaApiCaller"
)

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
