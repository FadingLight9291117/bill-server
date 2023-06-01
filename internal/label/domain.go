package label

import (
	"context"
	"gorm.io/gorm"
)

type Label struct {
	gorm.Model
	ID         int    `json:"id"`
	Type       string `json:"type"`
	Name       string `json:"name"`
	RelativeId int    `json:"relativeId"`
	Count      int    `json:"count"`
}

type LabelRepository interface {
	GetLabels(ctx context.Context) ([]Label, error)
	GetLabelById(ctx context.Context, id int) (*Label, error)
	CreateLabel(ctx context.Context, label *Label) error
	UpdateLabel(ctx context.Context, label *Label) error
	DeleteLabel(ctx context.Context, id int) error
}

type LabelService interface {
	GetLabels(ctx context.Context) ([]Label, error)
	GetLabelById(ctx context.Context, id int) (*Label, error)
	CreateLabel(ctx context.Context, label *Label) error
	UpdateLabel(ctx context.Context, label *Label) error
	DeleteLabel(ctx context.Context, id int) error
}
