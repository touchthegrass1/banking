package handlers

import (
	"net/http"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type KafkaHandler struct {
	consumer *kafka.Consumer
	log      *zap.Logger
}

func (handler KafkaHandler) GetNextMessage(c *gin.Context) {
	ev, err := handler.consumer.ReadMessage(100 * time.Millisecond)
	if err != nil {
		handler.log.Error("Error with event", zap.Error(err))
		c.JSON(http.StatusBadRequest, "No message now")
		return
	}
	c.JSON(http.StatusOK, string(ev.Value))
}
