# ⚠️ Error Codes Specification

This document defines the standard error codes used by the Akelni API.  
All error codes follow the naming convention: `<DOMAIN>_<REASON>`

> **Note:** Error codes are stable and must not be changed once released to ensure frontend consistency.

---

### 🔐 AUTH Errors
**Default HTTP Status:** `401 Unauthorized`

| Error Code | Description |
| :--- | :--- |
| `AUTH_UNAUTHORIZED` | User is not authenticated. |
| `AUTH_INVALID_TOKEN` | Provided authentication token is invalid. |
| `AUTH_TOKEN_EXPIRED` | Authentication token has expired. |
| `AUTH_MISSING_TOKEN` | Authentication token is missing from headers. |
| `AUTH_INVALID_CREDENTIALS` | Email/username or password is incorrect. |
| `AUTH_ACCOUNT_LOCKED` | Account is locked due to security reasons. |

---

### ✅ VALIDATION Errors
**Default HTTP Status:** `422 Unprocessable Entity`

| Error Code | Description |
| :--- | :--- |
| `VALIDATION_FAILED` | General validation failure for the request. |
| `VALIDATION_FIELD_MISSING` | A required field is missing from the payload. |
| `VALIDATION_INVALID_FORMAT` | Field format (e.g., Date, UUID) is invalid. |
| `VALIDATION_OUT_OF_RANGE` | Field value is outside the allowed range. |
| `VALIDATION_EMAIL_INVALID` | Email format is technically incorrect. |
| `VALIDATION_PASSWORD_WEAK` | Password does not meet complexity requirements. |
| `VALIDATION_DUPLICATE` | Value must be unique but already exists in the system. |

---

### 👤 USER Errors
**Default HTTP Status:** `404 Not Found` / `409 Conflict`

| Error Code | Description | Status |
| :--- | :--- | :--- |
| `USER_NOT_FOUND` | User does not exist in the database. | 404 |
| `USER_ALREADY_EXISTS` | User with this email/phone already exists. | 409 |
| `USER_DISABLED` | User account has been deactivated by admin. | 403 |
| `USER_NOT_VERIFIED` | User email or phone is not yet verified. | 403 |
| `USER_PROFILE_INCOMPLETE`| Required profile details are missing. | 400 |

---

### 🚫 PERMISSION Errors
**Default HTTP Status:** `403 Forbidden`

| Error Code | Description |
| :--- | :--- |
| `PERMISSION_DENIED` | General lack of permission for this action. |
| `PERMISSION_ROLE_LOW` | User role (e.g., Driver) cannot access this (e.g., Admin). |
| `PERMISSION_ACCESS_DENIED`| User cannot access this specific resource instance. |
