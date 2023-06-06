package bd

import (
	"context"

	"github.com/ptilotta/twitterGo/models"
)

func BorroRelacion(t models.Relacion) (bool, error) {
	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("relacion")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil
}
