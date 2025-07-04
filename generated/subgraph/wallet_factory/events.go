// Code generated by generator-mapper; DO NOT EDIT.

package wallet_factory

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
type WalletDeployedKey struct {
	Wallet string `json:"wallet"`
	Admin  string `json:"admin"`
}
type WalletDeployed struct {
	Key WalletDeployedKey `json:"key"`
}
type LogAddToWhitelistKey struct {
	User string `json:"user"`
}
type LogAddToWhitelist struct {
	Key LogAddToWhitelistKey `json:"key"`
}
type LogRemoveFromWhitelistKey struct {
	User string `json:"user"`
}
type LogRemoveFromWhitelist struct {
	Key LogRemoveFromWhitelistKey `json:"key"`
}
type LogSetDecentralizedModeData struct {
	IsOn bool `json:"is_on"`
}
type LogSetDecentralizedMode struct {
	Data LogSetDecentralizedModeData `json:"data"`
}
