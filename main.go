package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	web "mailspamer/web"
	mail "mailspamer/mail"
	database "mailspamer/db"
	"os"
)

type Config struct {
	ServerPort string `yaml:"serverPort"`
	ExternalServer string `yaml:"externalServer"`
	Mailconf struct {
		Mail     string `yaml:"mail"`
		Password string `yaml:"password"`
		SMTP     string `yaml:"smtp"`
	} `yaml:"mailconf"`
	Database struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Dbname   string `yaml:"dbname"`
		Sslmode  string `yaml:"sslmode"`
	} `yaml:"database"`
}

func mailInit(config Config) {
	mail.FromEmail = config.Mailconf.Mail
	mail.PasswordEmail = config.Mailconf.Password
	mail.HostEmail = config.Mailconf.SMTP
}

func dbInit(config Config) {

	user := "user=" + config.Database.User + " "
	password := "password=" + config.Database.Password + " "
	dbname := "dbname=" + config.Database.Dbname + " "
	sslmode := "sslmode=" + config.Database.Sslmode
	connStr := user + password + dbname + sslmode

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("sql.Open error: ", err)
	}
	database.CreateDB(db)
	web.Database = db
}

func readConfig(conFile string) (Config){
	yamlFile, err := ioutil.ReadFile(conFile)
	if err != nil {
		fmt.Println("ioutil.ReadFile error: ", err)
	}

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Println("yaml.Unmarshal error: ", err)
	}

	return config
}

func main() {
	pathToConfig := os.Args[1]
	config := readConfig(pathToConfig)
	dbInit(config)
	mailInit(config)
	web.ExternalServer = config.ExternalServer
	web.ServerPort = config.ServerPort
	web.WebServer()
}
