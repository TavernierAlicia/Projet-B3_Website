package main

import (
	_ "database/sql"
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

//Prepare struct to get send mail
type Mail struct {
	From      string ``
	Name      string ``
	Surname   string ``
	Subject   string ``
	CmdNumber string ``
	Message   string ``
}

func (r Results) String() string {
	return ""
}

//connecting to the database
func DbConnect(clientcode string, idconf string) (bool, bool) {

	//configure log
	log, _ = zap.NewProduction()
	defer log.Sync()

	//set variables
	hmanyrequest := false
	clirequest := false

	//converting string to int
	intidconf, er := strconv.ParseInt(clientcode, 10, 64)
	if er != nil {
		log.Fatal("idconf conversion failed")
	} else {
		log.Info("idconf conversion succeed")
	}

	//printing for now
	fmt.Println(intidconf)

	//Import config file with viper
	viper.SetConfigName("dbconf")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil {
		log.Error("Unable to load config file ")
		hmanyrequest = false
		clirequest = false
	} else {
		log.Info("config file loaded")
	}

	//Define config variables from JSON
	host := viper.GetString("database.host")
	port := viper.GetInt("database.port")
	user := viper.GetString("database.user")
	password := viper.GetString("database.password")
	dbname := viper.GetString("database.dbname")

	//Connect
	dbPath := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbname)
	db, err := sqlx.Connect("mysql", dbPath)

	if err != nil {
		log.Error("Failed to connect db", zap.String("database", err.Error()))
		clirequest = false
		hmanyrequest = false
	} else {
		log.Info("Connection etablished")
	}

	//Define requests
	clientrequest := "SELECT c.email,c.name,cc.isadmin,cc.ismuted FROM clients AS c,currentcalls AS cc,clientcode AS cco WHERE cco.clientcode = :intclientcode AND cco.confid = :intidconf AND cc.clientcode=cco.clientcode AND cc.confid=cco.confid AND c.id=cco.clientid "

	nbrequest := "SELECT COUNT(id) AS nb FROM currentcalls WHERE confid = :idconf"

	//Executing requests

	//client request
	err = db.Get(&Results{}, clientrequest)
	if err != nil {
		log.Error("Failed to request database", zap.String("database", err.Error()))
		clirequest = false
	} else {
		log.Info("Request Succeed")
		clirequest = true
	}

	//nb clients request
	err = db.Get(&Results{}, nbrequest)
	if err != nil {
		log.Error("Failed to request database", zap.String("database", err.Error()))
		hmanyrequest = false
	} else {
		log.Info("Request succeed")
		hmanyrequest = true
	}

	fmt.Println(clirequest)
	fmt.Println(hmanyrequest)
	return clirequest, hmanyrequest
}
