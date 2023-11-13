package env

import (
	"github.com/gookit/goutil/strutil"
)

func apply() {
	BotToken = readString("BOT_TOKEN")

	OwnerId = readInt64("OWNER_ID")

	AllowedUserIds = readStringSliceWithTrim("ALLOWED_USER_IDS", func(s string) string {
		return strutil.Trim(s)
	})
}
