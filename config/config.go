package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var conf Config

const confPath = "./config/config.json"

type Config struct {
	HttpAddress   string `json:"http_address"`
	Port          string `json:"port"`
	DbUsername    string `json:"db_username"`
	DbPass        string `json:"db_pass"`
	DbName        string `json:"db_name"`
	DbAddress     string `json:"db_address"`
	DbPort        string `json:"db_port"`
	CacheName     string `json:"cache_name"`
	CachePort     string `json:"cache_port"`
	CachePassword string `json:"cache_password"`
	CacheDuration string `json:"cache_duration"`
	AuthSecretKey string `json:"auth_secret_key"`
}

func init() {
	var err error
	conf, err = getConfig()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func GetConf() Config {
	return conf
}

func getConfig() (Config, error) {
	content, err := os.ReadFile(confPath)
	if err != nil {
		return Config{}, err
	}

	config := Config{}
	err = json.Unmarshal(content, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func GetHttpAddress() string {
	return fmt.Sprintf("%s:%s", conf.HttpAddress, conf.Port)
}

// GetStoreDsn returns dsn = data source name
func GetStoreDsn() string {
	// Example: db_username:db_password@tcp(host:port)/db_name?charset=utf8&parseTime=True&loc=Local
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.DbUsername,
		conf.DbPass,
		conf.DbAddress,
		conf.DbPort,
		conf.DbName,
	)
}

// GetRedisDsn returns dsn fore redis cache
func GetRedisDsn() string {
	return fmt.Sprintf(
		"%s:%s",
		conf.CacheName,
		conf.CachePort,
	)
}

func GetRedisCacheDuration() (time.Duration, error) {
	dur, err := strconv.Atoi(conf.CacheDuration)
	if err != nil {
		return 0, err
	}

	return time.Duration(dur) * time.Minute, nil
}

func GetSecretKey() string {
	return "fe655f46-c2eb-4560-a435-9673c4b919ac"
}

func GetDbName() string {
	return conf.DbName
}
