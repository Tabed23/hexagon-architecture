package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Rest) DeleteCandy(ctx *gin.Context) {
	id := ctx.Param("id")

	msg, err := r.s.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": msg})
}
