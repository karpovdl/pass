package main

import (
	"bufio"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"

	"net/http"
	_ "net/http/pprof"

	cli "github.com/urfave/cli/v2"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	initApp()

	if isPprof {
		http.ListenAndServe("localhost:"+strconv.Itoa(pprofPort), nil)
	}
}

func initApp() bool {
	app := cli.App{
		Name:     "pass",
		Version:  "v1.0.0.0",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "Denis Karpov",
				Email: "KarpovDL@hotmail.com",
			},
		},
		Copyright: "(c) 2020 Denis Karpov",
		HelpName:  "pass",
		Usage:     "application for pass data",

		EnableBashCompletion:   true,
		HideHelp:               false,
		HideVersion:            false,
		UseShortOptionHandling: true,

		Commands: []*cli.Command{
			&cli.Command{
				Name:        "run",
				Aliases:     []string{"r"},
				Usage:       "",
				Description: "",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "port",
						Aliases: []string{"p"},
						Value:   9000,
						Usage:   "`port` to use listen server",
					},
					&cli.BoolFlag{
						Name:    "pprof",
						Aliases: []string{"pf"},
						Value:   false,
						Usage:   "if `flag` [true], then active pprof, otherwise [false]",
					},
					&cli.IntFlag{
						Name:    "pprof_port",
						Aliases: []string{"pfp"},
						Value:   9001,
						Usage:   "`port` to use pprof listen server",
					},
				},
				Subcommands: []*cli.Command{
					telegramCommand(),
				},
				Before: func(c *cli.Context) error {
					var err error

					/* Read configuration files */

					if err = telegramReadConfiguration(); err != nil {
						return err
					}

					return nil
				},
				Action: func(c *cli.Context) error {
					cli.ShowVersion(c)

					var bufout = bufio.NewWriter(os.Stdout)

					port = c.Int("port")
					bufout.WriteString("Port '" + strconv.Itoa(port) + "'" + Endl)

					isPprof = c.Bool("pprof")
					if isPprof {
						bufout.WriteString("Mode is pprof '" + strconv.FormatBool(isPprof) + "'" + Endl)

						pprofPort = c.Int("pprof_port")
						bufout.WriteString("Pprof port '" + strconv.Itoa(pprofPort) + "'" + Endl)
					}

					bufout.Flush()

					if isPprof {
						go runServer()
					} else {
						runServer()
					}

					return nil
				},
			},
		},

		Action: func(c *cli.Context) error {
			cli.ShowVersion(c)

			return nil
		},
	}

	app.Run(os.Args)

	return true
}

func logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("server (net/http) method [%s] connection from [%v]", r.Method, r.RemoteAddr)
		next.ServeHTTP(w, r)
	}
}

func runServer() {
	handler := http.NewServeMux()
	handler.HandleFunc("/"+telegramAlias+"/", logger(telegramHandler))

	s := &http.Server{
		Addr:           ":" + strconv.Itoa(port),
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
