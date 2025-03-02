package app

import (
	"github.com/opensourceways/xihe-server/domain"
	"github.com/opensourceways/xihe-server/domain/repository"
)

type TagsService interface {
	List([]string) ([]DomainTagsDTO, error)
}

func NewTagsService(repo repository.Tags) TagsService {
	return tagsService{repo}
}

type tagsService struct {
	repo repository.Tags
}

type DomainTagsDTO = domain.DomainTags

func (s tagsService) List(domainNames []string) ([]DomainTagsDTO, error) {
	return s.repo.List(domainNames)
}
