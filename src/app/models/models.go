package models

type UserModel struct {
	Uuid     string `json:"uuid"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserAuthModel struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type PostModel struct {
	UserUUID string `json:"user_uuid"`
	Content  string `json:"content"`
	Author   string `json:"author"`
	Date     string `json:"date"`
	Likes    uint64 `json:"likes_count"`
}
