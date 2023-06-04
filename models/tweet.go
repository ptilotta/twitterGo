package models

type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}
