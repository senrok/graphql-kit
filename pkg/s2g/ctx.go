package s2g

import (
	"fmt"
	"strings"
)

type Ctx struct {
	ModelName  string
	ModelField []ModelField
}

type ModelField struct {
	Name            string
	FilterOperators []FilterOperator
	Scalar          string
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
			ModelName: c.ModelName,
			FieldName: field.Name,
			Operators: nil,
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
	return fmt.Sprintf(`@goTag(key: "bson",value:"%s")`, strings.Replace(tag, "_", "$", 1))
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
		},
	}
)
