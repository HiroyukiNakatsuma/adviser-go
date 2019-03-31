package presenters

import (
    "github.com/HiroyukiNakatsuma/adviser-go/app/domain/model"

    "github.com/line/line-bot-sdk-go/linebot"
)

const noContentMessage = "ごめんなさい。該当するレストランがありませんでした。。"
const noImageUrl = "https://adviser-go.herokuapp.com/public/images/noImage.jpg"
const gnaviCreditText = "Supported by ぐるなびWebService : https://api.gnavi.co.jp/api/scope/"
const altText = "レストラン情報を送信しました"
const detailLabel = "詳細を見る"
const imageComponentType = "image"
const boxComponentType = "box"
const textComponentType = "text"
const buttonComponentType = "button"
const bubbleContainerType = "bubble"
const carouselContainerType = "carousel"

type RestaurantPresenter struct{}

func NewRestaurantPresenter() *RestaurantPresenter {
    return &RestaurantPresenter{}
}

func (restaurantPresenter *RestaurantPresenter) BuildReplyContent(rests []*model.Restaurant) (reply linebot.SendingMessage) {
    if len(rests) == 0 {
        return linebot.NewTextMessage(noContentMessage)
    }

    var contents []*linebot.BubbleContainer
    for _, rest := range rests {
        hero := newHeroBlock(rest)
        body := newBodyBlock(rest)
        footer := newFooterBlock(rest)
        contents = append(contents, newBubbleContainer(hero, body, footer))
    }

    return linebot.NewFlexMessage(altText, newCarouselContainer(contents))
}

func newHeroBlock(rest *model.Restaurant) *linebot.ImageComponent {
    return &linebot.ImageComponent{
        Type:        imageComponentType,
        URL:         imageUrl(rest.ImageUrls),
        Size:        "full",
        AspectRatio: "4:3",
        AspectMode:  "cover",
        Action:      linebot.NewURIAction("image", rest.Url)}
}

func newBodyBlock(rest *model.Restaurant) *linebot.BoxComponent {
    return &linebot.BoxComponent{
        Type:    boxComponentType,
        Layout:  "vertical",
        Spacing: "sm",
        Contents: []linebot.FlexComponent{
            &linebot.TextComponent{
                Type:   textComponentType,
                Text:   rest.Name,
                Weight: "bold",
                Size:   "lg"},
            &linebot.BoxComponent{
                Type:    boxComponentType,
                Layout:  "vertical",
                Spacing: "sm",
                Contents: []linebot.FlexComponent{
                    newDefinition("Time", rest.OpenTime),
                    newDefinition("Credit", gnaviCreditText)}}}}
}

func newDefinition(title string, desc string) *linebot.BoxComponent {
    titleFlex := int(1)
    descriptionFlex := int(5)
    return &linebot.BoxComponent{
        Type:    boxComponentType,
        Layout:  "baseline",
        Spacing: "sm",
        Contents: []linebot.FlexComponent{
            &linebot.TextComponent{
                Type:  textComponentType,
                Text:  title,
                Color: "#aaaaaa",
                Flex:  &titleFlex,
                Size:  "xs"},
            &linebot.TextComponent{
                Type:  textComponentType,
                Text:  desc,
                Wrap:  true,
                Color: "#666666",
                Flex:  &descriptionFlex,
                Size:  "xs"}}}

}

func newFooterBlock(rest *model.Restaurant) *linebot.BoxComponent {
    return &linebot.BoxComponent{
        Type:    boxComponentType,
        Layout:  "vertical",
        Spacing: "sm",
        Contents: []linebot.FlexComponent{
            &linebot.ButtonComponent{
                Type:   buttonComponentType,
                Style:  "link",
                Height: "sm",
                Action: linebot.NewURIAction(detailLabel, rest.Url)}}}
}

func newBubbleContainer(hero *linebot.ImageComponent, body *linebot.BoxComponent, footer *linebot.BoxComponent) *linebot.BubbleContainer {
    return &linebot.BubbleContainer{
        Type:   bubbleContainerType,
        Hero:   hero,
        Body:   body,
        Footer: footer}
}

func newCarouselContainer(contents []*linebot.BubbleContainer) *linebot.CarouselContainer {
    return &linebot.CarouselContainer{
        Type:     carouselContainerType,
        Contents: contents}
}

func imageUrl(urls []string) string {
    for _, url := range urls {
        if url != "" {
            return url
        }
    }
    return noImageUrl
}
