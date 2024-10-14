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
	"reflect"
	"strconv"
)

func DeleteItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.TodoItemUpdate

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := storage.NewSQLStore(db)
		business := biz.NewDeleteItemBiz(store)

		if err := business.DeleteItemById(c.Request.Context(), id); err != nil {
			fmt.Println("error conver", reflect.TypeOf(err))
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}