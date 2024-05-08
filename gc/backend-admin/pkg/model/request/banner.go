package request

type BannerPriorities struct {
  BannerPriorities []BannerPriority `json:"banner_priorities"`
}

type BannerPriority struct {
  Id       int `json:"id"`
  Priority int `json:"priority"`
}
