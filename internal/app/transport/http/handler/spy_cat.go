package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sca/internal/service/spy_cat"
	"sca/pkg/response"
	"strconv"
)

func (h *Handler) CreateSpyCat(w http.ResponseWriter, r *http.Request) {
	var (
		err  error
		req  spy_cat.CreateSpyCatRequest
		resp *spy_cat.CreateSpyCatResponse
	)

	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if resp, err = h.services.SpyCat.CreateSpyCat(&req); err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Spy Cat with id %d was successfully created", resp.ID)))
}

func (h *Handler) GetSpyCat(w http.ResponseWriter, r *http.Request) {
	var (
		err   error
		intId int
		req   spy_cat.GetSpyCatRequest
		resp  *spy_cat.GetSpyCatResponse
	)

	id := r.PathValue("id")
	if intId, err = strconv.Atoi(id); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	req.ID = intId
	if resp, err = h.services.SpyCat.GetSpyCat(&req); err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseData := map[string]interface{}{
		"spy_cat": resp,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(responseData); err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, "Error encoding response")
	}
}

func (h *Handler) ListSpyCats(w http.ResponseWriter, r *http.Request) {
	var (
		err  error
		resp *spy_cat.ListSpyCatsResponse
	)

	if resp, err = h.services.SpyCat.ListSpyCats(&spy_cat.ListSpyCatsRequest{}); err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseData := map[string]interface{}{
		"spy_cats": resp,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(responseData); err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, "Error encoding response")
	}
}

func (h *Handler) UpdateSpyCatSalary(w http.ResponseWriter, r *http.Request) {
	var (
		err   error
		req   spy_cat.UpdateSpyCatSalaryRequest
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

	if _, err = h.services.SpyCat.UpdateSpyCatSalary(&req); err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Spy Cat salary updated successfully"))
}

func (h *Handler) DeleteSpyCat(w http.ResponseWriter, r *http.Request) {
	var (
		err   error
		intId int
		req   spy_cat.DeleteSpyCatRequest
	)

	id := r.PathValue("id")
	if intId, err = strconv.Atoi(id); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	req.ID = intId
	if _, err = h.services.SpyCat.DeleteSpyCat(&req); err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Spy Cat deleted successfully"))
}
