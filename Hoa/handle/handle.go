package handle

import (
	"Hoa/Module/Model"
	"Hoa/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func AddAdvise() func(c *gin.Context) {
	return func(c *gin.Context) {
		var db *gorm.DB
		var advise Model.Advise
		if err := c.ShouldBindJSON(&advise); err != nil {
			c.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
		}
		err := db.Table(Model.Advise{}.TableName()).Create(advise).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
		}
		c.JSON(http.StatusOK, common.NewResponse(0, "Insert advise success", nil))
	}
}
