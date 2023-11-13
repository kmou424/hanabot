package bot

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/hanakogo/exceptiongo"
	"github.com/kmou424/hanabot/internal/env"
	"github.com/kmou424/hanabot/internal/i18n"
	"github.com/kmou424/hanabot/internal/logger"
	"github.com/kmou424/hanabot/internal/types"
)

func sayHello(bot *gotgbot.Bot) {
	logger.Get().Infof(i18n.GetTr("bot_create_success", bot.User.Username))
	_, err := bot.SendMessage(env.OwnerId, i18n.GetTr("bot_message_hello"), &gotgbot.SendMessageOpts{})
	exceptiongo.ThrowErr[types.BotSendMessageError](err)
}
