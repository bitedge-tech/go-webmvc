package types

import (
	"go-webmvc/config"
	"time"
)

type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {
	loc, _ := time.LoadLocation(config.AppConfig.App.Timezone)
	formatted := "\"" + time.Time(t).In(loc).Format("2006-01-02 15:04:05") + "\""
	return []byte(formatted), nil
}
