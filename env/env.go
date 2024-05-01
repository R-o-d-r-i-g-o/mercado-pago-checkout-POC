package env

func Load() {
	loadEnv(&GeneralConfig, &EmailService, &Server, &Database, &Tracer, &Logger, &MercadoPago)
}

var (
	GeneralConfig GeneralConfigEnv
	EmailService  EmailServiceEnv
	Server        ServerEnv
	Database      DatabaseEnv
	Tracer        TracerEnv
	Logger        loggerEnv
	MercadoPago   MercadoPagoEnv
)

type GeneralConfigEnv struct {
	ServiceName        string `env:"SERVICE_NAME"`
	ServiceEnvironment string `env:"SERVICE_ENVIRONMENT"`
	SecretKey          string `env:"TOKEN_SECRET_KEY"`
}

type EmailServiceEnv struct {
	Host          string `env:"EMAIL_HOST"`
	Port          int    `env:"EMAIL_PORT"`
	Username      string `env:"EMAIL_USERNAME"`
	Password      string `env:"EMAIL_PASSWORD"`
	SenderAddress string `env:"EMAIL_SENDER_ADDRESS"`
}

type ServerEnv struct {
	Host string `env:"SERVER_HOST"`
	Port string `env:"SERVER_PORT"`
}

type DatabaseEnv struct {
	DSN string `env:"DB_DSN"`
}

type TracerEnv struct {
	IsTracingEnabled bool `env:"TRACE_ENABLED"`
}

type loggerEnv struct {
	LogLevel string `env:"LOG_LEVEL"`
}

type MercadoPagoEnv struct {
	Token   string `env:"MP_TOKEN"`
	BaseURL string `env:"MP_BASE_URL"`
}
