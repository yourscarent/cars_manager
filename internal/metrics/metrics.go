package metrics

import "time"

type Metrics interface {
	ObserveReq(name string, status int, dur time.Duration)
}
