package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/osalomon89/meli-items-w2/internal/entity"
)

type itemDB struct {
	Id          int
	Code        string
	Title       string
	Description string
	Price       int
	Stock       int
	Status      string
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func (b itemDB) toItemEntity() entity.Item {
	return entity.Item{
		Id:          b.Id,
		Code:        b.Code,
		Title:       b.Title,
		Description: b.Description,
		Price:       b.Price,
		Stock:       b.Stock,
		Status:      b.Status,
		CreatedAt:   b.CreatedAt,
		UpdatedAt:   b.UpdatedAt,
	}
}

type mysqlItemRepository struct {
	conn *sqlx.DB
}

func NewMysqlItemRepository(db *sqlx.DB) entity.ItemRepository {
	return &mysqlItemRepository{
		conn: db,
	}
}

func (r *mysqlItemRepository) AddItem(item *entity.Item) error {

	status := r.ValidateStatus(item.Stock)
	timeRequest := time.Now()

	result, err := r.conn.Exec(`INSERT INTO items
	(code, title, description, price, stock, status, created_at, updated_at)
	VALUES(?, ?, ?, ?, ?, ?, ?, ?)`, item.Code, item.Title, item.Description, item.Price, item.Stock, status, timeRequest, timeRequest)
	if err != nil {
		return fmt.Errorf("error inserting item: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error saving item: %w", err)
	}

	item.Id = int(id)
	item.Status = status
	item.CreatedAt = timeRequest
	item.UpdatedAt = timeRequest

	fmt.Println(result)

	return nil
}

func (r *mysqlItemRepository) UpdateItem(item *entity.Item, id int) (entity.Item, error) {

	status := r.ValidateStatus(item.Stock)
	update := time.Now()

	var itemDataB itemDB

	_, err := r.conn.Exec(`UPDATE items SET code=?, title=?, description=?, price=?, stock=?, status=?, updated_at=? WHERE id=?`, item.Code, item.Title, item.Description, item.Price, item.Stock, status, update, id)
	if err != nil {
		return entity.Item{}, fmt.Errorf("error updating item: %w", err)
	}

	err = r.conn.Get(&itemDataB, `SELECT * FROM items WHERE id=?`, id)
	if err != nil {
		return entity.Item{}, entity.ItemNotFound{
			Message: fmt.Sprintf("Item with id '%d' not found", id),
		}
	}

	return itemDataB.toItemEntity(), nil
}

func (r *mysqlItemRepository) GetItem(id int) (entity.Item, error) {

	var itemDataB itemDB

	err := r.conn.Get(&itemDataB, `SELECT * FROM items WHERE id=?`, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.Item{}, entity.ItemNotFound{
				Message: fmt.Sprintf("Item with id '%d' not found", id),
			}
		}
		return entity.Item{}, fmt.Errorf("error getting item: %w", err)
	}
	return itemDataB.toItemEntity(), nil
}

func (r *mysqlItemRepository) DeleteItem(id int) error {

	var itemDataB itemDB

	err := r.conn.Get(&itemDataB, `SELECT * FROM items WHERE id=?`, id)
	if err != nil {
		return entity.ItemNotFound{
			Message: fmt.Sprintf("Item with id '%d' not found", id),
		}
	}

	_, err = r.conn.Exec(`DELETE FROM items WHERE id=?`, id)
	if err != nil {
		return fmt.Errorf("error deleting item: %w", err)
	}

	return nil

}

func (r *mysqlItemRepository) GetItems(status string, limit int) ([]entity.Item, error) {

	if limit > 20 || limit < 1 {
		limit = 10
	}

	var itemsToshow []entity.Item
	var itemDataB []itemDB
	var err error

	if len(status) != 0 {
		err = r.conn.Select(&itemDataB, `SELECT * FROM items WHERE status=? ORDER BY updated_at DESC LIMIT ?`, status, limit)
	} else {
		err = r.conn.Select(&itemDataB, `SELECT * FROM items ORDER BY updated_at DESC LIMIT ?`, limit)
	}

	if err != nil {
		return []entity.Item{}, fmt.Errorf("error getting items: %w", err)
	}

	for _, v := range itemDataB {
		itemsToshow = append(itemsToshow, v.toItemEntity())
	}

	return itemsToshow, nil

}

func (r *mysqlItemRepository) ValidateCode(code string) (bool, error) {
	var isDuplicated bool
	err := r.conn.Get(&isDuplicated, `SELECT EXISTS(SELECT * FROM items WHERE code=?)`, code)
	if err != nil {
		return isDuplicated, fmt.Errorf("error validating item: %w", err)
	}

	return isDuplicated, nil
}

func (r *mysqlItemRepository) ValidateCodeUpdate(code string, id int) (bool, error) {
	var isDuplicated bool
	err := r.conn.Get(&isDuplicated, `SELECT EXISTS(SELECT * FROM items WHERE code=? AND NOT id=?)`, code, id)
	if err != nil {
		return isDuplicated, fmt.Errorf("error validating item: %w", err)
	}

	return isDuplicated, nil
}

func (r *mysqlItemRepository) ValidateStatus(stock int) string {
	if stock > 0 {
		return "ACTIVE"
	}
	return "INACTIVE"
}

// Query crear tabla
/* CREATE TABLE items (
    id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
	code varchar(200) DEFAULT NULL,
	title longtext DEFAULT NULL,
	description varchar(200) DEFAULT NULL,
	price bigint(20) DEFAULT NULL,
	stock bigint(20) DEFAULT NULL,
    status varchar(10) default NULL,
	created_at datetime(3) DEFAULT NULL,
	updated_at datetime(3) DEFAULT NULL,
	PRIMARY KEY (id),
	UNIQUE KEY code (code)
); */
