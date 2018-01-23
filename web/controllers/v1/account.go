package v1

import (
	"github.com/gin-gonic/gin"
	"edb/web/requests/v1"
	"edb/entities/account"
	accountDomain "edb/domain/account"
	"net/http"
	"strconv"
)

func StoreAccount(c *gin.Context) {
	var request v1.AccountCreationRequest
	c.BindJSON(&request)

	form := account.CreationForm{Firstname: request.Firstname, Lastname: request.Lastname, Email: request.Email, Password: request.Password}
	creator := accountDomain.MakeCreator()
	creator.Create(form)

	c.JSON(http.StatusCreated, gin.H{})
}

func CreditAccount(c *gin.Context) {
	idString := c.Param("id")

	id, err := strconv.ParseUint(idString, 10, 64)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Id is invalid!"})
		return
	}

	var request v1.AccountCreditingRequest
	c.BindJSON(&request)

	creditor := accountDomain.MakeCreditor()
	err = creditor.Credit(uint(id), request.Amount, request.Description)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

func DebitAccount(c *gin.Context) {
	idString := c.Param("id")

	id, err := strconv.ParseUint(idString, 10, 64)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Id is invalid!"})
		return
	}

	var request v1.AccountDebitingRequest
	c.BindJSON(&request)

	debitor := accountDomain.MakeDebitor()
	err = debitor.Debit(uint(id), request.Amount, request.Description)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
