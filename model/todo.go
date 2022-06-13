package model

type Todo struct {
	ID   int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Task string `json:"task" validate:"required"`
}

func (Todo) TableName() string {
	return "todos"
}
