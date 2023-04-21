package mysql

import (
	"time"

	"gigigarino/challengeMELI/internal/domain"
	"gigigarino/challengeMELI/internal/domain/port"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// type itemDTO struct {
// 	id        int
// 	StartTime string
// 	EndTime   string
// }
type itemDAO struct {
	ID        int
	Code      string
	Title     string
	Price     int
	Stock     int
	CreatedAt time.Time `db:"created_at"`
	UpdateAt time.Time `db:"updateAt_at"`
}

func (i itemDAO) toBookDomain() domain.Item {
	return domain.Item{
		ID:        i.ID,
		Code:      i.Code,
		Title:     i.Title,
		Price:     i.Price,
		Stock:     i.Stock,
		CreatedAt: i.CreatedAt,
		UpdateAt: i.UpdateAt,
	}
}
type itemMysqlRepository struct {
	conn *sqlx.DB
}

func NewItemSqlRepository(db *sqlx.DB) port.ItemRepository {
	return &itemMysqlRepository{
		conn: db,
	}
}

func (r *itemMysqlRepository) Index() []domain.Item {
return nil
}
func (r *itemMysqlRepository) GetListaInicial() []domain.Item {
return nil
}
func (r *itemMysqlRepository) GetAllItems() []domain.Item {
	/*query := "INSERT INTO `events` (start_time, end_time) VALUES (?, ?)"

	result, err := r.db.ExecContext(ctx, query, event.StartTime, event.EndTime)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	event.ID = int(id)

	return nil
	*/
	var items []domain.Item
	var itemsDB []itemDAO

	err := r.conn.Select(&itemsDB, "SELECT * FROM books LIMIT 10")
	if err != nil {
		return items
	}

	for _, b := range itemsDB {
		items = append(items, b.toBookDomain())
	}

	return items
}

func (r *itemMysqlRepository) GetItemById(id int) *domain.Item {

	

	return nil
}

func (r *itemMysqlRepository) AddItem(item domain.Item) *domain.Item {
	/*query := "UPDATE events_service SET start_time = ?, end_time = ? WHERE id = ?"

	_, err := r.conn.ExecContext(ctx, query, event.StartTime, event.EndTime, event.ID)
	if err != nil {
		return err
	}
	*/
	return nil
}

func (r *itemMysqlRepository) DeleteItem(id int) bool {
	/*
	query := "DELETE FROM events_service WHERE id = ?"

	_, err := r.conn.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
*/
	return true
}
func (r *itemMysqlRepository) UpdateItemNuevo(domain.Item) error{
 return nil

}
func (r *itemMysqlRepository) UpdateItem(id int) *domain.Item {
	/*
	query := "SELECT id, start_time, end_time FROM `events`"

	rows, err := r.conn.QueryContext(ctx, query)
	if err != nil {
		return nil, errors.New("rows - " + err.Error())
	}
	defer rows.Close()

	events := make([]*domain.Event, 0)

	for rows.Next() {
		var e eventDTO
		if err := rows.Scan(&e.id, &e.StartTime, &e.EndTime); err != nil {
			return nil, errors.New("error de tipo - " + err.Error())
		}

		event := convertEventDtoToDomainEvent(e)
		events = append(events, &event)
	}
*/
	return nil
}
