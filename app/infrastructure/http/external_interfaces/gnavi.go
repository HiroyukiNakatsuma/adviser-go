package external_interfaces

import (
    "net/http"
    "time"
    "fmt"
    "os"
    "log"
    "encoding/json"

    "github.com/HiroyukiNakatsuma/adviser-go/app/domain/model"
    "github.com/HiroyukiNakatsuma/adviser-go/app/usecase/external_interface"
)

const gnaviEndpoint = "https://api.gnavi.co.jp/RestSearchAPI/v3/"

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
    Id         string    `json:"id"`
    UpdateDate string    `json:"update_date"`
    Name       string    `json:"name"`
    NameKana   string    `json:"name_kana"`
    Latitude   string    `json:"latitude"`
    Longitude  string    `json:"longitude"`
    Category   string    `json:"category"`
    Url        string    `json:"url"`
    GnaviCodes GnaviCode `json:"code"`
}

type GnaviCode struct {
    CategoryCodeSmalls []string `json:"category_code_s"`
}

type gnavi struct{}

func NewGnavi() external_interface.RestaurantExternalInterface {
    return &gnavi{}
}

func (gnavi *gnavi) GetRestaurants(latitude float64, longitude float64, isLunch bool, isNoSmoking bool, amount int) (restaurants []*model.Restaurant) {
    var client = &http.Client{Timeout: 10 * time.Second}
    url := gnaviEndpoint + fmt.Sprintf("?keyid=%s&latitude=%f&longitude=%f&lunch=%d&no_smoking=%d&hit_per_page=%d", os.Getenv("GNAVI_ACCESS_KEY"), latitude, longitude, b2i(isLunch), b2i(isNoSmoking), amount)
    log.Printf("Start GET %s", gnaviEndpoint)
    log.Printf("Params latitude=%f, longitude=%f, is_lunch=%d, no_smoking=%d, amount=%d", latitude, longitude, b2i(isLunch), b2i(isNoSmoking), amount)
    res, err := client.Get(url)
    if err != nil {
        log.Fatal(err)
    }
    defer res.Body.Close()

    gnaviRestSearchResponse := new(GnaviRestSearchResponse)
    err = json.NewDecoder(res.Body).Decode(gnaviRestSearchResponse)
    if err != nil {
        log.Fatal(err)
    }

    for _, rest := range gnaviRestSearchResponse.Restaurants {
        if isLunch && gnavi.isNotLunchCategory(rest.GnaviCodes.CategoryCodeSmalls) {
            continue
        }

        restaurant := model.Restaurant{rest.Id, rest.Name, rest.NameKana, rest.Latitude, rest.Longitude, rest.Category, rest.UpdateDate, rest.UpdateDate}
        restaurants = append(restaurants, &restaurant)
    }

    return
}

func (gnavi *gnavi) isNotLunchCategory(gnaviCategoryCodeSmalls []string) bool {
    notLunchCategorySmalls := []string{"RSFST01013", "RSFST10012", "RSFST10013", "RSFST19001", "RSFST19004", "RSFST19005", "RSFST19006", "RSFST19007", "RSFST19008", "RSFST19009", "RSFST19010", "RSFST19011", "RSFST20001", "RSFST20002", "RSFST20003"}
    for _, notLunchCategorySmall := range notLunchCategorySmalls {
        if contains(gnaviCategoryCodeSmalls, notLunchCategorySmall) {
            return true
        }
    }
    return false
}

func contains(s []string, e string) bool {
    for _, v := range s {
        if e == v {
            return true
        }
    }
    return false
}

func b2i(b bool) int8 {
    i := int8(0)
    if b {
        i = 1
    }
    return i
}
