// Code generated by generator-mapper; DO NOT EDIT.

package collateral_adapter_strk

import "math/big"

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
type LogDepositKey struct {
	Caller          string `json:"caller"`
	PositionAddress string `json:"position_address"`
}
type LogDepositData struct {
	Amount big.Int `json:"amount"`
}
type LogDeposit struct {
	Key  LogDepositKey  `json:"key"`
	Data LogDepositData `json:"data"`
}
type LogWithdrawKey struct {
	Caller          string `json:"caller"`
	PositionAddress string `json:"position_address"`
}
type LogWithdrawData struct {
	Amount big.Int `json:"amount"`
}
type LogWithdraw struct {
	Key  LogWithdrawKey  `json:"key"`
	Data LogWithdrawData `json:"data"`
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
type LogEmergencyWithdrawKey struct {
	Caller string `json:"caller"`
	To     string `json:"to"`
}
type LogEmergencyWithdrawData struct {
	Amount big.Int `json:"amount"`
}
type LogEmergencyWithdraw struct {
	Key  LogEmergencyWithdrawKey  `json:"key"`
	Data LogEmergencyWithdrawData `json:"data"`
}
type LogCage struct {
}
