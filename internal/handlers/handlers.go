package handlers

import (
	"database/sql"
	"encoding/json"
	"log/slog"
	"math"
	"net/http"
	"os"
	"strings"

	"github.com/mhson281/backend-api/internal/auth"
	"github.com/mhson281/backend-api/internal/database"
	"github.com/mhson281/backend-api/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func HandleCalculation(w http.ResponseWriter, r *http.Request) {
	var req models.CalculationRequest

	// Initialize logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Decode the JSON request body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body.  Ensure the operands are valid numbers and requests are in JSON format", http.StatusBadRequest)
		return
	}

	// Declare valid operations
	validOperations := map[string]bool{
		"add":      true,
		"subtract": true,
		"multiply": true,
		"divide":   true,
	}

	if !validOperations[strings.ToLower(req.Operation)] {
		logger.Warn("Unsupported operation attempted", "operation", req.Operation)
		http.Error(w, "Unsupported operation attempted", http.StatusBadRequest)
	}

	if math.IsNaN(req.Operand1) || math.IsNaN(req.Operand2) {
		http.Error(w, "Operands must be valid numbers", http.StatusBadRequest)
		return
	}

	// Perform calculation based on request
	var res models.CalculationResponse
	res.Operand1 = req.Operand1
	res.Operand2 = req.Operand2
	res.Operation = req.Operation

	switch strings.ToLower(req.Operation) {
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
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	var req  models.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
	}

	// Insert new user into sql db
	_, err = database.DB.Exec("INSERT INTO users(username, password) VALUES (?, ?)", req.Username, string(hashedPassword))
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Username already exists", http.StatusConflict)
		} else {
			http.Error(w, "Failed to register user", http.StatusInternalServerError)
		}
		return
	}

	res := models.RegisterResponse{Message: "User registered successfully"}
	json.NewEncoder(w).Encode(res)
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusInternalServerError)
		return
	}

	var storedPassword string
	row := database.DB.QueryRow("SELECT password from users where username = ?", req.Username)
	if err := row.Scan(&storedPassword); err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Compare hashed password with stored Password
	if err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(req.Password)); err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token, err := auth.GenerateToken(req.Username)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	res := models.LoginResponse{Token: token}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

}

