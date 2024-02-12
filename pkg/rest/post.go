package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tabed23/hexagon-architecture/pkg/model"
)

func (r *Rest) PostCandy(ctx *gin.Context) {
	var candy model.Candy
	if err := ctx.ShouldBind(&candy); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "err": err.Error()})
		return
	}
	created, err := r.s.Post(candy)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "err": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "msg": created})

}
