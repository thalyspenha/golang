package model

//ResponseSuccess resposta de sucesso
type ResponseSuccess struct {
	Meta    Meta        `json:"meta"`
	Records interface{} `json:"records"`
}

//Meta resposta
type Meta struct {
	Limit       int `json:"limit"`
	Offset      int `json:"offset"`
	RecordCount int `json:"recordCount"`
}
