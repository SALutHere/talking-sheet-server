package logger

const (
	EnvLocal = "local"
	EnvDev   = "dev"
	EnvProd  = "prod"

	LvlInfo  = "info"
	LvlWarn  = "warn"
	LvlDebug = "debug"
	LvlError = "error"
)

type Config struct {
	Env   string
	Level string
}
