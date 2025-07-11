// Code generated by ent, DO NOT EDIT.

package protocolstat

import (
	"entgo.io/ent/dialect/sql"
	"github.com/tsisar/starknet-indexer/generated/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldContainsFold(FieldID, id))
}

// TotalSupply applies equality check predicate on the "totalSupply" field. It's identical to TotalSupplyEQ.
func TotalSupply(v string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldEQ(FieldTotalSupply, v))
}

// Tvl applies equality check predicate on the "tvl" field. It's identical to TvlEQ.
func Tvl(v string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldEQ(FieldTvl, v))
}

// Pools applies equality check predicate on the "pools" field. It's identical to PoolsEQ.
func Pools(v string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldEQ(FieldPools, v))
}

// TotalSupplyEQ applies the EQ predicate on the "totalSupply" field.
func TotalSupplyEQ(v string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldEQ(FieldTotalSupply, v))
}

// TotalSupplyNEQ applies the NEQ predicate on the "totalSupply" field.
func TotalSupplyNEQ(v string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldNEQ(FieldTotalSupply, v))
}

// TotalSupplyIn applies the In predicate on the "totalSupply" field.
func TotalSupplyIn(vs ...string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldIn(FieldTotalSupply, vs...))
}

// TotalSupplyNotIn applies the NotIn predicate on the "totalSupply" field.
func TotalSupplyNotIn(vs ...string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldNotIn(FieldTotalSupply, vs...))
}

// TotalSupplyGT applies the GT predicate on the "totalSupply" field.
func TotalSupplyGT(v string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldGT(FieldTotalSupply, v))
}

// TotalSupplyGTE applies the GTE predicate on the "totalSupply" field.
func TotalSupplyGTE(v string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldGTE(FieldTotalSupply, v))
}

// TotalSupplyLT applies the LT predicate on the "totalSupply" field.
func TotalSupplyLT(v string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldLT(FieldTotalSupply, v))
}

// TotalSupplyLTE applies the LTE predicate on the "totalSupply" field.
func TotalSupplyLTE(v string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldLTE(FieldTotalSupply, v))
}

// TotalSupplyContains applies the Contains predicate on the "totalSupply" field.
func TotalSupplyContains(v string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldContains(FieldTotalSupply, v))
}

// TotalSupplyHasPrefix applies the HasPrefix predicate on the "totalSupply" field.
func TotalSupplyHasPrefix(v string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldHasPrefix(FieldTotalSupply, v))
}

// TotalSupplyHasSuffix applies the HasSuffix predicate on the "totalSupply" field.
func TotalSupplyHasSuffix(v string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldHasSuffix(FieldTotalSupply, v))
}

// TotalSupplyEqualFold applies the EqualFold predicate on the "totalSupply" field.
func TotalSupplyEqualFold(v string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldEqualFold(FieldTotalSupply, v))
}

// TotalSupplyContainsFold applies the ContainsFold predicate on the "totalSupply" field.
func TotalSupplyContainsFold(v string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldContainsFold(FieldTotalSupply, v))
}

// TvlEQ applies the EQ predicate on the "tvl" field.
func TvlEQ(v string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldEQ(FieldTvl, v))
}

// TvlNEQ applies the NEQ predicate on the "tvl" field.
func TvlNEQ(v string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldNEQ(FieldTvl, v))
}

// TvlIn applies the In predicate on the "tvl" field.
func TvlIn(vs ...string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldIn(FieldTvl, vs...))
}

// TvlNotIn applies the NotIn predicate on the "tvl" field.
func TvlNotIn(vs ...string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldNotIn(FieldTvl, vs...))
}

// TvlGT applies the GT predicate on the "tvl" field.
func TvlGT(v string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldGT(FieldTvl, v))
}

// TvlGTE applies the GTE predicate on the "tvl" field.
func TvlGTE(v string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldGTE(FieldTvl, v))
}

// TvlLT applies the LT predicate on the "tvl" field.
func TvlLT(v string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldLT(FieldTvl, v))
}

// TvlLTE applies the LTE predicate on the "tvl" field.
func TvlLTE(v string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldLTE(FieldTvl, v))
}

// TvlContains applies the Contains predicate on the "tvl" field.
func TvlContains(v string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldContains(FieldTvl, v))
}

// TvlHasPrefix applies the HasPrefix predicate on the "tvl" field.
func TvlHasPrefix(v string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldHasPrefix(FieldTvl, v))
}

// TvlHasSuffix applies the HasSuffix predicate on the "tvl" field.
func TvlHasSuffix(v string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldHasSuffix(FieldTvl, v))
}

// TvlEqualFold applies the EqualFold predicate on the "tvl" field.
func TvlEqualFold(v string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldEqualFold(FieldTvl, v))
}

// TvlContainsFold applies the ContainsFold predicate on the "tvl" field.
func TvlContainsFold(v string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldContainsFold(FieldTvl, v))
}

// PoolsEQ applies the EQ predicate on the "pools" field.
func PoolsEQ(v string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldEQ(FieldPools, v))
}

// PoolsNEQ applies the NEQ predicate on the "pools" field.
func PoolsNEQ(v string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldNEQ(FieldPools, v))
}

// PoolsIn applies the In predicate on the "pools" field.
func PoolsIn(vs ...string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldIn(FieldPools, vs...))
}

// PoolsNotIn applies the NotIn predicate on the "pools" field.
func PoolsNotIn(vs ...string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldNotIn(FieldPools, vs...))
}

// PoolsGT applies the GT predicate on the "pools" field.
func PoolsGT(v string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldGT(FieldPools, v))
}

// PoolsGTE applies the GTE predicate on the "pools" field.
func PoolsGTE(v string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldGTE(FieldPools, v))
}

// PoolsLT applies the LT predicate on the "pools" field.
func PoolsLT(v string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldLT(FieldPools, v))
}

// PoolsLTE applies the LTE predicate on the "pools" field.
func PoolsLTE(v string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldLTE(FieldPools, v))
}

// PoolsContains applies the Contains predicate on the "pools" field.
func PoolsContains(v string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldContains(FieldPools, v))
}

// PoolsHasPrefix applies the HasPrefix predicate on the "pools" field.
func PoolsHasPrefix(v string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldHasPrefix(FieldPools, v))
}

// PoolsHasSuffix applies the HasSuffix predicate on the "pools" field.
func PoolsHasSuffix(v string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldHasSuffix(FieldPools, v))
}

// PoolsEqualFold applies the EqualFold predicate on the "pools" field.
func PoolsEqualFold(v string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldEqualFold(FieldPools, v))
}

// PoolsContainsFold applies the ContainsFold predicate on the "pools" field.
func PoolsContainsFold(v string) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.FieldContainsFold(FieldPools, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ProtocolStat) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ProtocolStat) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ProtocolStat) predicate.ProtocolStat {
	return predicate.ProtocolStat(sql.NotPredicates(p))
}
