package apps

import (
	"net/http"
	"net/http/httptest"
	"senao/pkg/domain"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
	"github.com/parnurzeal/gorequest"
	"github.com/stretchr/testify/assert"
)

func Test_accountHandler_CreateAccount(t *testing.T) {
	type fields struct {
		accountUsecase domain.AccountUsecase
	}
	type args struct {
		c *gin.Context
	}
	tcs := []struct {
		name          string
		inpBody       string
		expHttpStatus int
	}{
		{
			name:          "success",
			inpBody:       `{"name":"aaa","password":"SenaoPretest01"}`,
			expHttpStatus: http.StatusCreated,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := gorequest.New().Post("/subscribers").Send(tc.inpBody).MakeRequest()
			r.ServeHTTP(w, req)
			assert.Equal(t, tc.expHttpStatus, w.Code, w.Code)
			// fmt.Printf("%s\n", w.Body.String())
			assert.True(t, cmp.Equal(tc.expBody, w.Body.String()), cmp.Diff(tc.expBody, w.Body.String()))
		})
	}
}
