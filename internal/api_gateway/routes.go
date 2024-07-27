package apigateway

import (
	"github.com/gin-gonic/gin"
	"github.com/paper-assessment/internal/api_gateway/routes"
)

func RegisterRoutes(r *gin.Engine) *Client {
	svc := &Client{
	}

	r.POST("/disburse", svc.Disburse)

	return svc
}

func (svc *Client) Disburse(ctx *gin.Context) {
	routes.Disburse(ctx);
}