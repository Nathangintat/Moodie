package response

type UserResponse struct {
	ID           int64  `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	ProfileImage string `json:"profile_image"`
}
