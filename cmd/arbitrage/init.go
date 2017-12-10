package main

import (
	"runtime"
	"strings"
	"time"

	log "./sirupsen/logrus"
	"./utilgo/conflag"
	flags "github.com/jessevdk/go-flags"
)

// Opts configs
var Opts struct {
	Conf      string   `long:"conf" description:"config file"`
	Exchanges []string `long:"exchanges" description:"the exchanges list for arbitrage"`
	Base      string   `long:"base" default:"cny" description:"base currency"`
	Quote     string   `long:"quote" default:"eth" description:"quote currency"`
	SlackKey  string   `long:"slack-key" description:"slack key"`
	LogLevel  string   `long:"log-level" default:"info" description:"Adjust the log level. Valid options are: error, warn, info, debug"`
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func init() {
	parser := flags.NewParser(&Opts, flags.Default|flags.IgnoreUnknown)

	parser.Parse()

	if Opts.Conf != "" {
		conflag.LongHyphen = true
		conflag.BoolValue = false
		args, err := conflag.ArgsFrom(Opts.Conf)
		if err != nil {
			panic(err)
		}

		parser.ParseArgs(args)
	}

	log.Infof("opts: %+v", Opts)
}

func init() {
	if level, err := log.ParseLevel(strings.ToLower(Opts.LogLevel)); err != nil {
		log.SetLevel(level)
	}

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	})
}
