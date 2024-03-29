package stats

import (
	"github.com/tsuna/gohbase/filter"
	"github.com/lulunevermind/bioviewer/back/logger"
)

type PageableOperations struct {
	Operations []map[string]interface{} `json:"operations"`
	LastRowKey []byte `json:"last_row_key"`
	Error string `json:"error"`
}

func NewPageableOperations(opts ...func(*PageableOperations)) *PageableOperations {
	p := &PageableOperations{}
	for _, option := range opts {
		option(p)
	}
	return p
}

func OptionOperations(operations []map[string]interface{}) func(*PageableOperations) {
	return func(p *PageableOperations) {
		p.Operations = operations
	}
}

func OptionLastRowKey(lastRowKey []byte) func(*PageableOperations) {
	return func(p *PageableOperations) {
		p.LastRowKey = lastRowKey
	}
}

func OptionErrorString(err string) func(*PageableOperations) {
	return func(p *PageableOperations) {
		p.Error = err
	}
}

func NewStatsFilterList(opts ...func(*filter.List)) *filter.List {
	filterList := filter.NewList(1)
	for _, option := range opts {
		option(filterList)
	}
	return filterList
}

func OptionUserPrefixFilter(intKey []byte) func(*filter.List) {
	pf := filter.NewPrefixFilter(intKey)
	logger.Slog.Infow("Добавляем фильтр по префиксу ключа в список фильтров",
		"intKey", intKey)
	return func(fl *filter.List) {
		fl.AddFilters(pf)
	}
}

func OptionPageFilter(size int64) func(*filter.List) {
	pf := filter.NewPageFilter(size)
	logger.Slog.Infow("Добавляем фильтр пагинации в список фильтров",
		"size", size)
	return func(fl *filter.List) {
		fl.AddFilters(pf)
	}
}

func OptionRowKeyFromFilter(rowKeyFrom string) func(*filter.List) {
	var rrList []*filter.RowRange
	var rr *filter.RowRange
	if rowKeyFrom == "NULL" {
		return func(fl *filter.List) {}
	} else {
		rr = filter.NewRowRange([]byte(rowKeyFrom), nil, false, true)
	}
	rrList = append(rrList, rr)
	mrrFilter := filter.NewMultiRowRangeFilter(rrList)
	logger.Slog.Infow("Добавляем фильтр диапазона ключей, начинающийся с ключа",
		"rowKeyFrom", rowKeyFrom)
	return func(fl *filter.List) {
		fl.AddFilters(mrrFilter)
	}
}
