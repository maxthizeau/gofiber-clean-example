package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

const (
	defaultHTTPPort               = "8000"
	defaultHTTPRWTimeout          = 10 * time.Second
	defaultHTTPMaxHeaderMegabytes = 1
	defaultAccessTokenTTL         = 15 * time.Minute
	defaultRefreshTokenTTL        = 24 * time.Hour * 30
	defaultLimiterRPS             = 10
	defaultLimiterBurst           = 2
	defaultLimiterTTL             = 10 * time.Minute
	defaultVerificationCodeLength = 8
	frontendUrl                   = "http://localhost:5173"

	EnvLocal = "local"
	Prod     = "prod"
)

type (
	Config struct {
		Environment string
		PSQL        PSQLConfig
		HTTP        HTTPConfig
		Auth        AuthConfig
		Frontend    FrontendConfig
		// FileStorage FileStorageConfig
		// Email       EmailConfig
		// Payment     PaymentConfig
		// Limiter     LimiterConfig
		// CacheTTL    time.Duration `mapstructure:"ttl"`
		// SMTP        SMTPConfig
		// Cloudflare  CloudflareConfig
	}

	PSQLConfig struct {
		Host            string
		User            string
		Password        string
		DBName          string
		Port            string
		MaxPoolOpen     int
		MaxPoolIdle     int
		MaxPollLifeTime int
	}

	FrontendConfig struct {
		Url string `mapstructure:"url"`
	}

	AuthConfig struct {
		JWT                    JWTConfig
		PasswordSalt           string
		VerificationCodeLength int `mapstructure:"verificationCodeLength"`
	}

	JWTConfig struct {
		AccessTokenTTL  time.Duration `mapstructure:"accessTokenTTL"`
		RefreshTokenTTL time.Duration `mapstructure:"refreshTokenTTL"`
		SigningKey      string
	}

	HTTPConfig struct {
		Host               string        `mapstructure:"host"`
		Port               string        `mapstructure:"port"`
		ReadTimeout        time.Duration `mapstructure:"readTimeout"`
		WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
		MaxHeaderMegabytes int           `mapstructure:"maxHeaderBytes"`
	}
)

// Init populates Config struct with values from config file
// located at filepath and environment variables.
func Init(configsDir string) (*Config, error) {
	populateDefaults()

	if err := parseConfigFile(configsDir, os.Getenv("APP_ENV")); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	setFromEnv(&cfg)

	fmt.Println(cfg.PSQL)
	fmt.Println(cfg.HTTP)
	fmt.Println(cfg.Frontend)

	return &cfg, nil
}

func unmarshal(cfg *Config) error {

	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("auth", &cfg.Auth.JWT); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("frontend", &cfg.Frontend); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("auth.verificationCodeLength", &cfg.Auth.VerificationCodeLength); err != nil {
		return err
	}

	return nil
}

func setFromEnv(cfg *Config) {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("error loading env")
		panic(err)
	}

	// TODO use envconfig https://github.com/kelseyhightower/envconfig
	cfg.PSQL.Host = os.Getenv("DATASOURCE_HOST")
	cfg.PSQL.User = os.Getenv("DATASOURCE_USERNAME")
	cfg.PSQL.Password = os.Getenv("DATASOURCE_PASSWORD")
	cfg.PSQL.DBName = os.Getenv("DATASOURCE_DB_NAME")
	cfg.PSQL.Port = os.Getenv("DATASOURCE_PORT")
	cfg.PSQL.MaxPollLifeTime, err = strconv.Atoi(os.Getenv("DATASOURCE_POOL_LIFE_TIME"))
	if err != nil {
		panic("error while converting MaxPollLifeTime config")
	}
	cfg.PSQL.MaxPoolIdle, err = strconv.Atoi(os.Getenv("DATASOURCE_POOL_IDLE_CONN"))
	if err != nil {
		panic("error while converting MaxPoolIdle config")
	}
	cfg.PSQL.MaxPoolOpen, err = strconv.Atoi(os.Getenv("DATASOURCE_POOL_MAX_CONN"))
	if err != nil {
		panic("error while converting MaxPoolOpen config")
	}

	cfg.Auth.PasswordSalt = os.Getenv("PASSWORD_SALT")
	cfg.Auth.JWT.SigningKey = os.Getenv("JWT_SECRET_KEY")

	cfg.HTTP.Host = os.Getenv("HTTP_HOST")

	cfg.Environment = os.Getenv("APP_ENV")

}

func parseConfigFile(folder, env string) error {
	viper.AddConfigPath(folder)
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	// if env == EnvLocal {
	// 	return nil
	// }

	viper.SetConfigName(env)

	return viper.MergeInConfig()
}

func populateDefaults() {
	viper.SetDefault("http.port", defaultHTTPPort)
	viper.SetDefault("http.max_header_megabytes", defaultHTTPMaxHeaderMegabytes)
	viper.SetDefault("http.timeouts.read", defaultHTTPRWTimeout)
	viper.SetDefault("http.timeouts.write", defaultHTTPRWTimeout)
	viper.SetDefault("auth.accessTokenTTL", defaultAccessTokenTTL)
	viper.SetDefault("auth.refreshTokenTTL", defaultRefreshTokenTTL)
	viper.SetDefault("auth.verificationCodeLength", defaultVerificationCodeLength)
	viper.SetDefault("limiter.rps", defaultLimiterRPS)
	viper.SetDefault("limiter.burst", defaultLimiterBurst)
	viper.SetDefault("limiter.ttl", defaultLimiterTTL)
	viper.SetDefault("frontend.url", frontendUrl)
}
