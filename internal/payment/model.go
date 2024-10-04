package paymentmethods

type PaymentMethod struct {
	Id    int    `json:"id"`
	Image string `json:"image"`
}
type PaymentMethodResponse struct {
	PaymentMethods []PaymentMethod `json:"payments"`
}

var paymentmMthods = PaymentMethodResponse{
	PaymentMethods: []PaymentMethod{
		{
			Id:    1,
			Image: "/images/icon_ac.png",
		},
		{
			Id:    2,
			Image: "/images/icon_aba.png",
		},
		{
			Id:    3,
			Image: "/images/icon_wing.png",
		},
		{
			Id:    4,
			Image: "/images/icon_true_money.png",
		},
	},
}
