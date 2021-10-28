package server

import (
	"encoding/json"
	"git.cyradar.com/phinc/my-awesome-project/internal/model"
	"log"
	"net/http"

	"git.cyradar.com/phinc/my-awesome-project/internal/services"
	"git.cyradar.com/phinc/my-awesome-project/pkg/hxxp"
	"github.com/go-chi/chi/v5"
)

func addStaff(w http.ResponseWriter, r *http.Request) {
	staff := new(model.Staff)

	err := json.NewDecoder(r.Body).Decode(&staff)
	if err != nil {
		log.Println("Add Staff", "error", err.Error())
		hxxp.ResponseJson(w, http.StatusInternalServerError, nil)
		return
	}

	err = services.AddStaff(r.Context(), staff)
	if err != nil {
		log.Println("Add Staff", "error", err.Error())
		hxxp.ResponseJson(w, http.StatusInternalServerError, nil)
		return
	}

	hxxp.ResponseJson(w, http.StatusOK, nil)
}

func viewStaff(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	staff, err := services.ViewStaff(r.Context(), id)
	if err != nil {
		log.Println("View Staff", "error", err.Error())
		hxxp.ResponseJson(w, http.StatusInternalServerError, nil)
		return
	}

	hxxp.ResponseJson(w, http.StatusOK, staff)
}

func updateStaff(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	staff := new(model.Staff)

	err := json.NewDecoder(r.Body).Decode(&staff)
	if err != nil {
		log.Println("Add Staff", "error", err.Error())
		hxxp.ResponseJson(w, http.StatusInternalServerError, nil)
		return
	}

	err = services.UpdateStaff(r.Context(), staff, id)
	if err != nil {
		log.Println("Update Staff", "error", err.Error())
		hxxp.ResponseJson(w, http.StatusInternalServerError, nil)
		return
	}

	hxxp.ResponseJson(w, http.StatusOK, nil)
}

func removeStaff(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := services.RemoveStaff(r.Context(), id)
	if err != nil {
		log.Println("Remove Staff", "error", err.Error())
		hxxp.ResponseJson(w, http.StatusInternalServerError, nil)
		return
	}

	hxxp.ResponseJson(w, http.StatusOK, nil)
}
