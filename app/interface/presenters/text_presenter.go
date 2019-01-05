package presenters

import (
    "math/rand"
    "time"

    "github.com/HiroyukiNakatsuma/adviser-go/app/usecase/presenter"
)

const firstGreetingMessage = "さん、よろしくお願いします"

var emojiList = [8]string{"(^^)", "(^_^)", "(^-^)", "(*^^*)", "(^ ^)", "(^.^)", "(≧▽≦)", "！！"}

type textPresenter struct{}

func NewTextPresenter() presenter.TextPresenter {
    return &textPresenter{}
}

func (textPresenter *textPresenter) BuildFirstGreeting(userName string) string {
    rand.Seed(time.Now().UnixNano())
    return userName + firstGreetingMessage + emojiList[rand.Intn(8)]
}
