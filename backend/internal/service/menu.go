package service

import (
	"context"
	"errors"
	"time"

	"github.com/custard-technology/abakcus/backend/internal/models"
	"github.com/custard-technology/abakcus/backend/internal/repository/mongo"
	"github.com/google/uuid"
)

type MenuService struct {
	repo mongo.MenuRepositoryI
}

func NewMenuService(repo mongo.MenuRepositoryI) *MenuService {
	return &MenuService{repo: repo}
}

func (s *MenuService) CreateMenu(ctx context.Context, req *models.CreateMenuRequest, businessID string) (*models.Menu, error) {
	if req == nil {
		return nil, errors.New("request cannot be nil")
	}
	if req.Name == "" {
		return nil, errors.New("menu name is required")
	}
	if businessID == "" {
		return nil, errors.New("business_id is required")
	}

	menu := &models.Menu{
		MenuID:      uuid.New().String(),
		Name:        req.Name,
		Description: req.Description,
		BusinessID:  businessID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		IsActive:    true,
	}

	err := s.repo.CreateMenu(ctx, menu)
	if err != nil {
		return nil, err
	}

	return menu, nil
}

func (s *MenuService) GetMenu(ctx context.Context, menuID string) (*models.Menu, error) {
	if menuID == "" {
		return nil, errors.New("menu_id is required")
	}

	menu, err := s.repo.GetMenuByID(ctx, menuID)
	if err != nil {
		return nil, err
	}

	return menu, nil
}

func (s *MenuService) UpdateMenu(ctx context.Context, menuID string, req *models.UpdateMenuRequest) (*models.Menu, error) {
	if menuID == "" {
		return nil, errors.New("menu_id is required")
	}
	if req == nil {
		return nil, errors.New("request cannot be nil")
	}

	existing, err := s.repo.GetMenuByID(ctx, menuID)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		existing.Name = req.Name
	}
	if req.Description != "" {
		existing.Description = req.Description
	}
	if req.IsActive != nil {
		existing.IsActive = *req.IsActive
	}
	existing.UpdatedAt = time.Now()

	err = s.repo.UpdateMenu(ctx, menuID, existing)
	if err != nil {
		return nil, err
	}

	return existing, nil
}

func (s *MenuService) DeleteMenu(ctx context.Context, menuID string) error {
	if menuID == "" {
		return errors.New("menu_id is required")
	}

	return s.repo.DeleteMenu(ctx, menuID)
}

func (s *MenuService) ListMenusByBusiness(ctx context.Context, businessID string) ([]models.Menu, error) {
	if businessID == "" {
		return nil, errors.New("business_id is required")
	}

	menus, err := s.repo.ListMenusByBusiness(ctx, businessID)
	if err != nil {
		return nil, err
	}

	if menus == nil {
		menus = []models.Menu{}
	}

	return menus, nil
}
