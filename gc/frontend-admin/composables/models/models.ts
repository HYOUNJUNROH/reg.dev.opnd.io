export interface User {
  id: number;
  password: string;
  role: string;
  created_at: string;
  activated: boolean;
  user_id: string;
  gender: string;
  name: string;
  phone1: string;
  phone2: string;
  amount: number;
  paid_status: string;
  cancel_status: string;
  token: string;
}

export interface UsersResponse {
  success: boolean;
  data: {
    users: User[];
  };
}

export interface Investment {
  id: number;
  invest_id: string;
  created_at: string;
  start_date: string;
  status: string;
  amount: number;
  paid_status: string;
  user_count: number;
  current_invest: number;
  max_invest: number;
  title: string;
}

export interface GetInvestmentsResponse {
  success: boolean;
  data: {
    investments: Investment[];
  };
}

export interface CancelUserInvestmentRequest {
  user_ids: {
    id: number;
  }[];
}

export type Pagenation = {
  page: number;
  limit: number;
  type: string;
  search?: string;
  keyword?: string;
  id?: number;
  isCallBack?: boolean;
};

export type Banner = {
  id: number;
  created_at: string;
  updated_at: string;
  description: string;
  image: string;
  priority: number;
};

export type Banners = {
  banners: Banner[];
};

export type BannerPriority = {
  id: number;
  priority: number;
}

export type BannerPriorities = {
  banner_priorities: BannerPriority[];
}
