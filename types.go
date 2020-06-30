package designer

type command interface {
	Command() string
	Check() error
	Values() []string
}

const (
	COMMAND_DESIGNER             = "DESIGNER"
	COMMAND_CREATEINFOBASE       = "CREATEINFOBASE"
	COMMAND_ENTERPRISE           = "ENTERPRISE"
	DEFAULT_1SSERVER_PORT  int16 = 1541
)
