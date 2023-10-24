package openapi

import (
	"net/http"

	"github.com/Big-Vi/ticketInf/models"
	"github.com/swaggest/openapi-go"
	"github.com/swaggest/openapi-go/openapi3"
)


func buildUser(reflector *openapi3.Reflector) {
	postCreateOp, _ := reflector.NewOperationContext(http.MethodPost, "/api/user")
	postCreateOp.SetTags("user")

	postCreateOp.AddReqStructure(new(models.CreateUserReq))
	postCreateOp.AddRespStructure(new(models.User))
	postCreateOp.AddRespStructure(new([]models.User), openapi.WithHTTPStatus(http.StatusConflict))
	reflector.AddOperation(postCreateOp)

	postLoginOp, _ := reflector.NewOperationContext(http.MethodPost, "/api/user/login")
	postLoginOp.SetTags("user")

	postLoginOp.AddReqStructure(new(models.LoginReq))
	postLoginOp.AddRespStructure(new(models.LoginRes))
	postLoginOp.AddRespStructure(new([]models.LoginRes), openapi.WithHTTPStatus(http.StatusConflict))
	reflector.AddOperation(postLoginOp)

	getDashboardOp, _ := reflector.NewOperationContext(http.MethodGet, "/api/user/dashboard")
	getDashboardOp.SetTags("user")
	reflector.AddOperation(getDashboardOp)

	postLogoutOp, _ := reflector.NewOperationContext(http.MethodPost, "/api/user/logout")
	postLogoutOp.SetTags("user")
	reflector.AddOperation(postLogoutOp)
}
