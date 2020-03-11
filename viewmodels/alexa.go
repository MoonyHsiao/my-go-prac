package viewmodels

type RankQueryParam struct {
	Top int `form:"top" json:"top"`
}

type CountryQueryParam struct {
	Region string `form:"region" json:"region"`
}
