package database

import (
	"database/sql"
	"log"
	"net/url"
	"os"
	"os/signal"
	"project1/config"
	_ "github.com/go-sql-driver/mysql"
)

type DbConnections struct {
	Read  *sql.DB
	Write *sql.DB
}

var Connections DbConnections

type DbConfig struct {
	Host        string `yaml:"host" json:"host"`
	Port        string `yaml:"port" json:"port"`
	Db          string `yaml:"database" json:"database"`
	User        string `yaml:"user" json:"user"`
	Password    string `yaml:"password" json:"password"`
	MaxOpenCons int    `yaml:"max_open_connections" json:"max_open_connections"`
	MaxIdleCons int    `yaml:"max_idle_connections" json:"max_idle_connections"`
}

type confFile struct {
	Read     DbConfig `yaml:"read" json:"read"`
	Timezone string   `yaml:"timezone" json:"timezone"`
}

var dbConfFile confFile

func Init() {
	parseConfig()

	var err error
	Connections.Read, err = open(dbConfFile.Read)
	if err != nil {
		log.Fatal("Failed to initialize Read DB:", err)
	}

	Connections.Write, err = open(dbConfFile.Read)
	if err != nil {
		log.Fatal("Failed to initialize Write DB:", err)
	}

	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt)
		sig := <-signals
		Close(Connections.Read)
		Close(Connections.Write)
		log.Println("Mysql connections aborted:", sig)
	}()
}

func open(conf DbConfig) (*sql.DB, error) {
	dsn := conf.User + ":" + conf.Password + "@tcp(" + conf.Host + ":" + conf.Port + ")/" + conf.Db + "?loc=" + url.QueryEscape(dbConfFile.Timezone) + "&parseTime=true"
	con, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err := con.Ping(); err != nil {
		return nil, err
	}
	con.SetMaxIdleConns(conf.MaxIdleCons)
	con.SetMaxOpenConns(conf.MaxOpenCons)
	log.Println("Mysql connection established:", conf.Host+":"+conf.Port)
	return con, nil
}

func Close(db *sql.DB) {
	if db != nil {
		db.Close()
	}
}

func parseConfig() {
	config.DefaultConfigurator.Load("config/database", &dbConfFile, func(cfg interface{}) {
		conf := cfg.(*confFile)
		if conf.Timezone == "" {
			log.Fatal("timezone cannot be empty")
		}
	})
}
