package utils

const (
	SERVER_ERR        = "Internal Server Error."
	SERVER_ERR_DETAIL = "An nternal server error occurred."
	SERVER_ERR_CODE   = "500"
)

// User registration errors
const (
	EMPTY_NAME_FIELD        = "Error: Name field empty."
	EMPTY_NAME_FIELD_DETAIL = "Oops, you skipped your name!."
	EMPTY_NAME_FIELD_CODE   = "400"

	EMPTY_EMAIL_FIELD        = "Error: Email field empty."
	EMPTY_EMAIL_FIELD_DETAIL = "Oops, you skipped your email!."
	EMPTY_EMAIL_FIELD_CODE   = "400"

	BAD_EMAIL_FORMAT        = "Bad email format."
	BAD_EMAIL_FORMAT_DETAIL = "The email you entered has a wrong format."
	BAD_EMAIL_FORMAT_CODE   = "400"

	EMPTY_PASSWORD_FIELD        = "Error: Name field empty."
	EMPTY_PASSWORD_FIELD_DETAIL = "Oops, you skipped your name!."
	EMPTY_PASSWORD_FIELD_CODE   = "400"
)

// Template parsing errors
const (
	PAGE_PARSING_ERROR = "Page parse error."
	PAGE_PARSING_CODE  = "500"
)

// Request routes
const (
	LOGIN_ROUTE   = "/auth/login"
	SIGNUP_ROUTE  = "/auth/register"
	WELCOME_ROUTE = "/"
)

// Internal folder path
const (
	STYLES_PATH_PATTERN = "/web/static/styles/"
	STYLES_PATH         = "web/static/styles"
)
