package response

type Investments struct {
  Investments []Investment `json:"investments"`
}

type Investment struct {
  ID            int    `json:"id"`
  InvestID      string `json:"invest_id"`
  CreatedAt     string `json:"created_at"`
  StartDate     string `json:"start_date"`
  Status        string `json:"status"`
  Amount        int    `json:"amount,omitempty"`
  PaidStatus    string `json:"paid_status,omitempty"`
  UserCount     int    `json:"user_count,omitempty"`
  CurrentInvest int    `json:"current_invest,omitempty"`
  MaxInvest     int    `json:"max_invest,omitempty"`
  Title         string `json:"title"`
}
