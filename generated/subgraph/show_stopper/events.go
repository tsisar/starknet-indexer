// Code generated by generator-mapper; DO NOT EDIT.

package show_stopper

import "math/big"

type LogCageData struct {
	CageCoolDown big.Int `json:"cage_cool_down"`
}
type LogCage struct {
	Data LogCageData `json:"data"`
}
type LogCageCollateralPoolKey struct {
	CollateralPoolId string `json:"collateral_pool_id"`
}
type LogCageCollateralPool struct {
	Key LogCageCollateralPoolKey `json:"key"`
}
type LogAccumulateBadDebtKey struct {
	CollateralPoolId string `json:"collateral_pool_id"`
	PositionAddress  string `json:"position_address"`
}
type LogAccumulateBadDebtData struct {
	Amount    big.Int `json:"amount"`
	DebtShare big.Int `json:"debt_share"`
}
type LogAccumulateBadDebt struct {
	Key  LogAccumulateBadDebtKey  `json:"key"`
	Data LogAccumulateBadDebtData `json:"data"`
}
type LogRedeemLockedCollateralKey struct {
	CollateralPoolId string `json:"collateral_pool_id"`
	PositionAddress  string `json:"position_address"`
}
type LogRedeemLockedCollateralData struct {
	Amount big.Int `json:"amount"`
}
type LogRedeemLockedCollateral struct {
	Key  LogRedeemLockedCollateralKey  `json:"key"`
	Data LogRedeemLockedCollateralData `json:"data"`
}
type LogFinalizeDebt struct {
}
type LogFinalizeCashPriceKey struct {
	CollateralPoolId string `json:"collateral_pool_id"`
}
type LogFinalizeCashPrice struct {
	Key LogFinalizeCashPriceKey `json:"key"`
}
type LogAccumulateStablecoinKey struct {
	OwnerAddress string `json:"owner_address"`
}
type LogAccumulateStablecoinData struct {
	Amount big.Int `json:"amount"`
}
type LogAccumulateStablecoin struct {
	Key  LogAccumulateStablecoinKey  `json:"key"`
	Data LogAccumulateStablecoinData `json:"data"`
}
type LogRedeemStablecoinKey struct {
	CollateralPoolId string `json:"collateral_pool_id"`
	OwnerAddress     string `json:"owner_address"`
}
type LogRedeemStablecoinData struct {
	Amount big.Int `json:"amount"`
}
type LogRedeemStablecoin struct {
	Key  LogRedeemStablecoinKey  `json:"key"`
	Data LogRedeemStablecoinData `json:"data"`
}
type LogSetBookKeeperKey struct {
	Caller string `json:"caller"`
}
type LogSetBookKeeperData struct {
	Data string `json:"data"`
}
type LogSetBookKeeper struct {
	Key  LogSetBookKeeperKey  `json:"key"`
	Data LogSetBookKeeperData `json:"data"`
}
type LogSetLiquidationEngineKey struct {
	Caller string `json:"caller"`
}
type LogSetLiquidationEngineData struct {
	Data string `json:"data"`
}
type LogSetLiquidationEngine struct {
	Key  LogSetLiquidationEngineKey  `json:"key"`
	Data LogSetLiquidationEngineData `json:"data"`
}
type LogSetSystemDebtEngineKey struct {
	Caller string `json:"caller"`
}
type LogSetSystemDebtEngineData struct {
	Data string `json:"data"`
}
type LogSetSystemDebtEngine struct {
	Key  LogSetSystemDebtEngineKey  `json:"key"`
	Data LogSetSystemDebtEngineData `json:"data"`
}
type LogSetPriceConsumerKey struct {
	Caller string `json:"caller"`
}
type LogSetPriceConsumerData struct {
	Data string `json:"data"`
}
type LogSetPriceConsumer struct {
	Key  LogSetPriceConsumerKey  `json:"key"`
	Data LogSetPriceConsumerData `json:"data"`
}
type LogSetCageCoolDownKey struct {
	Caller string `json:"caller"`
}
type LogSetCageCoolDownData struct {
	Data big.Int `json:"data"`
}
type LogSetCageCoolDown struct {
	Key  LogSetCageCoolDownKey  `json:"key"`
	Data LogSetCageCoolDownData `json:"data"`
}
