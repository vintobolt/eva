package config

import (
	"fmt"
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Database struct {
		Ip       string `env:"PG_IP" env-required:"true" env-description:"database ip address"`
		Port     string `env:"PG_PORT" env-required:"true" env-description:"database port"`
		User     string `env:"PG_USER" env-required:"true" env-description:"database user"`
		Password string `env:"PG_PASSWD" env-required:"true" env-description:"database passwd"`
		Dbname   string `env:"PG_DBNAME" env-required:"true" env-description:"database name"`
	}
	Server struct {
		Ip           string `env:"SRV_IP" env-required:"true" env-description:"EVA server ip address"`
		Port         string `env:"SRV_PORT" env-required:"true" env-description:"EVA server port"`
		LogPath      string `env:"SRV_LOGPATH" env-required:"true" env-description:"EVA servers log path"`
		Timeout      int    `env:"SRV_GRACEFUL_TIMEOUT" env-required="true" env-description:"TIMEOUT for GRACEFUL shutdown"`
		LogLevel     string `env:"SRV_LOGLEVEL" env-description:"Levels: Debug, Info, Warning, Error, DPanic, Panic, and Fatal."`
		ReadTimeout  int    `env:"SRV_READ_TIMEOUT" env-description:"HTTP read timeout"`
		WriteTimeout int    `env:"SRV_WRITE_TIMEOUT" env-description:"HTTP write timeout"`
	}
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		log.Println("Gather config")
		instance = &Config{}
		if err := cleanenv.ReadEnv(instance); err != nil {
			var helpHeaderText string = "EVA server wrong configured.."
			helpText, _ := cleanenv.GetDescription(instance, &helpHeaderText)
			log.Println(helpText)
			log.Fatalln(err)
		}
	})
	return instance
}

// urlExample := "postgres://username:password@localhost:5432/database_name"
func (c *Config) GetPostgresConnectionString() string {
	postgres_connection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		c.Database.User,
		c.Database.Password,
		c.Database.Ip,
		c.Database.Port,
		c.Database.Dbname)
	return postgres_connection
}

func (c *Config) GetServeString() string {
	serverString := fmt.Sprintf("%s:%s", c.Server.Ip, c.Server.Port)
	return serverString
}
