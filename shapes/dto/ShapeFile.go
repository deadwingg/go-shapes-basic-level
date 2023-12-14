package dto

type ShapeFile struct {
	Type string  `json:"type"`
	A    float64 `json:"a"`
	B    float64 `json:"b"`
	Id   string  `json:"id"`
}
