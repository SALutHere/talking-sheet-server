package log

type Environment string

const (
	envLocal Environment = "local"
	envDev   Environment = "dev"
	envProd  Environment = "prod"
)

type Config struct {
	Env Environment
}
