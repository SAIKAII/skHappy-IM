package dao

import (
	"fmt"
	"github.com/SAIKAII/skHappy-IM/infra/base"
	"github.com/jinzhu/gorm"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGroupDao_Test(t *testing.T) {
	Convey("测试创建一个群组", t, func() {
		var (
			id  uint64
			err error
		)
		db := base.Database()
		db.Transaction(func(tx *gorm.DB) error {
			groupDao := &GroupDao{DB: db}
			group := &Group{
				GroupName:    "TestGroup1",
				CreateUser:   "ituen,vlos1",
				Owner:        "ituen,vlos1",
				Announcement: "Hello world",
				UserNum:      0,
				IsDeleted:    0,
			}
			id, err = groupDao.InsertOne(group)

			So(err, ShouldBeNil)
			So(id, ShouldNotEqual, 0)

			return nil
		})

		groupDao := &GroupDao{DB: db}
		g, err := groupDao.GetOne(id)
		So(err, ShouldBeNil)
		So(g, ShouldNotBeNil)

		fmt.Println(g)
	})
}

func TestGroupUserDao_InsertOne(t *testing.T) {
	Convey("测试创建群组用户", t, func() {
		db := base.Database()
		groupUserDao := &GroupUserDao{DB: db}
		groupUser := &GroupUser{
			GroupId:  10,
			Username: "qffqwrtb231",
		}
		err := groupUserDao.InsertOne(groupUser)
		So(err, ShouldBeNil)
	})
}

func TestGroupUserDao_GetAll(t *testing.T) {
	Convey("测试获取群组所有群员", t, func() {
		db := base.Database()
		groupUserDao := &GroupUserDao{DB: db}
		users, err := groupUserDao.GetAll(10)
		So(err, ShouldBeNil)
		So(users, ShouldNotBeNil)

		for _, v := range users {
			fmt.Println(v)
		}
	})
}

func TestGroupDao_UserNum(t *testing.T) {
	Convey("测试获取不存在群组的人数", t, func() {
		db := base.Database()
		groupDao := &GroupDao{DB: db}
		num, err := groupDao.UserNum(123)
		So(err, ShouldNotBeNil)
		So(num, ShouldEqual, 0)
	})
}
