package user

import (
	_ "github.com/go-sql-driver/mysql" //导入mysql驱动
	"github.com/gone-io/gone"
	"github.com/gone-io/gone/goner"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"template_module/internal/interface/entity"
	"template_module/internal/interface/mock"
	"testing"
)

func Test_iUser_Register(t *testing.T) {
	gone.RunTest(func(in struct {
		iUser      *iUser               `gone:"*"` //inject iUser for test
		iDependent *mock.MockIDependent `gone:"*"` //inject iDependent for mock
	}) {
		err := gone.ToError("err")
		in.iDependent.EXPECT().DoSomething().Return(err)

		register, err2 := in.iUser.Register(&entity.RegisterParam{
			Username: "test",
			Password: "test",
		})
		assert.Nil(t, register)
		assert.Equal(t, err2, err)
	}, func(cemetery gone.Cemetery) error {
		controller := gomock.NewController(t)

		//load all mocked components
		mock.MockPriest(cemetery, controller)

		_ = goner.XormPriest(cemetery)

		//bury the tested component
		cemetery.Bury(&iUser{})

		return nil
	})
}
