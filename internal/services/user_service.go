func GetAllUsers() []models.User {
	return repositories.FindAllUsers()
}
