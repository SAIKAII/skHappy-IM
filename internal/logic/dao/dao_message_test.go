package dao

import (
	"fmt"
	"github.com/SAIKAII/skHappy-IM/infra/base"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMessageDao_GetAllRecvByLastSeqId(t *testing.T) {
	db := base.Database()
	Convey("测试获取大于seqId的信息", t, func() {
		messageDao := &MessageDao{DB: db}
		username := "qffqwrtb231"
		var lastSeqId uint64
		lastSeqId = 69
		res, err := messageDao.GetAllRecvByLastSeqId(username, lastSeqId)
		ShouldBeNil(err)
		ShouldNotBeNil(res)
		for _, v := range res {
			fmt.Println("[From]", v.Sender, "[To]", v.Receiver, "[Time]", v.SendTime, "[SeqId]", v.SeqId)
			fmt.Println("[Content]", v.Content)
		}
	})
}
