package projectstruct

type ProjectStruct struct {
	Name     string          `json:"name"`
	Type     EntityTypeEnum  `json:"type"`
	Children []ProjectStruct `json:"children"`
	Tag      string          `json:"tag"`
}

type EntityTypeEnum string

const (
	Dir  = EntityTypeEnum("dir")
	File = EntityTypeEnum("file")
)
