package utils

import (
	"context"

	config2 "github.com/marco79423/sai-portal/service/config"
)

func GetCtxLogger(ctx context.Context) LoggerInterface {
	return ctx.Value("logger").(LoggerInterface)
}

func GetCtxConfig(ctx context.Context) config2.IConfig {
	return ctx.Value("config").(config2.IConfig)
}
