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

	PROCESS_TEXT_ERR        = "Process text error."
	PROCESS_TEXT_ERR_DETAIL = "Unable to process text."
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
	LOGIN_ROUTE                 = "/auth/login"
	SIGNUP_ROUTE                = "/auth/register"
	SESSION_EXPIRED_ROUTE       = "/auth/session-expired"
	LOGOUT_ROUTE                = "/auth/log-out"
	WELCOME_ROUTE               = "/"
	HOME_ROUTE                  = "/home"
	ASCII_ROUTE                 = "/home/ascii-art"
	HISTORY_ROUTE               = "/home/ascii-art/history"
	ALL_HISTORY_QUERY           = "?font=all&page=1"
	STANDARD_HISTORY_QUERY      = "?font=standard&page=1"
	SHADOW_HISTORY_QUERY        = "?font=shadow&page=1"
	TINKERTOY_HISTORY_QUERY     = "?font=thinkertoy&page=1"
	ABOUT_US_ROUTE              = "/home#about"
	CONTRIBUTORS_ROUTE          = "/home#contributors"
	HELP_ROUTE                  = "/home#help"
	SAVE_ASCII_ROUTE            = "/home/ascii-art/save"
	DELETE_ROUTE                = "/home/ascii-art/delete"
	STANDARD_FILTER_ROUTE       = "/home/ascii-art/history/standard-filter"
	TINKERTOY_FILTER_ROUTE      = "/home/ascii-art/history/tinkertoy-filter"
	SHADOW_FILTER_ROUTE         = "/home/ascii-art/history/shadow-filter"
	ALL_ASCII_FILTER_ROUTE      = "/home/ascii-art/history/all-filter"
	CLEAR_ALL_ROUTE             = "/home/ascii-art/history/clear-all"
	COPY_ASCII_ROUTE            = "/history/ascii-art/copy"
	DOWNLOAD_ASCII_AS_TXT_ROUTE = "/home/ascii-art/download-txt"
	BEGIN_DOWNLOAD_TXT          = "/home/ascii-art/begin-download"
)

// Internal folder path
const (
	STYLES_PATH_PATTERN = "/web/static/styles/"
	STYLES_PATH         = "web/static/styles"
)

// Repo Query statements
const (
	// Users table
	INSERT_INTO_USERS                = "INSERT INTO users (id,name,email,hashed_password) VALUES (?,?,?,?)"
	CHECK_USER_EXISTS                = "SELECT EXISTS (SELECT 1 FROM users WHERE email=?)"
	GET_HPASS_ID_AND_NAME_WITH_EMAIL = "SELECT id, name, hashed_password FROM users WHERE email=? LIMIT 1"
	GET_HPASS_ID_AND_NAME_WITH_ID    = "SELECT id, name, hashed_password FROM users WHERE id=? LIMIT 1"

	// AsciiTexts Table
	INSERT_INTO_ASCII_TEXTS      = "INSERT INTO ascii_outputs (id,user_id,input_text,font,ascii_text) VALUES (?,?,?,?,?)"
	GET_ALL_USER_SAVED_ASCII     = "SELECT id, input_text, font, ascii_text, created_at FROM ascii_outputs WHERE user_id=? ORDER BY created_at DESC LIMIT ? OFFSET ?"
	CHECK_ASCII_EXISTS           = "SELECT EXISTS (SELECT 1 FROM ascii_outputs WHERE id=?)"
	DELETE_ASCII                 = "DELETE FROM ascii_outputs WHERE id=?"
	FILTER_ASCII                 = "SELECT id, input_text, font, ascii_text, created_at FROM ascii_outputs WHERE user_id=? AND font=? ORDER BY created_at DESC LIMIT ? OFFSET ?"
	CLEAR_ALL_USER_DATA          = "DELETE FROM ascii_outputs WHERE user_id=?"
	GET_TABLE_LENGHT_WITH_FONT   = "SELECT COUNT(*) FROM ascii_outputs WHERE user_id=? AND font=?"
	GET_TABLE_LENGHT             = "SELECT COUNT(*) FROM ascii_outputs WHERE user_id=?"
	UPDATE_ASCII_DOWNLOAD_AS_TXT = "UPDATE ascii_outputs SET downloaded_as_txt = true where id = ?"
)

// Ascii keys
const (
	TEXT_KEY                  = "text"
	BANNER_KEY                = "banner"
	DOWNLOAD_ASCII_AS_TXT_COL = "downloaded_as_txt"
)
