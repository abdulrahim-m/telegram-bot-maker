package repositories

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/jmoiron/sqlx"
)

type QueryOptions struct {
	Limit  int
	Offset int
	SortBy string
}

type BaseRepository[T any] struct {
	DB    sqlx.DB
	Table string
}

// ======== FETCH ========

func (b *BaseRepository[T]) GetAll() ([]T, error) {
	var item []T
	query := fmt.Sprintf(`SELECT * FROM %s`, b.Table)

	err := b.DB.Select(&item, query)
	return item, err
}

func (b *BaseRepository[T]) FindByID(id int) (*T, error) {
	var item T
	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = ?`, b.Table)

	err := b.DB.Get(&item, query, id)
	return &item, err
}

func (b *BaseRepository[T]) FindByField(field, value string) (*T, error) {
	var item T
	query := fmt.Sprintf(`SELECT * FROM %s WHERE %s = %s`, b.Table, field, value)

	err := b.DB.Get(&item, query)
	return &item, err
}

func (b *BaseRepository[T]) GetWithOpts(opts QueryOptions) ([]T, error) {
	var items []T
	query := fmt.Sprintf(`SELECT * FROM %s ORDER BY %s LIMIT %d OFFSET %d`,
		b.Table, opts.SortBy, opts.Limit, opts.Offset)

	err := b.DB.Select(&items, query)
	return items, err
}

func (b *BaseRepository[T]) GetCount() (int64, error) {
	var count int64
	err := b.DB.Get(&count, `SELECT COUNT(*) FROM ?`, b.Table)
	return count, err
}

// ======== TRANSACTIONS ========

func (b *BaseRepository[T]) Create(item *T) error {
	cols, ph, _ := getColumns(*new(T))

	query := fmt.Sprintf(`INSERT INTO %s (%s) VALUES (%s)`, b.Table, cols, ph)

	_, err := b.DB.NamedExec(query, item)
	return err
}

func (b *BaseRepository[T]) Update(item *T, id int64) error {
	_, _, list := getColumns(*new(T))

	query := fmt.Sprintf(`UPDATE %s SET %s WHERE id = %d`, b.Table, list, id)

	_, err := b.DB.NamedExec(query, item)
	return err
}

// ======== HELPERS ========

func getColumns(s interface{}) (string, string, string) {
	t := reflect.TypeOf(s)
	var columns []string
	var placeholders []string
	var updateString []string

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("db")

		if tag == "" || tag == "id" || tag == "-" {
			continue
		}

		columns = append(columns, tag)
		placeholders = append(placeholders, ":"+tag)
		updateString = append(updateString, tag+" = :"+tag)
	}

	return strings.Join(columns, ", "), strings.Join(placeholders, ", "), strings.Join(updateString, ", ")
}
