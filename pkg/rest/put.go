package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tabed23/hexagon-architecture/pkg/model"
)

func (r *Rest) PuteCandy(ctx *gin.Context) {
	id := ctx.Param("id")

	var candy model.Candy
	if err := ctx.BindJSON(&candy); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedCandy, err := r.s.Put(id, candy)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "updated_candy": updatedCandy})
}
