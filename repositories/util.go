package repositories

import (
	"github.com/Masterminds/squirrel"
	"strconv"
)

func applyFilterAndPageSize(builder squirrel.SelectBuilder, params map[string][]string) (squirrel.SelectBuilder, error) {
	page, limit := 0, 0
	if len(params) > 0 {
		for k, v := range params {
			if k == "page" && len(v) > 0 {
				page, _ = strconv.Atoi(v[0])
				continue
			}
			if k == "limit" && len(v) > 0 {
				limit, _ = strconv.Atoi(v[0])
				continue
			}
			builder = builder.Where(squirrel.Eq{k: v})
		}
	}
	if limit != 0 && page != 0 {
		builder = builder.Limit(uint64(limit)).Offset(uint64(limit * (page - 1)))
	}
	return builder, nil
}
