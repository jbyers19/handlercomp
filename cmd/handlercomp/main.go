package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/jbyers19/handlercomp"

	"github.com/urfave/cli/v2"
)

var (
	// Create a new structured logger.
	logOptions = &slog.HandlerOptions{Level: slog.LevelInfo}
	logger     = slog.New(slog.NewTextHandler(os.Stdout, logOptions))

	listenAddr = ":8080"
)

func main() {
	// Update if a different log level is needed.
	logOptions.Level = slog.LevelInfo
	slog.SetDefault(logger)

	app := cli.NewApp()
	app.Name = "handlercomp"
	app.Usage = "Compare different web frameworks in Go."
	app.Args = true
	app.ArgsUsage = "Web framework to use: echo, gin, chi, http (stdlib net/http). Defaults to net/http router."
	app.Action = func(cCtx *cli.Context) error {
		selectFramework(cCtx.Args().First())
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func selectFramework(framework string) {
	switch framework {
	case "echo":
		e := handlercomp.NewEchoServer()
		handlercomp.StartEchoServer(e, listenAddr)
	case "gin":
		g := handlercomp.NewGinServer(listenAddr)
		handlercomp.StartHTTPServer(g)
	case "chi":
		c := handlercomp.NewChiServer(listenAddr)
		handlercomp.StartHTTPServer(c)
	default:
		s := handlercomp.NewHTTPServer(listenAddr)
		handlercomp.StartHTTPServer(s)
	}
}
