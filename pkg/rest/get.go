package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Rest) GetCandy(ctx *gin.Context) {
	candies, err := r.s.Get()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "err": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"candy": candies})
}
