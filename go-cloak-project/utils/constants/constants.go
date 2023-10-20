package constants

const (
	//HTTP responses
	ErrorStatus   = "error"
	SuccessStatus = "success"

	// Validation error messages
	RequiredError  = "required"
	InvalidEmail   = "email"
	EmailDomainErr = "emaildomain"

	// Validation Error Types
	InvalidJSON = "invalid_json"
	Unknown     = "unknown"

	// Ports
	//DefaultPort = ":8081" // this should come from config

	// User structure fields
	FirstName     = "FirstName"
	UserName      = "UserName"
	UserEmail     = "Email"
	EmailVerified = "EmailVerified"
	Enabled       = "Enabled"
	Password      = "Password"
	Realm         = "Realm"

	// Validation error messages
	AgeRange    = "10-150"
	NotProvided = "Not provided"
	//EmailDomain = "@domain.com" // this should come from config

	// Keycloak details
	CLIENT_ID     = "MCX"
	CLIENT_SECRET = "OFyNKbP4g6sYtR4nACtS4V30ILsruzY1"
	KEYCLOAK_URL  = "http://0.0.0.0:8080/"
	REALM         = "remiges-tech"
)
