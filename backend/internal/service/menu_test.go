package service

import (
	"context"
	"testing"

	"github.com/custard-technology/abakcus/backend/internal/models"
)

func TestCreateMenu(t *testing.T) {
	mockRepo := NewMockMenuRepository()
	svc := NewMenuService(mockRepo)

	req := &models.CreateMenuRequest{
		Name:        "Lunch",
		Description: "Daily lunch menu",
	}

	menu, err := svc.CreateMenu(context.Background(), req, "biz-123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if menu.Name != "Lunch" {
		t.Errorf("expected Lunch, got %s", menu.Name)
	}
	if menu.BusinessID != "biz-123" {
		t.Errorf("expected biz-123, got %s", menu.BusinessID)
	}
}

func TestGetMenu(t *testing.T) {
	mockRepo := NewMockMenuRepository()
	svc := NewMenuService(mockRepo)

	menu := &models.Menu{MenuID: "m1", Name: "Test", BusinessID: "b1"}
	mockRepo.SetMenu("m1", menu)

	retrieved, err := svc.GetMenu(context.Background(), "m1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if retrieved.MenuID != "m1" {
		t.Errorf("expected m1, got %s", retrieved.MenuID)
	}
}

func TestDeleteMenu(t *testing.T) {
	mockRepo := NewMockMenuRepository()
	svc := NewMenuService(mockRepo)

	mockRepo.SetMenu("m1", &models.Menu{MenuID: "m1"})

	err := svc.DeleteMenu(context.Background(), "m1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if _, ok := mockRepo.menus["m1"]; ok {
		t.Error("menu should have been deleted")
	}
}
