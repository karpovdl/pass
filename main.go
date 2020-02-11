package main

import (
	"bufio"
	"os"
	"runtime"
	"strconv"
	"time"

	"net/http"
	_ "net/http/pprof"

	cli "github.com/urfave/cli"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	initApp()

	if isPprof {
		http.ListenAndServe("localhost:8080", nil)
	}
}

func initApp() bool {
	app := cli.App{
		Name:      "pass",
		Version:   "v1.0.0.0",
		Compiled:  time.Now(),
		Author:    "Denis Karpov",
		Email:     "KarpovDL@hotmail.com",
		Copyright: "(c) 2020 Denis Karpov",
		HelpName:  "pass",
		UsageText: "pass - application for pass data",

		Commands: []cli.Command{
			cli.Command{
				Name: "run",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:  "pprof, pf",
						Usage: "If [true], then active pprof, otherwise [false]",
					},
				},
				Action: func(c *cli.Context) error {
					cli.ShowVersion(c)

					var bufout = bufio.NewWriter(os.Stdout)

					isPprof = c.Bool("pprof")
					bufout.WriteString("Mode is pprof '" + strconv.FormatBool(isPprof) + "'" + Endl)

					bufout.Flush()

					//run()
					return nil
				},
			},
		},

		Action: func(c *cli.Context) error {
			cli.ShowVersion(c)

			return nil
		},

		EnableBashCompletion: true,
		HideHelp:             false,
		HideVersion:          false,
	}

	app.Run(os.Args)

	return true
}
