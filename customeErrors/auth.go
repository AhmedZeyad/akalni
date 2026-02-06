package customeErrors

const (
	AUTH_UNAUTHORIZED        = "AUTH_UNAUTHORIZED"
	AUTH_INVALID_TOKEN       = "AUTH_INVALID_TOKEN"
	AUTH_TOKEN_EXPIRED       = "AUTH_TOKEN_EXPIRED"
	AUTH_MISSING_TOKEN       = "AUTH_MISSING_TOKEN"
	AUTH_INVALID_CREDENTIALS = "AUTH_INVALID_CREDENTIALS"
	AUTH_ACCOUNT_LOCKED      = "AUTH_ACCOUNT_LOCKED"
	AUTH_EMAIL_NOT_VERIFIED  = "AUTH_EMAIL_NOT_VERIFIED"
	AUTH_EMAIL_NOT_FOUND     = "AUTH_EMAIL_NOT_FOUND"
	AUTH_USER_NOT_FOUND      = "AUTH_USER_NOT_FOUND"
	AUTH_PASSWORD_INCORRECT  = "AUTH_PASSWORD_INCORRECT"
)

// `VALIDATION_FAILED` | General validation failure for the request. |
// | `VALIDATION_FIELD_MISSING` | A required field is missing from the payload. |
// | `VALIDATION_INVALID_FORMAT` | Field format (e.g., Date, UUID) is invalid. |
// | `VALIDATION_OUT_OF_RANGE` | Field value is outside the allowed range. |
// | `VALIDATION_EMAIL_INVALID` | Email format is technically incorrect. |
// | `VALIDATION_PASSWORD_WEAK` | Password does not meet complexity requirements. |
// | `VALIDATION_DUPLICATE
