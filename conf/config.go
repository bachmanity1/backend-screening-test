package conf

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// DefaultConf ...
type DefaultConf struct {
	EnvServerDEV   string
	EnvServerSTAGE string
	EnvServerPROD  string

	ConfServerPORT    int
	ConfServerTIMEOUT int
	ConfAPILOGLEVEL   string

	ConfDBHOST string
	ConfDBPORT int
	ConfDBUSER string
	ConfDBPASS string
	ConfDBNAME string
}

var defaultConf = DefaultConf{
	EnvServerDEV:      ".env.dev",
	EnvServerSTAGE:    ".env.stage",
	EnvServerPROD:     ".env",
	ConfServerPORT:    10811,
	ConfServerTIMEOUT: 30,
	ConfAPILOGLEVEL:   "debug",
	ConfDBHOST:        "infra_mysqldb",
	ConfDBPORT:        3306,
	ConfDBUSER:        "cbteam",
	ConfDBPASS:        "cbteampass",
	ConfDBNAME:        "pandita",
}

// ViperConfig ...
type ViperConfig struct {
	*viper.Viper
}

// Pandita ...
var Pandita *ViperConfig

func init() {
	pflag.BoolP("version", "v", false, "Show version number and quit")
	pflag.IntP("port", "p", defaultConf.ConfServerPORT, "pandita Port")

	pflag.String("db_host", defaultConf.ConfDBHOST, "pandita's DB host")
	pflag.Int("db_port", defaultConf.ConfDBPORT, "pandita's DB port")
	pflag.String("db_user", defaultConf.ConfDBUSER, "pandita's DB user")
	pflag.String("db_pass", defaultConf.ConfDBPASS, "pandita's DB password")
	pflag.String("db_name", defaultConf.ConfDBNAME, "pandita's DB name")

	pflag.Parse()

	var err error
	Pandita, err = readConfig(map[string]interface{}{
		"debug_route":  false,
		"debug_sql":    false,
		"port":         defaultConf.ConfServerPORT,
		"loglevel":     defaultConf.ConfAPILOGLEVEL,
		"profile":      false,
		"profilePort":  6060,
		"db_retry":     true,
		"db_maxopen":   100,
		"db_maxlife":   600,
		"env":          "devel",
		"swagger_host": "localhost:10811",
	})
	if err != nil {
		fmt.Printf("Error when reading config: %v\n", err)
		os.Exit(1)
	}

	Pandita.BindPFlags(pflag.CommandLine)
}

func readConfig(defaults map[string]interface{}) (*ViperConfig, error) {
	// Read Sequence (will overloading)
	// defaults -> config file -> env -> cmd flag
	v := viper.New()
	for key, value := range defaults {
		v.SetDefault(key, value)
	}
	v.AddConfigPath("./")
	v.AddConfigPath("./conf")
	v.AddConfigPath("../conf")
	v.AddConfigPath("../../conf")
	v.AddConfigPath("$HOME/.pandita")

	v.AutomaticEnv()

	stage := strings.ToLower(v.GetString("ENV"))
	switch stage {
	case "devel":
		v.SetConfigName(defaultConf.EnvServerDEV)
	case "stage":
		v.SetConfigName(defaultConf.EnvServerSTAGE)
	case "prod":
		v.SetConfigName(defaultConf.EnvServerPROD)
	default:
		v.SetConfigName(fmt.Sprintf(".env.%s", stage))
	}

	err := v.ReadInConfig()
	switch err.(type) {
	default:
		fmt.Println("error ", err)
		return &ViperConfig{}, err
	case nil:
		break
	case viper.ConfigFileNotFoundError:
		fmt.Printf("Warn: %s\n", err)
	}

	return &ViperConfig{
		Viper: v,
	}, nil
}
