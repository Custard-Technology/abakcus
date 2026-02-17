package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/custard-technology/abakcus/backend/internal/models"
	"github.com/custard-technology/abakcus/backend/internal/service"
)

func TestCreateMenuHandler(t *testing.T) {
	mockRepo := service.NewMockMenuRepository()
	svc := service.NewMenuService(mockRepo)
	handler := NewMenuHandler(svc)

	body := models.CreateMenuRequest{
		Name:        "Test Menu",
		Description: "Test",
	}
	bodyBytes, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/menus", bytes.NewReader(bodyBytes))
	req.Header.Set("X-Business-ID", "biz-1")
	w := httptest.NewRecorder()

	handler.CreateMenu(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("expected 201, got %d", w.Code)
	}
}

func TestGetMenuHandler(t *testing.T) {
	mockRepo := service.NewMockMenuRepository()
	mockRepo.SetMenu("m1", &models.Menu{MenuID: "m1", Name: "Test", BusinessID: "b1"})

	svc := service.NewMenuService(mockRepo)
	handler := NewMenuHandler(svc)

	req := httptest.NewRequest(http.MethodGet, "/menus/m1", nil)
	w := httptest.NewRecorder()

	handler.GetMenu(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestDeleteMenuHandler(t *testing.T) {
	mockRepo := service.NewMockMenuRepository()
	mockRepo.SetMenu("m1", &models.Menu{MenuID: "m1"})

	svc := service.NewMenuService(mockRepo)
	handler := NewMenuHandler(svc)

	req := httptest.NewRequest(http.MethodDelete, "/menus/m1", nil)
	w := httptest.NewRecorder()

	handler.DeleteMenu(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("expected 204, got %d", w.Code)
	}
}

func TestListMenusHandler(t *testing.T) {
	mockRepo := service.NewMockMenuRepository()
	mockRepo.SetMenu("m1", &models.Menu{MenuID: "m1", BusinessID: "b1"})
	mockRepo.SetMenu("m2", &models.Menu{MenuID: "m2", BusinessID: "b1"})

	svc := service.NewMenuService(mockRepo)
	handler := NewMenuHandler(svc)

	req := httptest.NewRequest(http.MethodGet, "/menus", nil)
	req.Header.Set("X-Business-ID", "b1")
	w := httptest.NewRecorder()

	handler.ListMenus(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}
