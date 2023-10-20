package models

// Request represents the standard structure of a request to the web service.
type Request struct {
	Ver       int    `json:"ver" binding:"required"`
	AuthToken string `json:"authtoken"`
	Data      any    `json:"data" binding:"required"`
}

// Response represents the standard structure of a response of the web service.
type Response struct {
	Status   string         `json:"status"`
	Data     any            `json:"data"`
	Messages []ErrorMessage `json:"messages"`
}

// ErrorMessage defines the format of error part of the standard response object
// See: https://redmine.bquanta.xyz/projects/mail-doc/wiki/Websvcgeneral#Web-service-response-format
type ErrorMessage struct {
	Field   string   `json:"field"`
	Code    string   `json:"code"`
	Msgcode int      `json:"msgcode"`
	Vals    []string `json:"vals,omitempty"` // omit if Vals is empty
}
