package mojito

import "time"

const (
	// Default Log Keys
	DEFAULT_SEVERITY_KEY = "severity"
	DEFAULT_TIME_KEY     = "time"
	DEFAULT_MESSAGE_KEY  = "message"
	DEFAULT_TRACE_KEY    = "trace"

	// Default Log Values
	DEFAULT_DELIMITER   = '\n'
	DEFAULT_MIN_LEVEL   = INFO
	DEFAULT_TIME_FORMAT = time.RFC3339
)
