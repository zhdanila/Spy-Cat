package http

import (
	"go.uber.org/zap"
	"net/http"
	"sca/internal/middleware"
	"sca/internal/service"
)

type Handler struct {
	services *service.Service
	logger   *zap.Logger
}

func NewHandler(services *service.Service, logger *zap.Logger) *Handler {
	return &Handler{services: services, logger: logger}
}

func (h *Handler) InitRoutes() http.Handler {
	mux := http.NewServeMux()

	//cats endpoints
	mux.HandleFunc("POST /api/cats/create", h.createCat)
	mux.HandleFunc("DELETE /api/cats/delete/{id}", h.deleteCat)
	mux.HandleFunc("PUT /api/cats/update/{id}", h.updateCat)
	mux.HandleFunc("GET /api/cats/", h.getAllCats)
	mux.HandleFunc("GET /api/cats/{id}", h.getByIdCat)

	//mission endpoints
	mux.HandleFunc("POST /api/missions/create", h.createMission)
	mux.HandleFunc("POST /api/missions/{mission_id}/create", h.createTarget)
	mux.HandleFunc("DELETE /api/missions/delete/{id}", h.deleteMission)
	mux.HandleFunc("DELETE /api/missions/{mission_id}/delete/target/{target_id}", h.deleteTargetInMission)
	mux.HandleFunc("PUT /api/missions/update/{id}", h.updateMission)
	mux.HandleFunc("GET /api/missions", h.getAllMissions)
	mux.HandleFunc("GET /api/missions/{id}", h.getByIdMission)

	handler := middleware.Logging(h.logger, mux)

	return handler
}
