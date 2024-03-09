package handler

import (
	"github.com/woxQAQ/im-service/config"
)

type Option func(opt *config.HandlerConfig)

type conf struct {
}

func defaultOptions() {

}

func loadOptions(opt []Option, config *config.HandlerConfig) {

}
