package handlers


import (
	"encoding/json"
	"net/http"
	
	"github.com/mhson281/backend-api/internal/models"
)

func HandleCalculation(w http.ResponseWriter, r *http.Request) {
	var req models.CalculationRequest

	// Decode the JSON request body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Perform calculation based on request
	var res models.CalculationResponse

	switch req.Operation {
	case "add":
	  res.Result = req.Operand1 + req.Operand2
	case "subtract":
	  res.Result = req.Operand1 - req.Operand2
	case "multiply":
	  res.Result = req.Operand1 * req.Operand2
	case "divide":
	  if req.Operand2 == 0 {
			res.Error = "Unable to divide by zero"
			w.WriteHeader(http.StatusBadRequest)
		} else {
			res.Result = req.Operand1 / req.Operand2
		}
	default:
	  http.Error(w, "Unsupported operation", http.StatusBadRequest)
	  return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
