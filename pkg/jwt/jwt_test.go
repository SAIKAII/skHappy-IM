package jwt

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test(t *testing.T) {
	Convey("测试JWT生成与验证", t, func() {
		m := make(map[string]interface{})
		m["username"] = "SAIKAII"
		jwtString, err := NewJWT(m)
		ShouldBeNil(err)
		ShouldNotEqual(jwtString, "")

		fmt.Println(jwtString)

		res, err := VerifyJWT(jwtString)
		So(err, ShouldBeNil)
		So(res, ShouldNotBeNil)
		fmt.Println(res)
	})
}
