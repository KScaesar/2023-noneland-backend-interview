package pkg

import (
	"time"

	"noneland/backend/interview/pkg/errors"
)

func NewListResponse[T any](list []T) ListResponse[T] {
	if list == nil {
		list = make([]T, 0)
	}
	return ListResponse[T]{
		Rows:  list,
		Total: len(list),
	}
}

type ListResponse[T any] struct {
	Rows  []T `json:"rows"`
	Total int `json:"total"`
}

//

type PageParam struct {
	Page uint64 `json:"current" form:"current"`
	Size uint64 `json:"size" form:"size"`
}

func (p PageParam) Validate() error {
	if p.Page > 0 && p.Size > 0 {
		return nil
	}
	return errors.WrapWithMessage(errors.ErrInvalidParams, "page dto is invalid")
}

func (p PageParam) IsPagination() bool {
	return !(p.Page == 0 && p.Size == 0)
}

func (p *PageParam) SetWithoutPagination() {
	p.Page = 0
	p.Size = 0
}

func (p *PageParam) SetDefaultIfInvalid() {
	const (
		defaultSize = 10
		defaultMax  = 100
	)
	p.SetDefaultAndMaxSizeIfInvalid(defaultSize, defaultMax)
}

func (p *PageParam) SetDefaultAndMaxSizeIfInvalid(defaultSize, maxPageSize uint64) {
	const (
		defaultPage = 1
	)

	err := p.Validate()
	if err != nil {
		p.Page = defaultPage
		p.Size = defaultSize
		return
	}

	if p.Size > maxPageSize {
		p.Size = maxPageSize
	}
}

func (p PageParam) OffsetOrSkip() int64 {
	return int64((p.Page - 1) * (p.Size))
}

//

type TimeRangeEndTimeLessThan struct {
	StartTime time.Time `json:"startTime" form:"startTime"`
	EndTime   time.Time `json:"endTime" form:"endTime"`
}

func (t *TimeRangeEndTimeLessThan) ToMilliTimestamp() TimestampRangeEndTimeLessThan {
	return TimestampRangeEndTimeLessThan{
		StartTime: t.StartTime.UnixMilli(),
		EndTime:   t.EndTime.UnixMilli(),
	}
}

type TimestampRangeEndTimeLessThan struct {
	StartTime int64 `json:"startTime" form:"startTime"`
	EndTime   int64 `json:"endTime" form:"endTime"`
}

func (t *TimestampRangeEndTimeLessThan) ToTime() TimeRangeEndTimeLessThan {
	return TimeRangeEndTimeLessThan{
		StartTime: time.UnixMilli(t.StartTime),
		EndTime:   time.UnixMilli(t.EndTime),
	}
}

type TimeRangeEndTimeLessThanEqual struct {
	StartTime time.Time `json:"startTime" form:"startTime"`
	EndTime   time.Time `json:"endTime" form:"endTime"`
}

func (t *TimeRangeEndTimeLessThanEqual) ToMilliTimestamp() TimestampRangeEndTimeLessThanEqual {
	return TimestampRangeEndTimeLessThanEqual{
		StartTime: t.StartTime.UnixMilli(),
		EndTime:   t.EndTime.UnixMilli(),
	}
}

type TimestampRangeEndTimeLessThanEqual struct {
	StartTime int64 `json:"startTime" form:"startTime"`
	EndTime   int64 `json:"endTime" form:"endTime"`
}

func (t *TimestampRangeEndTimeLessThanEqual) ToTime() TimeRangeEndTimeLessThanEqual {
	return TimeRangeEndTimeLessThanEqual{
		StartTime: time.UnixMilli(t.StartTime),
		EndTime:   time.UnixMilli(t.EndTime),
	}
}
