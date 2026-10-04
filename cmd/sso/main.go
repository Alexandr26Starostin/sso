package main

import (
	"log/slog"
	"os"
	"sso/internal/config"
	"sso/internal/lib/logger/handlers/slogpretty"
	"sso/internal/lib/logger/sl"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

// точка входа в приложение
func main() {
	//----------------------------------------------
	//инициализировать объект конфига
	cfg := config.MustLoad()

	//fmt.Println(cfg)

	//----------------------------------------------
	//инициализировать логгер
	log := setupLogger(cfg.Env)

	log.Info("starting application",
		slog.String("env", cfg.Env),
		slog.Any("cfg", cfg),
		slog.Int("port", cfg.GRPC.Port),
	)

	log.Debug("debug message")
	log.Error("error message")
	log.Warn("warn message")

	//----------------------------------------------
	//Использование sl.go
	var err error

	if err != nil {
		log.Error("error message", slog.String("error", err.Error()))
		//or:
		log.Error("error message", sl.Err(err))
	}
	//----------------------------------------------

	//TODO: инициализировать приложение (app)
	//приложение будет запускаться в пакете app
	//точка запуска не обязательно main; это может быть тест

	//TODO: запустить gRPC-сервер приложения, которое мы до этого инициализировали
}

// slog -- wrapper for any logger
func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		// log = slog.New(
		// 	slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		// )

		//or: with own slogpretty.go:
		log = setupPrettySlog()

	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)

	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
