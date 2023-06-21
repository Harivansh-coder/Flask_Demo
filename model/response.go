package model

type Response struct {
	Name string `json:"name"`
	Href string `json:"href"`
	Aim  string `json:"aim"`
	Pdf  string `json:"pdf"`
	Txt  string `json:"txt"`
}
