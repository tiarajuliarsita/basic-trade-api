package request

type Variant struct {
	VariantName string `json:"variant_name" form:"variant_name" valid:"required~variant~name is required"`
	Quantity    int    `json:"quantity" form:"quantity" valid:"required~entity is required"`
	UUID        string   `json:"product_id" form:"product_id"`
}
