package app

import (
	"golang.org/x/net/context"
	aelog "google.golang.org/appengine/log"
	"github.com/line/line-bot-sdk-go/linebot"
)

var imgUrl = "https://line-bot-159514.appspot.com/img/ai.png"

func SendSelectMessage(bot *linebot.Client, replyToken string, ctx context.Context) {

	title := "title"
	subscription := "subscriotion"
	titleSub := "titleSub"
	template := linebot.NewButtonsTemplate(
		imgUrl,
		title,
		titleSub,
		linebot.NewMessageTemplateAction(CarouselMessage, CarouselMessage),
		linebot.NewMessageTemplateAction(ImageMessage, ImageMessage),
		linebot.NewPostbackTemplateAction(LocationMessage, LocationMessage, LocationMessage),
		linebot.NewMessageTemplateAction(MoreSelect, MoreSelect),
	)

	messageTitle := "message"
	message1 := linebot.NewTextMessage(messageTitle)
	message2 := linebot.NewTemplateMessage(subscription, template)

	if _, err := bot.ReplyMessage(replyToken, message1, message2).WithContext(ctx).Do(); err != nil {
		aelog.Errorf(ctx, "%v", err)
	}
}

func SendMoreSelectMessage(bot *linebot.Client, replyToken string, ctx context.Context) {

	title := "title"
	subscription := "subscriotion"
	titleSub := "titleSub"
	template := linebot.NewButtonsTemplate(
		imgUrl,
		title,
		titleSub,
		linebot.NewMessageTemplateAction(NormalMessage, NormalMessage),
		linebot.NewURITemplateAction("link1", "https://www.google.co.jp"),
	)

	messageTitle := "message"
	message1 := linebot.NewTextMessage(messageTitle)
	message2 := linebot.NewTemplateMessage(subscription, template)

	if _, err := bot.ReplyMessage(replyToken, message1, message2).WithContext(ctx).Do(); err != nil {
		aelog.Errorf(ctx, "%v", err)
	}
}

func SendlocationMessage(bot *linebot.Client, replyToken string, ctx context.Context) {
	var locationTitle = "locationTitle"
	var locationSubTitle = "locationSubTitle"

	if _, err := bot.ReplyMessage(replyToken, linebot.NewLocationMessage(locationTitle, locationSubTitle, 35.681298, 139.766247)).WithContext(ctx).Do(); err != nil {
		aelog.Errorf(ctx, "%v", err)
	}
}

func SendNormalMessage(bot *linebot.Client, replyToken string, ctx context.Context) {
	if _, err := bot.ReplyMessage(replyToken, linebot.NewTextMessage("Hi! I am onga bot!!")).WithContext(ctx).Do(); err != nil {
		aelog.Errorf(ctx, "%v", err)
	}

}

func SendImageMessage(bot *linebot.Client, replyToken string, ctx context.Context) {
	if _, err := bot.ReplyMessage(replyToken, linebot.NewImageMessage(imgUrl, imgUrl)).WithContext(ctx).Do(); err != nil {
		aelog.Errorf(ctx, "%v", err)
	}
}

func SendCarouselMessage(bot *linebot.Client, replyToken string, ctx context.Context) {

	template := linebot.NewCarouselTemplate(
		linebot.NewCarouselColumn(
			imgUrl, "A", "text text text text text text text ",
			linebot.NewPostbackTemplateAction("選択", CarouselMessageSelectA, CarouselMessageSelectA),
		),
		linebot.NewCarouselColumn(
			imgUrl, "B", "text text text text text text text ",
			linebot.NewPostbackTemplateAction("選択", CarouselMessageSelectB, CarouselMessageSelectB),
		),
		linebot.NewCarouselColumn(
			imgUrl, "C", "text text text text text text text ",
			linebot.NewPostbackTemplateAction("選択", CarouselMessageSelectC, CarouselMessageSelectC),
		),
		linebot.NewCarouselColumn(
			imgUrl, "D", "text text text text text text text ",
			linebot.NewPostbackTemplateAction("選択", CarouselMessageSelectD, CarouselMessageSelectD),
		),
	)

	message1 := linebot.NewTextMessage("please select")
	message2 := linebot.NewTemplateMessage("itmes", template)

	if _, err := bot.ReplyMessage(replyToken, message1, message2).WithContext(ctx).Do(); err != nil {
		aelog.Errorf(ctx, "%v", err)
	}

}
