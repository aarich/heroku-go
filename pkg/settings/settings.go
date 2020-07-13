package settings

import (
	"log"

	"github.com/go-ini/ini"
)

type AppSettings struct {
	RuntimeRootPath  string
	ImageSavePath    string
	ImageAllowedExts []string
	ImageMaxSize     int
	PrefixUrl        string
}

var App = &AppSettings{}
var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")

	if err != nil {
		log.Fatalf("Settings setup failed. Failed to parse app.ini: %v", err)
	}

	mapTo("app", App)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.mapto failed for section %s: %v", section, err)
	}
}
