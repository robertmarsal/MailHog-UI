package config

import (
	"flag"

	"github.com/ian-kent/envconf"
)

func DefaultConfig() *Config {
	return &Config{
		APIHost:    "",
		UIBindAddr: "0.0.0.0:8025",
		DarkMode:   false,
		WebPath:    "",
	}
}

type Config struct {
	APIHost    string
	UIBindAddr string
	DarkMode   bool
	WebPath    string
}

var cfg = DefaultConfig()

func Configure() *Config {
	return cfg
}

func RegisterFlags() {
	flag.StringVar(&cfg.APIHost, "api-host", envconf.FromEnvP("MH_API_HOST", "").(string), "API URL for MailHog UI to connect to, e.g. http://some.host:1234")
	flag.StringVar(&cfg.UIBindAddr, "ui-bind-addr", envconf.FromEnvP("MH_UI_BIND_ADDR", "0.0.0.0:8025").(string), "HTTP bind interface and port for UI, e.g. 0.0.0.0:8025 or just :8025")
	flag.BoolVar(&cfg.DarkMode, "dark-mode", envconf.FromEnvP("MH_DARK_MODE", false).(bool), "Enable dark mode")
}
