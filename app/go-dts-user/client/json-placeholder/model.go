package jsonplaceholder

type JsonPlaceholderResp struct {
	UserID    uint64 `json:"userId"`
	ID        uint64 `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
