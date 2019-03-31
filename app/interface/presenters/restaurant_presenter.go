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
const defaultImageSize = "full"
const defaultImageAspectRatio = "4:3"
const defaultImageAspectMode = "cover"
const boxComponentType = "box"
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
        hero := newHero(restaurantPresenter.imageUrl(rest.ImageUrls), rest.Url)
        body := newBody(rest.Name)
        footer := newFooter(rest.Url)
        contents = append(contents, newBubbleContainer(hero, body, footer))
    }

    return linebot.NewFlexMessage(altText, newCarouselContainer(contents))
}

func newHero(imageUrl string, shopUrl string) *linebot.ImageComponent {
    return &linebot.ImageComponent{
        Type:        imageComponentType,
        URL:         imageUrl,
        Size:        defaultImageSize,
        AspectRatio: defaultImageAspectRatio,
        AspectMode:  defaultImageAspectMode,
        Action:      linebot.NewURIAction("image", shopUrl)}
}

func newBody(name string) *linebot.BoxComponent {
    return &linebot.BoxComponent{
        Type:    boxComponentType,
        Layout:  "vertical",
        Spacing: "sm",
        Contents: []linebot.FlexComponent{
            &linebot.TextComponent{Type: "text", Text: name, Weight: "bold", Size: "lg"},
            &linebot.BoxComponent{
                Type:    boxComponentType,
                Layout:  "vertical",
                Spacing: "sm",
                Contents: []linebot.FlexComponent{
                    &linebot.TextComponent{Type: "text", Text: gnaviCreditText, Wrap: true, Color: "#666666", Size: "xs"}}}}}
}

func newFooter(shopUrl string) *linebot.BoxComponent {
    return &linebot.BoxComponent{
        Type:    boxComponentType,
        Layout:  "vertical",
        Spacing: "sm",
        Contents: []linebot.FlexComponent{
            &linebot.ButtonComponent{
                Type:   "button",
                Style:  "link",
                Height: "sm",
                Action: linebot.NewURIAction(detailLabel, shopUrl)}}}
}

func newBubbleContainer(hero *linebot.ImageComponent, body *linebot.BoxComponent, footer *linebot.BoxComponent) *linebot.BubbleContainer {
    return &linebot.BubbleContainer{Type: bubbleContainerType, Hero: hero, Body: body, Footer: footer}
}

func newCarouselContainer(contents []*linebot.BubbleContainer) *linebot.CarouselContainer {
    return &linebot.CarouselContainer{Type: carouselContainerType, Contents: contents}
}

func (restaurantPresenter *RestaurantPresenter) imageUrl(urls []string) string {
    for _, url := range urls {
        if url != "" {
            return url
        }
    }
    return noImageUrl
}
