package models

type Product struct {
	Plans []Plan
}
type Plan struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Cycle
}

type Cycles struct {
	Annually     GenericCycle `json:"annually"`
	Biennially   GenericCycle `json:"biennially"`
	Monthly      GenericCycle `json:"monthly"`
	Quarterly    GenericCycle `json:"quarterly"`
	Semiannually GenericCycle `json:"semiannually"`
	Triennially  GenericCycle `json:"triennially"`
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
