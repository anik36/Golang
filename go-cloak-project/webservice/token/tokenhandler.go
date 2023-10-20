package token

import (
	"example/goframework/models"
	"example/goframework/utils"
	"example/goframework/utils/constants"
	"net/http"

	"github.com/Nerzal/gocloak/v13"
	"github.com/gin-gonic/gin"
)

type RealmAdminTokenVo struct {
	UserName string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Realm    string `json:"realm" validate:"required"`
}

var (
	client = gocloak.NewClient(constants.KEYCLOAK_URL)
)

// RegisterTokenHandlers registers all the token-related handlers
func RegisterTokenHandlers(router *gin.Engine) { // we create a function to register all the token-related handlers
	router.POST("/login-admin", getRealmAdminToken)
}

// getTokenForRealmAdmin handles the POST /token request
func getRealmAdminToken(c *gin.Context) {
	var realmAdminToken RealmAdminTokenVo

	// step 1: bind request body to struct
	err := utils.BindJson(c, &realmAdminToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(constants.ErrorStatus, nil, []models.ErrorMessage{utils.BuildValidationError(err.Error(), "invalid_req_body")}))
		return
	}

	// step 2: validate request body
	validationErrors := validate(realmAdminToken)

	// step 3: if there are validation errors, add them to response and send it
	if len(validationErrors) > 0 {
		c.JSON(http.StatusBadRequest, utils.NewResponse(constants.ErrorStatus, nil, validationErrors))
		return
	}

	// step 4: process the request
	token, err := client.LoginAdmin(c, realmAdminToken.UserName, realmAdminToken.Password, realmAdminToken.Realm)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(constants.ErrorStatus, nil, []models.ErrorMessage{utils.BuildValidationError("invalid_user_credentials", "invalid_req_body")}))
		return
	}

	// step 5: if there are no errors, send success response
	c.JSON(http.StatusOK, utils.NewResponse(constants.SuccessStatus, token, []models.ErrorMessage{}))
}

// validate validates the request body
func validate(realmAdminToken RealmAdminTokenVo) []models.ErrorMessage {
	// step 2.1: validate request body using standard validator
	validationErrors := utils.WscValidate(realmAdminToken)

	// step 2.2: add request-specific vals to validation errors
	// NOTE: it mutates validationErrors
	validationErrors = addVals(validationErrors, realmAdminToken)

	// if there are standard validation errors, return
	// do not execute custom validations
	if len(validationErrors) > 0 {
		return validationErrors
	}

	return validationErrors
}

// addVals adds request-specific values to a slice of ErrorMessage returned by standard validator
// This is required because vals for different requests could be different.
func addVals(validationErrors []models.ErrorMessage, realmAdminToken RealmAdminTokenVo) []models.ErrorMessage {
	for i, err := range validationErrors {
		switch err.Field {
		case constants.Password:
			inputValue := constants.NotProvided
			validationErrors[i].Vals = []string{inputValue}
		case constants.UserName:
			inputValue := constants.NotProvided
			validationErrors[i].Vals = []string{inputValue}
		case constants.Realm:
			inputValue := constants.NotProvided
			validationErrors[i].Vals = []string{inputValue}
		}
	}

	return validationErrors
}
