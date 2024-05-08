package response

type Data struct {
  Success bool        `json:"success"`
  Data    interface{} `json:"data"`
}
