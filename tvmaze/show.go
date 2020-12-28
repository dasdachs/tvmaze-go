package tvmaze

import "time"

type Search struct {
	Shows []*ShowJSON
}

type ShowJSON struct {
	Score           float64         `json:"score"`
	Show            Show            `json:"show"`
	Status          string          `json:"status"`
	Runtime         uint16          `json:"runtime"`
	Premiered       time.Time       `json:"premiered"`
	OfficialSite    string          `json:"officialSite"`
	Schedule        Schedule        `json:"schedule"`
	Rating          Rating          `json:"rating"`
	Weight          uint16          `json:"weight"`
	WebChannel      string          `json:"webChannel"`
	ExternalRateing ExternalRateing `json:"externals"`
	Image           Image           `json:"image"`
	Summary         string          `json:"summary"`
	Updated         uint16          `json:"updated"`
	Links           Links           `json:"_links"`
	Previousepisode Previousepisode `json:"previousepisode"`
}

type Show struct {
	Id        uint16   `json:"id"`
	Url       string   `json:"url"`
	Name      string   `json:"name"`
	Show_type string   `json:"type"`
	Language  string   `json:"language"`
	Genres    []string `json:"genres"`
}

type Schedule struct {
	Time time.Time
	Days []string
}

type Rating struct {
	Average float32
}

type Network struct {
	Id      uint16  `json:"id"`
	Name    string  `json:"name"`
	Country Country `json:"country"`
}

type Country struct {
	Name     string `json:"name"`
	Code     string `json:"code"`
	Timezone string `json:"timezone"`
}

type ExternalRateing struct {
	Tvrage  uint32 `json:"tvrage"`
	Thetvdb uint32 `json:"thetvdb"`
	Imdb    string `json:"imdb"`
}

type Image struct {
	Medium   string `json:"medium"`
	Original string `json:"original"`
}

type Links struct {
	Self SelfLink `json:"self"`
}

type Previousepisode struct {
	Href string `json:"href"`
}

type SelfLink struct {
	Href string `json:"href"`
}
