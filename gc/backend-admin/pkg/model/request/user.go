package request

type UserIds struct {
  UserIds []UserId `json:"user_ids"`
}

type UserId struct {
  ID int `json:"id"`
}
