package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"

	advertproto "github.com/s21platform/advert-proto/advert-proto"
	chatproto "github.com/s21platform/chat-proto/chat-proto"
	logger_lib "github.com/s21platform/logger-lib"
	societyproto "github.com/s21platform/society-proto/society-proto"
	userproto "github.com/s21platform/user-proto/user-proto"

	"github.com/s21platform/gateway-service/internal/config"
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
			nil,
		)

		s.MyProfile(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestApi_CreateSociety(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)

	ctx := context.Background()
	mockLogger := logger_lib.NewMockLoggerInterface(ctrl)
	t.Run("should_ok", func(t *testing.T) {
		mockSocietyService := NewMockSocietyService(ctrl)

		in := &societyproto.SetSocietyIn{
			Name:             "test",
			PostPermissionID: 1,
			FormatID:         2,
			IsSearch:         true,
		}
		body, _ := json.Marshal(in)

		// Создаем HTTP-запрос
		req := httptest.NewRequest(http.MethodPost, "/society", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		ctx = context.WithValue(ctx, config.KeyLogger, mockLogger)
		req = req.WithContext(ctx)

		out := &societyproto.SetSocietyOut{SocietyUUID: "societyUUID"}

		mockLogger.EXPECT().AddFuncName("CreateSociety")
		mockSocietyService.EXPECT().CreateSociety(req).Return(out, nil)

		w := httptest.NewRecorder()

		s := New(
			nil,
			nil,
			nil,
			nil,
			nil,
			mockSocietyService,
			nil,
			nil,
			nil,
		)

		s.CreateSociety(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

		var responseBody societyproto.SetSocietyOut
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		assert.NoError(t, err)
		assert.Equal(t, out.SocietyUUID, responseBody.SocietyUUID)
	})
	t.Run("should_return_internal_server_error_if_CreateSociety_fails", func(t *testing.T) {
		mockSocietyService := NewMockSocietyService(ctrl)

		in := &societyproto.SetSocietyIn{
			Name:             "test",
			PostPermissionID: 1,
			FormatID:         2,
			IsSearch:         true,
		}
		body, _ := json.Marshal(in)

		req := httptest.NewRequest(http.MethodPost, "/society", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		ctx = context.WithValue(ctx, config.KeyLogger, mockLogger)
		req = req.WithContext(ctx)

		expectedError := errors.New("database error")
		mockLogger.EXPECT().AddFuncName("CreateSociety")
		mockLogger.EXPECT().Error("failed to create society error: database error")
		mockSocietyService.EXPECT().CreateSociety(req).Return(nil, expectedError)

		w := httptest.NewRecorder()

		s := New(
			nil,
			nil,
			nil,
			nil,
			nil,
			mockSocietyService,
			nil,
			nil,
			nil,
		)

		s.CreateSociety(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestApi_GetSocietyInfo(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	mockLogger := logger_lib.NewMockLoggerInterface(ctrl)

	t.Run("should_get_society_info_successfully", func(t *testing.T) {
		mockSocietyService := NewMockSocietyService(ctrl)

		societyUUID := "test-uuid"
		req := httptest.NewRequest(http.MethodGet, "/society/"+societyUUID, nil)
		req.Header.Set("Content-Type", "application/json")

		ctx = context.WithValue(ctx, config.KeyLogger, mockLogger)
		req = req.WithContext(ctx)

		expectedResult := &societyproto.GetSocietyInfoOut{
			Name:           "Test Society",
			Description:    "Test Description",
			OwnerUUID:      "owner-uuid",
			PhotoURL:       "https://example.com/photo.jpg",
			FormatID:       1,
			PostPermission: 2,
			IsSearch:       true,
		}

		mockLogger.EXPECT().AddFuncName("GetSocietyInfo")
		mockSocietyService.EXPECT().GetSocietyInfo(req).Return(expectedResult, nil)

		w := httptest.NewRecorder()

		s := New(
			nil,
			nil,
			nil,
			nil,
			nil,
			mockSocietyService,
			nil,
			nil,
			nil,
		)

		s.GetSocietyInfo(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

		var responseBody societyproto.GetSocietyInfoOut
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		assert.NoError(t, err)
		assert.Equal(t, expectedResult, &responseBody)
	})

	t.Run("should_return_internal_server_error_if_GetSocietyInfo_fails", func(t *testing.T) {
		mockSocietyService := NewMockSocietyService(ctrl)

		societyUUID := "test-uuid"
		req := httptest.NewRequest(http.MethodGet, "/society/"+societyUUID, nil)
		req.Header.Set("Content-Type", "application/json")

		ctx = context.WithValue(ctx, config.KeyLogger, mockLogger)
		req = req.WithContext(ctx)

		expectedError := errors.New("database error")
		mockLogger.EXPECT().AddFuncName("GetSocietyInfo")
		mockLogger.EXPECT().Error("failed to get society info error: database error") // Ожидаем вызов Error
		mockSocietyService.EXPECT().GetSocietyInfo(req).Return(nil, expectedError)

		w := httptest.NewRecorder()

		s := New(
			nil,
			nil,
			nil,
			nil,
			nil,
			mockSocietyService,
			nil,
			nil,
			nil,
		)

		s.GetSocietyInfo(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestApi_GetAdverts(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)

	ctx := context.Background()
	mockLogger := logger_lib.NewMockLoggerInterface(ctrl)

	t.Run("should_ok", func(t *testing.T) {
		mockAdvertService := NewMockAdvertService(ctrl)
		mockLogger.EXPECT().AddFuncName("GetAdverts")

		ctx = context.WithValue(ctx, config.KeyLogger, mockLogger)
		r := &http.Request{}
		w := httptest.NewRecorder()
		r = r.WithContext(ctx)

		expectedAdverts := &advertproto.GetAdvertsOut{
			Adverts: []*advertproto.AdvertText{
				{
					TextContent: "test",
					ExpiredAt:   timestamppb.New(time.Now()),
				},
			},
		}

		mockAdvertService.EXPECT().GetAdverts(r).Return(expectedAdverts, nil)

		s := New(
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			mockAdvertService,
		)

		s.GetAdverts(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should_err_us_fail_response", func(t *testing.T) {
		mockAdvertService := NewMockAdvertService(ctrl)
		mockLogger.EXPECT().AddFuncName("GetAdverts")
		mockLogger.EXPECT().Error(gomock.Any())

		ctx = context.WithValue(ctx, config.KeyLogger, mockLogger)
		r := &http.Request{}
		w := httptest.NewRecorder()
		r = r.WithContext(ctx)

		mockErr := errors.New("some error")

		mockAdvertService.EXPECT().GetAdverts(r).Return(nil, mockErr)

		s := New(
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			mockAdvertService,
		)

		s.GetAdverts(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestApi_CreatePrivateChat(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)

	ctx := context.Background()
	mockLogger := logger_lib.NewMockLoggerInterface(ctrl)

	t.Run("should_ok", func(t *testing.T) {
		mockChatService := NewMockChatService(ctrl)
		mockLogger.EXPECT().AddFuncName("CreatePrivateChat")

		ctx = context.WithValue(ctx, config.KeyLogger, mockLogger)
		r := &http.Request{}
		w := httptest.NewRecorder()
		r = r.WithContext(ctx)

		expectedChat := &chatproto.CreatePrivateChatOut{
			NewChatUuid: uuid.New().String(),
		}

		mockChatService.EXPECT().CreatePrivateChat(r).Return(expectedChat, nil)

		s := New(
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			mockChatService,
			nil,
		)

		s.CreatePrivateChat(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should_err_us_fail_response", func(t *testing.T) {
		mockChatService := NewMockChatService(ctrl)
		mockLogger.EXPECT().AddFuncName("CreatePrivateChat")
		mockLogger.EXPECT().Error(gomock.Any())

		ctx = context.WithValue(ctx, config.KeyLogger, mockLogger)
		r := &http.Request{}
		w := httptest.NewRecorder()
		r = r.WithContext(ctx)

		mockErr := errors.New("some error")

		mockChatService.EXPECT().CreatePrivateChat(r).Return(nil, mockErr)

		s := New(
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			mockChatService,
			nil,
		)

		s.CreatePrivateChat(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestApi_GetPrivateRecentMessages(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)

	ctx := context.Background()
	mockLogger := logger_lib.NewMockLoggerInterface(ctrl)

	t.Run("should_ok", func(t *testing.T) {
		mockChatService := NewMockChatService(ctrl)
		mockLogger.EXPECT().AddFuncName("GetPrivateRecentMessages")

		ctx = context.WithValue(ctx, config.KeyLogger, mockLogger)
		ctx = context.WithValue(ctx, config.KeyUUID, "test-uuid")

		r := &http.Request{}
		w := httptest.NewRecorder()
		r = r.WithContext(ctx)

		expectedMessages := &chatproto.GetPrivateRecentMessagesOut{
			Messages: []*chatproto.Message{
				{
					Uuid:       uuid.New().String(),
					Content:    "test",
					SentAt:     time.Now().String(),
					UpdatedAt:  "",
					RootUuid:   uuid.New().String(),
					ParentUuid: uuid.New().String(),
				},
				{
					Uuid:       uuid.New().String(),
					Content:    "testing",
					SentAt:     time.Now().String(),
					UpdatedAt:  time.Now().Add(1 * time.Hour).String(),
					RootUuid:   "",
					ParentUuid: "",
				},
			},
		}

		mockChatService.EXPECT().GetPrivateRecentMessages(r).Return(expectedMessages, nil)

		s := New(
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			mockChatService,
			nil,
		)

		s.GetPrivateRecentMessages(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should_err_us_fail_response", func(t *testing.T) {
		mockChatService := NewMockChatService(ctrl)
		mockLogger.EXPECT().AddFuncName("GetPrivateRecentMessages")
		mockLogger.EXPECT().Error(gomock.Any())

		ctx = context.WithValue(ctx, config.KeyLogger, mockLogger)
		ctx = context.WithValue(ctx, config.KeyUUID, "test-uuid")

		r := &http.Request{}
		w := httptest.NewRecorder()
		r = r.WithContext(ctx)

		mockErr := errors.New("some error")

		mockChatService.EXPECT().GetPrivateRecentMessages(r).Return(nil, mockErr)

		s := New(
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			mockChatService,
			nil,
		)

		s.GetPrivateRecentMessages(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestHandler_UpdateSociety(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	mockLogger := logger_lib.NewMockLoggerInterface(ctrl)

	t.Run("should_update_society_successfully", func(t *testing.T) {
		mockSocietyService := NewMockSocietyService(ctrl)

		req := httptest.NewRequest(http.MethodPatch, "/society", nil)
		req.Header.Set("Content-Type", "application/json")

		ctx = context.WithValue(ctx, config.KeyLogger, mockLogger)
		req = req.WithContext(ctx)

		mockLogger.EXPECT().AddFuncName("UpdateSociety")

		mockSocietyService.EXPECT().UpdateSociety(req).Return(nil)

		h := &Handler{
			sS: mockSocietyService,
		}

		w := httptest.NewRecorder()

		h.UpdateSociety(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should_return_internal_server_error_if_UpdateSociety_fails", func(t *testing.T) {
		mockSocietyService := NewMockSocietyService(ctrl)

		req := httptest.NewRequest(http.MethodPatch, "/society", nil)
		req.Header.Set("Content-Type", "application/json")

		ctx = context.WithValue(ctx, config.KeyLogger, mockLogger)
		req = req.WithContext(ctx)

		mockLogger.EXPECT().AddFuncName("UpdateSociety")
		expectedError := errors.New("database error")
		mockLogger.EXPECT().Error(fmt.Sprintf("failed to update society error: %v", expectedError))

		mockSocietyService.EXPECT().UpdateSociety(req).Return(expectedError)

		h := &Handler{
			sS: mockSocietyService,
		}

		w := httptest.NewRecorder()

		h.UpdateSociety(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
