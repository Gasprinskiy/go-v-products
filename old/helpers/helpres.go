package helpers

import (
	"fmt"
	"jun2/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParseId(c *gin.Context) (res int, err error) {
	res, err = ParseString(c.Param("id"))
	return
}

func GetConfig(c *structs.Config) string {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.DBName)

	return psqlInfo
}

func ParseString(str string) (res int, err error) {
	res, err = strconv.Atoi(str)
	return
}
