package service

import (
    "strings"
    "math/rand"
    "time"
)

const firstGreetingMessage = "さん、よろしくお願いします"

func includeFirstGreeting(inputMes string) (includeFlag bool) {
    var firstGreeting = [5]string{"よろしく", "よろしこ", "宜しく", "初めまして", "はじめまして"}
    for _, g := range firstGreeting {
        i := strings.Index(inputMes, g)
        if i >= 0 {
            return true
        }
    }
    return false
}

func getFirstGreeting(userName string) (mes string) {
    var emojiList = [8]string{"(^^)", "(^_^)", "(^-^)", "(*^^*)", "(^ ^)", "(^.^)", "(≧▽≦)", "！！"}
    rand.Seed(time.Now().UnixNano())
    return userName + firstGreetingMessage + emojiList[rand.Intn(8)]
}

func ReplyContent4PlaneMessage(inputMes string, userName string) (reply string) {
    if includeFirstGreeting(inputMes) {
        return getFirstGreeting(userName)
    }
    return inputMes
}
