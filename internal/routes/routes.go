func RegisterRoutes(r *gin.Engine) {
	user := r.Group("/users")
	{
		user.GET("/", controllers.GetUsers)
	}
}
