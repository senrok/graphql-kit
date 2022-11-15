package s2g

import (
	"fmt"
	"reflect"
	"strings"
)

type Ctx struct {
	ModelName  string
	ModelField []ModelField
}

type ModelField struct {
	reflect.StructField
	Name            string
	FilterOperators []FilterOperator
	Scalar          string
	SortScalar      string
}

type FilterOperator struct {
	Name            string
	ScalarGenerator ScalarGenerator
	TagGenerator    TagGenerator
}

func (c Ctx) ToTplCtx() ModelFilterTplCtx {
	tc := ModelFilterTplCtx{
		ModelName: c.ModelName,
		Fields:    nil,
	}
	for _, field := range c.ModelField {
		ftc := FieldFilterTplCtx{
			ModelName:      c.ModelName,
			FieldName:      field.Name,
			SortScalar:     "SortDirection",
			CustomTags:     ExportOriginBsonTag(field),
			OriginBsonName: GetOriginBsonName(field),
			Operators:      nil,
		}
		for _, operator := range field.FilterOperators {
			o := Operator{
				Name:       operator.Name,
				ScalarName: operator.ScalarGenerator(field.Scalar),
				CustomTags: operator.TagGenerator(operator.Name),
			}
			ftc.Operators = append(ftc.Operators, o)
		}
		tc.Fields = append(tc.Fields, ftc)
	}
	return tc
}

type TagGenerator func(name string) string

func DefaultTagGenerator(tag string) string {
	return fmt.Sprintf(`@goTag(key: "bson",value:"%s,omitempty")`, strings.Replace(tag, "_", "$", 1))
}

func GetOriginBsonName(f ModelField) string {
	if value, ok := f.Tag.Lookup("bson"); ok {
		return strings.Split(value, ",")[0]
	}
	return ""
}
func ExportOriginBsonTag(f ModelField) string {
	if value, ok := f.Tag.Lookup("bson"); ok {
		return fmt.Sprintf(`@goTag(key: "bson",value:"%s,omitempty")`, strings.Split(value, ",")[0])
	}
	return ""
}

type ScalarGenerator func(scalar string) string

func DefaultGenerator(scalar string) string {
	return scalar
}

func NullableListGenerator(scalar string) string {
	return fmt.Sprintf(`[%s!]`, scalar)
}

var (
	FilterMap = map[string][]FilterOperator{
		"eq": []FilterOperator{
			{
				Name:            "_eq",
				ScalarGenerator: DefaultGenerator,
				TagGenerator:    DefaultTagGenerator,
			},
		},
		"elemMatch": []FilterOperator{
			{
				Name:            "_elemMatch",
				ScalarGenerator: DefaultGenerator,
				TagGenerator:    DefaultTagGenerator,
			},
		},
		"comparable": []FilterOperator{
			{
				Name:            "_eq",
				ScalarGenerator: DefaultGenerator,
				TagGenerator:    DefaultTagGenerator,
			},
			{
				Name:            "_gt",
				ScalarGenerator: DefaultGenerator,
				TagGenerator:    DefaultTagGenerator,
			},
			{
				Name:            "_gte",
				ScalarGenerator: DefaultGenerator,
				TagGenerator:    DefaultTagGenerator,
			},
			{
				Name:            "_lt",
				ScalarGenerator: DefaultGenerator,
				TagGenerator:    DefaultTagGenerator,
			},
			{
				Name:            "_lte",
				ScalarGenerator: DefaultGenerator,
				TagGenerator:    DefaultTagGenerator,
			},
			{
				Name:            "_in",
				ScalarGenerator: NullableListGenerator,
				TagGenerator:    DefaultTagGenerator,
			},
			{
				Name:            "_regex",
				ScalarGenerator: DefaultGenerator,
				TagGenerator:    DefaultTagGenerator,
			},
		},
		"range": []FilterOperator{
			{
				Name:            "_gt",
				ScalarGenerator: DefaultGenerator,
				TagGenerator:    DefaultTagGenerator,
			},
			{
				Name:            "_gte",
				ScalarGenerator: DefaultGenerator,
				TagGenerator:    DefaultTagGenerator,
			},
			{
				Name:            "_lt",
				ScalarGenerator: DefaultGenerator,
				TagGenerator:    DefaultTagGenerator,
			},
			{
				Name:            "_lte",
				ScalarGenerator: DefaultGenerator,
				TagGenerator:    DefaultTagGenerator,
			},
		},
	}
)
