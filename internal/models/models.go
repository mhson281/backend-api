package models

type CalculationRequest struct {
	Operation string  `json:"operation"`
	Operand1  float64 `json:"operand1"`
	Operand2  float64 `json:"operand2"`
}

type CalculationResponse struct {
	Result    float64 `json:"result"`
	Operation string  `json:"operation"`
	Operand1  float64 `json:"operand1"`
	Operand2  float64 `json:"operand2"`
	Error     string  `json:"error,omitempty"`
}
