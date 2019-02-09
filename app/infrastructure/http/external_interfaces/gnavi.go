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

const gnabiEndpoint = "https://api.gnavi.co.jp/RestSearchAPI/v3/"
const gnaviHitPerPage = 5

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

type gnavi struct{}

func NewGnavi() external_interface.RestaurantExternalInterface {
    return &gnavi{}
}

func (gnavi *gnavi) GetRestaurants(latitude float64, longitude float64, isLunch bool, isNoSmoking bool) (restaurants []*model.Restaurant) {
    var client = &http.Client{Timeout: 10 * time.Second}
    url := gnabiEndpoint + fmt.Sprintf("?keyid=%s&latitude=%f&longitude=%f&no_smoking=%d&lunch=%d&hit_per_page=%d", os.Getenv("GNAVI_ACCESS_KEY"), latitude, longitude, b2i(isLunch), b2i(isNoSmoking), gnaviHitPerPage)
    log.Printf("Start GET %s", gnabiEndpoint)
    log.Printf("Params latitude=%f, longitude=%f, no_smoking=%d", latitude, longitude, b2i(isNoSmoking))
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
        restaurant := model.Restaurant(rest)
        restaurants = append(restaurants, &restaurant)
    }

    return
}

func b2i(b bool) int8 {
    i := int8(0)
    if b {
        i = 1
    }
    return i
}
