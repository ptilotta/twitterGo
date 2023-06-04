package routers

import (
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/ptilotta/twitterGo/bd"
	"github.com/ptilotta/twitterGo/models"
)

func LeoTweets(request events.APIGatewayProxyRequest) models.RespApi {
	var r models.RespApi
	r.Status = 400

	ID := request.QueryStringParameters["id"]
	pagina := request.QueryStringParameters["pagina"]

	if len(ID) < 1 {
		r.Message = "El parámetro ID es obligatorio"
		return r
	}

	if len(pagina) < 1 {
		pagina = "1"
	}

	pag, err := strconv.Atoi(pagina)
	if err != nil {
		r.Message = "Debe enviar el parámetro Pagina como un valor mayor a 0"
		return r
	}

	tweets, correcto := bd.LeoTweets(ID, int64(pag))
	if !correcto {
		r.Message = "Error al leer los Tweets"
		return r
	}

	respJson, err := json.Marshal(tweets)
	if err != nil {
		r.Status = 500
		r.Message = "Error al formatear los datos de los usuarios como JSON"
		return r
	}

	r.Status = 200
	r.Message = string(respJson)
	return r
}
