package Controllers

import (
	"capydemy/netless"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetRoomToken(c *gin.Context) {
	ctx := netless.RoomContent{
		Role: netless.AdminRole,
		Uuid: "802e8a20936f11ed86a26bf46008d7bd",
	}
	netlessRoomToken := netless.RoomToken(
		os.Getenv("AK"),
		os.Getenv("SK"),
		0, &ctx,
	)
	c.JSON(http.StatusAccepted, gin.H{
		"room_token": netlessRoomToken,
		"room_uuid":  "802e8a20936f11ed86a26bf46008d7bd",
	})
}
