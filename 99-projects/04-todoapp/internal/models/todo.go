package models

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"runtime/debug"
	"time"
)

type Todo struct {
	Id          int64     `json:"id"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	ModifiedAt  time.Time `json:"modified_at"`
}

func (t *Todo) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(t)
}

func (t *Todo) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(t)
}

func (t *Todo) GetAll(page int64, per_page int64) ([]*Todo, error) {
	var todos []*Todo

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	stmt := `select * from todos LIMIT $1 OFFSET $2`

	rows, err := db.QueryContext(ctx, stmt, per_page, (page-1)*per_page)

	if err != nil {
		return todos, err
	}

	defer rows.Close()

	for rows.Next() {
		var todo Todo
		err := rows.Scan(
			&todo.Id,
			&todo.Title,
			&todo.Description,
			&todo.CreatedAt,
			&todo.ModifiedAt,
		)

		if err != nil {
			return todos, err
		}

		todos = append(todos, &todo)
	}

	return todos, nil

}

func (t *Todo) GetOne(id int) (*Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	todo := Todo{}

	stmt := `select * from todos where id=$1`

	err := db.QueryRowContext(ctx, stmt, id).Scan(&todo.Id, &todo.Title, &todo.Description, &todo.CreatedAt, &todo.ModifiedAt)

	if err != nil {
		err = fmt.Errorf("%w\n%s", err, debug.Stack())
		return &todo, err
	}

	return &todo, nil
}

func (t *Todo) Update(id int, todo Todo) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `update todos set title=$2, description=$3, modified_at=$4 where id=$1 `

	err := db.QueryRowContext(ctx, stmt, id,
		todo.Title, todo.Description, todo.ModifiedAt).Err()

	if err != nil {
		err = fmt.Errorf("%w\n%s", err, debug.Stack())
		return err
	}

	return nil
}

func (t *Todo) DeleteByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from todos where id=$1`

	result, err := db.ExecContext(ctx, stmt, id)

	if err != nil {
		return err
	}

	num, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if num != 1 {
		return errors.New("RowsAffected != 1")
	}

	return nil
}

func (t *Todo) Insert(todo Todo) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var newID int
	stmt := `insert into todos (title, description, created_at, modified_at)
		values ($1, $2, $3, $4) returning id`

	err := db.QueryRowContext(ctx, stmt,
		todo.Title,
		todo.Description,
		todo.CreatedAt,
		todo.ModifiedAt,
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}
