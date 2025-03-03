package api

import (
	"net/http"

	"github.com/MACSEC-Proje-Gelistirme/Community-Portal-Api/internal/models"
	"github.com/MACSEC-Proje-Gelistirme/Community-Portal-Api/internal/repository"
	"github.com/MACSEC-Proje-Gelistirme/Community-Portal-Api/pkg/utils"
)

func (ro *Router) CreateUser(w http.ResponseWriter, r *http.Request) {
	var payload models.CreateUserPayload
	if err := utils.DecodeRequestBody(r, &payload); err != nil {
		utils.JSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	user := models.User{
		UserID:               payload.UserID,
		FirstName:            payload.FirstName,
		LastName:             payload.LastName,
		Email:                payload.Email,
		TelephoneNumber:      payload.TelephoneNumber,
		EmailPreferences:     true,
		MarketingPreferences: true,
		CreatedAt:            utils.GetCurrentTime(),
		UpdatedAt:            utils.GetCurrentTime(),
	}

	userRepository := repository.NewUserRepository(ro.db)
	result, err := userRepository.CreateUser(user)
	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusCreated, result)
}
func (ro *Router) GetUser(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	if userID == "" {
		utils.JSONError(w, http.StatusBadRequest, "User ID is required")
		return
	}

	userRepository := repository.NewUserRepository(ro.db)
	user, err := userRepository.GetUserByID(userID)
	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, user)
}
