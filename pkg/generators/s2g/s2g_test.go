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
	Name         string `bson:"name,omitempty" filter:"comparable" filterScalar:"String" sort:""`
}

type Device struct {
	DefaultModel `bson:",inline" s2g:"nested"`
	DeviceType   string             `bson:"type" filter:"comparable" filterScalar:"DeviceType" sort:""`
	Protocol     string             `bson:"protocol" filter:"comparable" filterScalar:"DeviceProtocol" sort:""`
	Secret       string             `bson:"secret"`
	Activated    bool               `bson:"activated"  filter:"comparable" filterScalar:"Boolean"`
	Merchant     primitive.ObjectID `bson:"merchant,omitempty" filter:"comparable" filterScalar:"ID" sort:""`
	Store        primitive.ObjectID `bson:"store,omitempty" filter:"comparable" filterScalar:"ID" sort:""`
}

type Account struct {
	DefaultModel `bson:",inline" s2g:"nested"`
	Type         string `bson:"type" filter:"comparable" filterScalar:"AccountType" sort:""`
	Suspended    bool   `bson:"suspended" filter:"eq" filterScalar:"Boolean" sort:""`

	// balances
	Total          string `bson:"total"`
	ConsumedAmount string `bson:"consumedAmount"`
	InvoiceAmount  string `bson:"invoiceAmount"`

	Name   string             `bson:"name" filter:"comparable" filterScalar:"String" sort:""`
	Tenant primitive.ObjectID `bson:"tenant"`
}

type Merchant struct {
	DefaultModel `bson:",inline" s2g:"nested"`
	Name         string `bson:"name" filter:"comparable" filterScalar:"String" sort:""`
	Activated    bool   `bson:"activated" filter:"eq" filterScalar:"Boolean"`
}

type Store struct {
	DefaultModel `bson:",inline" s2g:"nested"`
	Name         string             `bson:"name"  filter:"comparable" filterScalar:"String" sort:""`
	Tenant       primitive.ObjectID `bson:"tenant"  filter:"comparable" filterScalar:"ID" sort:""`
	Merchant     primitive.ObjectID `bson:"merchant"  filter:"comparable" filterScalar:"ID" sort:""`
}

type FullProfile struct {
	DefaultModel   `bson:",inline" s2g:"nested"`
	Type           string             `bson:"type" filter:"comparable" filterScalar:"ProfileType" sort:""`
	Tenant         primitive.ObjectID `bson:"tenant" filter:"comparable" filterScalar:"ID" sort:""`
	Account        primitive.ObjectID `bson:"account" filter:"comparable" filterScalar:"ID" sort:""`
	Identification primitive.ObjectID `bson:"identification" filter:"comparable" filterScalar:"ID" sort:""`
	Card           primitive.ObjectID `bson:"card" filter:"comparable" filterScalar:"ID" sort:""`
	User           primitive.ObjectID `bson:"user" filter:"comparable" filterScalar:"ID" sort:""`
	MasterBook     primitive.ObjectID `bson:"masterBook" filter:"comparable" filterScalar:"ID" sort:""`
	Store          primitive.ObjectID `bson:"store" filter:"comparable" filterScalar:"ID" sort:""`
	Merchant       primitive.ObjectID `bson:"merchant" filter:"comparable" filterScalar:"ID" sort:""`
}

type Identification struct {
	DefaultModel `bson:",inline" s2g:"nested"`
	Name         string             `bson:"name" filter:"comparable" filterScalar:"String" sort:""`
	Role         string             `bson:"role" filter:"comparable" filterScalar:"RoleType" sort:""`
	UniqueID     string             `bson:"uniqueID,omitempty" filter:"comparable" filterScalar:"String" sort:""`
	IdNo         string             `bson:"idNo,omitempty" filter:"comparable" filterScalar:"String" sort:""`
	Phone        string             `bson:"phone,omitempty" filter:"comparable" filterScalar:"String" sort:""`
	Birthdate    time.Time          `bson:"birthdate,omitempty" filter:"comparable" filterScalar:"Date" sort:""`
	Gender       string             `bson:"gender,omitempty" filter:"comparable" filterScalar:"GenderType" sort:""`
	Tenant       primitive.ObjectID `bson:"tenant" filter:"comparable" filterScalar:"ID" sort:""`
}

type Card struct {
	DefaultModel `bson:",inline" s2g:"nested"`
	SerialNo     string             `bson:"serialNo" filter:"comparable" filterScalar:"String" sort:""`
	CardNo       string             `bson:"cardNo" filter:"comparable" filterScalar:"String" sort:""`
	Protocol     string             `bson:"protocol" filter:"comparable" filterScalar:"CardProtocol" sort:""`
	LastUsedAt   time.Time          `bson:"lastUsedAt,omitempty" filter:"comparable" filterScalar:"Time" sort:""`
	Activated    bool               `bson:"activated" filter:"comparable" filterScalar:"Boolean" sort:""`
	Profile      primitive.ObjectID `bson:"profile,omitempty" filter:"comparable" filterScalar:"ID" sort:""`
}

type Masterbook struct {
	DefaultModel `bson:",inline" s2g:"nested"`
	Tenant       primitive.ObjectID `bson:"tenant" filter:"eq" filterScalar:"ID" sort:""`
	Income       string             `bson:"income" filter:"comparable" filterScalar:"Decimal" sort:""`
	Outcome      string             `bson:"outcome" filter:"comparable" filterScalar:"Decimal" sort:""`
}

type Order struct {
	DefaultModel `bson:",inline" s2g:"nested"`
	Type         string `bson:"type" filter:"comparable" filterScalar:"OrderType" sort:""`
	State        string `bson:"state"  filter:"comparable" filterScalar:"OrderState" sort:""`

	Transactions []primitive.ObjectID `bson:"transactions,omitempty"`

	Payment     string `bson:"paymentMethod" filter:"comparable" filterScalar:"PaymentMethod" sort:""`
	Amount      string `bson:"amount" filter:"comparable" filterScalar:"Decimal" sort:""`
	Description string `bson:"description"`

	PayerProfile    primitive.ObjectID `bson:"payerProfile,omitempty" filter:"comparable" filterScalar:"ID" sort:""`
	ReceiverProfile primitive.ObjectID `bson:"receiverProfile,omitempty" filter:"comparable" filterScalar:"ID" sort:""`
	Merchant        primitive.ObjectID `bson:"merchant" filter:"comparable" filterScalar:"ID" sort:""`
	Store           primitive.ObjectID `bson:"store" filter:"comparable" filterScalar:"ID" sort:""`
	Device          primitive.ObjectID `bson:"device,omitempty" filter:"comparable" filterScalar:"ID" sort:""`

	FinishedAt time.Time `bson:"finishedAt,omitempty" bson:"store" filter:"comparable" filterScalar:"Time" sort:""`
}

type Charge struct {
	DefaultModel `bson:",inline" s2g:"nested"`

	Operator primitive.ObjectID `bson:"operator" filter:"comparable" filterScalar:"ID" sort:""`
	Receiver primitive.ObjectID `bson:"receiver" filter:"comparable" filterScalar:"ID" sort:""`
	Tenant   primitive.ObjectID `bson:"tenant" filter:"comparable" filterScalar:"ID" sort:""`

	State       string `bson:"state" filter:"comparable" filterScalar:"ChargeState" sort:""`
	Type        string `bson:"type" filter:"comparable" filterScalar:"ChargeType" sort:""`
	Amount      string `bson:"amount" filter:"comparable" filterScalar:"Decimal" sort:""`
	Sign        int    `bson:"sign"`
	Description string `bson:"description"`

	Withdrawal bool `bson:"withdrawal" filter:"comparable" filterScalar:"Boolean" sort:""`

	Transaction  primitive.ObjectID `bson:"transaction" filter:"comparable" filterScalar:"ID" sort:""`
	Confirmation primitive.ObjectID `bson:"confirmation" filter:"comparable" filterScalar:"ID" sort:""`
}

type Confirmation struct {
	DefaultModel `bson:",inline" s2g:"nested"`
	Tenant       primitive.ObjectID `bson:"tenant" filter:"comparable" filterScalar:"ID" sort:""`
	Masterbook   primitive.ObjectID `bson:"masterbook" filter:"comparable" filterScalar:"ID" sort:""`

	Amount string `bson:"amount" filter:"comparable" filterScalar:"Decimal" sort:""`
	Sign   int    `bson:"sign"`
	State  string `bson:"state" filter:"comparable" filterScalar:"ConfirmationState" sort:""`
	Type   string `bson:"type" filter:"comparable" filterScalar:"ConfirmationType" sort:""`
	// charge receiver's transaction
	Transaction primitive.ObjectID `bson:"transaction" filter:"comparable" filterScalar:"ID" sort:""`

	// Profile
	ConfirmedBy primitive.ObjectID `bson:"confirmedBy,omitempty" filter:"comparable" filterScalar:"ID" sort:""`

	Withdrawal bool `bson:"withdrawal" filter:"comparable" filterScalar:"Boolean" sort:""`
}

type Transaction struct {
	DefaultModel `bson:",inline" s2g:"nested"`
	Order        primitive.ObjectID `bson:"parent,omitempty" filter:"comparable" filterScalar:"ID" sort:""`
	Amount       string             `bson:"amount" filter:"comparable" filterScalar:"Decimal" sort:""`
	Sign         int                `bson:"sign"`
	Type         string             `bson:"type" filter:"comparable" filterScalar:"TransactionType" sort:""`
	State        string             `bson:"state" filter:"comparable" filterScalar:"TxnState" sort:""`

	// belong to
	Account         primitive.ObjectID `bson:"account" filter:"comparable" filterScalar:"ID" sort:""`
	PayerAccount    primitive.ObjectID `bson:"payerAccount,omitempty" filter:"comparable" filterScalar:"ID" sort:""`
	ReceiverAccount primitive.ObjectID `bson:"receiverAccount,omitempty" filter:"comparable" filterScalar:"ID" sort:""`

	// for charges, incomes, whether amount able to withdrawal
	Withdrawal bool `bson:"withdrawal"`
	// confirmation, for cash charge, cash withdrawal transaction
	Confirmation primitive.ObjectID `bson:"confirmation,omitempty" filter:"comparable" filterScalar:"ID" sort:""`
}

func TestNewGenerator(t *testing.T) {
	g := NewGenerator(&Card{})
	s, err := g.Generate()
	assert.Nil(t, err)
	fmt.Println(s)
}

func TestNewBatchGenerator(t *testing.T) {
	err := NewBatchGenerator("tmp/", Config{
		Model: &Card{},
	})
	assert.Nil(t, err, err)
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
