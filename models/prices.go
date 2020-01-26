package models

// type Product struct {
// 	PlanoP     PlanInfo `json:"planoP,omitempty"`
// 	PlanoM     PlanInfo `json:"planoM,omitempty"`
// 	PlanoTurbo PlanInfo `json:"PlanoTurbo,omitempty"`
// }

type Plan struct {
	ID        int         `json:"id"`
	Name      string      `json:"name"`
	CleanName string      `json:"cleanName,omitempty"`
	Cycle     interface{} `json:"cycle,omitempty"`
}

type Cycle struct {
	Annually     CycleInfo `json:"annually"`
	Biennially   CycleInfo `json:"biennially"`
	Monthly      CycleInfo `json:"monthly"`
	Quarterly    CycleInfo `json:"quarterly"`
	Semiannually CycleInfo `json:"semiannually"`
	Triennially  CycleInfo `json:"triennially"`
}

type CycleInfo struct {
	Type       string `json:"type,omitempty"`
	PriceRenew string `json:"priceRenew"`
	PriceOrder string `json:"priceOrder"`
	Months     int    `json:"months"`
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
