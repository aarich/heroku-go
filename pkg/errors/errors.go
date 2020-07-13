package errors

const (
	SUCCESS                         = 2000
	INVALID_PARAMS                  = 4000
	UNKNOWN                         = 5000
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT = 5100
	ERROR_UPLOAD_CHECK_IMAGE_FAIL   = 5101
	ERROR_UPLOAD_SAVE_IMAGE_FAIL    = 5102
)

var messages = map[int]string{
	INVALID_PARAMS:                  "invalid parameters",
	UNKNOWN:                         "unknown",
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "bad image format",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL:   "image check failed",
	ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "Failed saving image",
	SUCCESS:                         "",
}

func GetMessage(errorCode int) string {
	message, ok := messages[errorCode]
	if ok {
		return message
	}

	return messages[UNKNOWN]
}
