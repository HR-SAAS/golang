package request

type PageInfo struct {
	Page  int `json:"page" binding:"int"`
	Limit int `json:"limit" binding:"int"`
}
