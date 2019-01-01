package service

import (
    "strings"
    "math/rand"
    "time"
    "net/http"
    "encoding/json"
    "log"
    "fmt"
    "os"
)

const firstGreetingMessage = "さん、よろしくお願いします"
const noContentMessage = "ごめんなさい。該当するコンテンツがありませんでした。。"
const gnabiEndpoint = "https://api.gnavi.co.jp/RestSearchAPI/v3/"
const gnaviHitPerPage = 5
const gnaviCreditText = "Supported by ぐるなびWebService : https://api.gnavi.co.jp/api/scope/"

type GnaviRestSearchResponse struct {
    Errors        []GnaviError `json:"error"`
    TotalHitCount int          `json:"total_hit_count"`
    HitPerPage    int          `json:"hit_per_page"`
    PageOffset    int          `json:"page_offset"`
    Restaurants   []GnaviRest  `json:"rest"`
}

type GnaviError struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}

type GnaviRest struct {
    Id         string `json:"id"`
    UpdateDate string `json:"update_date"`
    Name       string `json:"name"`
    NameKana   string `json:"name_kana"`
    Latitude   string `json:"latitude"`
    Longitude  string `json:"longitude"`
    Category   string `json:"category"`
    Url        string `json:"url"`
}

func b2i(b bool) int8 {
    i := int8(0)
    if b {
        i = 1
    }
    return i
}

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

func getRestaurants(latitude float64, longitude float64, isNoSmoking bool) (reply string) {
    var client = &http.Client{Timeout: 10 * time.Second}
    url := gnabiEndpoint + fmt.Sprintf("?keyid=%s&latitude=%f&longitude=%f&no_smoking=%d&hit_per_page=%d", os.Getenv("GNAVI_ACCESS_KEY"), latitude, longitude, b2i(isNoSmoking), gnaviHitPerPage)
    log.Printf("Start GET %s", gnabiEndpoint)
    log.Printf("Params latitude=%f, longitude=%f, no_smoking=%d", latitude, longitude, b2i(isNoSmoking))
    res, err := client.Get(url)
    if err != nil {
        log.Fatal(err)
    }
    defer res.Body.Close()

    gnavi := new(GnaviRestSearchResponse)
    err = json.NewDecoder(res.Body).Decode(gnavi)
    if err != nil {
        log.Fatal(err)
    }

    for _, rest := range gnavi.Restaurants {
        reply += fmt.Sprintf("%s\n", rest.Url)
    }

    if len(gnavi.Restaurants) == 0 {
        reply = fmt.Sprintf("%s\n", noContentMessage)
    }

    return reply + gnaviCreditText
}

func getNoSmokingRestaurants(latitude float64, longitude float64) (reply string) {
    return getRestaurants(latitude, longitude, true)
}

func ReplyContent4Location(latitude float64, longitude float64) (reply string) {
    return getNoSmokingRestaurants(latitude, longitude)
}
