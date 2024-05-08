package model

//{
//"id":0,
//"created_at":"2023-05-09 00:43:21.433000 +00:00",
//"updated_at":"2023-05-09 00:43:21.433000 +00:00",
//"description":"Test Text",
//"image":"e9d0624a-9ed0-4c7a-9472-000f0b4deead.jpg",
//"priority":1
//}

type Banners struct {
  Banners []Banner `json:"banners"`
}

type Banner struct {
  Id          int    `json:"id"`
  CreatedAt   string `json:"created_at"`
  UpdatedAt   string `json:"updated_at"`
  Description string `json:"description"`
  Image       string `json:"image"`
  Priority    int    `json:"priority"`
}
