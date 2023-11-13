package main

import (
	"github.com/hanakogo/exceptiongo"
	"github.com/kmou424/hanabot/bot"
	"github.com/kmou424/hanabot/internal/env"
	"github.com/kmou424/hanabot/internal/i18n"
	"os"
)

func preload() {
	env.Load()
	i18n.Load()
}

func main() {
	defer exceptiongo.NewExceptionHandler(func(e *exceptiongo.Exception) {
		e.PrintStackTrace()
		os.Exit(1)
	}).Deploy()

	preload()
	bot.Start()
}
