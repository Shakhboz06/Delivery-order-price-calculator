package models

type Response struct {
	TotalPrice          int      `json:"total_price"`
	SmallOrderSurcharge int      `json:"small_order_surcharge"`
	CartValue           int      `json:"cart_value"`
	Delivery            Delivery `json:"delivery"`
}

type Delivery struct {
	Fee      int `json:"fee"`
	Distance int `json:"distance"`
}
