package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"memoryCache/models"
	"strconv"
)

func getItem(ctx *gin.Context) (models.Item, error) {
	var item models.Item
	if err := ctx.ShouldBindJSON(&item); err != nil {
		return models.Item{}, errors.New(fmt.Sprintf("failed BindJSON: %s", err))
	}

	return item, nil
}

func getKey(ctx *gin.Context) string {
	key := ctx.Query("key")

	return key
}

func getValue(ctx *gin.Context) (int64, error) {
	str := ctx.Query("value")

	value, err := strconv.Atoi(str)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("value is not int: %s", err))
	}

	return int64(value), nil
}

func getN(ctx *gin.Context) (int64, error) {
	str := ctx.DefaultQuery("N", "1")

	n, err := strconv.Atoi(str)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("N is not int: %s", err))
	}

	return int64(n), nil
}
