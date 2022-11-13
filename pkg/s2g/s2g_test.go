package s2g

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

type DefaultModel struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" odm:"primaryID" gqlgen:"id" filter:"comparable" filter:"comparable" filterScalar:"ID" sort:""`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty" odm:"autoCreateTime" gqlgen:"createdAt" filter:"range" filterScalar:"Time" sort:""`
	UpdatedAt time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty" odm:"autoUpdateTime" gqlgen:"updatedAt" filter:"range" filterScalar:"Time" sort:""`
	DeletedAt time.Time          `json:"deletedAt,omitempty" bson:"deletedAt,omitempty" odm:"deleteTime" gqlgen:"deletedAt" filter:"range" filterScalar:"Time" sort:""`
}

type Tenant struct {
	DefaultModel `s2g:"nested"`
	Name         string `bson:"name" filter:"comparable" filterScalar:"String" sort:""`
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
