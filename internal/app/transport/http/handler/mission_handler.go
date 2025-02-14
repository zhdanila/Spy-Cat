package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sca/internal/domain"
	"sca/pkg/response"
	"strconv"
)

func (h *Handler) createMission(w http.ResponseWriter, r *http.Request) {
	//unmarshal mission from request
	var mission domain.Mission

	err := json.NewDecoder(r.Body).Decode(&mission)
	if err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Mission.Create(mission)
	if err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("id - %d, error - %s", id, err.Error())))
}

func (h *Handler) deleteMission(w http.ResponseWriter, r *http.Request) {
	//unmarshal id from request
	missionId := r.PathValue("id")
	missionIntId, err := strconv.Atoi(missionId)
	if err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Mission.Delete(missionIntId)
	if err != nil {
		response.NewErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("status: ok"))
}

func (h *Handler) deleteTargetInMission(w http.ResponseWriter, r *http.Request) {
	//unmarshal mission id from request
	missionId := r.PathValue("mission_id")
	missionIntId, err := strconv.Atoi(missionId)
	if err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	//unmarshal target id from request
	targetId := r.PathValue("target_id")
	targetIntId, err := strconv.Atoi(targetId)
	if err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Mission.DeleteTarget(missionIntId, targetIntId)
	if err != nil {
		response.NewErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("status: ok"))
}

func (h *Handler) updateMission(w http.ResponseWriter, r *http.Request) {
	//unmarshal id from request
	missionId := r.PathValue("id")
	missionIntId, err := strconv.Atoi(missionId)
	if err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	//unmarshal mission from request
	var updatedMission domain.UpdatedMission
	err = json.NewDecoder(r.Body).Decode(&updatedMission)
	if err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.Mission.Update(missionIntId, updatedMission); err != nil {
		response.NewErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("status: updated"))
}

func (h *Handler) createTarget(w http.ResponseWriter, r *http.Request) {
	//unmarshal mission id from request
	missionId := r.PathValue("mission_id")
	missionIntId, err := strconv.Atoi(missionId)
	if err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	//unmarshal target from request
	var target domain.Target
	err = json.NewDecoder(r.Body).Decode(&target)
	if err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Mission.CreateTarget(missionIntId, target)
	if err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("id - %d", id)))
}

func (h *Handler) getAllMissions(w http.ResponseWriter, r *http.Request) {
	missions, err := h.services.Cat.GetAll()
	if err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	marshalledMissions, err := json.Marshal(&missions)
	if err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(marshalledMissions)
}

func (h *Handler) getByIdMission(w http.ResponseWriter, r *http.Request) {
	//unmarshal id from request
	missionId := r.PathValue("id")
	missionIntId, err := strconv.Atoi(missionId)
	if err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	mission, err := h.services.Mission.GetByID(missionIntId)
	if err != nil {
		response.NewErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	marshalledMission, err := json.Marshal(&mission)
	if err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(marshalledMission)

}
