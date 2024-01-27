package handlers

import (
	"meetspace_backend/client/types"
	"meetspace_backend/config"
	"meetspace_backend/utils"

	"github.com/gin-gonic/gin"
)

// ClientAddUser godoc
//	@Summary		ClientAddUser account
//	@Description	ClientAddUser account
//	@Tags			Client-User
//	@Produce		json
// @Param user body types.ClientAddUser true "User registration details"
//	@Router			/v1/client/users [post]
func ClientAddUser(c *gin.Context){
	currentUser, exists := utils.GetUserFromContext(c)
    if !exists{
        return 
    }


	currentClient, err  := config.ClientService.GetClientById(currentUser.ClientID.String())
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var reqData types.ClientAddUser
	if err := utils.BindJsonData(c, &reqData); err != nil {
		utils.HandleError(c, err)
		return
	}
	reqData.ClientID = currentUser.ClientID
	reqData.CreatedBy = &currentClient
	reqData.UpdatedBy = &currentClient

	user, err := config.ClientUserService.AddClientUser(reqData)
	if err != nil {  
		utils.HandleError(c, err)
		return 
	}

	resp := utils.SuccessResponse("success", user)
	c.JSON(resp.StatusCode, resp)
	return
}

// GetClientUsers godoc
//	@Summary		GetClientUsers account
//	@Description	GetClientUsers account
//	@Tags			Client-User
//	@Produce		json
//	@Router			/v1/client/users [get]
func GetClientUsers(c *gin.Context){
	currentUser, _ := utils.GetUserFromContext(c)

	users , _  := config.ClientUserService.GetClientUsers(currentUser.ClientID.String())
	resp := utils.SuccessResponse("success", users)
	c.JSON(resp.StatusCode, resp)
	return
}
