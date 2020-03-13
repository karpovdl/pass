package main

import (
	"bufio"
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spf13/afero"
	cli "github.com/urfave/cli/v2"
)

/* Telegram block */

/* Configuration file JSON */
type telegramConfiguration struct {
	BotToken  string `json:"bot_token"`
	ChannelID int64  `json:"channel_id"`
	Message   string `json:"message"`
}

/* Constant for telegram */
const (
	telegram               = "telegram"
	telegramAlias          = "tm"
	telegramBotTokenName   = "bot_token"
	telegramBotTokenAlias  = "bt"
	telegramChannelIDName  = "channel_id"
	telegramChannelIDAlias = "cid"
	telegramMessageName    = "message"
	telegramMessageAlias   = "m"
)

/* Variable for telegram */
var (
	telegramBotToken  string
	telegramСhannelID int64
	telegramMessage   string
)

func telegramReadConfiguration() error {
	var (
		file    afero.File
		conf    telegramConfiguration
		decoder *json.Decoder
		err     error
	)

	if file, err = appFs.Open("conf/telegram.json"); err != nil {
		return err
	}

	defer file.Close()

	decoder = json.NewDecoder(file)
	conf = telegramConfiguration{}

	if err = decoder.Decode(&conf); err != nil {
		return err
	}

	telegramBotToken = conf.BotToken
	telegramСhannelID = conf.ChannelID
	telegramMessage = conf.Message

	return nil
}

func telegramCommand() *cli.Command {
	return &cli.Command{
		Name:        telegram,
		Aliases:     []string{telegramAlias},
		Usage:       "",
		Description: "pass to telegram chanel",
		Flags: []cli.Flag{
			telegramFlagBotToken(),
			telegramFlagChannelID(),
			telegramFlagMessage(),
		},
		Action: telegramAction,
	}
}

func telegramFlagBotToken() *cli.StringFlag {
	return &cli.StringFlag{
		Name:    telegramBotTokenName,
		Aliases: []string{telegramBotTokenAlias},
		Value:   "",
		Usage:   "bot token `uuid`",
	}
}

func telegramFlagChannelID() *cli.Int64Flag {
	return &cli.Int64Flag{
		Name:    telegramChannelIDName,
		Aliases: []string{telegramChannelIDAlias},
		Usage:   "channel `id`",
	}
}

func telegramFlagMessage() *cli.StringFlag {
	return &cli.StringFlag{
		Name:    telegramMessageName,
		Aliases: []string{telegramMessageAlias},
		Value:   "",
		Usage:   "message `text`",
	}
}

func telegramAction(c *cli.Context) error {
	cli.ShowVersion(c)

	var bufout = bufio.NewWriter(os.Stdout)

	telegramBotToken = c.String(telegramBotTokenName)
	bufout.WriteString("Bot token '" + telegramBotToken + "'" + Endl)

	telegramСhannelID = c.Int64(telegramChannelIDName)
	bufout.WriteString("Channel ID '" + strconv.FormatInt(telegramСhannelID, 10) + "'" + Endl)

	telegramMessage = c.String(telegramMessageName)
	bufout.WriteString("Message '" + telegramMessage + "'" + Endl)

	bufout.Flush()

	telegramSend()

	return nil
}

func telegramHandler(w http.ResponseWriter, r *http.Request) {
	var (
		u    *url.URL
		resp Resp
		err  error
	)

	w.Header().Set("Control-Type", "application/json")

	if u, err = url.Parse(r.URL.String()); err != nil {
		resp = Resp{
			Message: strings.Join([]string{"Path:", u.Path, "Raw query:", u.RawQuery}, " "),
			Error:   err.Error(),
		}
	} else {
		for _, s := range strings.Split(u.RawQuery, "&") {
			v := strings.Split(s, "=")
			if len(v) == 2 {
				switch v[0] {
				case telegramBotTokenName:
				case telegramBotTokenAlias:
					telegramBotToken = v[1]
				case telegramChannelIDName:
				case telegramChannelIDAlias:
					telegramСhannelID, _ = strconv.ParseInt(v[1], 10, 64)
				case telegramMessageName:
				case telegramMessageAlias:
					telegramMessage = v[1]
				default:
				}
			}
		}

		resp = Resp{
			Message: "Bot token '" + telegramBotToken + "'" + Endl + "Channel ID '" + strconv.FormatInt(telegramСhannelID, 10) + "'" + Endl + "Message '" + telegramMessage + "'" + Endl,
		}
		w.WriteHeader(http.StatusOK)
	}

	respJSON, _ := json.Marshal(resp)
	w.Write(respJSON)

	telegramSend()
}

func telegramSend() (tgbotapi.Message, error) {
	var (
		bot    *tgbotapi.BotAPI
		msg    tgbotapi.Message
		msgCfg tgbotapi.MessageConfig
		err    error
	)

	if bot, err = tgbotapi.NewBotAPI(telegramBotToken); err != nil {
		return tgbotapi.Message{}, err
	}

	msgCfg = tgbotapi.NewMessage(telegramСhannelID, telegramMessage)

	if msg, err = bot.Send(msgCfg); err != nil {
		return tgbotapi.Message{}, err
	}

	return msg, nil
}
