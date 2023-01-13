package Controllers

import (
	"capydemy/netless"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetRoomToken(c *gin.Context) {
	room_name := c.Param("room_name")
	ctx := netless.RoomContent{
		Role: netless.AdminRole,
		Uuid: room_name,
	}
	netlessRoomToken := netless.RoomToken(
		os.Getenv("AK"),
		os.Getenv("SK"),
		0, &ctx,
	)
	c.JSON(http.StatusAccepted, gin.H{
		"room_token": netlessRoomToken,
		"room_uuid":  room_name,
	})
}
