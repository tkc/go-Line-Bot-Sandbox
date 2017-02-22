package app

import (
	"os"
	"log"
	"net/http"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
	aelog "google.golang.org/appengine/log"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/line/line-bot-sdk-go/linebot/httphandler"
)

func init() {

	handler, err := httphandler.New(
		os.Getenv("LINE_API_SECRET"),
		os.Getenv("LINE_API_TOKEN"),
	)

	if err != nil {
		log.Fatal(err)
	}

	handler.HandleEvents(func(events []*linebot.Event, r *http.Request) {

		ctx := appengine.NewContext(r)
		bot, err := handler.NewClient(linebot.WithHTTPClient(urlfetch.Client(ctx)))

		if err != nil {
			aelog.Errorf(ctx, "%v", err)
			return
		}

		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).WithContext(ctx).Do(); err != nil {
						aelog.Errorf(ctx, "%v", err)
					}
				}
			}
		}
	})

	http.Handle("/callback", handler)
}

