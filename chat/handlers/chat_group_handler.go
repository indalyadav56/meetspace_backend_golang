package handlers

import (
	"meetspace_backend/chat/models"
	"meetspace_backend/chat/types"
	"meetspace_backend/config"
	userModel "meetspace_backend/user/models"
	"meetspace_backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddChatGroup godoc
//	@Summary		UserLogin User account
//	@Description	UserLogin User account
//	@Tags			Chat-Group
//	@Produce		json
// @Param user body types.LoginRequest true "User login details"
//	@Router			/v1/chat/room/groups [post]
func AddChatGroup(ctx *gin.Context){
    currentUser, exists := utils.GetUserFromContext(ctx)
    if !exists{
        return 
    }
	var reqData types.AddChatGroup

	utils.BindJsonData(ctx, &reqData)
	
	var chatRoom models.ChatRoom
	var roomUsers []*userModel.User

	chatRoom.IsGroup = true
	chatRoom.RoomName = reqData.Title

	for _, userId := range reqData.UserIds {
		user, err := config.ChatGroupService.UserService.UserRepository.GetUserByID(userId)
		if err == nil {
			roomUsers = append(roomUsers, &user)
		}
	}

	roomUsers = append(roomUsers, currentUser)
	chatRoom.RoomUsers = roomUsers
	chatRoom.RoomOwner = currentUser

	chatGroup, _ := config.ChatGroupService.CreateChatGroup(chatRoom)
	ctx.JSON(http.StatusOK, utils.SuccessResponse("success", chatGroup))
	return
}