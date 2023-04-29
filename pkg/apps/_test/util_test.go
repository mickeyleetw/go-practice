package apps

import (
	"os"
	"testing"

	"senao/pkg/apps"
	"senao/pkg/domain/mock"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func setup() {
	accountUsecase := mock.NewAccountUsecase()
	accountHandler := apps.NewAccountHandler(accountUsecase)
	r = gin.Default()

	r.POST("/account", accountHandler.CreateAccount)
}

func shutdown() {
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}
