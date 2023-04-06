package models

type TodoInterface interface {
	GetAll(page int64, per_page int64) ([]*Todo, error)
	GetOne(id int) (*Todo, error)
	Update(id int, todo Todo) error
	DeleteByID(id int) error
	Insert(todo Todo) (int, error)
}
