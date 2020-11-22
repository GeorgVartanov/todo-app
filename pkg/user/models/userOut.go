package models


type UserOut struct {
	Id            uint64 `json:"id"`
	Email         string `json:"email"`
	DisplayName   string `json:"displayName"`
}
