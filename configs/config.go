package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

type conf struct {
	DBHost        string `pastructure:"DB_DRIVER"`
	DBDriver      string `pastructure:"DB_HOST"`
	DBPort        string `pastructure:"DB_PORT"`
	DBUser        string `pastructure:"DB_USER"`
	DBPassword    string `pastructure:"DB_PASSWORD"`
	DBName        string `pastructure:"DB_NAME"`
	WebServerPort string `pastructure:"WEB_SERVER_PORT"`
	JWTSecret     string `pastructure:"JWT_SECRET"`
	JwtExpiresIn  int    `pastructure:"JWT_EXPIRESIN"`
	TokenAuth     *jwtauth.JWTAuth
}

func LoadConfig(path string) (*conf, error) {
	var cfg *conf

	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(*cfg)
	if err != nil {
		panic(err)
	}
	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)
	return cfg, nil
}
