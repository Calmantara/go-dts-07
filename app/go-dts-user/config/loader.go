package config

import (
	"flag"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	MODE_PG   = "PG"
	MODE_GORM = "GORM"
	MODE_MAP  = "MAP"

	ENV_LOCAL      = "local"
	ENV_PRODUCTION = "production"
)

var (
	Load structure

	availableMode = map[string]bool{
		MODE_MAP:  true,
		MODE_PG:   true,
		MODE_GORM: true,
	}

	searchPath = []string{
		".",
		"./config",
	}
)

type (
	structure struct {
		Server     server     `mapstructure:"server"`
		DataSource dataSource `mapstructure:"dataSource"`
	}
	server struct {
		Name string `mapstructure:"name"`
		Env  string `mapstructure:"env"`
		Http struct {
			Port uint `mapstructure:"port"`
		} `mapstructure:"http"`
	}
	dataSource struct {
		Mode     string `mapstructure:"mode"`
		Migrate  bool   `mapstructure:"migrate"`
		Postgres struct {
			Master PostgresConfig `mapstructure:"master"`
		}
	}
)

// init config to load all
// config sturcture
func init() {
	var configName = flag.String("configName", "local", "config name for service, default local")
	var source = flag.String("source", "MAP", "data source mode for service, default MAP")

	v := viper.New()
	if err := initialiseFileAndEnv(v, *configName); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Warning("No config file found on search paths")
		} else {
			log.Fatalf("Error occurred during loading config: %s", err.Error())
		}
		panic(err)
	}

	if err := v.Unmarshal(&Load); err != nil {
		log.Fatal("cannot unmarshal config")
		panic(err)
	}

	if availableMode[*source] {
		Load.DataSource.Mode = *source
	}
}

func initialiseFileAndEnv(v *viper.Viper, configName string) error {
	v.SetConfigName(configName)
	for _, path := range searchPath {
		v.AddConfigPath(path)
	}
	return v.ReadInConfig()
}
