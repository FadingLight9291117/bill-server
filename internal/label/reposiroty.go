package label

import (
	"context"
	"gorm.io/gorm"
)

type sqliteRepository struct {
	*gorm.DB
}

func NewLabelRepository(db *gorm.DB) LabelRepository {
	return &sqliteRepository{
		db,
	}
}

func (s *sqliteRepository) GetLabels(ctx context.Context) ([]Label, error) {
	var labels []Label
	res := s.Find(&labels)
	if res.Error != nil {
		return nil, res.Error
	}
	return labels, nil
}

func (s *sqliteRepository) GetLabelById(ctx context.Context, id int) (*Label, error) {
	var label Label
	res := s.First(&label, id)
	if res.Error != nil {
		return nil, res.Error
	}

	return &label, nil
}

func (s *sqliteRepository) CreateLabel(ctx context.Context, label *Label) error {
	return s.Create(label).Error
}

func (s *sqliteRepository) UpdateLabel(ctx context.Context, label *Label) error {
	return s.Save(label).Error
}

func (s *sqliteRepository) DeleteLabel(ctx context.Context, id int) error {
	return s.Delete(&Label{}, id).Error
}
