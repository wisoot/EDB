package routes

import (
	controller "edb/web/controllers/v1"
	"github.com/gin-gonic/gin"
)

func InitV1Routes(router *gin.Engine) {
	v1Accounts := router.Group("v1/accounts")
	{
		v1Accounts.POST("/", controller.StoreAccount)
		v1Accounts.POST("/:id/credit", controller.CreditAccount)
		v1Accounts.POST("/:id/debit", controller.DebitAccount)
	}
}
