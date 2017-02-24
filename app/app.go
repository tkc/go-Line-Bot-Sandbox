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

				log.Print(event.Source.UserID)
				log.Print(event.Source.RoomID)
				log.Print(event.Source.GroupID)
				log.Print(event.Source.GroupID)
				log.Print(event.Message)
				log.Print(event.Postback)

				switch message := event.Message.(type) {
				case *linebot.TextMessage:

					var selectText = message.Text
					switch selectText {
					case NormalMessage:SendNormalMessage(bot, event.ReplyToken, ctx)
					case CarouselMessage:SendCarouselMessage(bot, event.ReplyToken, ctx)
					case ImageMessage:SendImageMessage(bot, event.ReplyToken, ctx)
					case LocationMessage:SendlocationMessage(bot, event.ReplyToken, ctx)
					case MoreSelect:SendMoreSelectMessage(bot, event.ReplyToken, ctx)
					default: SendSelectMessage(bot, event.ReplyToken, ctx)
					}

				}
			}
		}
	})
	
	http.Handle("/callback", handler)
}

