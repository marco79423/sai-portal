package dto

import "time"

// MPSPingEvent Ping 事件
type MPSPingEvent struct {
	Timestamp int64 `json:"timestamp"` // 當前時間 (毫秒)
	Delayed   bool  `json:"delayed"`   // 是否延遲
}

// MPSMatchUpdatedEvent 賽事更新事件
type MPSMatchUpdatedEvent struct {
	RequestID  int64     `json:"requestID"`  // RequestID
	SourceTime time.Time `json:"sourceTime"` // 數據源更新時間 (毫秒)
	Phase      int       `json:"phase"`      // 賽事類型 1 (滾球) 或 2 (早盤)

	// 除錯用的額外資訊
	SportID               int64 `json:"sportID"`
	MatchID               int64 `json:"matchID"`
	IsCoverageOrCancelled bool  `json:"isCoverageOrCancelled"`
}
