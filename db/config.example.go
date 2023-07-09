package db

import "github.com/go-sql-driver/mysql"

// Rename to getConfig()
func getConfigExample() mysql.Config {
	conf := mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Addr:                 "127.0.1.1:3306",
		DBName:               "mnogoclinica",
		Net:                  "tcp",
		AllowNativePasswords: true,
	}

	return conf
}
