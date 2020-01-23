package models

// Product - Struct model for products
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Cycle []Cycle `json:"cycle"`
}

//Cycle - Struct model for cycle
type Cycle struct {
	TypeCycle TypeCycle `json:"priceRenew"`
}
type GenericCycle struct {
	PriceRenew string `json:"priceRenew"`
	PriceOrder string `json:"priceOrder"`
	Months     int    `json:"months"`
}

type MyJsonName struct {
	Shared struct {
		Products struct {
			PlanoM struct {
				ID    int    `json:"id"`
				Name  string `json:"name"`
				Cycle struct {
					Annually struct {
						Months     int    `json:"months"`
						PriceOrder string `json:"priceOrder"`
						PriceRenew string `json:"priceRenew"`
					} `json:"annually"`
					Biennially struct {
						Months     int    `json:"months"`
						PriceOrder string `json:"priceOrder"`
						PriceRenew string `json:"priceRenew"`
					} `json:"biennially"`
					Monthly struct {
						Months     int    `json:"months"`
						PriceOrder string `json:"priceOrder"`
						PriceRenew string `json:"priceRenew"`
					} `json:"monthly"`
					Quarterly struct {
						Months     int    `json:"months"`
						PriceOrder string `json:"priceOrder"`
						PriceRenew string `json:"priceRenew"`
					} `json:"quarterly"`
					Semiannually struct {
						Months     int    `json:"months"`
						PriceOrder string `json:"priceOrder"`
						PriceRenew string `json:"priceRenew"`
					} `json:"semiannually"`
					Triennially struct {
						Months     int    `json:"months"`
						PriceOrder string `json:"priceOrder"`
						PriceRenew string `json:"priceRenew"`
					} `json:"triennially"`
				} `json:"cycle"`
			} `json:"planoM"`
		} `json:"products"`
	} `json:"shared"`
}

// teste := "shared": {
// 	"products": {
// 		"planoP": {
// 			"name": "Plano P",
// 			"id": 5,
// 			"cycle": {
// 				"monthly": {
// 					"priceRenew": "24.19",
// 					"priceOrder": "24.19",
// 					"months": 1
// 				},
// 				"semiannually": {
// 					"priceRenew": "128.34",
// 					"priceOrder": "128.34",
// 					"months": 6
// 				},
// 				"biennially": {
// 					"priceRenew": "393.36",
// 					"priceOrder": "393.36",
// 					"months": 24
// 				},
// 				"triennially": {
// 					"priceRenew": "561.13",
// 					"priceOrder": "561.13",
// 					"months": 36
// 				},
// 				"quarterly": {
// 					"priceRenew": "67.17",
// 					"priceOrder": "67.17",
// 					"months": 3
// 				},
// 				"annually": {
// 					"priceRenew": "220.66",
// 					"priceOrder": "220.66",
// 					"months": 12
// 				}
// 			}
// 		},
// 	}
// }
