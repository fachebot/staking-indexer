package config

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/zeromicro/go-zero/rest"
	"staking-indexer/pkg/pgsql"
)

type Config struct {
	rest.RestConf
	Postgres        pgsql.PgConfig
	StakingContract common.Address
}
