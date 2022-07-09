package handler

import (
	"github.com/gin-gonic/gin"
	"memoryCache/models"
	"net/http"
)

func (h *Handler) setItem(ctx *gin.Context) {
	item, err := getItem(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	h.services.MemoryCache.Set(item.Key, item.Value, h.config.Duration)

	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
	})
}

func (h *Handler) getValue(ctx *gin.Context) {
	key := getKey(ctx)

	value, err := h.services.MemoryCache.Get(key)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	item := models.Item{
		Key:   key,
		Value: value,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": &item,
	})
}

func (h *Handler) deleteItem(ctx *gin.Context) {
	key := getKey(ctx)

	err := h.services.MemoryCache.Delete(key)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
	})
}

func (h *Handler) searchKey(ctx *gin.Context) {
	value, err := getValue(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	keys, err := h.services.MemoryCache.Search(value)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	items := make([]models.Item, 0)

	for _, key := range keys {
		items = append(items, models.Item{
			Key:   key,
			Value: value,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": &items,
	})
}

func (h *Handler) increaseValue(ctx *gin.Context) {
	key := getKey(ctx)
	n, err := getN(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.MemoryCache.Increase(key, n)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
	})
}

func (h *Handler) reduceValue(ctx *gin.Context) {
	key := getKey(ctx)
	n, err := getN(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.MemoryCache.Reduce(key, n)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
	})
}
