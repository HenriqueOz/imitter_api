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
	Id       uint64 `json:"id"`
	UserUUID string `json:"user_uuid"`
	Content  string `json:"content"`
	Author   string `json:"author"`
	Date     string `json:"date"`
	Likes    uint64 `json:"likes_count"`
	IsLiked  bool   `json:"is_liked"`
}

type UserProfileModel struct {
	Uuid           string `json:"uuid"`
	Name           string `json:"name"`
	FollowersCount uint64 `json:"followers_count"`
	FollowingCount uint64 `json:"following_count"`
	IsFollowing    bool   `json:"is_following"`
}
