package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sca/internal/models"
	util "sca/internal/util"
	"strconv"
)

func(h *Handler) createCat(w http.ResponseWriter, r *http.Request) {
	//unmarshal cat from request
	var cat models.Cat

	err := json.NewDecoder(r.Body).Decode(&cat)
	if err != nil {
		util.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Cat.Create(cat)
	if err != nil {
		util.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("id - %d", id)))
}

func(h *Handler) deleteCat(w http.ResponseWriter, r *http.Request) {
	//unmarshal id from request
	catId := r.PathValue("id")
	catIntId, err := strconv.Atoi(catId)
	if err != nil {
		util.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Cat.Delete(catIntId)
	if err != nil {
		util.NewErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("status: ok"))
}

func(h *Handler) updateCat(w http.ResponseWriter, r *http.Request) {
	//unmarshal id from request
	catId := r.PathValue("id")
	catIntId, err := strconv.Atoi(catId)
	if err != nil {
		util.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	//unmarshal cat from request
	var updatedCat models.UpdatedCat
	err = json.NewDecoder(r.Body).Decode(&updatedCat)
	if err != nil {
		util.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.Cat.Update(catIntId, updatedCat); err != nil {
		util.NewErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("status: updated"))
}

func(h *Handler) getAllCats(w http.ResponseWriter, r *http.Request) {
	cats, err := h.services.Cat.GetAll()
	if err != nil {
		util.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	marshalledCats, err := json.Marshal(&cats)
	if err != nil {
		util.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(marshalledCats)
}

func(h *Handler) getByIdCat(w http.ResponseWriter, r *http.Request) {
	//unmarshal id from request
	catId := r.PathValue("id")
	catIntId, err := strconv.Atoi(catId)
	if err != nil {
		util.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	cat, err := h.services.Cat.GetById(catIntId)
	if err != nil {
		util.NewErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	marshalledCat, err := json.Marshal(&cat)
	if err != nil {
		util.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(marshalledCat)
}