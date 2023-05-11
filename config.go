package mojito

type config struct {
	severityKey  string
	messageKey   string
	timestampKey string
	delimiter    byte
	timeFormat   string
	minLevel     Level
}

func (c *config) GetLevel() Level {
	return c.minLevel
}

func (c *config) SetLevel(level Level) {
	c.minLevel = level
}

func (c *config) SetTimeFormat(format string) {
	c.timeFormat = format
}

func (c *config) SetDelimiter(delimiter byte) {
	c.delimiter = delimiter
}

func (c *config) SetMessageKey(key string) {
	c.messageKey = key
}

func (c *config) SetSeverityKey(key string) {
	c.severityKey = key
}

func (c *config) SetTimeKey(key string) {
	c.timestampKey = key
}
