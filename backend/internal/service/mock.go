package service

import (
	"context"

	"github.com/custard-technology/abakcus/backend/internal/models"
	"github.com/custard-technology/abakcus/backend/internal/repository/mongo"
)

// Verify that MockMenuRepository implements MenuRepositoryI
var _ mongo.MenuRepositoryI = (*MockMenuRepository)(nil)

type MockMenuRepository struct {
	menus map[string]*models.Menu
}

func NewMockMenuRepository() *MockMenuRepository {
	return &MockMenuRepository{
		menus: make(map[string]*models.Menu),
	}
}

// SetMenu adds a menu to the mock repository for testing
func (m *MockMenuRepository) SetMenu(menuID string, menu *models.Menu) {
	m.menus[menuID] = menu
}

func (m *MockMenuRepository) CreateMenu(ctx context.Context, menu *models.Menu) error {
	if menu != nil {
		m.menus[menu.MenuID] = menu
	}
	return nil
}

func (m *MockMenuRepository) GetMenuByID(ctx context.Context, menuID string) (*models.Menu, error) {
	if menu, ok := m.menus[menuID]; ok {
		return menu, nil
	}
	return nil, nil
}

func (m *MockMenuRepository) UpdateMenu(ctx context.Context, menuID string, updates *models.Menu) error {
	if updates != nil && menuID != "" {
		m.menus[menuID] = updates
	}
	return nil
}

func (m *MockMenuRepository) DeleteMenu(ctx context.Context, menuID string) error {
	delete(m.menus, menuID)
	return nil
}

func (m *MockMenuRepository) ListMenusByBusiness(ctx context.Context, businessID string) ([]models.Menu, error) {
	var result []models.Menu
	for _, menu := range m.menus {
		if menu.BusinessID == businessID {
			result = append(result, *menu)
		}
	}
	return result, nil
}
