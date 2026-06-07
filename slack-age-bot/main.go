package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/slack-io/slacker"
)

func main() {

	os.Setenv("SLACK_BOT_TOKEN", "GET YOUR OWN TOKEN :)")
	os.Setenv("SLACK_APP_TOKEN", ";)")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	bot.AddCommand(&slacker.CommandDefinition{
		Command: "ping",
		Handler: func(ctx *slacker.CommandContext) {
			ctx.Response().Reply("pong")
		},
	})

	bot.AddCommand(&slacker.CommandDefinition{
		Command:     "My yob is <year>",
		Description: "YOB Calculator",
		Handler: func(cc *slacker.CommandContext) {
			year := cc.Request().Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				log.Printf("Error", err)
				return
			}
			age := 2026 - yob
			r := fmt.Sprintf("Your age is %d", age)
			cc.Response().Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
