package s2g

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Tenant struct {
	Name string `bson:"name" filter:"comparable" filterScalar:"String"`
}

// input TenantNameFilter {
//    _eq: String
//    _gt: ID
//    _lt: ID
//    _gte: ID
//    _lte: ID
//    _in: [ID!]
// }
//
// input TenantFilter {
//    name: TenantNameFilter
// }

func TestNewGenerator(t *testing.T) {
	g := NewGenerator(&Tenant{})
	s, err := g.Generate()
	assert.Nil(t, err)
	fmt.Println(s)
}
