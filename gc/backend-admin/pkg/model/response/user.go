package response

type Users struct {
  Users []User `json:"users"`
}

type User struct {
  ID           int    `json:"id"`
  Password     string `json:"password,omitempty"`
  Role         string `json:"role,omitempty"`
  CreatedAt    string `json:"created_at,omitempty"`
  Activated    bool   `json:"activated,omitempty"`
  UserID       string `json:"user_id"`
  Gender       string `json:"gender"`
  Name         string `json:"name"`
  Phone1       string `json:"phone1"`
  Phone2       string `json:"phone2"`
  Amount       int    `json:"amount,omitempty"`
  PaidStatus   string `json:"paid_status,omitempty"`
  CancelStatus string `json:"cancel_status,omitempty"`
  Token        string `json:"token,omitempty"`
}
