package config

func GetSecretKey() string {
	return "fe655f46-c2eb-4560-a435-9673c4b919ac"
}

func GetMySqlDSN() string {
	return "root:@tcp(127.0.0.1:3306)/fidibo?charset=utf8mb4&parseTime=True&loc=Local"
}
