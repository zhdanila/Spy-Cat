package handler

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"sca/internal/app/transport/http/middleware"
	"sca/internal/service"
)

type Handler struct {
	services  *service.Service
	validator *validator.Validate
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services:  services,
		validator: validator.New(),
	}
}

func (h *Handler) InitRoutes() http.Handler {
	mux := http.NewServeMux()
	handler := middleware.Logging(zap.L(), mux)

	// Spy Cats
	mux.Handle("POST /spycat", http.HandlerFunc(h.CreateSpyCat))
	mux.Handle("GET /spycat/{id}", http.HandlerFunc(h.GetSpyCat))
	mux.Handle("GET /spycat", http.HandlerFunc(h.ListSpyCats))
	mux.Handle("PUT /spycat/{id}/salary", http.HandlerFunc(h.UpdateSpyCatSalary))
	mux.Handle("DELETE /spycat/{id}", http.HandlerFunc(h.DeleteSpyCat))

	// Missions
	mux.Handle("POST /mission", http.HandlerFunc(h.CreateMission))
	mux.Handle("GET /mission/{id}", http.HandlerFunc(h.GetMission))
	mux.Handle("GET /mission", http.HandlerFunc(h.ListMissions))
	mux.Handle("PUT /mission/{id}/completion", http.HandlerFunc(h.UpdateMissionCompletion))
	mux.Handle("DELETE /mission/{id}", http.HandlerFunc(h.DeleteMission))
	mux.Handle("POST /mission/{id}/assign", http.HandlerFunc(h.AssignSpyCatToMission))

	// Targets
	mux.Handle("POST /mission/{id}/targets", http.HandlerFunc(h.AddTargetsToMission))
	mux.Handle("PUT /target/{id}/completion", http.HandlerFunc(h.UpdateTargetCompletion))
	mux.Handle("PUT /target/{id}/notes", http.HandlerFunc(h.UpdateTargetNotes))
	mux.Handle("DELETE /target/{id}", http.HandlerFunc(h.DeleteTarget))

	return handler
}
