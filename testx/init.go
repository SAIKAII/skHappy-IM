package testx

import (
	"github.com/SAIKAII/skHappy-IM/infra"
	"github.com/SAIKAII/skHappy-IM/infra/base"
)

func init() {
	b := infra.New()

	infra.Register(&base.DatabaseStarter{})
	//infra.Register(&base.RedisStarter{})
	//infra.Register(&base.ConnStarter{})
	b.Start()
}
