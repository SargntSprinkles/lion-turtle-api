package db

import (
	"database/sql"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/golang-migrate/migrate/v4"
	migratepg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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
		logrus.Error("could not connect to db")
		panic(connErr)
	}
	pingErr := conn.Ping()
	if pingErr != nil {
		logrus.Fatalf("could not ping db - %s", pingErr.Error())
	}
	conn.SetConnMaxLifetime(time.Minute * 5)
	conn.SetMaxIdleConns(100)
	conn.SetMaxOpenConns(100)

	gormDB, gormErr := gorm.Open(gormpg.New(gormpg.Config{Conn: conn}), &gorm.Config{})
	if gormErr != nil {
		logrus.Fatalf("could not convert db conn to gorm - %s", gormErr.Error())
	}

	p.connection = gormDB
	p.connected = true
	logrus.Info("connected to db")

	logrus.Info("migrating up")
	driver, driverErr := migratepg.WithInstance(conn, &migratepg.Config{})
	if driverErr != nil {
		logrus.Fatalf("could not create migration driver - %s", driverErr.Error())
	}
	m, mErr := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if mErr != nil {
		logrus.Fatalf("could not instantiate Migrate object - %s", mErr.Error())
	}
	if upErr := m.Up(); upErr != nil {
		switch upErr {
		case migrate.ErrNoChange:
			logrus.Info("no change to schema")
		default:
			logrus.Fatalf("could not migrate up - %s", upErr.Error())
		}
	}
	logrus.Info("finished migrating up")

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
