package config

import (
	"encoding/json"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"otus.ru/tbw/msa-25/src/deps/log"
)

// Настройки приложения
type Config struct {
	// Настройки сервера
	Server struct {
		// Хост сервера
		Host string `yaml:"host" env:"HOST" env-default:"127.0.0.1"`

		// Порт сервера
		Port string `yaml:"port" env:"PORT" env-default:"8000"`

		// Таймауты сервера
		Timeout struct {
			// Общий таймаут, используемый для безопасного закрытия сервера
			Server time.Duration `yaml:"server" env:"SERVER" env-default:"30s"`

			// Таймаут после которых будет прервана операция записи HTTP
			Write time.Duration `yaml:"write" env:"WRITE" env-default:"10s"`

			// Таймаут после которых будет прервана операция чтения HTTP
			Read time.Duration `yaml:"read" env:"READ" env-default:"15s"`

			// Таймаут после которых будет прервана idle-сессия HTTP
			Idle time.Duration `yaml:"idle" env:"IDLE" env-default:"5s"`
		} `yaml:"timeout" env-prefix:"TIMEOUT_"`
	} `yaml:"server" env-prefix:"SERVER_"`
	DB struct {
		// Хост сервера БД
		Host string `yaml:"host" env:"HOST" env-default:"localhost"`

		// Порт сервера БД
		Port string `yaml:"port" env:"PORT" env-default:"5432"`

		// Отключение проверки SSL БД
		SSLInsecure bool `yaml:"ssl_insecure" env:"SSL_INSECURE" env-default:"true"`

		// Название БД
		Database string `yaml:"database" env:"DATABASE" env-default:"dev"`

		// Имя пользователя БД
		Username string `yaml:"username" env:"USERNAME" env-default:"postgres"`

		// Пароль пользователя БД
		Password secretString `yaml:"password" env:"PASSWORD" env-default:"postgres"`

		// Название приложения
		ApplicationName string `yaml:"application_name" env:"APPLICATION_NAME" env-default:"2025-msa-project"`

		// Таймауты БД
		Timeout struct {
			// Таймаут после которых будет прервана операция записи БД
			Write time.Duration `yaml:"write" env:"WRITE" env-default:"5s"`

			// Таймаут после которых будет прервана операция чтения БД
			Read time.Duration `yaml:"read" env:"READ" env-default:"5s"`

			// Таймаут создания новой сессии БД
			Dial time.Duration `yaml:"dial" env:"DIAL" env-default:"5s"`
		} `yaml:"timeout" env-prefix:"TIMEOUT_"`
	} `yaml:"db" env-prefix:"DB_"`
}

func New(log *log.Logger) *Config {
	config := &Config{}

	config_path, exists := os.LookupEnv("CONFIG_PATH")
	if exists == false {
		// По-умолчанию конфигурация задана стандартными значениями
		// и может быть перезаписана в .env
		config_path = ".env"
	}

	var err error
	_, err = os.Stat(config_path)
	if os.IsNotExist(err) {
		// В случае если нет никакого файла конфигурации,
		// то читаем переменные окружения
		err = cleanenv.ReadEnv(config)
		if err != nil {
			log.Fatalf("Error reading config from environment, %v", err)
		}
	} else {
		err = cleanenv.ReadConfig(config_path, config)
		if err != nil {
			log.Fatalf("Error reading config file, %v", err)
		}
	}

	configJSON, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		log.Panicf("Something went wrong while cofig structure: %v", err)
	}
	log.Debugf("Log config structure for convenience:\n%s", string(configJSON))

	return config
}
