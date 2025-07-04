// Code generated by generator-mapper; DO NOT EDIT.

package price_consumer

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
type LogSetPriceKey struct {
	CollateralPoolId string `json:"collateral_pool_id"`
}
type LogSetPriceData struct {
	RawPrice              big.Int `json:"raw_price"`
	PriceWithSafetyMargin big.Int `json:"price_with_safety_margin"`
}
type LogSetPrice struct {
	Key  LogSetPriceKey  `json:"key"`
	Data LogSetPriceData `json:"data"`
}
type LogSetPriceForBatchData struct {
	CollateralPoolIds string `json:"collateral_pool_ids"`
}
type LogSetPriceForBatch struct {
	Data LogSetPriceForBatchData `json:"data"`
}
type LogSetStableCoinReferencePriceKey struct {
	Caller string `json:"caller"`
}
type LogSetStableCoinReferencePriceData struct {
	Data big.Int `json:"data"`
}
type LogSetStableCoinReferencePrice struct {
	Key  LogSetStableCoinReferencePriceKey  `json:"key"`
	Data LogSetStableCoinReferencePriceData `json:"data"`
}
type LogSetBookKeeperData struct {
	NewAddress string `json:"new_address"`
}
type LogSetBookKeeper struct {
	Data LogSetBookKeeperData `json:"data"`
}
type LogCage struct {
}
