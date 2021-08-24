package bo

import "time"

type MPSPingEvent struct {
	Timestamp time.Time // 當前時間
	Delayed   bool      // 是否延遲
}

type MPSMatchUpdatedEvent struct {
	SourceTime  time.Time // 數據源更新時間
	CurrentTime time.Time // 事件接收時間
	Phase       int       // 賽事類型 1 (滾球) 或 2 (早盤)

	// 除錯用的額外資訊
	RequestID             int64
	SportID               int64
	MatchID               int64
	IsCoverageOrCancelled bool
}

func (event *MPSMatchUpdatedEvent) IsInPlay() bool {
	return event.Phase == 1
}

func (event *MPSMatchUpdatedEvent) IsPreMatch() bool {
	return event.Phase == 2
}
