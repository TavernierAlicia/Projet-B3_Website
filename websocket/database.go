package main

import (
	"fmt"
	"os/exec"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func RunDb() (db *sqlx.DB, err error) {
	log, _ = zap.NewProduction()

	defer log.Sync()
	//// IMPORT CONFIG ////
	viper.SetConfigName("conf")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err = viper.ReadInConfig()
	if err != nil {
		log.Error("Unable to load config file", zap.Int("attempt", 3), zap.Duration("backoff", time.Second))
	}

	//// DEFINE CONFIG VARIABLES FROM JSON FILE ////
	host := viper.GetString("database.host")
	port := viper.GetInt("database.port")
	user := viper.GetString("database.user")
	pass := viper.GetString("database.pass")
	dbname := viper.GetString("database.dbname")

	//// DB CONNECTION ////
	pathSQL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, pass, host, port, dbname)
	db, err = sqlx.Connect("mysql", pathSQL)

	if err != nil {
		return db, err

	} else {
		log.Info("Connexion etablished ", zap.String("database", dbname),
			zap.Int("attempt", 3), zap.Duration("backoff", time.Second))
	}
	return db, err
}

func insertDb(mail string, name string, surname string, subject string, message string, pro bool, tel string, entname string) (err error) {

	db, err := RunDb()

	attempt := 1
	for attempt <= 3 && err != nil {
		//printerr
		log.Error("failed to connect database", zap.String("database", "orderndrink"),
			zap.Int("attempt", attempt), zap.Duration("backoff", time.Second))

		//restart mysql
		exec.Command("sudo", "service", "mysqld", "restart").Output()

		//wait
		time.Sleep(4 * time.Second)

		//retry
		db, err = RunDb()
		attempt++
	}

	if pro == true {
		_, err = db.Exec("INSERT INTO promessages(mail, name, surname, subject, message, tel, entname) VALUES(?, ?, ?, ?, ?, ?, ?)", mail, name, surname, subject, message, tel, entname)
		_, err = db.Exec("INSERT INTO proinfos(mail, name, surname, tel, entname) SELECT * FROM ( SELECT ? AS mail, ? AS name, ? AS surname, ? AS tel, ? AS entname)  AS ifexists WHERE NOT EXISTS (SELECT mail FROM proinfos WHERE mail = ?) LIMIT 1", mail, name, surname, tel, entname, mail)
	} else {
		_, err = db.Exec("INSERT INTO partmessages VALUES()", mail, name, surname, subject, message, tel)
		_, err = db.Exec("INSERT INTO partinfos(mail, name, surname, tel) SELECT * FROM ( SELECT ? AS mail, ? AS name, ? AS surname, ? AS tel) AS ifexists WHERE NOT EXISTS (SELECT mail FROM partinfos WHERE mail = ?)", mail, name, surname, tel, mail)
	}

	if err != nil {
		fmt.Println(err)
	}
	return err
}
