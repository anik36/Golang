package token

import (
	"go-framework/internal/wscutils"
	"net/http"
	"strings"

	"github.com/Nerzal/gocloak/v13"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RealmAdminTokenVo struct {
	Realm    string `json:"realm" validate:"required"`
	UserName string `json:"username" validate:"required,min=4"`
	Password string `json:"password" validate:"required,min=4"`
}

var (
	client = gocloak.NewClient(wscutils.KEYCLOAK_URL)
)

// RegisterTokenHandlers registers all the token-related handlers
func RegisterTokenHandlers(router *gin.Engine) { // we create a function to register all the token-related handlers
	router.POST("/login-admin", getRealmAdminToken)
}

// getTokenForRealmAdmin handles the POST /login-admin request
func getRealmAdminToken(c *gin.Context) {
	var requestVo RealmAdminTokenVo

	// step 1: bind request body to struct
	err := wscutils.BindJson(c, &requestVo)
	if err != nil {
		c.JSON(http.StatusOK, wscutils.NewResponse(wscutils.ErrorStatus, nil, []wscutils.ErrorMessage{wscutils.BuildErrorMessage("invalid_req_body", "invalid_req_body")}))
		return
	}

	// step 2: validate request body
	validationErrors := validate(requestVo)

	// step 3: if there are validation errors, add them to response and send it
	if len(validationErrors) > 0 {
		c.JSON(http.StatusOK, wscutils.NewResponse(wscutils.SuccessStatus, nil, validationErrors))
		return
	}

	// step 4: process the request
	token, err := client.LoginAdmin(c, requestVo.UserName, requestVo.Password, requestVo.Realm)
	if err != nil {
		errAry := strings.Split(err.Error(), " ")
		ercode := errAry[0]
		ermsg := strings.Join(errAry[2:], " ")
		c.JSON(http.StatusOK, wscutils.NewResponse(wscutils.ErrorStatus, nil, []wscutils.ErrorMessage{wscutils.BuildErrorMessage(ermsg, ercode)}))
		return
	}

	// step 5: if there are no errors, send success response
	c.JSON(http.StatusOK, wscutils.NewResponse(wscutils.SuccessStatus, token, []wscutils.ErrorMessage{}))
}

// validate validates the request body
func validate(requestVo RealmAdminTokenVo) []wscutils.ErrorMessage {
	// step 2.1: validate request body using standard validator
	validationErrors := wscutils.WscValidate(requestVo, requestVo.getValsForTokenError)
	// var validationErrors []wscutils.ErrorMessage
	// step 2.2: add request-specific vals to validation errors
	// NOTE: it mutates validationErrors
	validationErrors = addVals(validationErrors, requestVo)

	// if there are standard validation errors, return
	// do not execute custom validations
	if len(validationErrors) > 0 {
		return validationErrors
	}

	return validationErrors
}

// addVals adds request-specific values to a slice of ErrorMessage returned by standard validator
// This is required because vals for different requests could be different.
func addVals(validationErrors []wscutils.ErrorMessage, realmAdminToken RealmAdminTokenVo) []wscutils.ErrorMessage {
	var nilSlice []string
	for i, err := range validationErrors {
		switch err.Field {
		case PASSWORD:
			inputValue := nilSlice
			if len(err.Vals) > 0 {
				inputValue = err.Vals
			}
			validationErrors[i].Vals = inputValue
		case USER_NAME:
			inputValue := nilSlice
			if len(err.Vals) > 0 {
				inputValue = err.Vals
			}
			validationErrors[i].Vals = inputValue
		case REALM:
			inputValue := nilSlice
			if len(err.Vals) > 0 {
				inputValue = err.Vals
			}
			validationErrors[i].Vals = inputValue
		}
	}
	return validationErrors
}

// getValsForUserError returns a slice of strings to be used as vals for a validation error.
// The vals are determined based on the field and the validation rule that failed.
func (requestVo *RealmAdminTokenVo) getValsForTokenError(err validator.FieldError) []string {
	var vals []string
	switch err.Field() {
	case "Password":
		switch err.Tag() {
		case "min":
			vals = append(vals, "4")                  // Minimum valid lenght is 8
			vals = append(vals, *&requestVo.Password) // provided value that failed validation
			// case "max":
			// 	vals = append(vals, "150")                  // Maximum valid age is 150
			// 	vals = append(vals, strconv.Itoa(user.Age)) // provided value that failed validation
		}
		// Add more cases as needed
	}
	return vals
}
