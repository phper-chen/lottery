package dao

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
)
type config struct {
	Admin string
	Pwd string
	Host string
	Port string
	DataBase string
	MaxIdelConns int
	MaxOpenConns int
}
type Dao struct {
	db *sql.DB
	Logger *log.Logger
	RdPool *redis.Pool
}


func New() *Dao {
	d := &Dao{
		db: newMysql(&config{
			Admin:        "root",
			Pwd:		  "root",
			Host:         "127.0.0.1",
			Port:         "3306",
			DataBase:     "test_db",
			MaxIdelConns: 20,
			MaxOpenConns: 20,
		}),
	}
	d.initLog()
	d.RdPool = NewRedisPool()
	return d
}

func newMysql(c *config) (db *sql.DB) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", c.Admin, c.Pwd, c.Host, c.Port, c.DataBase))
	if err != nil{
		log.Fatalln(err)
	}
	db.SetMaxIdleConns(c.MaxIdelConns)
	db.SetMaxOpenConns(c.MaxOpenConns)
	checkPanicErr(db.Ping())
	return
}

func (d *Dao) Close() {
	_ = d.db.Close()
	_ = d.RdPool.Close()
}

// 初始化日志
func (d *Dao) initLog() {
	currentPa, _ := os.Getwd()
	f, err := os.Create(currentPa + "/lottery.log")
	checkPanicErr(err)
	d.Logger = log.New(f, "", log.Ldate|log.Lmicroseconds)
}

func checkPanicErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
