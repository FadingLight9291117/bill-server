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
