package wscutils

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
	//DefaultPort = ":8080" // this should come from config

	// Keycloak details
	CLIENT_ID     = "MCX"
	CLIENT_SECRET = "OFyNKbP4g6sYtR4nACtS4V30ILsruzY1"
	KEYCLOAK_URL  = "http://0.0.0.0:8080/"
)
