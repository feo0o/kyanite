package server

import (
	"flag"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type config struct {
	ConfigFile   string `mapstructure:"config_file"`
	BindAddr     string `mapstructure:"bind_addr"`
	BindPort     int    `mapstructure:"bind_port"`
	EnableTLS    bool   `mapstructure:"enable_tls"`
	CertFile     string `mapstructure:"cert_file"`
	KeyFile      string `mapstructure:"key_file"`
	EnableACCLog bool   `mapstructure:"enable_acc_log"`
	AccLogOutput string `mapstructure:"acc_log_output"`
	AccLogFile   string `mapstructure:"acc_log_file"`
}

var cfg config

func init() {
	flag.StringVar(&cfg.ConfigFile, "config-file", "", "path to config file")
	flag.StringVar(&cfg.BindAddr, "bind-addr", "", "local IP address to use, default is empty to bind on all network interface")
	flag.IntVar(&cfg.BindPort, "bind-port", 8080, "local net port to use, default: 8080")
	flag.BoolVar(&cfg.EnableTLS, "enable-tls", false, "enable TLS to use HTTPS, default: false")
	flag.StringVar(&cfg.CertFile, "cert-file", "", "path to TLS cert file, needed when enable tls")
	flag.StringVar(&cfg.KeyFile, "key-file", "", "path to TLS key file, needed wht enable tls")
	flag.BoolVar(&cfg.EnableACCLog, "enable-acc-log", true, "enable access log")
	flag.StringVar(&cfg.AccLogOutput, "acc-log-output", "stdout", "location to logging access log, valid value: 'stdout','file'. default: stdout")
	flag.StringVar(&cfg.AccLogFile, "acc-log-file", "", "path to access log file, needed when --acc-log-output is 'file'")

	flag.Parse()

	if cfg.ConfigFile != "" {
		viper.SetConfigFile(cfg.ConfigFile)
		if err := viper.ReadInConfig(); err != nil {
			log.Fatal(err)
		}
		if err := viper.Unmarshal(&cfg); err != nil {
			log.Fatal(err)
		}
	}
}

func Run() (err error) {
	svr := gin.Default()
	if cfg.EnableACCLog {
		// todo: set svr to use logger with config
		svr.Use(gin.Logger())
	}
	route(svr)

	addr := fmt.Sprintf("%s:%d", cfg.BindAddr, cfg.BindPort)
	if cfg.EnableTLS {
		return svr.RunTLS(addr, cfg.CertFile, cfg.KeyFile)
	}
	return svr.Run(addr)
}
