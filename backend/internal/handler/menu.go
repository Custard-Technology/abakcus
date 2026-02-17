package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/custard-technology/abakcus/backend/internal/models"
	"github.com/custard-technology/abakcus/backend/internal/service"
)

type MenuHandler struct {
	service *service.MenuService
}

func NewMenuHandler(svc *service.MenuService) *MenuHandler {
	return &MenuHandler{service: svc}
}

func getBusinessIDFromRequest(r *http.Request) string {
	return r.Header.Get("X-Business-ID")
}

func extractMenuIDFromPath(path string, prefix string) string {
	parts := strings.Split(strings.TrimPrefix(path, prefix), "/")
	if len(parts) > 0 && parts[0] != "" {
		return parts[0]
	}
	return ""
}

func respondJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Printf("error encoding response: %v", err)
		}
	}
}

func respondError(w http.ResponseWriter, statusCode int, message string) {
	respondJSON(w, statusCode, map[string]string{"error": message})
}

func (h *MenuHandler) CreateMenu(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	businessID := getBusinessIDFromRequest(r)
	if businessID == "" {
		respondError(w, http.StatusBadRequest, "X-Business-ID header is required")
		return
	}

	var req models.CreateMenuRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	menu, err := h.service.CreateMenu(r.Context(), &req, businessID)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if strings.Contains(err.Error(), "required") {
			statusCode = http.StatusBadRequest
		}
		respondError(w, statusCode, err.Error())
		return
	}

	respondJSON(w, http.StatusCreated, menu)
}

func (h *MenuHandler) GetMenu(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	menuID := extractMenuIDFromPath(r.URL.Path, "/menus/")
	if menuID == "" {
		respondError(w, http.StatusBadRequest, "menu_id is required")
		return
	}

	menu, err := h.service.GetMenu(r.Context(), menuID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			respondError(w, http.StatusNotFound, err.Error())
			return
		}
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, menu)
}

func (h *MenuHandler) UpdateMenu(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		respondError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	menuID := extractMenuIDFromPath(r.URL.Path, "/menus/")
	if menuID == "" {
		respondError(w, http.StatusBadRequest, "menu_id is required")
		return
	}

	var req models.UpdateMenuRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	menu, err := h.service.UpdateMenu(r.Context(), menuID, &req)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			respondError(w, http.StatusNotFound, err.Error())
			return
		}
		statusCode := http.StatusInternalServerError
		if strings.Contains(err.Error(), "required") {
			statusCode = http.StatusBadRequest
		}
		respondError(w, statusCode, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, menu)
}

func (h *MenuHandler) DeleteMenu(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		respondError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	menuID := extractMenuIDFromPath(r.URL.Path, "/menus/")
	if menuID == "" {
		respondError(w, http.StatusBadRequest, "menu_id is required")
		return
	}

	err := h.service.DeleteMenu(r.Context(), menuID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			respondError(w, http.StatusNotFound, err.Error())
			return
		}
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *MenuHandler) ListMenus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	businessID := getBusinessIDFromRequest(r)
	if businessID == "" {
		respondError(w, http.StatusBadRequest, "X-Business-ID header is required")
		return
	}

	menus, err := h.service.ListMenusByBusiness(r.Context(), businessID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, menus)
}
