package handler

import (
	"encoding/json"
	"net/http"
	"sca/internal/service/mission"
	"sca/pkg/response"
	"strconv"
)

func (h *Handler) CreateMission(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		req mission.CreateMissionRequest
	)

	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
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

func (h *Handler) GetMission(w http.ResponseWriter, r *http.Request) {
	var (
		err   error
		intId int
		req   mission.GetMissionRequest
		resp  *mission.GetMissionResponse
	)

	id := r.PathValue("id")
	if intId, err = strconv.Atoi(id); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	req.ID = intId
	if resp, err = h.services.Mission.GetMission(&req); err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseData := map[string]interface{}{
		"missions": resp,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseData)
}

func (h *Handler) ListMissions(w http.ResponseWriter, r *http.Request) {
	var (
		err  error
		resp *mission.ListMissionsResponse
	)

	if resp, err = h.services.Mission.ListMissions(&mission.ListMissionsRequest{}); err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseData := map[string]interface{}{
		"missions": resp,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseData)
}

func (h *Handler) UpdateMissionCompletion(w http.ResponseWriter, r *http.Request) {
	var (
		err   error
		req   mission.UpdateMissionCompletionRequest
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

	if _, err = h.services.Mission.UpdateMissionCompletion(&req); err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Mission completion updated successfully"))
}

func (h *Handler) DeleteMission(w http.ResponseWriter, r *http.Request) {
	var (
		err   error
		intId int
		req   mission.DeleteMissionRequest
	)

	id := r.PathValue("id")
	if intId, err = strconv.Atoi(id); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	req.ID = intId
	if _, err = h.services.Mission.DeleteMission(&req); err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Mission deleted successfully"))
}

func (h *Handler) AssignSpyCatToMission(w http.ResponseWriter, r *http.Request) {
	var (
		err   error
		req   mission.AssignSpyCatToMissionRequest
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

	if _, err = h.services.Mission.AssignSpyCatToMission(&req); err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Spy Cat assigned to mission successfully"))
}
