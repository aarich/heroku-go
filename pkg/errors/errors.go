package errors

const (
	INVALID_PARAMS                  = 4000
	UNKNOWN                         = 5000
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT = 5100
	ERROR_UPLOAD_CHECK_IMAGE_FAIL   = 5101
	ERROR_UPLOAD_SAVE_IMAGE_FAIL    = 5102
)

var messages = map[int]string{
	INVALID_PARAMS:                  "Invalid parameters.",
	UNKNOWN:                         "Unknown error.",
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "Bad image format.",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL:   "Image check failed.",
	ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "Failed saving image.",
}

func GetMessage(errorCode int) string {
	message, ok := messages[errorCode]
	if ok {
		return message
	}

	return messages[UNKNOWN]
}
