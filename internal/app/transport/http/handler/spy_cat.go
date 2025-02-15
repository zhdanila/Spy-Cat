package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sca/internal/service/spy_cat"
	"sca/pkg/response"
	"strconv"
)

// CreateSpyCat @Summary Create a new SpyCat
// @Description Creates a new SpyCat with the provided details
// @Tags SpyCats
// @ID create-spycat
// @Accept  json
// @Produce  json
// @Param request body spy_cat.CreateSpyCatRequest true "SpyCat Creation Request"
// @Success 201 {string} string "Spy Cat with id {id} was successfully created"
// @Router /spycat [post]
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

// GetSpyCat @Summary Get SpyCat by ID
// @Description Retrieves a SpyCat by its ID
// @Tags SpyCats
// @ID get-spycat
// @Produce  json
// @Param id path int true "SpyCat ID"
// @Success 200 {object} spy_cat.GetSpyCatResponse
// @Router /spycat/{id} [get]
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

// ListSpyCats @Summary List all SpyCats
// @Description Retrieves a list of all SpyCats
// @Tags SpyCats
// @ID list-spycats
// @Produce  json
// @Success 200 {object} spy_cat.ListSpyCatsResponse
// @Router /spycat [get]
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

// UpdateSpyCatSalary @Summary Update SpyCat Salary
// @Description Updates the salary of a SpyCat
// @Tags SpyCats
// @ID update-spycat-salary
// @Accept  json
// @Produce  json
// @Param id path int true "SpyCat ID"
// @Param request body spy_cat.UpdateSpyCatSalaryRequest true "SpyCat Salary Update Request"
// @Success 200 {string} string "Spy Cat salary updated successfully"
// @Router /spycat/{id}/salary [put]
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

// DeleteSpyCat @Summary Delete a SpyCat
// @Description Deletes a SpyCat by its ID
// @Tags SpyCats
// @ID delete-spycat
// @Produce  json
// @Param id path int true "SpyCat ID"
// @Success 200 {string} string "Spy Cat deleted successfully"
// @Router /spycat/{id} [delete]
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
