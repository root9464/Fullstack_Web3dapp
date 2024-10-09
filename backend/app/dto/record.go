package dto

type Record struct {
	Data []Item `json:"data"`
}

type Item struct {
	Total   float64 `json:"total,omitempty"`
	Percent float64 `json:"percent,omitempty"`
}
