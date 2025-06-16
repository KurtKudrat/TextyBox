package mysql

import (
	"database/sql"

	"github.com/KurtKudrat/TextyBox/pkg/models"
)

type TextyModel struct {
	DB *sql.DB
}

func (m *TextyModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

func (m *TextyModel) Get(id int) (*models.Texty, error) {
	return nil, nil
}

func (m *TextyModel) Latest() ([]*models.Texty, error) {
	return nil, nil
}
