package card

import "time"

type Card struct {
	ID         uint       `json:"id" bson:"id"`
	CardId     string     `json:"card_id" bson:"card_id"`
	CreatedAt  time.Time  `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at" bson:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
	Name       string     `json:"name" bson:"name"`
	Collection string     `json:"collection" bson:"collection"`
	Rarity     string     `json:"rarity" bson:"rarity"`
	ImageUrl   string     `json:"image_url" bson:"image_url"`
}

type Cards struct {
	Cards []Card `json:"cards" bson:"cards"`
}

type CardRequestBody struct {
	Name       string `json:"name" bson:"name"`
	CardId     string `json:"card_id" bson:"card_id"`
	Collection string `json:"collection" bson:"collection"`
	Rarity     string `json:"rarity" bson:"rarity"`
}
