package handler

import (
	"encoding/json"
	"net/http"
	"sca/internal/service/mission"
	"sca/pkg/response"
	"strconv"
)

// CreateMission @Summary Create a new mission
// @Description Creates a new mission with the specified details
// @Tags Missions
// @Accept json
// @Produce json
// @Param request body mission.CreateMissionRequest true "Mission details"
// @Success 201 {object} mission.CreateMissionResponse
// @Router /mission [post]
func (h *Handler) CreateMission(w http.ResponseWriter, r *http.Request) {
	var req mission.CreateMissionRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.Mission.CreateMission(&req)
	if err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

// GetMission @Summary Get a mission by ID
// @Description Retrieves a specific mission by its ID
// @Tags Missions
// @Produce json
// @Param id path int true "Mission ID"
// @Success 200 {object} mission.GetMissionResponse
// @Router /mission/{id} [get]
func (h *Handler) GetMission(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.Mission.GetMission(&mission.GetMissionRequest{ID: intId})
	if err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"mission": resp})
}

// ListMissions @Summary List all missions
// @Description Retrieves a list of all missions
// @Tags Missions
// @Produce json
// @Success 200 {object} mission.ListMissionsResponse
// @Router /mission [get]
func (h *Handler) ListMissions(w http.ResponseWriter, r *http.Request) {
	resp, err := h.services.Mission.ListMissions(&mission.ListMissionsRequest{})
	if err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"missions": resp})
}

// UpdateMissionCompletion @Summary Update mission completion status
// @Description Updates the completion status of a mission
// @Tags Missions
// @Accept json
// @Produce json
// @Param id path int true "Mission ID"
// @Param request body mission.UpdateMissionCompletionRequest true "Update completion status"
// @Success 200 {string} string "Mission completion updated successfully"
// @Router /mission/{id}/completion [put]
func (h *Handler) UpdateMissionCompletion(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var req mission.UpdateMissionCompletionRequest
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	req.ID = intId
	if _, err = h.services.Mission.UpdateMissionCompletion(&req); err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Mission completion updated successfully"))
}

// DeleteMission @Summary Delete a mission
// @Description Deletes a mission by ID
// @Tags Missions
// @Produce json
// @Param id path int true "Mission ID"
// @Success 200 {string} string "Mission deleted successfully"
// @Router /mission/{id} [delete]
func (h *Handler) DeleteMission(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if _, err = h.services.Mission.DeleteMission(&mission.DeleteMissionRequest{ID: intId}); err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Mission deleted successfully"))
}

// AssignSpyCatToMission @Summary Assign a Spy Cat to a mission
// @Description Assigns a Spy Cat to a specified mission
// @Tags Missions
// @Accept json
// @Produce json
// @Param id path int true "Mission ID"
// @Param request body mission.AssignSpyCatToMissionRequest true "Assignment details"
// @Success 200 {string} string "Spy Cat assigned to mission successfully"
// @Router /mission/{id}/assign [post]
func (h *Handler) AssignSpyCatToMission(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var req mission.AssignSpyCatToMissionRequest
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	req.MissionID = intId
	if _, err = h.services.Mission.AssignSpyCatToMission(&req); err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Spy Cat assigned to mission successfully"))
}
