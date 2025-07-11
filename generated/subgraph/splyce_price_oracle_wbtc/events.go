// Code generated by generator-mapper; DO NOT EDIT.

package splyce_price_oracle_wbtc

type RoleGrantedData struct {
	Role    string `json:"role"`
	Account string `json:"account"`
	Sender  string `json:"sender"`
}
type RoleGranted struct {
	Data RoleGrantedData `json:"data"`
}
type RoleRevokedData struct {
	Role    string `json:"role"`
	Account string `json:"account"`
	Sender  string `json:"sender"`
}
type RoleRevoked struct {
	Data RoleRevokedData `json:"data"`
}
type RoleAdminChangedData struct {
	Role              string `json:"role"`
	PreviousAdminRole string `json:"previous_admin_role"`
	NewAdminRole      string `json:"new_admin_role"`
}
type RoleAdminChanged struct {
	Data RoleAdminChangedData `json:"data"`
}
type PausedData struct {
	Account string `json:"account"`
}
type Paused struct {
	Data PausedData `json:"data"`
}
type UnpausedData struct {
	Account string `json:"account"`
}
type Unpaused struct {
	Data UnpausedData `json:"data"`
}
type LogSetOracleAggregatorData struct {
	OracleAggregator string `json:"oracle_aggregator"`
}
type LogSetOracleAggregator struct {
	Data LogSetOracleAggregatorData `json:"data"`
}
