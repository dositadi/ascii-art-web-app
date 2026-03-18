package utils

// Error code strings
const (
	SERVER_ERR        = "Internal Server Error."
	SERVER_ERR_DETAIL = "An nternal server error occurred."
	SERVER_ERR_CODE   = "500"

	NOT_FOUND_ERR    = "Not found"
	NOT_FOUND_DETAIL = "Resource not found."
	NOT_FOUND_CODE   = "404"

	UNAUTHORIZED_ERR        = "Unauthorized Error."
	UNAUTHORIZED_ERR_DETAIL = "Access denied."
	UNAUTHORIZED_ERR_CODE   = "401"

	CONFLICT_ERR        = "Conflict Error."
	CONFLICT_ERR_DETAIL = "Resource already exists"
	CONFLICT_ERR_CODE   = "409"

	CONN_LOST_ERR        = "Connection lost."
	CONN_LOST_ERR_DETAIL = "Connection has been lost."
	CONN_LOST_ERR_CODE   = "500"
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
	ASCII_ROUTE = "/home/ascii"
)

// Internal folder path
const (
	STYLES_PATH_PATTERN = "/web/static/styles/"
	STYLES_PATH         = "web/static/styles"
)

// Repo Query statements
const (
	INSERT_INTO_USERS     = "INSERT INTO users (id,name,email,hashed_password) VALUES (?,?,?,?)"
	CHECK_USER_EXISTS     = "SELECT EXISTS (SELECT 1 FROM users WHERE email=?)"
	GET_HPASS_ID_AND_NAME = "SELECT id, name, hashed_password FROM users WHERE email=? LIMIT 1"
)
