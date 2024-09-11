package bins

import (
	"errors"
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

func NewBin(id string, private bool, name string) (*Bin, error) {
	if id == "" {
		return nil, errors.New("Пустой ID")
	}
	return &Bin{
		id:        id,
		private:   private,
		createdAt: time.Now(),
		name:      name,
	}, nil
}

func newBinList(length int) []*Bin {
	return make([]*Bin, length)
}
