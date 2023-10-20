package utils

import (
	"errors"
	"example/goframework/models"
	"example/goframework/utils/constants"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// errorTypes is a map from error type to its corresponding attributes.
var errorTypes map[string]models.ErrorMessage

func init() {
	errorTypes = make(map[string]models.ErrorMessage)
	// loadErrorTypes()
}

// loadErrorTypes is used for loading predefined error types
// func loadErrorTypes() {
// 	// todo: this will come from config
// 	errorTypesPath := os.Getenv("ERROR_TYPES_FILE_PATH")
// 	errorTypesFile, err := os.Open(errorTypesPath)
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	defer func(yamlFile *os.File) {
// 		err := yamlFile.Close()
// 		if err != nil {
// 			log.Printf("Error closing file: %v", err)
// 		}
// 	}(errorTypesFile)

// 	byteValue, _ := io.ReadAll(errorTypesFile)

// 	err = yaml.Unmarshal(byteValue, &errorTypes)
// 	if err != nil {
// 		log.Panic(err)
// 	}
// }

// WscValidate is a generic function that accepts any data structure,
// validates it according to struct tag-provided validation rules
// and returns a slice of ErrorMessage in case of validation errors.
// This function will not add `vals` that's required as per the specifications
// because it does not know the request-specific values.
// `vals` will be added to ErrorMessage by the caller.
func WscValidate[T any](data T) []models.ErrorMessage {
	var validationErrors []models.ErrorMessage

	validate := validator.New()

	err := validate.Struct(data)

	if err != nil {
		var validationErrs validator.ValidationErrors
		if errors.As(err, &validationErrs) {
			for _, err := range validationErrs {
				// We handle validation error creation for developers.
				vErr := BuildValidationError(err.Field(), err.Tag())
				validationErrors = append(validationErrors, vErr)
			}
		}
	}
	return validationErrors
}

// BuildValidationError generates a ErrorMessage which includes
// the required validation error information such as code, msgcode
// It encapsulates the process of building an error message for consistency.
func BuildValidationError(fieldName, errorType string) models.ErrorMessage {
	errorMessage, exists := errorTypes[errorType]

	// if the provided errorType doesn't exist, use attributes of "unknown" error type
	// and set the code (string) to the provided errorType for debugging
	if !exists {
		errorMessage = errorTypes[constants.Unknown]
		errorMessage.Code = errorType
	}
	errorMessage.Field = fieldName
	return errorMessage
}

// NewResponse is a helper function to create a new web service response
// and any error messages that might need to be sent back to the client. It allows
// for a consistent structure in all API responses
func NewResponse(status string, data any, messages []models.ErrorMessage) *models.Response {
	return &models.Response{
		Status:   status,
		Data:     data,
		Messages: messages,
	}
}

// BindJson provides a standard way of binding incoming JSON data to a
// given request data structure. It incorporates error handling .
func BindJson(c *gin.Context, data interface{}) error {
	req := models.Request{Data: data}
	if err := c.ShouldBindJSON(&req); err != nil {
		//invalidJsonError := BuildValidationError("", "invalid_json")
		invalidJsonError := BuildValidationError("", constants.InvalidJSON)
		c.JSON(http.StatusBadRequest, NewResponse(constants.ErrorStatus, nil, []models.ErrorMessage{invalidJsonError}))
		return err
	}
	return nil
}


