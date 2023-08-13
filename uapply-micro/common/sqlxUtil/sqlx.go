package sqlxUtil

import mysqlx "github.com/jmoiron/sqlx"

func InitMySqlx(dns string) *mysqlx.DB {
	connect, _ := mysqlx.Connect("mysql", dns)
	return connect
}
