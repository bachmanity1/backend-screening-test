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
	ConfAPILOGLEVEL:   "log_debug",
	ConfDBHOST:        "terra",
	ConfDBPORT:        3306,
	ConfDBUSER:        "bachmanity1",
	ConfDBPASS:        "bachmanity1",
	ConfDBNAME:        "terra",
}

// ViperConfig ...
type ViperConfig struct {
	*viper.Viper
}

// Terra ...
var Terra *ViperConfig

func init() {
	pflag.BoolP("version", "v", false, "Show version number and quit")
	pflag.IntP("port", "p", defaultConf.ConfServerPORT, "terra Port")

	pflag.String("db_host", defaultConf.ConfDBHOST, "terra's DB host")
	pflag.Int("db_port", defaultConf.ConfDBPORT, "terra's DB port")
	pflag.String("db_user", defaultConf.ConfDBUSER, "terra's DB user")
	pflag.String("db_pass", defaultConf.ConfDBPASS, "terra's DB password")
	pflag.String("db_name", defaultConf.ConfDBNAME, "terra's DB name")

	pflag.Parse()

	var err error
	Terra, err = readConfig(map[string]interface{}{
		"debug_route": false,
		"debug_sql":   false,
		"port":        defaultConf.ConfServerPORT,
		"loglevel":    defaultConf.ConfAPILOGLEVEL,
		"db_retry":    true,
		"db_maxopen":  100,
		"db_maxlife":  600,
		"env":         "devel",
	})
	if err != nil {
		fmt.Printf("Error when reading config: %v\n", err)
		os.Exit(1)
	}

	Terra.BindPFlags(pflag.CommandLine)
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
	v.AddConfigPath("$HOME/.terra")

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
