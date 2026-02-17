package mongo

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/custard-technology/abakcus/backend/internal/models"
)

// MenuRepositoryI defines the interface for menu repository operations.
type MenuRepositoryI interface {
	CreateMenu(ctx context.Context, menu *models.Menu) error
	GetMenuByID(ctx context.Context, menuID string) (*models.Menu, error)
	UpdateMenu(ctx context.Context, menuID string, updates *models.Menu) error
	DeleteMenu(ctx context.Context, menuID string) error
	ListMenusByBusiness(ctx context.Context, businessID string) ([]models.Menu, error)
}

type MenuRepository struct {
	client *mongo.Client
	dbName string
}

func NewMenuRepository(client *mongo.Client, dbName string) *MenuRepository {
	return &MenuRepository{client: client, dbName: dbName}
}

func (r *MenuRepository) CreateMenu(ctx context.Context, menu *models.Menu) error {
	if menu == nil {
		return errors.New("menu cannot be nil")
	}
	if menu.MenuID == "" {
		return errors.New("menu_id is required")
	}
	if menu.Name == "" {
		return errors.New("menu name is required")
	}
	if menu.BusinessID == "" {
		return errors.New("business_id is required")
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	coll := r.client.Database(r.dbName).Collection("menus")
	_, err := coll.InsertOne(ctx, menu)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return errors.New("menu with this ID already exists")
		}
		return err
	}

	return nil
}

func (r *MenuRepository) GetMenuByID(ctx context.Context, menuID string) (*models.Menu, error) {
	if menuID == "" {
		return nil, errors.New("menu_id is required")
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	coll := r.client.Database(r.dbName).Collection("menus")
	var menu models.Menu
	err := coll.FindOne(ctx, bson.M{"_id": menuID}).Decode(&menu)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("menu not found")
		}
		return nil, err
	}

	return &menu, nil
}

func (r *MenuRepository) UpdateMenu(ctx context.Context, menuID string, updates *models.Menu) error {
	if menuID == "" {
		return errors.New("menu_id is required")
	}
	if updates == nil {
		return errors.New("updates cannot be nil")
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	updateFields := bson.M{}
	if updates.Name != "" {
		updateFields["name"] = updates.Name
	}
	if updates.Description != "" {
		updateFields["description"] = updates.Description
	}
	updateFields["updated_at"] = time.Now()
	updateFields["is_active"] = updates.IsActive

	coll := r.client.Database(r.dbName).Collection("menus")
	result := coll.FindOneAndUpdate(ctx, bson.M{"_id": menuID}, bson.M{"$set": updateFields})
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return errors.New("menu not found")
		}
		return result.Err()
	}

	return nil
}

func (r *MenuRepository) DeleteMenu(ctx context.Context, menuID string) error {
	if menuID == "" {
		return errors.New("menu_id is required")
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	coll := r.client.Database(r.dbName).Collection("menus")
	result, err := coll.DeleteOne(ctx, bson.M{"_id": menuID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("menu not found")
	}

	return nil
}

func (r *MenuRepository) ListMenusByBusiness(ctx context.Context, businessID string) ([]models.Menu, error) {
	if businessID == "" {
		return nil, errors.New("business_id is required")
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	coll := r.client.Database(r.dbName).Collection("menus")
	cursor, err := coll.Find(ctx, bson.M{"business_id": businessID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var menus []models.Menu
	if err := cursor.All(ctx, &menus); err != nil {
		return nil, err
	}

	return menus, nil
}
