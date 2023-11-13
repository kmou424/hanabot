package env

import (
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/strutil"
	"github.com/gookit/ini/v2/dotenv"
	"github.com/hanakogo/exceptiongo"
	"github.com/kmou424/hanabot/internal/i18n"
	"github.com/kmou424/hanabot/internal/logger"
	"github.com/kmou424/hanabot/internal/types"
	"os"
)

const envFile = "hanabot.env"

func Load() {
	targetEnvFile := envFile
	if !fsutil.FileExists(envFile) {
		exceptiongo.ThrowMsg[types.DotEnvFileNotFoundError](i18n.GetTr("file_not_found", envFile))
	}

	envDotEnvPath := os.Getenv("DOTENV_PATH")
	if !strutil.IsEmpty(envDotEnvPath) {
		if fsutil.FileExists(envDotEnvPath) {
			targetEnvFile = envDotEnvPath
		} else {
			logger.Get().Warn(i18n.GetTr("file_not_found", envDotEnvPath))
		}
	}

	err := dotenv.LoadFiles(targetEnvFile)
	if err != nil {
		logger.Get().Warn(i18n.GetTr("file_load_error", targetEnvFile))
		return
	}

	apply()
}
