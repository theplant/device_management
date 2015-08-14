package config

import (
	"log"
	"os"
	"strings"
)

// this is configured from env variables
var (
	Env               string
	WebDir            string
	MySQLHost         string
	MySQLPort         string
	MySQLDatabase     string
	MySQLRootPassword string
	Verbose           bool
)

func init() {
	Env = envOrPanic("PROFILE_ENV", false)
	WebDir = envOrPanic("PROFILE_WEBDIR", false)

	MySQLHost = envOrPanic("PROFILE_MYSQL_PORT_3306_TCP_ADDR", false)
	MySQLPort = envOrPanic("PROFILE_MYSQL_PORT_3306_TCP_PORT", false)
	MySQLRootPassword = envOrPanic("PROFILE_MYSQL_ENV_MYSQL_ROOT_PASSWORD", true)

	MySQLDatabase = envOrPanic("PROFILE_MYSQL_DATABASE", false)
	Verbose = (envOrPanic("PROFILE_VERBOSE", true) != "")
}

func envOrPanic(key string, allowEmpty bool) (r string) {
	r = os.Getenv(key)
	if r == "" && !allowEmpty {
		panic("env " + key + " is not set")
	}
	logValue := r
	if strings.Contains(logValue, "PASSWORD") || strings.Contains(logValue, "SECRET") {
		logValue = "<HIDDEN>"
	}
	log.Printf("Configure: %s = %s\n", key, logValue)
	return
}
