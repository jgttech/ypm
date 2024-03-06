package utils

// For ease of use, the status codes that should be
// returned from the spinner function should be one of
// the valid HTTP status codes listed here. I decided
// on using HTTP status codes because it leverages
// existing web development knowledge.
//
// Status Codes:
// 199: Miscellaneous Warning
// 200: Success
// 501: Internal Server Error
const SUCCESS = 200
const WARNING = 199
const FAILURE = 501

// General binary status codes.
const YES = 1
const NO = -1
const UNKNOWN = 0
