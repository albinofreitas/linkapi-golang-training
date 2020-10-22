package orders

import "go.mongodb.org/mongo-driver/bson/primitive"

// Item from order
type Item struct {
	Sku          string `json:"sku" bson:"sku"`
	SourceItemID string `json:"sourceItemId" bson:"sourceItemId"`
	Price        int32  `json:"price" bson:"price"`
}

// Shipment from order
type Shipment struct {
	Name        string `json:"name" bson:"name"`
	CompanyName string `json:"companyName" bson:"companyName"`
	Address     string `json:"address" bson:"address"`
	Town        string `json:"town" bson:"town"`
	PostCode    string `json:"postCode" bson:"postCode"`
	IsoCountry  string `json:"isoCountry" bson:"isoCountry"`
}

// OrderStatus from order
var OrderStatus = struct {
	PENDING    string
	INTEGRATED string
	FAILED     string
}{
	PENDING:    "pending",
	INTEGRATED: "integrated",
	FAILED:     "failed",
}

// Order struct
type Order struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	SourceOrderID string             `json:"sourceOrderId" bson:"sourceOrderId"`
	Items         []Item             `json:"items" bson:"items"`
	Shipments     []Shipment         `json:"shipments" bson:"shipments"`
	Status        string             `json:"status" bson:"status"`
}
