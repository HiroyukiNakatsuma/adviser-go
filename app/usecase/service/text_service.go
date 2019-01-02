package service

import (
    "strings"

    "github.com/HiroyukiNakatsuma/adviser-go/app/usecase/presenter"
)

type textService struct {
    txtPre presenter.TextPresenter
}

func NewTextService(txtPre presenter.TextPresenter) *textService {
    return &textService{txtPre}
}

func (textService *textService) includeFirstGreeting(inputMes string) (includeFlag bool) {
    var firstGreeting = [5]string{"よろしく", "よろしこ", "宜しく", "初めまして", "はじめまして"}
    for _, g := range firstGreeting {
        i := strings.Index(inputMes, g)
        if i >= 0 {
            return true
        }
    }
    return false
}

func (textService *textService) ReplyContent4PlaneMessage(inputMes string, userName string) (reply string) {
    if textService.includeFirstGreeting(inputMes) {
        return textService.txtPre.BuildFirstGreeting(userName)
    }
    return inputMes
}
