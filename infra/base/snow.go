package base

import (
	"github.com/GUAIK-ORG/go-snowflake/snowflake"
	"github.com/SAIKAII/skHappy-IM/infra"
)

var snow *snowflake.Snowflake

func Snow() *snowflake.Snowflake {
	return snow
}

type SnowflakeStarter struct {
	infra.BaseStarter
}

func (s *SnowflakeStarter) Init(ctx infra.StarterContext) {
	var err error
	snow, err = snowflake.NewSnowflake(int64(0), int64(0))
	if err != nil {
		panic(err)
	}
}
