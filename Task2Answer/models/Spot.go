package models

type Spot struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Website     string  `json:"website"`
	Coordinates string  `json:"coordinates"`
	Description string  `json:"description"`
	Rating      float64 `json:"rating"`
}
