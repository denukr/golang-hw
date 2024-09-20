package bins

import (
	"errors"
	"time"
)

type Bin struct {
	ID        string    `json: "id"`
	Private   bool      `json: "private"`
	CreatedAt time.Time `json: "createdAt"`
	Name      string    `json: "name"`
}

func NewBin(id string, private bool, name string) (*Bin, error) {
	if id == "" {
		return nil, errors.New("Пустой ID")
	}
	newBin := &Bin{
		ID:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
	return newBin, nil
}
