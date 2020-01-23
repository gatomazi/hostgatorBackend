package models

// Price - Price model struct
type Price struct {
	ID                 int    `json:"id" bson:"id"`
	Codigo             string `json:"codigo" bson:"codigo" binding:"required"`
	GlosaAmigavel      string `json:"glosaAmigavel,omitempty" bson:"glosaAmigavel,omitempty"`
	Termo              string `json:"termo" bson:"termo" binding:"required"`
	GlosaSemFormatacao string `json:"glosaSemFormatacao,omitempty" bson:"glosaSemFormatacao,omitempty"`
	Delete             bool   `json:"delete,omitempty" bson:"delete,omitempty"`
}

"shared": {
	"products": {
	"planoP": {
	"name": "Plano P",
	"id": 5,
	"cycle": {
	"monthly": {
	"priceRenew": "24.19",
	"priceOrder": "24.19",
	"months": 1
	},
	"semiannually": {
	"priceRenew": "128.34",
	"priceOrder": "128.34",
	"months": 6
	},
	"biennially": {
	"priceRenew": "393.36",
	"priceOrder": "393.36",
	"months": 24
	},
	"triennially": {
	"priceRenew": "561.13",
	"priceOrder": "561.13",
	"months": 36
	},
	"quarterly": {
	"priceRenew": "67.17",
	"priceOrder": "67.17",
	"months": 3
	},
	"annually": {
	"priceRenew": "220.66",
	"priceOrder": "220.66",
	"months": 12
	}
	}
	},
}