package label

import (
	"context"
)

type labelService struct {
	labelRepository LabelRepository
}

func NewLabelService(labelRepository LabelRepository) LabelService {
	return &labelService{labelRepository: labelRepository}
}

func (l *labelService) GetLabels(ctx context.Context) ([]Label, error) {
	return l.labelRepository.GetLabels(ctx)
}

func (l *labelService) GetLabelById(ctx context.Context, id int) (*Label, error) {
	//TODO implement me
	panic("implement me")
}

func (l *labelService) CreateLabel(ctx context.Context, label *Label) error {
	//TODO implement me
	panic("implement me")
}

func (l *labelService) UpdateLabel(ctx context.Context, label *Label) error {
	//TODO implement me
	panic("implement me")
}

func (l *labelService) DeleteLabel(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}
