package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"rvgdb-service/pkg/company"
)

func CompanyRouter(router *gin.Engine, service company.Service) {
	c := router.Group(fmt.Sprintf("%s/company", pathPrefix))
	c.GET("", getCompanyList(service))
}

func getCompanyList(service company.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		companies, err := service.GetCompanyList()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.Error{
				Err: err,
			})
		}

		c.JSON(http.StatusOK, companies)
	}
}
