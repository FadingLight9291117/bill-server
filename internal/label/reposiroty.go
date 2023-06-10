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
	if err := s.DB.Find(&labels).Error; err != nil {
		return nil, err
	}
	return labels, nil
}

func (s *sqliteRepository) GetLabelById(ctx context.Context, id int) (*Label, error) {
	var label Label
	if err := s.DB.Where("id =?", id).First(&label).Error; err != nil {
		return nil, err
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
	return s.Where("id =?", id).Delete(&Label{}).Error
}
