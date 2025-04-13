package middleware

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

// ValidationError represents a validation error with field and message
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ValidationResponse represents the response structure for validation errors
type ValidationResponse struct {
	Errors []ValidationError `json:"errors"`
}

// ValidatorMiddleware creates a middleware that validates request bodies
func ValidatorMiddleware(v *validator.Validate) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Skip validation for GET requests
			if r.Method == http.MethodGet {
				next.ServeHTTP(w, r)
				return
			}

			// Get the content type
			contentType := r.Header.Get("Content-Type")
			if !strings.Contains(contentType, "application/json") {
				http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
				return
			}

			// Decode the request body into a map
			var body map[string]interface{}
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				http.Error(w, "Invalid JSON body", http.StatusBadRequest)
				return
			}

			// Validate the request body
			if err := v.Struct(body); err != nil {
				validationErrors := err.(validator.ValidationErrors)
				response := ValidationResponse{
					Errors: make([]ValidationError, len(validationErrors)),
				}

				for i, err := range validationErrors {
					response.Errors[i] = ValidationError{
						Field:   err.Field(),
						Message: getValidationMessage(err),
					}
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(response)
				return
			}

			// If validation passes, continue to the next handler
			next.ServeHTTP(w, r)
		})
	}
}

// getValidationMessage returns a user-friendly validation message
func getValidationMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return "Value must be at least " + err.Param()
	case "max":
		return "Value must be at most " + err.Param()
	case "len":
		return "Length must be exactly " + err.Param()
	case "gte":
		return "Value must be greater than or equal to " + err.Param()
	case "lte":
		return "Value must be less than or equal to " + err.Param()
	default:
		return "Invalid value for " + err.Field()
	}
}

// NewValidator creates a new validator instance with custom validation rules
func NewValidator() *validator.Validate {
	v := validator.New()

	// Add custom validation rules here
	// Example: v.RegisterValidation("custom_rule", customValidationFunction)

	return v
}
