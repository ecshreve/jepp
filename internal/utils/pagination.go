package utils

import (
	"entgo.io/contrib/entgql"
	gqlserver "github.com/ecshreve/jepp/internal/gqlserver/gen"
	"github.com/sirupsen/logrus"
)

func JeppPagerToEntPager(pager *gqlserver.PaginationParams) (*entgql.Cursor[int], *int, *entgql.Cursor[int], *int) {
	limit := 10
	if pager == nil {
		return nil, &limit, nil, nil
	}

	if pager.Limit != nil && *pager.Limit > 0 && *pager.Limit <= 50 {
		limit = *pager.Limit
	} else {
		logrus.Info("invalid limit provided, defaulting to 10")
	}

	if pager.Before != nil {
		return nil, nil, pager.Before, &limit
	}

	return nil, &limit, nil, nil
}
