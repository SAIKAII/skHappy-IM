package dao

import (
	"github.com/SAIKAII/skHappy-IM/infra/base"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRelationShipDao_Insert(t *testing.T) {
	db := base.Database()
	Convey("测试关系表插入记录", t, func() {
		dao := RelationShipDao{DB: db}
		rel := Relationship{
			UserA: "Scatt",
			UserB: "saikaii",
		}
		err := dao.Insert(&rel)
		So(err, ShouldBeNil)
	})
}

func TestRelationShipDao_GetAll(t *testing.T) {
	db := base.Database()
	Convey("测试关系表查询指定用户所有联系人", t, func() {
		dao := RelationShipDao{DB: db}
		users, err := dao.GetAll("saikaii")
		So(err, ShouldBeNil)
		So(users, ShouldNotBeNil)

		for _, u := range users {
			Println(u)
		}
	})
}
