package handler

import (
	"encoding/json"
	"net/http"
	"sca/internal/service/target"
	"sca/pkg/response"
	"strconv"
)

// AddTargetsToMission adds multiple targets to a mission.
// @Summary Add targets to a mission
// @Description Adds multiple targets to a specific mission.
// @Tags Targets
// @Accept json
// @Produce json
// @Param id path int true "Mission ID"
// @Param request body target.AddTargetsToMissionRequest true "Targets Data"
// @Success 200 {string} string "Targets added to mission successfully"
// @Router /mission/{id}/targets [post]
func (h *Handler) AddTargetsToMission(w http.ResponseWriter, r *http.Request) {
	var (
		err   error
		req   target.AddTargetsToMissionRequest
		intId int
	)

	id := r.PathValue("id")
	if intId, err = strconv.Atoi(id); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	req.MissionID = intId
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if _, err = h.services.Target.AddTargetsToMission(&req); err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Targets added to mission successfully"))
}

// UpdateTargetCompletion updates the completion status of a target.
// @Summary Update target completion
// @Description Updates the completion status of a target.
// @Tags Targets
// @Accept json
// @Produce json
// @Param id path int true "Target ID"
// @Param request body target.UpdateTargetCompletionRequest true "Completion Data"
// @Success 200 {string} string "Target completion updated successfully"
// @Router /target/{id}/completion [put]
func (h *Handler) UpdateTargetCompletion(w http.ResponseWriter, r *http.Request) {
	var (
		err   error
		req   target.UpdateTargetCompletionRequest
		intId int
	)

	id := r.PathValue("id")
	if intId, err = strconv.Atoi(id); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	req.ID = intId

	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if _, err = h.services.Target.UpdateTargetCompletion(&req); err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Target completion updated successfully"))
}

// UpdateTargetNotes updates notes for a specific target.
// @Summary Update target notes
// @Description Updates the notes of a target.
// @Tags Targets
// @Accept json
// @Produce json
// @Param id path int true "Target ID"
// @Param request body target.UpdateTargetNotesRequest true "Notes Data"
// @Success 200 {string} string "Target notes updated successfully"
// @Router /target/{id}/notes [put]
func (h *Handler) UpdateTargetNotes(w http.ResponseWriter, r *http.Request) {
	var (
		err   error
		req   target.UpdateTargetNotesRequest
		intId int
	)

	id := r.PathValue("id")
	if intId, err = strconv.Atoi(id); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	req.ID = intId

	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if _, err = h.services.Target.UpdateTargetNotes(&req); err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Target notes updated successfully"))
}

// DeleteTarget deletes a target.
// @Summary Delete a target
// @Description Deletes a specific target.
// @Tags Targets
// @Produce json
// @Param id path int true "Target ID"
// @Success 200 {string} string "Target deleted successfully"
// @Router /target/{id} [delete]
func (h *Handler) DeleteTarget(w http.ResponseWriter, r *http.Request) {
	var (
		err   error
		intId int
		req   target.DeleteTargetRequest
	)

	id := r.PathValue("id")
	if intId, err = strconv.Atoi(id); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	req.ID = intId
	if _, err = h.services.Target.DeleteTarget(&req); err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Target deleted successfully"))
}
