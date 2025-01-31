package api

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	society_proto "github.com/s21platform/society-proto/society-proto"

	"github.com/golang/mock/gomock"
	"github.com/s21platform/gateway-service/internal/config"
	logger_lib "github.com/s21platform/logger-lib"
	userproto "github.com/s21platform/user-proto/user-proto"
	"github.com/stretchr/testify/assert"
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

func TestApi_GetSocietyInfo(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)

	ctx := context.Background()
	mockLogger := logger_lib.NewMockLoggerInterface(ctrl)

	t.Run("should_ok", func(t *testing.T) {
		MockSocietyService := NewMockSocietyService(ctrl)
		mockLogger.EXPECT().AddFuncName("SocietyInfo")

		ctx = context.WithValue(ctx, config.KeyLogger, mockLogger)
		r := &http.Request{}
		w := httptest.NewRecorder()
		r = r.WithContext(ctx)

		societyInfo := &society_proto.GetSocietyInfoOut{
			Name:             "test",
			Description:      "test",
			OwnerUUID:        "test",
			PhotoUrl:         "test",
			IsPrivate:        true,
			CountSubscribers: 0,
		}

		MockSocietyService.EXPECT().GetSocietyInfo(r).Return(societyInfo, nil)

		s := New(
			nil,
			nil,
			nil,
			nil,
			nil,
			MockSocietyService,
			nil,
			nil,
		)

		s.GetSocietyInfo(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should_err_us_fail_response", func(t *testing.T) {
		MockSocietyService := NewMockSocietyService(ctrl)
		mockLogger.EXPECT().AddFuncName("SocietyInfo")
		mockLogger.EXPECT().Error(gomock.Any())

		ctx = context.WithValue(ctx, config.KeyLogger, mockLogger)
		r := &http.Request{}
		w := httptest.NewRecorder()
		r = r.WithContext(ctx)

		mockErr := errors.New("some error")

		MockSocietyService.EXPECT().GetSocietyInfo(r).Return(nil, mockErr)

		s := New(
			nil,
			nil,
			nil,
			nil,
			nil,
			MockSocietyService,
			nil,
			nil,
		)

		s.GetSocietyInfo(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
