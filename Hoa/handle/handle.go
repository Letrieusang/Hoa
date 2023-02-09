package handle

import (
	"Hoa/Module/Model"
	"Hoa/common"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func DbConnect() *gorm.DB {
	dsn := "root:Iamspectre16@tcp(127.0.0.1:3306)/hoa?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// AddAdvise event godoc
// @Summary add advise
// @Tags content
// @Param user body Model.AdviseRequest true "add advise"
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response{data=map[int64]string}
// @Failure 400,401,500 {object} common.Response
// @Router /content/advise [post]
func AddAdvise() func(c *gin.Context) {
	return func(c *gin.Context) {
		db := DbConnect()
		var rq Model.AdviseRequest
		if err := c.ShouldBindJSON(&rq); err != nil {
			c.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
		}
		advise := rq.ConvertToAdvise()
		err := db.Table(Model.Advise{}.TableName()).Create(&advise).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
		}
		c.JSON(http.StatusOK, common.NewResponse(0, "Insert advise success", nil))
	}
}

// AddConfess event godoc
// @Summary add advise
// @Tags content
// @Param user body Model.ConfessionRequest true "add advise"
// @Accept json
// @Produce json
// @Success 200 {object}{object} common.Response{data=map[int64]string}
// @Failure 400,401,500 {object} common.Response
// @Router /content/advise [post]
func AddConfess() func(c *gin.Context) {
	return func(c *gin.Context) {
		db := DbConnect()
		var rq Model.ConfessionRequest
		if err := c.ShouldBindJSON(&rq); err != nil {
			c.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
		}
		confess := rq.ConvertToConfess()
		err := db.Table(Model.Confession{}.TableName()).Create(&confess).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.NewResponse(1, "Internal Server Error", err.Error()))
		}
		c.JSON(http.StatusOK, common.NewResponse(0, "Insert confess success", nil))
	}
}
