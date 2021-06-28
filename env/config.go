package env

const (
	envPort = "PORT"
)

var (
	Port = GetEnv(envPort, "1988")
)
