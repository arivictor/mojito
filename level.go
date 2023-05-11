package mojito

type Level int8

const (
	Debug Level = iota
	Info
	Warn
	Error
	Fatal
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
	case Debug:
		return DebugString
	case Info:
		return InfoString
	case Warn:
		return WarnString
	case Error:
		return ErrorString
	case Fatal:
		return FatalString
	default:
		return ""
	}
}
