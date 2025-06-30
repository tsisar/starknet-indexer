package resolvers

import (
	"entgo.io/ent/dialect/sql"
	"github.com/tsisar/starknet-indexer/generated/ent/position"
)

// Order by position_id as bigint (numeric order)
func OrderByPositionIDAsBigInt(desc bool) position.OrderOption {
	return func(s *sql.Selector) {
		order := "ASC"
		if desc {
			order = "DESC"
		}
		s.OrderBy("CAST(position_id AS BIGINT) " + order)
	}
}
