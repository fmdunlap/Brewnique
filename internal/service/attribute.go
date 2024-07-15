package service

import "brewnique.fdunlap.com/internal/data"

type AttributeProvider interface {
	GetAttribute(id int64) (*data.AttributeDefinition, error)
	ListAttributes() ([]*data.AttributeDefinition, error)
}

type AttributeService struct {
	attributeProvider AttributeProvider
}

func NewAttributeService(attributeProvider AttributeProvider) *AttributeService {
	return &AttributeService{
		attributeProvider: attributeProvider,
	}
}

func (s *AttributeService) GetAttribute(id int64) (*data.AttributeDefinition, error) {
	return s.attributeProvider.GetAttribute(id)
}

func (s *AttributeService) ListAttributes() ([]*data.AttributeDefinition, error) {
	return s.attributeProvider.ListAttributes()
}
