package dto

type Record struct {
	Data []Item `json:"data"`
}

type Item struct {
	Total   []int64 `json:"total,omitempty"`
	Percent int64   `json:"percent,omitempty"`
}

type ItemRes struct {
	Name    string `json:"name"`
	Total   any    `json:"total"`
	Percent int64  `json:"percent"`
}
