func GetUsers(c *gin.Context) {
	users := services.GetAllUsers()
	c.JSON(200, users)
}
