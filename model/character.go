package model

type Character struct {
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	Height    string `json:"height"`
	Homeworld string `json:"home_world"`
	Mass      string `json:"mass"`
	Created   string `json:"created"`
	Url       string `json:"url"`
}
