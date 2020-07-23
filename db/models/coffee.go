package models

import "graphql/graph/model"

type Coffee struct {
	ID        *string `json:"id" bson:"_id,omitempty"`
	Name      *string `json:"name" bson:"name"`
	Key       *string `json:"key" bson:"key"`
	CreatedAt int     `json:"created_at" bson:"created_at"`
	UpdateAt  int     `json:"update_at" bson:"updated_at"`
	DeletedAt int     `json:"deleted_at" bson:"deleted_at"`
}

func (c *Coffee) Serialize() *model.Coffee {
	coff:=&model.Coffee{
		ID: c.ID,
		Name: c.Name,
		Key: c.Key,
		UpdateAt: c.UpdateAt,
		CreatedAt: c.CreatedAt,
		DeletedAt: c.DeletedAt,
	}

	return coff
}
