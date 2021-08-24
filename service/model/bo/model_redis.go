package bo

import "time"

// ServiceStatusRedisValue 服務狀態的 Redis 值
type ServiceStatusRedisValue struct {
	Timestamp time.Time `json:"timestamp"`
	Error     bool      `json:"error"`
	Reason    string    `json:"reason"`
}
