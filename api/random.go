package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/feo0o/iridium"
	"github.com/gin-gonic/gin"
)

// GetRandomString return a random string with given length, default length is 8
// @Summary generate random string
// @Description generate random string with given length, default length is 8
// @Tags random
// @Param length query int false "random string length"
// @Produce plain
// @Success 200 {string} string "OK"
// @Router /random/string [get]
func GetRandomString(c *gin.Context) {
	len := c.DefaultQuery("length", "8")
	l, err := strconv.Atoi(len)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("invalid query value: %s, %s\n", len, err))
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("%s\n", iridium.RandomString(l)))
}

// GetRandonInt return a random int with given range, default range is 0~10
// @Summary generate random int number
// @Description generate random int number with given range, default is 0~10
// @Tags random
// @Param min query int false "minmal number"
// @Param max query int false "maxmum number"
// @Produce plain
// @Success 200 {string} string "'random number'"
// @Router /random/int [get]
func GetRandomInt(c *gin.Context) {
	min := c.DefaultQuery("min", "0")
	iMin, err := strconv.Atoi(min)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("invalid min value: %s, %s\n", min, err))
		return
	}

	max := c.DefaultQuery("max", "10")
	iMax, err := strconv.Atoi(max)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("invalid max value: %s, %s\n", max, err))
		return
	}

	v, err := iridium.RandomInt(iMin, iMax)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("%s\n", err))
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("%d\n", v))
}

// GetRandomdecimal return a decimal number with given precision, default precision is 6
// @Summary generate random decimal
// @Description generate a random decimal with given number range and precision, default range is 0~10, and defualt precision is 6
// @Tags random
// @Param min query int false "minmal number"
// @Param max query int false "maxmum number"
// @Param precision query int false "decimal precision"
// @Produce plain
// @Success 200 {string} string "'random decimal'"
// @Router /random/decimal [get]
func GetRandomdecimal(c *gin.Context) {
	min := c.DefaultQuery("min", "0")
	iMin, err := strconv.Atoi(min)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("invalid min value: %s, %s\n", min, err))
		return
	}

	max := c.DefaultQuery("max", "9")
	iMax, err := strconv.Atoi(max)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("invalid max value: %s, %s\n", max, err))
		return
	}
	iMax--
	vi, err := iridium.RandomInt(iMin, iMax)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("%s\n", err))
		return
	}

	precision := c.DefaultQuery("precision", "6")
	iPrecision, err := strconv.Atoi(precision)
	if err != nil {
		c.String(http.StatusBadGateway, fmt.Sprintf("invalid precision value: %s, %s\n", precision, err))
		return
	}
	vd, err := iridium.RandomFloat(iPrecision)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("%s\n", err))
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("%f\n", float64(vi)+vd))
}
