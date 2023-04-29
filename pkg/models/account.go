package models

type CreateAccountReq struct {
	Name     string `json:"name" binding:"required,lte=32,gte=3"`
	Password string `json:"password" binding:"required,lte=32,gte=8,contains=1Upper,contains=1Lower,contains=1Digit"`
}

type AccountResp struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
