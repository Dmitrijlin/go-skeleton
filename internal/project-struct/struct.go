package projectstruct

import "github.com/Dmitrijlin/go-skeleton/internal/generator/tag"

type ProjectStruct struct {
	Name     string          `json:"name"`
	Type     EntityTypeEnum  `json:"type"`
	Children []ProjectStruct `json:"children"`
	Tag      tag.Tag         `json:"tag"`
}

type EntityTypeEnum string

const (
	Dir  = EntityTypeEnum("dir")
	File = EntityTypeEnum("file")
)
