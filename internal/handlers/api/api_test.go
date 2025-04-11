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
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	advertproto "github.com/s21platform/advert-proto/advert-proto"
	chatproto "github.com/s21platform/chat-proto/chat-proto"
	logger_lib "github.com/s21platform/logger-lib"
	societyproto "github.com/s21platform/society-proto/society-proto"
	userproto "github.com/s21platform/user-proto/user-proto"

	"github.com/s21platform/gateway-service/internal/config"
	"github.com/s21platform/gateway-service/internal/model"
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

		expectedResult := &model.SocietyInfo{
			SocietyUUID:      "test-uuid",
			Name:             "Test Society",
			Description:      "Test Description",
			OwnerUUID:        "owner-uuid",
			PhotoURL:         "https://example.com/photo.jpg",
			FormatID:         1,
			PostPermissionID: 2,
			IsSearch:         true,
			CountSubscribe:   0,
			Tags:             []int64{},
			CanEditSociety:   false,
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

		var responseBody model.SocietyInfo
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
		mockLogger.EXPECT().Error("failed to get society info error: database error")
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

func TestApi_CreateAdvert(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)

	ctx := context.Background()
	mockLogger := logger_lib.NewMockLoggerInterface(ctrl)

	t.Run("should_ok", func(t *testing.T) {
		mockAdvertService := NewMockAdvertService(ctrl)
		mockLogger.EXPECT().AddFuncName("CreateAdvert")

		ctx = context.WithValue(ctx, config.KeyLogger, mockLogger)
		r := &http.Request{}
		w := httptest.NewRecorder()
		r = r.WithContext(ctx)

		expected := &advertproto.AdvertEmpty{}

		mockAdvertService.EXPECT().CreateAdvert(r).Return(expected, nil)

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

		s.CreateAdvert(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should_err_us_fail_response", func(t *testing.T) {
		mockAdvertService := NewMockAdvertService(ctrl)
		mockLogger.EXPECT().AddFuncName("CreateAdvert")
		mockLogger.EXPECT().Error(gomock.Any())

		ctx = context.WithValue(ctx, config.KeyLogger, mockLogger)
		r := &http.Request{}
		w := httptest.NewRecorder()
		r = r.WithContext(ctx)

		mockErr := errors.New("some error")

		mockAdvertService.EXPECT().CreateAdvert(r).Return(nil, mockErr)

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

		s.CreateAdvert(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestApi_GetChats(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)

	ctx := context.Background()
	mockLogger := logger_lib.NewMockLoggerInterface(ctrl)

	t.Run("should_ok", func(t *testing.T) {
		mockChatService := NewMockChatService(ctrl)
		mockLogger.EXPECT().AddFuncName("GetChats")

		ctx = context.WithValue(ctx, config.KeyLogger, mockLogger)
		r := &http.Request{}
		w := httptest.NewRecorder()
		r = r.WithContext(ctx)

		expected := &chatproto.GetChatsOut{
			Chats: []*chatproto.Chat{
				{
					LastMessage:          "last message",
					ChatName:             "name",
					AvatarUrl:            "url",
					LastMessageTimestamp: "23.01.25",
					ChatUuid:             "test-uuid",
				},
				{
					LastMessage:          "second last message",
					ChatName:             "other name",
					AvatarUrl:            "url_2",
					LastMessageTimestamp: "03.05.25",
					ChatUuid:             "test-uuid-2",
				},
			},
		}

		mockChatService.EXPECT().GetChats(r).Return(expected, nil)

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

		s.GetChats(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should_err_us_fail_response", func(t *testing.T) {
		mockChatService := NewMockChatService(ctrl)
		mockLogger.EXPECT().AddFuncName("GetChats")
		mockLogger.EXPECT().Error(gomock.Any())

		ctx = context.WithValue(ctx, config.KeyLogger, mockLogger)
		r := &http.Request{}
		w := httptest.NewRecorder()
		r = r.WithContext(ctx)

		mockErr := errors.New("some error")

		mockChatService.EXPECT().GetChats(r).Return(nil, mockErr)

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

		s.GetChats(w, r)

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

func TestApi_CancelAdvert(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)

	ctx := context.Background()
	mockLogger := logger_lib.NewMockLoggerInterface(ctrl)

	t.Run("should_cancel_advert_successfully", func(t *testing.T) {
		mockAdvertService := NewMockAdvertService(ctrl)
		mockLogger.EXPECT().AddFuncName("CancelAdvert")

		ctx = context.WithValue(ctx, config.KeyLogger, mockLogger)
		r := &http.Request{}
		w := httptest.NewRecorder()
		r = r.WithContext(ctx)

		expected := &advertproto.AdvertEmpty{}

		mockAdvertService.EXPECT().CancelAdvert(r).Return(expected, nil)

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

		s.CancelAdvert(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should_return_internal_server_error_if_CancelAdvert_fails", func(t *testing.T) {
		mockAdvertService := NewMockAdvertService(ctrl)
		mockLogger.EXPECT().AddFuncName("CancelAdvert")
		mockLogger.EXPECT().Error(gomock.Any())

		ctx = context.WithValue(ctx, config.KeyLogger, mockLogger)
		r := &http.Request{}
		w := httptest.NewRecorder()
		r = r.WithContext(ctx)

		mockErr := errors.New("some error")

		mockAdvertService.EXPECT().CancelAdvert(r).Return(nil, mockErr)

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

		s.CancelAdvert(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestApi_RestoreAdvert(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)

	ctx := context.Background()
	mockLogger := logger_lib.NewMockLoggerInterface(ctrl)

	t.Run("should_restore_advert_successfully", func(t *testing.T) {
		mockAdvertService := NewMockAdvertService(ctrl)
		mockLogger.EXPECT().AddFuncName("RestoreAdvert")

		ctx = context.WithValue(ctx, config.KeyLogger, mockLogger)
		r := &http.Request{}
		w := httptest.NewRecorder()
		r = r.WithContext(ctx)

		expected := &advertproto.AdvertEmpty{}

		mockAdvertService.EXPECT().RestoreAdvert(r).Return(expected, nil)

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

		s.RestoreAdvert(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should_return_internal_server_error_if_RestoreAdvert_fails", func(t *testing.T) {
		mockAdvertService := NewMockAdvertService(ctrl)
		mockLogger.EXPECT().AddFuncName("RestoreAdvert")
		mockLogger.EXPECT().Error(gomock.Any())

		ctx = context.WithValue(ctx, config.KeyLogger, mockLogger)
		r := &http.Request{}
		w := httptest.NewRecorder()
		r = r.WithContext(ctx)

		mockErr := errors.New("some error")

		mockAdvertService.EXPECT().RestoreAdvert(r).Return(nil, mockErr)

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

		s.RestoreAdvert(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestHandler_GetOptionRequests(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	mockLogger := logger_lib.NewMockLoggerInterface(ctrl)

	t.Run("should_get_option_requests_successfully", func(t *testing.T) {
		mockOptionService := NewMockOptionService(ctrl)

		req := httptest.NewRequest(http.MethodGet, "/api/option_requests", nil)
		req.Header.Set("Content-Type", "application/json")

		ctx = context.WithValue(ctx, config.KeyLogger, mockLogger)
		req = req.WithContext(ctx)

		expectedResult := model.OptionRequestsList{
			Items: []model.OptionRequest{
				{
					ID:             1,
					AttributeID:    2,
					AttributeValue: "city",
					Value:          "Москва",
					UserUuid:       uuid.New().String(),
					CreatedAt:      time.Now(),
				},
			},
		}

		mockLogger.EXPECT().AddFuncName("GetOptionRequests")
		mockOptionService.EXPECT().GetOptionRequests(req).Return(expectedResult, nil)

		w := httptest.NewRecorder()

		h := New(
			nil,
			nil,
			nil,
			nil,
			mockOptionService,
			nil,
			nil,
			nil,
			nil,
		)

		h.GetOptionRequests(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

		var responseBody model.OptionRequestsList
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		assert.NoError(t, err)
		assert.Equal(t, len(expectedResult.Items), len(responseBody.Items))
		assert.Equal(t, expectedResult.Items[0].AttributeValue, responseBody.Items[0].AttributeValue)
		assert.Equal(t, expectedResult.Items[0].Value, responseBody.Items[0].Value)
		assert.Equal(t, expectedResult.Items[0].UserUuid, responseBody.Items[0].UserUuid)
	})

	t.Run("should_return_internal_server_error_if_GetOptionRequests_fails", func(t *testing.T) {
		mockOptionService := NewMockOptionService(ctrl)

		req := httptest.NewRequest(http.MethodGet, "/api/option_requests", nil)
		req.Header.Set("Content-Type", "application/json")

		ctx = context.WithValue(ctx, config.KeyLogger, mockLogger)
		req = req.WithContext(ctx)

		expectedError := errors.New("database error")
		mockLogger.EXPECT().AddFuncName("GetOptionRequests")
		mockLogger.EXPECT().Error("failed to get option requests: database error")
		mockOptionService.EXPECT().GetOptionRequests(req).Return(model.OptionRequestsList{}, expectedError)

		w := httptest.NewRecorder()

		h := New(
			nil,
			nil,
			nil,
			nil,
			mockOptionService,
			nil,
			nil,
			nil,
			nil,
		)

		h.GetOptionRequests(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestHandler_MarkNotificationAsRead(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	mockLogger := logger_lib.NewMockLoggerInterface(ctrl)

	tests := []struct {
		name           string
		requestBody    interface{}
		setupMock      func(mock *MockNotificationService)
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "successful_mark_as_read",
			requestBody: model.MarkNotificationsAsReadRequest{
				Data: struct {
					IDs []int64 `json:"ids"`
				}{
					IDs: []int64{1, 2, 3},
				},
			},
			setupMock: func(mock *MockNotificationService) {
				mockLogger.EXPECT().AddFuncName("MarkNotificationAsRead")
				mock.EXPECT().
					MarkNotificationsAsRead(gomock.Any()).
					Return(&emptypb.Empty{}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{}` + "\n",
		},
		{
			name:        "invalid_request_body",
			requestBody: "invalid json",
			setupMock: func(mock *MockNotificationService) {
				mockLogger.EXPECT().AddFuncName("MarkNotificationAsRead")
				mockLogger.EXPECT().Error(gomock.Any())
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "",
		},
		{
			name: "empty_ids",
			requestBody: model.MarkNotificationsAsReadRequest{
				Data: struct {
					IDs []int64 `json:"ids"`
				}{
					IDs: []int64{},
				},
			},
			setupMock: func(mock *MockNotificationService) {
				mockLogger.EXPECT().AddFuncName("MarkNotificationAsRead")
				mockLogger.EXPECT().Error("notification IDs are required")
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "",
		},
		{
			name: "service_error",
			requestBody: model.MarkNotificationsAsReadRequest{
				Data: struct {
					IDs []int64 `json:"ids"`
				}{
					IDs: []int64{1},
				},
			},
			setupMock: func(mock *MockNotificationService) {
				mockLogger.EXPECT().AddFuncName("MarkNotificationAsRead")
				mockLogger.EXPECT().Error("failed to mark notification as read: service error")
				mock.EXPECT().
					MarkNotificationsAsRead(gomock.Any()).
					Return(nil, errors.New("service error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockNotificationService := NewMockNotificationService(ctrl)
			tt.setupMock(mockNotificationService)

			ctx = context.WithValue(ctx, config.KeyLogger, mockLogger)

			var body []byte
			var err error
			if str, ok := tt.requestBody.(string); ok {
				body = []byte(str)
			} else {
				body, err = json.Marshal(tt.requestBody)
				require.NoError(t, err)
			}

			r := httptest.NewRequest(http.MethodPatch, "/api/notification", bytes.NewReader(body))
			r = r.WithContext(ctx)
			w := httptest.NewRecorder()

			handler := New(nil, nil, mockNotificationService, nil, nil, nil, nil, nil, nil)
			handler.MarkNotificationAsRead(w, r)

			assert.Equal(t, tt.expectedStatus, w.Code)
			if tt.expectedBody != "" {
				assert.JSONEq(t, tt.expectedBody, w.Body.String())
			}
		})
	}
}
