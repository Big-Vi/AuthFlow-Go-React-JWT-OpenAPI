package openapi

import (
	"net/http"

	"github.com/Big-Vi/ticketInf/models"
	"github.com/swaggest/openapi-go"
	"github.com/swaggest/openapi-go/openapi3"
)

func buildUser(reflector *openapi3.Reflector) {
	postOp, _ := reflector.NewOperationContext(http.MethodPost, "/api/user")

	postOp.AddReqStructure(new(models.CreateUserReq))
	postOp.AddRespStructure(new(models.User))
	postOp.AddRespStructure(new([]models.User), openapi.WithHTTPStatus(http.StatusConflict))
	reflector.AddOperation(postOp)
}