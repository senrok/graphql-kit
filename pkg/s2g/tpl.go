package s2g

import "github.com/iancoleman/strcase"

type Operator struct {
	Name       string
	ScalarName string
	CustomTags string
}

type FieldFilterTplCtx struct {
	ModelName  string
	FieldName  string
	SortScalar string
	Operators  []Operator
}

func (f FieldFilterTplCtx) LowerCamelFieldName() string {
	return strcase.ToLowerCamel(f.FieldName)
}

type ModelFilterTplCtx struct {
	ModelName string
	Fields    []FieldFilterTplCtx
}
