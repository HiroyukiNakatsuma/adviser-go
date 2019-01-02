package presenter

import (
    "math/rand"
    "time"

    "github.com/HiroyukiNakatsuma/adviser-go/app/usecase/presenter"
)

const firstGreetingMessage = "さん、よろしくお願いします"

type textPresenter struct{}

func NewTextPresenter() presenter.TextPresenter {
    return &textPresenter{}
}

func (textPresenter *textPresenter) BuildFirstGreeting(userName string) string {
    var emojiList = [8]string{"(^^)", "(^_^)", "(^-^)", "(*^^*)", "(^ ^)", "(^.^)", "(≧▽≦)", "！！"}
    rand.Seed(time.Now().UnixNano())
    return userName + firstGreetingMessage + emojiList[rand.Intn(8)]
}
