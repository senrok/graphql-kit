package s2g

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"text/template"
)

var (
	tc = FieldFilterTplCtx{
		ModelName: "Tenant",
		FieldName: "Name",
		Operators: []Operator{
			{
				Name:       "_eq",
				ScalarName: "String",
				CustomTags: `@goTag(key: "bson",value:"$eq")`,
			},
		},
	}
	mc = ModelFilterTplCtx{
		ModelName: "Tenant",
		Fields: []FieldFilterTplCtx{
			{
				ModelName: "Tenant",
				FieldName: "Name",
				Operators: []Operator{
					{
						Name:       "_eq",
						ScalarName: "String",
						CustomTags: `@goTag(key: "bson",value:"$eq")`,
					},
				},
			},
		},
	}
)

func TestGenFieldFilterTpl(t *testing.T) {
	tpl, err := template.ParseFiles("./field_filter.tpl")
	assert.Nil(t, err)
	sb := strings.Builder{}
	err = tpl.Execute(&sb, tc)
	assert.Nil(t, err)
	fmt.Println(sb.String())
}

func TestGenModelFilterTpl(t *testing.T) {
	tpl, err := template.ParseFiles("./model_filter.tpl")
	assert.Nil(t, err)
	sb := strings.Builder{}
	err = tpl.Execute(&sb, mc)
	assert.Nil(t, err)
	fmt.Println(sb.String())
}
