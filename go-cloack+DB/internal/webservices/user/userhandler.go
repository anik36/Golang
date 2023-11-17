package user

import (
	"fmt"
	"go-framework/internal/wscutils"
	"net/http"
	"strings"

	"github.com/Nerzal/gocloak/v13"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	client = gocloak.NewClient(wscutils.KEYCLOAK_URL)
)

type UserVo struct {
	Realm    string                     `json:"realm" validate:"required,min=4"`
	User     gocloak.User               `json:"user,omitempty"`
	Params   gocloak.GetUsersParams     `json:"params,omitempty"`
	Password gocloak.SetPasswordRequest `json:"password,omitempty"`
}

type User struct {
	Username      *string `json:"username,omitempty" validate:"required"`
	EmailVerified *bool   `json:"emailVerified,omitempty"`
	FirstName     *string `json:"firstName,omitempty" validate:"required"`
	LastName      *string `json:"lastName,omitempty" validate:"required"`
	Email         *string `json:"email,omitempty"`
}

// RegisterUserHandlers registers all the user-related handlers
func RegisterUserHandlers(router *gin.Engine) { // we create a function to register all the user-related handlers
	router.POST("/user", createUser)
	router.GET("/user/:id", getUserByID)
	router.GET("/user/list", fetchAllUsers)
	router.PUT("/user/:id", updateUserByID)
	router.DELETE("/user/:id", deleteUserByID)
	router.PUT("/user/password/:id", setPassword)
}

// setPassword handles the PUT /user/reset-password request
func setPassword(c *gin.Context) {
	var request UserVo
	id := strings.TrimSpace(c.Param("id"))
	if request.User.ID == nil {
		request.User.ID = &id
	}

	// step 1: bind request body to struct
	err := wscutils.BindJson(c, &request)
	if err != nil {
		c.JSON(http.StatusOK, wscutils.NewResponse(wscutils.SuccessStatus, nil, []wscutils.ErrorMessage{wscutils.BuildErrorMessage("invalid_req_body", "invalid_req_body")}))
		return
	}
	// step 2: validate request body
	validationErrors := validate(request)

	// step 3: if there are validation errors, add them to response and send it
	if len(validationErrors) > 0 {
		c.JSON(http.StatusOK, wscutils.NewResponse(wscutils.SuccessStatus, nil, validationErrors))
		return
	}
	token := separateTokenFromHeader(c) // separate "Bearer_" word from token

	// step 4: process the request
	err = client.SetPassword(c, token, id, request.Realm, *request.Password.Password, *request.Password.Temporary)
	if err != nil {
		errAry := strings.Split(err.Error(), " ")
		ercode := errAry[0]
		ermsg := strings.Join(errAry[2:], " ")
		c.JSON(http.StatusOK, wscutils.NewResponse(wscutils.ErrorStatus, nil, []wscutils.ErrorMessage{wscutils.BuildErrorMessage(ermsg, ercode)}))
		return
	}
	// step 5: if there are no errors, send success response
	c.JSON(http.StatusOK, wscutils.NewResponse(wscutils.SuccessStatus, "password set successfully", []wscutils.ErrorMessage{}))
}

// separateTokenFromHeader is used to separate token from header takes context as input and return token string
func separateTokenFromHeader(c *gin.Context) string {
	token := c.GetHeader("Authorization")
	return token[7:]
}

// updateUserByID handles the PUT /user/update request
func updateUserByID(c *gin.Context) {
	var request UserVo
	id := strings.TrimSpace(c.Param("id"))
	if request.User.ID == nil {
		request.User.ID = &id
	}

	// step 1: bind request body to struct
	err := wscutils.BindJson(c, &request)
	if err != nil {
		c.JSON(http.StatusOK, wscutils.NewResponse(wscutils.SuccessStatus, nil, []wscutils.ErrorMessage{wscutils.BuildErrorMessage("invalid_req_body", "invalid_req_body")}))
		return
	}

	// step 2: validate request body
	validationErrors := validate(request)

	// step 3: if there are validation errors, add them to response and send it
	if len(validationErrors) > 0 {
		c.JSON(http.StatusOK, wscutils.NewResponse(wscutils.SuccessStatus, nil, validationErrors))
		return
	}

	token := separateTokenFromHeader(c) // separate "Bearer_" word from token
	// step 4: process the request
	err = client.UpdateUser(c, token, request.Realm, request.User)
	if err != nil {
		errAry := strings.Split(err.Error(), " ")
		ercode := errAry[0]
		ermsg := strings.Join(errAry[2:], " ")
		c.JSON(http.StatusOK, wscutils.NewResponse(wscutils.ErrorStatus, nil, []wscutils.ErrorMessage{wscutils.BuildErrorMessage(ermsg, ercode)}))
		return
	}

	// step 5: if there are no errors, send success response
	c.JSON(http.StatusOK, wscutils.NewResponse(wscutils.SuccessStatus, request.User, []wscutils.ErrorMessage{}))
}

// getUserByID handles the GET /user request
func getUserByID(c *gin.Context) {
	var request UserVo
	id := strings.TrimSpace(c.Param("id"))
	if request.User.ID == nil {
		request.User.ID = &id
	}

	// step 1: bind request body to struct
	err := wscutils.BindJson(c, &request)
	if err != nil {
		c.JSON(http.StatusOK, wscutils.NewResponse(wscutils.SuccessStatus, nil, []wscutils.ErrorMessage{wscutils.BuildErrorMessage("invalid_req_body", "invalid_req_body")}))
		return
	}

	// step 2: validate request body
	validationErrors := validate(request)

	// step 3: if there are validation errors, add them to response and send it
	if len(validationErrors) > 0 {
		c.JSON(http.StatusOK, wscutils.NewResponse(wscutils.SuccessStatus, nil, validationErrors))
		return
	}
	token := separateTokenFromHeader(c) // separate "Bearer_" word from token
	// step 4: process the request
	user, err := client.GetUserByID(c, token, request.Realm, *request.User.ID)
	if err != nil {
		errAry := strings.Split(err.Error(), " ")
		ercode := errAry[0]
		ermsg := strings.Join(errAry[2:], " ")
		c.JSON(http.StatusOK, wscutils.NewResponse(wscutils.ErrorStatus, nil, []wscutils.ErrorMessage{wscutils.BuildErrorMessage(ermsg, ercode)}))
		return
	}

	// step 5: if there are no errors, send success response
	c.JSON(http.StatusOK, wscutils.NewResponse(wscutils.SuccessStatus, user, []wscutils.ErrorMessage{}))
}

// deleteUserByID handles the DELETE /user/delete request
func deleteUserByID(c *gin.Context) {
	var request UserVo
	id := strings.TrimSpace(c.Param("id"))
	if request.User.ID == nil {
		request.User.ID = &id
	}

	// step 1: bind request body to struct
	err := wscutils.BindJson(c, &request)
	if err != nil {
		field := strings.Split(err.Error(), ":")
		c.JSON(http.StatusOK, wscutils.NewResponse(wscutils.SuccessStatus, nil, []wscutils.ErrorMessage{wscutils.BuildErrorMessage(strings.Trim(field[1], " "), "invalid_req_body")}))
		return
	}

	// step 2: validate request body
	validationErrors := validate(request)

	// step 3: if there are validation errors, add them to response and send it
	if len(validationErrors) > 0 {
		c.JSON(http.StatusOK, wscutils.NewResponse(wscutils.ErrorStatus, nil, validationErrors))
		return
	}
	token := separateTokenFromHeader(c) // separate "Bearer_" word from token
	// step 4: process the request
	err = client.DeleteUser(c, token, request.Realm, *request.User.ID)
	if err != nil {
		errAry := strings.Split(err.Error(), " ")
		ercode := errAry[0]
		ermsg := strings.Join(errAry[2:], " ")
		c.JSON(http.StatusOK, wscutils.NewResponse(wscutils.ErrorStatus, nil, []wscutils.ErrorMessage{wscutils.BuildErrorMessage(ermsg, ercode)}))
		return
	}

	// step 5: if there are no errors, send success response
	c.JSON(http.StatusOK, wscutils.NewResponse(wscutils.SuccessStatus, "user_delete_successful", []wscutils.ErrorMessage{}))
}

// fetchAllUsers handles the GET /user/list request
func fetchAllUsers(c *gin.Context) {
	var request UserVo

	// step 1: bind request body to struct
	err := wscutils.BindJson(c, &request)
	if err != nil {
		field := strings.Split(err.Error(), ":")
		c.JSON(http.StatusOK, wscutils.NewResponse(wscutils.ErrorStatus, nil, []wscutils.ErrorMessage{wscutils.BuildErrorMessage(strings.Trim(field[1], " "), "invalid_req_body")}))
		return
	}

	// step 2: validate request body
	validationErrors := validate(request)

	// step 3: if there are validation errors, add them to response and send it
	if len(validationErrors) > 0 {
		c.JSON(http.StatusOK, wscutils.NewResponse(wscutils.ErrorStatus, nil, validationErrors))
		return
	}
	fmt.Println("fetchAllUsers 4::")

	token := separateTokenFromHeader(c) // separate "Bearer_" word from token

	// step 4: process the request
	params := request.Params
	userID, err := client.GetUsers(c, token, request.Realm, params)
	fmt.Println("fetchAllUsers 4.5::")

	if err != nil {
		errAry := strings.Split(err.Error(), " ")
		ercode := errAry[0]
		ermsg := strings.Join(errAry[2:], " ")
		c.JSON(http.StatusOK, wscutils.NewResponse(wscutils.ErrorStatus, nil, []wscutils.ErrorMessage{wscutils.BuildErrorMessage(ermsg, ercode)}))
		return
	}
	fmt.Println("fetchAllUsers 5::")

	// step 5: if there are no errors, send success response
	c.JSON(http.StatusOK, wscutils.NewResponse(wscutils.SuccessStatus, userID, []wscutils.ErrorMessage{}))
}

// createUser handles the POST /user request
func createUser(c *gin.Context) {

	var request UserVo

	// step 1: bind request body to struct
	err := wscutils.BindJson(c, &request)
	if err != nil {
		c.JSON(http.StatusOK, wscutils.NewResponse(wscutils.ErrorStatus, nil, []wscutils.ErrorMessage{wscutils.BuildErrorMessage("invalid_req_body", "invalid_req_body")}))
		return
	}
	// step 2: validate request body
	validationErrors := validate(request)

	// step 3: if there are validation errors, add them to response and send it
	if len(validationErrors) > 0 {
		c.JSON(http.StatusOK, wscutils.NewResponse(wscutils.ErrorStatus, nil, validationErrors))
		return
	}
	token := separateTokenFromHeader(c) // separate "Bearer_" word from token

	// step 4: process the request
	userID, err := client.CreateUser(c, token, request.Realm, request.User)
	if err != nil {
		errAry := strings.Split(err.Error(), " ")
		ercode := errAry[0]
		ermsg := strings.Join(errAry[2:], " ")
		c.JSON(http.StatusOK, wscutils.NewResponse(wscutils.ErrorStatus, nil, []wscutils.ErrorMessage{wscutils.BuildErrorMessage(ermsg, ercode)}))
		return
	}
	// save userId to user object
	request.User.ID = &userID

	// step 5: if there are no errors, send success response
	c.JSON(http.StatusOK, wscutils.NewResponse(wscutils.SuccessStatus, request.User, []wscutils.ErrorMessage{}))
}

// validate validates the request body
func validate(userVo UserVo) []wscutils.ErrorMessage {
	// step 2.1: validate request body using standard validator
	validationErrors := wscutils.WscValidate(userVo, userVo.getValsForUserError)

	// step 2.2: add request-specific vals to validation errors
	// NOTE: it mutates validationErrors
	validationErrors = addVals(validationErrors, userVo.User)

	// if there are standard validation errors, return
	// do not execute custom validations
	if len(validationErrors) > 0 {
		return validationErrors
	}
	// step 2.3: check request specific custom validations and add errors
	validationErrors = addCustomValidationErrors(validationErrors, userVo.User)

	return validationErrors
}

// addVals adds request-specific values to a slice of ErrorMessage returned by standard validator
// This is required because vals for different requests could be different.
func addVals(validationErrors []wscutils.ErrorMessage, user gocloak.User) []wscutils.ErrorMessage {
	var nilSlice []string
	for i, err := range validationErrors {
		switch err.Field {
		case FIRST_NAME:
			inputValue := nilSlice
			if len(err.Vals) > 0 {
				inputValue = err.Vals
			}
			validationErrors[i].Vals = inputValue
		case USER_NAME:
			inputValue := NotProvided
			validationErrors[i].Vals = []string{inputValue}
		case EMAIL_VERIFIED:
			inputValue := NotProvided
			validationErrors[i].Vals = []string{inputValue}
		case ENABLED:
			inputValue := NotProvided
			validationErrors[i].Vals = []string{inputValue}
		case USER_EMAIL:
			if err.Code == wscutils.RequiredError {
				inputValue := NotProvided
				validationErrors[i].Vals = []string{inputValue}
			} else if err.Code == wscutils.InvalidEmail {
				inputValue := user.Email
				validationErrors[i].Vals = []string{*inputValue}
			}
		}
	}
	return validationErrors
}

// addCustomValidationErrors adds custom validation errors to the validationErrors slice.
// This is required because request specific custom validators are not supported by wscvalidation.
func addCustomValidationErrors(validationErrors []wscutils.ErrorMessage, user gocloak.User) []wscutils.ErrorMessage {
	// Example of a custom validation for email domains
	if user.Email != nil && !strings.Contains(*user.Email, "@domain.com") {
		emailDomainError := wscutils.BuildErrorMessage(USER_EMAIL, wscutils.EmailDomainErr)
		emailDomainError.Vals = []string{*user.Email, "@domain.com"}
		validationErrors = append(validationErrors, emailDomainError)
	}
	return validationErrors
}

// getValsForUserError returns a slice of strings to be used as vals for a validation error.
// The vals are determined based on the field and the validation rule that failed.
func (user *UserVo) getValsForUserError(err validator.FieldError) []string {
	var vals []string
	switch err.Field() {
	case "Password":
		switch err.Tag() {
		case "min":
			vals = append(vals, "4")                     // Minimum valid lenght is 8
			vals = append(vals, *user.Password.Password) // provided value that failed validation
			// case "max":
			// 	vals = append(vals, "150")                  // Maximum valid age is 150
			// 	vals = append(vals, strconv.Itoa(user.Age)) // provided value that failed validation
		}
		// Add more cases as needed
	}
	return vals
}
