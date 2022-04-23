package model

type Chapter struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Category *Category `json:"category"`
	Course   *Course   `json:"course"`
}
