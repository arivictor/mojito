package mojito

type Level int8

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
	FATAL
)

const (
	DebugString = "DEBUG"
	InfoString  = "INFO"
	WarnString  = "WARN"
	ErrorString = "ERROR"
	FatalString = "FATAL"
)

func (l Level) String() string {
	switch l {
	case DEBUG:
		return DebugString
	case INFO:
		return InfoString
	case WARN:
		return WarnString
	case ERROR:
		return ErrorString
	case FATAL:
		return FatalString
	default:
		return ""
	}
}
