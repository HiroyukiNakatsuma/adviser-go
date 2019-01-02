package service

import (
    "strings"

    "github.com/HiroyukiNakatsuma/adviser-go/app/usecase/presenter"
)

type TextService struct {
    txtPre presenter.TextPresenter
}

func NewTextService(txtPre presenter.TextPresenter) *TextService {
    return &TextService{txtPre}
}

func (textService *TextService) includeFirstGreeting(inputMes string) (includeFlag bool) {
    var firstGreeting = [5]string{"よろしく", "よろしこ", "宜しく", "初めまして", "はじめまして"}
    for _, g := range firstGreeting {
        i := strings.Index(inputMes, g)
        if i >= 0 {
            return true
        }
    }
    return false
}

func (textService *TextService) ReplyContent4PlaneMessage(inputMes string, userName string) (reply string) {
    if textService.includeFirstGreeting(inputMes) {
        return textService.txtPre.BuildFirstGreeting(userName)
    }
    return inputMes
}
