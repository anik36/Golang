package users

import (
	"example/goframework/models"
	"example/goframework/utils"
	"example/goframework/utils/constants"
	"net/http"
	"strings"

	"github.com/Nerzal/gocloak/v13"
	"github.com/gin-gonic/gin"
)

var (
	client = gocloak.NewClient(constants.KEYCLOAK_URL)
)

type UserVo struct {
	Realm    string                     `json:"realm" validate:"required"`
	User     gocloak.User               `json:"user,omitempty"`
	Params   gocloak.GetUsersParams     `json:"params,omitempty"`
	Password gocloak.SetPasswordRequest `json:"password,omitempty"`
}

// RegisterUserHandlers registers all the user-related handlers
func RegisterUserHandlers(router *gin.Engine) { // we create a function to register all the user-related handlers
	router.POST("/user", createUser)
	router.GET("/user", getUserByID)
	router.GET("/user/list", fetchAllUsers)
	router.PUT("/user/update", updateUserByID)
	router.DELETE("/user/delete", deleteUserByID)
	router.PUT("/user/reset-password", setPassword)
}

// setPassword handles the PUT /user/reset-password request
func setPassword(c *gin.Context) {
	var request UserVo
	// step 1: bind request body to struct
	err := utils.BindJson(c, &request)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(constants.ErrorStatus, nil, []models.ErrorMessage{utils.BuildValidationError("invalid_req_body", "invalid_req_body")}))
		return
	}

	token := separateTokenFromHeader(c) // separate "Bearer_" word from token

	// step 4: process the request
	err = client.SetPassword(c, token, *request.User.ID, request.Realm, *request.Password.Password, *request.Password.Temporary)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(constants.ErrorStatus, nil, []models.ErrorMessage{utils.BuildValidationError("invalid_credentials"+err.Error(), "invalid_req_body")}))
		return
	}
	// step 5: if there are no errors, send success response
	c.JSON(http.StatusOK, utils.NewResponse(constants.SuccessStatus, "password set successfully", []models.ErrorMessage{}))
}

// separateTokenFromHeader is used to separate token from header takes context as input and return token string
func separateTokenFromHeader(c *gin.Context) string {
	token := c.GetHeader("Authorization")
	return token[7:]
}

// updateUserByID handles the PUT /user/update request
func updateUserByID(c *gin.Context) {
	var request UserVo
	// step 1: bind request body to struct
	err := utils.BindJson(c, &request)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(constants.ErrorStatus, nil, []models.ErrorMessage{utils.BuildValidationError("invalid_req_body", "invalid_req_body")}))
		return
	}

	token := separateTokenFromHeader(c) // separate "Bearer_" word from token
	// step 4: process the request
	err = client.UpdateUser(c, token, request.Realm, request.User)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(constants.ErrorStatus, nil, []models.ErrorMessage{utils.BuildValidationError("invalid_credentials", "invalid_req_body")}))
		return
	}

	// step 5: if there are no errors, send success response
	c.JSON(http.StatusOK, utils.NewResponse(constants.SuccessStatus, request.User, []models.ErrorMessage{}))
}

// getUserByID handles the GET /user request
func getUserByID(c *gin.Context) {
	var request UserVo
	// step 1: bind request body to struct
	err := utils.BindJson(c, &request)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(constants.ErrorStatus, nil, []models.ErrorMessage{utils.BuildValidationError("invalid_req_body", "invalid_req_body")}))
		return
	}

	token := separateTokenFromHeader(c) // separate "Bearer_" word from token
	// step 4: process the request
	user, err := client.GetUserByID(c, token, request.Realm, *request.User.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(constants.ErrorStatus, nil, []models.ErrorMessage{utils.BuildValidationError("invalid_credentials", "invalid_req_body")}))
		return
	}

	// step 5: if there are no errors, send success response
	c.JSON(http.StatusOK, utils.NewResponse(constants.SuccessStatus, user, []models.ErrorMessage{}))
}

// deleteUserByID handles the DELETE /user/delete request
func deleteUserByID(c *gin.Context) {
	var request UserVo
	// step 1: bind request body to struct
	err := utils.BindJson(c, &request)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(constants.ErrorStatus, nil, []models.ErrorMessage{utils.BuildValidationError("invalid_req_body", "invalid_req_body")}))
		return
	}

	token := separateTokenFromHeader(c) // separate "Bearer_" word from token
	// step 4: process the request
	err = client.DeleteUser(c, token, request.Realm, *request.User.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(constants.ErrorStatus, nil, []models.ErrorMessage{utils.BuildValidationError("invalid_credentials", "invalid_req_body")}))
		return
	}

	// step 5: if there are no errors, send success response
	c.JSON(http.StatusOK, utils.NewResponse(constants.SuccessStatus, "user_delete_successful", []models.ErrorMessage{}))
}

// fetchAllUsers handles the GET /user/list request
func fetchAllUsers(c *gin.Context) {
	// var user gocloak.User
	var request UserVo
	// step 1: bind request body to struct
	err := utils.BindJson(c, &request)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(constants.ErrorStatus, nil, []models.ErrorMessage{utils.BuildValidationError("invalid_req_body", "invalid_req_body")}))
		return
	}

	token := separateTokenFromHeader(c) // separate "Bearer_" word from token

	// step 4: process the request
	params := request.Params
	userID, err := client.GetUsers(c, token, request.Realm, params)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(constants.ErrorStatus, nil, []models.ErrorMessage{utils.BuildValidationError("invalid_credentials", "invalid_req_body")}))
		return
	}

	// step 5: if there are no errors, send success response
	c.JSON(http.StatusOK, utils.NewResponse(constants.SuccessStatus, userID, []models.ErrorMessage{}))
}

// createUser handles the POST /user request
func createUser(c *gin.Context) {
	// var user gocloak.User
	var request UserVo
	// step 1: bind request body to struct
	err := utils.BindJson(c, &request)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(constants.ErrorStatus, nil, []models.ErrorMessage{utils.BuildValidationError("invalid_req_body", "invalid_req_body")}))
		return
	}
	// step 2: validate request body
	validationErrors := validate(request.User)
	// step 3: if there are validation errors, add them to response and send it
	if len(validationErrors) > 0 {
		c.JSON(http.StatusBadRequest, utils.NewResponse(constants.ErrorStatus, nil, validationErrors))
		return
	}
	token := separateTokenFromHeader(c) // separate "Bearer_" word from token

	// step 4: process the request
	userID, err := client.CreateUser(c, token, request.Realm, request.User)
	if err != nil {
		//"user_exists_with_same_credentials"+
		c.JSON(http.StatusConflict, utils.NewResponse(constants.ErrorStatus, nil, []models.ErrorMessage{utils.BuildValidationError(err.Error(), "invalid_req_body")}))
		return
	}
	// save userId to user object
	request.User.ID = &userID

	// step 5: if there are no errors, send success response
	c.JSON(http.StatusOK, utils.NewResponse(constants.SuccessStatus, request.User, []models.ErrorMessage{}))
}

// validate validates the request body
func validate(user gocloak.User) []models.ErrorMessage {
	// step 2.1: validate request body using standard validator
	validationErrors := utils.WscValidate(user)
	// step 2.2: add request-specific vals to validation errors
	// NOTE: it mutates validationErrors
	validationErrors = addVals(validationErrors, user)

	// if there are standard validation errors, return
	// do not execute custom validations
	if len(validationErrors) > 0 {
		return validationErrors
	}
	// step 2.3: check request specific custom validations and add errors
	validationErrors = addCustomValidationErrors(validationErrors, user)

	return validationErrors
}

// addVals adds request-specific values to a slice of ErrorMessage returned by standard validator
// This is required because vals for different requests could be different.
func addVals(validationErrors []models.ErrorMessage, user gocloak.User) []models.ErrorMessage {
	for i, err := range validationErrors {
		switch err.Field {
		case constants.FirstName:
			inputValue := constants.NotProvided
			validationErrors[i].Vals = []string{inputValue}
		case constants.UserName:
			inputValue := constants.NotProvided
			validationErrors[i].Vals = []string{inputValue}
		case constants.EmailVerified:
			inputValue := constants.NotProvided
			validationErrors[i].Vals = []string{inputValue}
		case constants.Enabled:
			inputValue := constants.NotProvided
			validationErrors[i].Vals = []string{inputValue}
		case constants.UserEmail:
			if err.Code == constants.RequiredError {
				inputValue := constants.NotProvided
				validationErrors[i].Vals = []string{inputValue}
			} else if err.Code == constants.InvalidEmail {
				inputValue := user.Email
				validationErrors[i].Vals = []string{*inputValue}
			}
		}
	}
	return validationErrors
}

// addCustomValidationErrors adds custom validation errors to the validationErrors slice.
// This is required because request specific custom validators are not supported by wscvalidation.
func addCustomValidationErrors(validationErrors []models.ErrorMessage, user gocloak.User) []models.ErrorMessage {
	// Example of a custom validation for email domains
	if user.Email != nil && !strings.Contains(*user.Email, "@domain.com") {
		emailDomainError := utils.BuildValidationError(constants.UserEmail, constants.EmailDomainErr)
		emailDomainError.Vals = []string{*user.Email, "@domain.com"}
		validationErrors = append(validationErrors, emailDomainError)
	}
	return validationErrors
}
