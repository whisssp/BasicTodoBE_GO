package ginitem

import (
	"ToDoAppBackEnd/common"
	"ToDoAppBackEnd/modules/item/biz"
	"ToDoAppBackEnd/modules/item/model"
	"ToDoAppBackEnd/modules/item/storage"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func CreateItem(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data model.TodoItemCreation
		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := storage.NewSQLStore(db)
		business := biz.NewCreateItemBiz(store)
		if err := business.CreateNewItem(ctx.Request.Context(), &data); err != nil {
			fmt.Println("error", err)
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}