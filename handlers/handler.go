package handler

import (
	"context"
	"net/http"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jasonbirchall/tools-api-poc/models"
	"gopkg.in/mgo.v2/bson"
	"gorm.io/gorm"
)

type ToolsHandler struct {
	collection  *gorm.DB
	ctx         context.Context
	redisClient *redis.Client
}

func NewToolsHandler(ctx context.Context, collection *gorm.DB, redisClient *redis.Client) *ToolsHandler {
	return &ToolsHandler{
		collection:  collection,
		ctx:         ctx,
		redisClient: redisClient,
	}
}

func (handler *ToolsHandler) ListToolsHandler(c *gin.Context) {
	val, err := handler.redisClient.Get("tools").Result()
	if err == redis.Nil {
		cur, err := handler.collection.Find(handler.ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		defer cur.Close(handler.ctx)

		tools := make([]models.Tool, 0)
		for cur.Next(handler.ctx) {
			var tool models.Tool
			if err := cur.Decode(&tool); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			tools = append(tools, tool)
		}
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tools := make([]models.Tool, 0)
	json.Unmarshall(val, &tools)
	c.JSON(http.StatusOK, tools)
}
