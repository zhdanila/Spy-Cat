package handler

import (
	"encoding/json"
	"net/http"
	"sca/internal/service/target"
	"sca/pkg/response"
	"strconv"
)

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
