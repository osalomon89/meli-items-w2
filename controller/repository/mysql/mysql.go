package mysql

import (
	"database/sql"
	ports "github.com/osalomon89/neocamp-meli-w2/domain/ports"
	dom "github.com/osalomon89/neocamp-meli-w2/domain"
)

type itemMysqlRepository struct {
	db *sql.DB
}

func NewItemSqlRepository(db *sql.DB) ports.Repository {
	return &itemMysqlRepository{
		db: db,
	}

}



func (i *itemMysqlRepository) UpdateItem(dom.Item) error{
	i.db.Exec("SELECT * FROM")
	return nil
}