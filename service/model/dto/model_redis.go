package dto

// ServiceStatusRedisValue 服務狀態的 Redis 值
type ServiceStatusRedisValue struct {
	Timestamp string `json:"timestamp"`
	Error     string `json:"error"`
	Reason    string `json:"reason"`
}

// GlobalLockRedisValue 全局鎖用的 Redis 值
type GlobalLockRedisValue struct {
	Timestamp int64 `json:"timestamp"`

	OddsDelayedLock *string `json:"oddsDelayedLock,omitempty"`
	ErrorLock       *string `json:"errorLock,omitempty"`
	ErrorReason     *string `json:"errorReason,omitempty"`
}

// GlobalLockRedisChannel 全局鎖用的 Redis Channel 值
type GlobalLockRedisChannel struct {
	Timestamp int64 `json:"timestamp"`

	OddsDelayedLock *bool   `json:"oddsDelayedLock,omitempty"`
	ErrorLock       *bool   `json:"errorLock,omitempty"`
	ErrorReason     *string `json:"errorReason,omitempty"`
}
