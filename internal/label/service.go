package label

import (
	"context"
	"github.com/samber/lo"
)

type labelService struct {
	labelRepository LabelRepository
}

func NewLabelService(labelRepository LabelRepository) LabelService {
	return &labelService{labelRepository: labelRepository}
}

func (l *labelService) GetLabels(ctx context.Context) ([]interface{}, error) {
	labels, err := l.labelRepository.GetLabels(ctx)

	return lo.Map(labels, func(l Label, _ int) interface{} { return l.Dto() }), err
}

func (l *labelService) GetLabelById(ctx context.Context, id int) (interface{}, error) {
	label, err := l.labelRepository.GetLabelById(ctx, id)
	if err != nil {
		return nil, err
	}

	return label.Dto(), nil
}

func (l *labelService) CreateLabel(ctx context.Context, label *Label) error {
	return l.labelRepository.CreateLabel(ctx, label)
}

func (l *labelService) UpdateLabel(ctx context.Context, label *Label) error {
	return l.labelRepository.UpdateLabel(ctx, label)
}

func (l *labelService) DeleteLabel(ctx context.Context, id int) error {
	return l.labelRepository.DeleteLabel(ctx, id)
}
