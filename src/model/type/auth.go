package Datatype

type LoginPostBody struct {
	EmailId  string `json:"emailId"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken          string `json:"accessToken"`
	RefreshToken         string `json:"refreshToken"`
	AccessTokenValidity  int64  `json:"accessTokenValidity"`
	RefreshTokenValidity int64  `json:"refreshTokenValidity"`
}

type JwtPayload struct {
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName"`
	UserId     string `json:"userId"`
	EmailId    string `json:"emailId"`
	Gender     string `json:"gender"`
}

type RefreshAccessTokenQuery struct {
	RefreshToken string `form:"refreshToken"`
}
