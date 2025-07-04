// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent/dialect/sql"
	"github.com/tsisar/starknet-indexer/generated/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldID, id))
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAddress, v))
}

// ActivePositionsCount applies equality check predicate on the "activePositionsCount" field. It's identical to ActivePositionsCountEQ.
func ActivePositionsCount(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldActivePositionsCount, v))
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAddress, v))
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldAddress, v))
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldAddress, vs...))
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldAddress, vs...))
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldAddress, v))
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldAddress, v))
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldAddress, v))
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldAddress, v))
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldAddress, v))
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldAddress, v))
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldAddress, v))
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldAddress, v))
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldAddress, v))
}

// ActivePositionsCountEQ applies the EQ predicate on the "activePositionsCount" field.
func ActivePositionsCountEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldActivePositionsCount, v))
}

// ActivePositionsCountNEQ applies the NEQ predicate on the "activePositionsCount" field.
func ActivePositionsCountNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldActivePositionsCount, v))
}

// ActivePositionsCountIn applies the In predicate on the "activePositionsCount" field.
func ActivePositionsCountIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldActivePositionsCount, vs...))
}

// ActivePositionsCountNotIn applies the NotIn predicate on the "activePositionsCount" field.
func ActivePositionsCountNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldActivePositionsCount, vs...))
}

// ActivePositionsCountGT applies the GT predicate on the "activePositionsCount" field.
func ActivePositionsCountGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldActivePositionsCount, v))
}

// ActivePositionsCountGTE applies the GTE predicate on the "activePositionsCount" field.
func ActivePositionsCountGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldActivePositionsCount, v))
}

// ActivePositionsCountLT applies the LT predicate on the "activePositionsCount" field.
func ActivePositionsCountLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldActivePositionsCount, v))
}

// ActivePositionsCountLTE applies the LTE predicate on the "activePositionsCount" field.
func ActivePositionsCountLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldActivePositionsCount, v))
}

// ActivePositionsCountContains applies the Contains predicate on the "activePositionsCount" field.
func ActivePositionsCountContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldActivePositionsCount, v))
}

// ActivePositionsCountHasPrefix applies the HasPrefix predicate on the "activePositionsCount" field.
func ActivePositionsCountHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldActivePositionsCount, v))
}

// ActivePositionsCountHasSuffix applies the HasSuffix predicate on the "activePositionsCount" field.
func ActivePositionsCountHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldActivePositionsCount, v))
}

// ActivePositionsCountEqualFold applies the EqualFold predicate on the "activePositionsCount" field.
func ActivePositionsCountEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldActivePositionsCount, v))
}

// ActivePositionsCountContainsFold applies the ContainsFold predicate on the "activePositionsCount" field.
func ActivePositionsCountContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldActivePositionsCount, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(sql.NotPredicates(p))
}
