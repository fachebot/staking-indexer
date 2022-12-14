// Code generated by goctl. DO NOT EDIT.
package types

type Token struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals uint8  `json:"decimals"`
	Contract string `json:"contract"`
}

type PoolInfo struct {
	StakeToken          Token  `json:"stake_token"`
	RewardToken         Token  `json:"reward_token"`
	TokenPerBlock       string `json:"token_per_block"`
	StartBlock          uint64 `json:"start_block"`
	EndBlock            uint64 `json:"end_block"`
	StartBlockTimestamp *int64 `json:"start_block_timestamp,omitempty"`
	EndBlockTimestamp   *int64 `json:"start_block_timestamp,omitempty"`
}

type GetPoolRequest struct {
}

type GetPoolResponse struct {
	Pool PoolInfo `json:"pool"`
}
