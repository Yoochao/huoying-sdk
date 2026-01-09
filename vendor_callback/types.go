package vendor_callback

// RewardInfo 奖励信息
type RewardInfo struct {
	Type      string `json:"type"`       // 奖励类型 (e.g., duration)
	Value     int64  `json:"value"`      // 奖励数值
	Unit      string `json:"unit"`       // 单位 (e.g., day)
	ProductID int64  `json:"product_id"` // 产品ID
}

// ProcessCodeRequest 统一码处理请求
type ProcessCodeRequest struct {
	Code       string `json:"code"`        // 邀请码或兑换码
	UserID     string `json:"user_id"`     // 厂商端用户ID
	ActionTime int64  `json:"action_time"` // 操作时间戳 (秒)
}

// ProcessCodeResponseData 统一码处理响应数据
type ProcessCodeResponseData struct {
	Success bool        `json:"success"` // 是否成功
	Type    string      `json:"type"`    // 码类型: invitation (邀请码), redemption (兑换码)
	Reward  *RewardInfo `json:"reward"`  // 奖励信息
	Reason  string      `json:"reason"`  // 失败原因
}

// PaymentCallbackRequest 支付上报请求
type PaymentCallbackRequest struct {
	UserID    string  `json:"user_id"`    // 厂商端用户ID
	OrderNo   string  `json:"order_no"`   // 厂商订单号
	Amount    float64 `json:"amount"`     // 支付金额
	ProductID int64   `json:"product_id"` // 产品ID
}

// ApiResponse 通用API响应
type ApiResponse[T any] struct {
	Code int    `json:"code"` // 0 表示成功，非 0 表示错误码
	Msg  string `json:"msg"`  // 提示信息或错误描述
	Data T      `json:"data"` // 具体的数据结构
}
