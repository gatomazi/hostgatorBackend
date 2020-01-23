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
