package dao

import (
	"fmt"
	"github.com/SAIKAII/skHappy-IM/infra/base"
	_ "github.com/SAIKAII/skHappy-IM/testx"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAccountDao_GetOne(t *testing.T) {
	db := base.Database()
	Convey("测试查询用户信息", t, func() {
		dao := &UserDao{DB: db}
		u1, err := dao.GetOne("fsnod10fn")
		ShouldBeNil(err)
		ShouldNotBeNil(u1)
		fmt.Println(u1)

		u2, err := dao.GetOne("ss")
		So(err, ShouldNotBeNil)
		So(u2, ShouldBeNil)
	})
}

func TestAccountDao_Insert(t *testing.T) {
	db := base.Database()
	Convey("测试插入用户", t, func() {
		dao := &UserDao{DB: db}
		u := &User{
			Username: "fsnod1ancp1",
			Nickname: "Santa",
			Password: "spqn21r-v",
			Salt:     "1v-1nv",
			Sex:      4,
		}
		err := dao.Insert(u)
		So(err, ShouldBeNil)
	})
}
