package bot

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/gookit/goutil/strutil"
	"github.com/hanakogo/exceptiongo"
	"github.com/hanakogo/hanakoutilgo"
	"github.com/kmou424/hanabot/internal/env"
	"github.com/kmou424/hanabot/internal/i18n"
	"github.com/kmou424/hanabot/internal/logger"
	"github.com/kmou424/hanabot/internal/modules"
	"github.com/kmou424/hanabot/internal/types"
	"log"
	"net/http"
	"time"
)

func Start() {
	bot, err := gotgbot.NewBot(env.BotToken, &gotgbot.BotOpts{
		BotClient: &gotgbot.BaseBotClient{
			Client: http.Client{},
			DefaultRequestOpts: &gotgbot.RequestOpts{
				Timeout: gotgbot.DefaultTimeout, // Customise the default request timeout here
				APIURL:  gotgbot.DefaultAPIURL,  // As well as the Default API URL here (in case of using local bot API servers)
			},
		},
	})
	if err != nil {
		exceptiongo.ThrowMsgF[types.FatalError]("%s: %s", i18n.GetTr("bot_create_failed"), err.Error())
	}

	updater := registerUpdater(bot)

	updater.Idle()
}

func registerDispatcher() *ext.Dispatcher {
	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		// If an error is returned by a handler, log it and continue going.
		Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) (action ext.DispatcherAction) {
			log.Println("an error occurred while handling update:", err.Error())
			return ext.DispatcherActionNoop
		},
		Panic: func(b *gotgbot.Bot, ctx *ext.Context, r interface{}) {
			if hanakoutilgo.Is[*exceptiongo.Exception](r) {
				stackTraceMessage := hanakoutilgo.CastTo[*exceptiongo.Exception](r).GetStackTraceMessage()
				logger.Get().Warn(strutil.FirstLine(stackTraceMessage))
			}
		},
		UnhandledErrFunc: func(err error) {
			logger.Get().Error(err.Error())
		},
		MaxRoutines: ext.DefaultMaxRoutines,
	})

	for _, handler := range modules.GetModulesHandler() {
		dispatcher.AddHandler(handler)
	}

	return dispatcher
}

func registerUpdater(bot *gotgbot.Bot) *ext.Updater {
	updater := ext.NewUpdater(&ext.UpdaterOpts{
		Dispatcher: registerDispatcher(),
	})

	err := updater.StartPolling(bot, &ext.PollingOpts{
		DropPendingUpdates: true,
		GetUpdatesOpts: &gotgbot.GetUpdatesOpts{
			Timeout: 9,
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: time.Second * 10,
			},
		},
	})
	if err != nil {
		exceptiongo.ThrowMsg[types.FatalError](i18n.GetTr("bot_start_polling_failed"))
	}

	sayHello(bot)

	return updater
}
