package i18n

import (
	"embed"
	"fmt"
	"github.com/Xuanwo/go-locale"
	"github.com/gookit/goutil/strutil"
	"github.com/gookit/i18n"
	"github.com/hanakogo/exceptiongo"
	"github.com/kmou424/hanabot/internal/logger"
	"github.com/kmou424/hanabot/internal/types"
)

const langDir = "lang"

//go:embed lang
var embedLang embed.FS

var i18nGetter *i18n.I18n

func readLanguages() {
	dir, err := embedLang.ReadDir(langDir)
	if err != nil {
		exceptiongo.ThrowMsg[types.FatalError]("failed to load language files")
	}
	for _, entry := range dir {
		langName := entry.Name()
		if strutil.IsEndOf(langName, ".ini") {
			langName = langName[:len(langName)-4]
		}
		content, err := embedLang.ReadFile(fmt.Sprintf("lang/%s.ini", langName))
		if err != nil {
			exceptiongo.ThrowMsgF[types.FatalError]("failed to read language <%s> configuration: %s", langName, err.Error())
		}
		i18nGetter.AddLang(langName, langName)
		err = i18nGetter.LoadString(langName, string(content))
		if err != nil {
			exceptiongo.ThrowMsgF[types.FatalError]("failed to load language <%s>: %s", langName, err.Error())
		}
	}
}

func loadI18n(lang string) {
	i18nGetter = i18n.NewEmpty()
	i18nGetter.DefaultLang = lang
	i18nGetter.FallbackLang = "en"
	readLanguages()
}

func Load() {
	defLang := "en"

	loadI18n(defLang)

	detectLang, err := locale.Detect()
	if err == nil {
		if detectLang := detectLang.String(); HasLang(detectLang) {
			defLang = detectLang
		} else {
			logger.Get().Warn(GetTr("language_not_support", detectLang))
		}
	} else {
		logger.Get().Warn(fmt.Sprintf("failed to detect language: %s", err.Error()))
	}

	if defLang != "en" {
		loadI18n(defLang)
	}
}

func HasLang(lang string) bool {
	return i18nGetter.HasLang(lang)
}

func GetTr(key string, args ...any) string {
	if len(args) > 0 {
		return i18nGetter.Dt(key, args)
	}
	return i18nGetter.Dt(key)
}

func GetTrl(lang string, key string, args ...any) string {
	if len(args) > 0 {
		return i18nGetter.T(lang, key, args)
	}
	return i18nGetter.T(lang, key)
}
