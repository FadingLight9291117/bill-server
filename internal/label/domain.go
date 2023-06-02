package label

import (
	"context"
	"gorm.io/gorm"
)

type Label struct {
	gorm.Model
	Type  string
	Name  string
	Count int
}

type LabelDTO struct {
	ID    int    `json:"id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

func (l *Label) Dto() interface{} {
	return LabelDTO{
		ID:    int(l.ID),
		Type:  l.Type,
		Name:  l.Name,
		Count: l.Count,
	}
}

type LabelRepository interface {
	GetLabels(ctx context.Context) ([]Label, error)
	GetLabelById(ctx context.Context, id int) (*Label, error)
	CreateLabel(ctx context.Context, label *Label) error
	UpdateLabel(ctx context.Context, label *Label) error
	DeleteLabel(ctx context.Context, id int) error
}

type LabelService interface {
	GetLabels(ctx context.Context) ([]interface{}, error)
	GetLabelById(ctx context.Context, id int) (interface{}, error)
	CreateLabel(ctx context.Context, label *Label) error
	UpdateLabel(ctx context.Context, label *Label) error
	DeleteLabel(ctx context.Context, id int) error
}
