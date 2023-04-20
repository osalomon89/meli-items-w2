package mysql_infrastructure

import "gomysql/db"

func main() {
	db.connect()
	db.close()
}
