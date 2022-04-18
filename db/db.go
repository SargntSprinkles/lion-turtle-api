package db

import (
	"database/sql"
	"fmt"
	"os"
	"sync"
	"time"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	gormpg "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var pgdb *postgres

func PGDB() *postgres {
	pgdb = &postgres{
		mutex:     &sync.Mutex{},
		connected: false,
		url:       os.Getenv("DATABASE_URL"),
		host:      os.Getenv("DATABASE_HOST"),
		port:      os.Getenv("DATABASE_PORT"),
		user:      os.Getenv("DATABASE_USER"),
		password:  os.Getenv("DATABASE_PASS"),
		db:        os.Getenv("DATABASE_DB"),
	}
	return pgdb.DB()
}

type postgres struct {
	connection *gorm.DB
	mutex      *sync.Mutex
	connected  bool
	url        string
	host       string
	port       string
	user       string
	password   string
	db         string
}

func (p *postgres) ToString() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s connected=%v", p.host, p.port, p.user, p.db, p.connected)
}

func (p *postgres) DB() *postgres {
	if p.connected {
		return p
	}

	p.mutex.Lock()
	defer p.mutex.Unlock()

	// sanity check in case multiple processes try to connect at the same time
	if p.connected {
		return p
	}

	uri := p.url
	if uri == "" {
		uri = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", p.host, p.port, p.user, p.password, p.db)
	}

	conn, connErr := sql.Open("postgres", uri)
	if connErr != nil {
		logrus.Errorf("could not connect to db")
		panic(connErr)
	}
	pingErr := conn.Ping()
	if pingErr != nil {
		logrus.Errorf("could not ping db")
		panic(pingErr)
	}
	conn.SetConnMaxLifetime(time.Minute * 5)
	conn.SetMaxIdleConns(100)
	conn.SetMaxOpenConns(100)

	gormDB, gormErr := gorm.Open(gormpg.New(gormpg.Config{Conn: conn}), &gorm.Config{})
	if gormErr != nil {
		logrus.Errorf("could not convert db conn to gorm")
		panic(gormErr)
	}

	p.connection = gormDB
	p.connected = true
	logrus.Info("connected to db")
	return p
}

func (p *postgres) Disconnect() {
	if !p.connected {
		logrus.Info("cannot disconnect db - already disconnected")
		return
	}

	conn, connErr := p.connection.DB()
	if connErr != nil {
		logrus.Errorf("could not get generic db from gorm")
	}
	conn.Close()
	p.connected = false
	logrus.Info("db disconnected")
}
