package admin

import (
	"git.eways.dev/eways/service/config"
)

type CoreConfig struct {
	*config.Config
}

type Config struct {
	ServiceName   string        `yaml:"service_name"`
	Core          core          `yaml:"core"`
	Environment   string        `yaml:"environment"`
	Postgres      Postgres      `yaml:"postgres"`
	Redis         redis         `yaml:"redis"`
	ElasticSearch elasticsearch `yaml:"elasticsearch"`
	Services      services      `yaml:"services"`
	Server        server        `yaml:"server"`
	Storage       Storage       `yaml:"storage"`
	Payment       Payment       `yaml:"payment"`
}

type Payment struct {
	Provider struct {
		Mellat struct {
			TerminalId        string `yaml:"terminal_id"`
			UrlWebService     string `yaml:"url_web_service"`
			UrlPaymentGateway string `yaml:"url_payment_gateway"`
			Username          string `yaml:"username"`
			Password          string `yaml:"password"`
			UrlCallback       string `yaml:"url_callback"`
			WalletUrlCallback string `yaml:"wallet_url_callback"`
		} `yaml:"mellat"`
	} `yaml:"provider"`
	PaymentTime string `yaml:"payment_time"`
	OrderTime   string `yaml:"order_time"`
}

type core struct {
	LogRoute string `yaml:"log_route"`
	Sentry   struct {
		DSN string `yaml:"dsn"`
	} `yaml:"sentry"`
	OTP struct {
		Engine string `yaml:"engine"`
		ExpSec string `yaml:"exp_sec"`
	} `yaml:"otp"`
	SMS struct {
		Provider string `yaml:"provider"`
	} `yaml:"sms"`
	PhoneVerifier struct {
		Provider string `yaml:"provider"`
		Config   struct {
			Response bool `yaml:"response"`
		} `yaml:"config"`
	} `yaml:"phone_verifier"`
	RateLimiter struct {
		MaxCustomerRegreqs          int `yaml:"max_customer_regreqs"`
		CustomerRegResetIntervalHrs int `yaml:"customer_regreq_reset_interval_hrs"`
	} `yaml:"rate_limiter"`
	Parsimap struct {
		Token string `yaml:"token"`
	} `yaml:"parsimap"`
	JWT struct {
		AuthTokenSecret string `yaml:"auth_token_secret"`
	} `yaml:"jwt"`
}

type Postgres struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `mask:"filled" yaml:"password"`
	Name     string `yaml:"name"`
	Driver   string `yaml:"driver"`
	Ssl      string `yaml:"ssl"`
}

type redis struct {
	Address  string `yaml:"address"`
	Password string `mask:"filled" yaml:"password"`
	DB       int    `yaml:"db"`
}

type elasticsearch struct {
	Addresses []string `yaml:"addresses"`
	Username  string   `yaml:"username"`
	Password  string   `yaml:"password"`
}

type server struct {
	Address string `yaml:"address"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
}

type Storage struct {
	Minio struct {
		Host      string `mask:"filled"`
		AccessKey string `mask:"filled" yaml:"access_key"`
		SecretKey string `mask:"filled" yaml:"secret_key"`
		Bucket    string `yaml:"bucket"`
		Secure    bool   `yaml:"secure"`
		Token     string `yaml:"token" mask:"filled"`
	} `yaml:"minio"`
}

type services struct {
	Auth Auth `yaml:"auth"`
	Policy Policy `yaml:"policy"`
}

type Auth struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Policy struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func LoadConfig() *Config {
	conf := config.NewConfig("./config/admin_tmp_todo.yml", &Config{})
	internalConfig := CoreConfig{conf}
	return internalConfig.GetInternalSection()
}

func (c *CoreConfig) GetInternalSection() *Config {
	return c.Internal.(*Config)
}
