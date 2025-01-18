package api

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/s21platform/gateway-service/internal/config"
	logger_lib "github.com/s21platform/logger-lib"
	userproto "github.com/s21platform/user-proto/user-proto"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApi_GetProfile(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)

	ctx := context.Background()
	mockLogger := logger_lib.NewMockLoggerInterface(ctrl)

	t.Run("should_ok", func(t *testing.T) {
		mockUserService := NewMockUserService(ctrl)
		mockLogger.EXPECT().AddFuncName("MyProfile")

		ctx = context.WithValue(ctx, config.KeyLogger, mockLogger)
		r := &http.Request{}
		w := httptest.NewRecorder()
		r = r.WithContext(ctx)

		userInfo := &userproto.GetUserInfoByUUIDOut{
			Nickname: "",
			Avatar:   "",
		}

		mockUserService.EXPECT().GetInfoByUUID(r).Return(userInfo, nil)

		s := New(
			mockUserService,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
		)

		s.MyProfile(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should_err_us_fail_response", func(t *testing.T) {
		mockUserService := NewMockUserService(ctrl)
		mockLogger.EXPECT().AddFuncName("MyProfile")
		mockLogger.EXPECT().Error(gomock.Any())

		ctx = context.WithValue(ctx, config.KeyLogger, mockLogger)
		r := &http.Request{}
		w := httptest.NewRecorder()
		r = r.WithContext(ctx)

		mockErr := errors.New("some error")

		mockUserService.EXPECT().GetInfoByUUID(r).Return(nil, mockErr)

		s := New(
			mockUserService,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
		)

		s.MyProfile(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

	})
}
