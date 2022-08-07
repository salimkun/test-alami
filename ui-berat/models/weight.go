package models

type Weight struct {
	Date  string `json:"date"`
	Max   int32  `json:"max"`
	Min   int32  `json:"min"`
	Range int32  `json:"range"`
}

type Meta struct {
	Page      int32 `json:"page"`
	Limit     int32 `json:"limit"`
	TotalData int32 `json:"total_data"`
	TotalPage int32 `json:"total_page"`
}

type ListWeight struct {
	Meta  Meta      `json:"meta"`
	Data  []*Weight `json:"data"`
	Error string    `json:"error"`
}
