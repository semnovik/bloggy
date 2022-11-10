package bloggy

type MessagesList struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UserList struct {
	Id     int `json:"id"`
	UserId int `json:"user_id"`
	ListId int `json:"list_id"`
}
type MessageItem struct {
	Id      int    `json:"id"`
	Message string `json:"message"`
}

type ListsItem struct {
	Id     int `json:"id"`
	ListId int `json:"list_id"`
	ItemId int `json:"item_id"`
}
