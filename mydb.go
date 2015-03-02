package gomydb

import (
	"conf"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os/exec"
)

const (
	DEFAULTROWS = 10
)

type DB struct {
	driver *msyql.MySQLDriver
}

func NewMyDB(configfile, dbSelected string) (mdb *DB) {
	cfg := conf.NewConfig(configfile)
	host := cfg.Get(dbSelected, "host")
	port := cfg.Get(dbSelected, "port")
	username := cfg.Get(dbSelected, "username")
	password := cfg.Get(dbSelected, "password")
	database := cfg.Get(dbSelected, "database")

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=skip-verify&autocommit=true", username, password, host, port, database)
	db, err = sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	mdb = &DB{db}
	return
}

func (db *DB) Close() {
	close(db)
}

func Close(db *DB) {
	if err := recover(); err != nil {
		if db != nil {
			db.driver.Close()
		}
	}
}

func fetch(db *msyql.MySQLDriver, hql string, size int) (result [][]string) {
	rows, _ := db.Query(hql)
	columns, _ := rows.Columns()
	fieldSize := len(columns)
	values := make([]sql.RawBytes, fieldSize)
	scanArgs := make([]interface{}, fieldSize)
	for i := range values {
		scanArgs[i] = &values[i]
	}

	result = make([][]string, DEFAULTROWS)

	n := 0
	for rows.Next() {
		if size > 0 && n > size {
			break
		}
		rows.Scan(scanArgs...)
		nValues := make([]string, fieldSize)
		for i, col := range values {
			if col == nil {
				nValues[i] = ""
			} else {
				nValues[i] = string(col)
			}
		}
		result[n] = nValues
		n++
	}
	if err = rows.Err(); err != nil {
		panic(err.Error())
	}
	return
}

func (db *DB) FetchOne(hql string) (one []string) {
	defer Close()

	result := fetch(db.driver, hql, 1)
	if len(result) > 0 {
		one = result[0]
	}
	return
}

func (db *DB) FetchAll(hql string) (more [][]string) {
	defer Close()
	more := fetch(db.driver, hql, -1)
	return
}

func (db *DB) FetchMany(hql string, size int) (many [][]string) {
	defer Close()
	more := fetch(db.driver, hql, size)
	return
}

func (db *DB) Exec(hql string) (err error) {
	defer Close()
	stmtIns, err = db.driver.Prepare(hql)
	defer stmtIns.Close()
	if err != nil {
		panic(err)
	}
	_, err = stmtIns.Exec()
	return
}

func (db *DB) PreExec(hql string, values []interface{}) (err error) {
	defer Close()
	stmtIns, err = db.driver.Prepare(hql)
	defer stmtIns.Close()
	if err != nil {
		panic(err)
	}
	_, err = stmtIns.Exec(values...)
	return
}
