package auth

import (
	"time"

	"portfolio/services/infrastructure/config"
)

type CoreConfig struct {
	*config.Config
}

type Config struct {
	Environment                     string `yaml:"environment"`
	IsActiveMessageGeneratorWorker  bool   `mapstructure:"is_active_message_generator_worker"`
	PeriodCronMessageGenerator      string `mapstructure:"period_cron_message_generator"`
	Server                          server `yaml:"server"`
	EwaysShopPhoneVerifierAuthToken string `mapstructure:"eways_shop_phone_verifier_auth_token"`
	EwaysShopSmsAtiyeUsername       string `mapstructure:"eways_shop_sms_atiye_username"`
	EwaysShopSmsAtiyePassword       string `mapstructure:"eways_shop_sms_atiye_password"`

	AuthTokenSecret             string `mapstructure:"eways_shop_auth_token_secret"`
	StocksElasticsearchUsername string `mapstructure:"stocks_elasticsearch_username"`
	StocksElasticsearchPassword string `mapstructure:"stocks_elasticsearch_password"`
	FileManagerServeUrl         string `mapstructure:"file_manager_serve_url"`
	Http                        struct {
		Port string `yaml:"port"`
	} `yaml:"http"`
	Parsimap struct {
		Token string `yaml:"token"`
	} `mapstructure:"parsimap"`
	OTP struct {
		Engine string `yaml:"engine"`
		ExpSec uint8  `mapstructure:"exp_sec"`
		Config struct {
			Code string `yaml:"code"`
		} `yaml:"config"`
	} `yaml:"otp"`
	SMS struct {
		Provider string `yaml:"provider"`
		Org      string `mapstructure:"org"`
		Sender   string `yaml:"sender"`
	} `yaml:"sms"`
	PhoneVerifier struct {
		Provider string `yaml:"provider"`
		Config   struct {
			Response bool   `yaml:"response"`
			Url      string `yaml:"url"`
		} `yaml:"config"`
	} `mapstructure:"phone_verifier"`
	RateLimiter struct {
		MaxCustomerRegreqs          int `mapstructure:"max_customer_regreqs"`
		CustomerRegResetIntervalHrs int `mapstructure:"customer_regreq_reset_interval_hrs"`
	} `mapstructure:"rate_limiter"`
	Redis  redis `yaml:"redis"`
	Mellat struct {
		TerminalId        int32  `yaml:"terminal_id" mapstructure:"terminal_id"`
		UrlWebService     string `yaml:"url_web_service" mapstructure:"url_web_service"`
		UrlPaymentGateway string `yaml:"url_payment_gateway" mapstructure:"url_payment_gateway"`
		Username          string `yaml:"username" mapstructure:"username"`
		Password          string `yaml:"password" mapstructure:"password"`
		UrlCallback       string `yaml:"url_callback" mapstructure:"url_callback"`
		WalletUrlCallback string `yaml:"wallet_url_callback" mapstructure:"wallet_url_callback"`
	} `yaml:"mellat" mapstructure:"mellat"`
	Db      DB      `yaml:"db"`
	Payment Payment `yaml:"payment"`
	Sentry  struct {
		DSN string `yaml:"dsn"`
	} `yaml:"sentry"`
	Minio Minio `yaml:"minio"`

	Core          core          `yaml:"core"`
	ElasticSearch Elasticsearch `yaml:"elasticsearch"`
	Services      services      `yaml:"services"`
}

type Payment struct {
	PaymentTime time.Duration `yaml:"payment_time" mapstructure:"payment_time"`
	OrderTime   time.Duration `yaml:"order_time" mapstructure:"order_time"`
}

type core struct {
	LogRoute string `mapstructure:"log_route"`
}

type DB struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `mask:"filled" yaml:"password"`
	Name     string `yaml:"name"`
	Driver   string `yaml:"driver"`
	SslMode  string `mapstructure:"ssl_mode"`
}

type redis struct {
	Address  string `yaml:"address"`
	Password string `mask:"filled" yaml:"password"`
	DB       int    `yaml:"db"`
}

type Elasticsearch struct {
	Addresses []string `yaml:"addresses"`
	Username  string   `yaml:"username"`
	Password  string   `yaml:"password"`
}

type server struct {
	Address string `yaml:"address"`
}

type Minio struct {
	Host      string `mask:"filled"`
	AccessKey string `mask:"filled" mapstructure:"access_key"`
	SecretKey string `mask:"filled" mapstructure:"secret_key"`
	Bucket    string `yaml:"bucket"`
	Secure    bool   `yaml:"secure"`
	Token     string `yaml:"token" mask:"filled"`
}

type services struct {
	Auth struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"auth"`

	Policy struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"policy"`
}

func LoadConfig() *Config {
	conf := config.NewConfig("./config.yml", &Config{})
	internalConfig := CoreConfig{conf}
	return internalConfig.GetInternalSection()
}

func (c *CoreConfig) GetInternalSection() *Config {
	return c.Internal.(*Config)
}
