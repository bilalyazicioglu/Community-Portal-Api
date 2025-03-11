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

func GetFirstNameFromClaims(claims map[string]any) (string, bool) {
	firstNameInterface, ok := claims["first_name"]
	if !ok || firstNameInterface == nil {
		return "", false
	}
	FirstName, ok := firstNameInterface.(string)
	return FirstName, ok
}

func GetLastNameFromClaims(claims map[string]any) (string, bool) {
	lastNameInterface, ok := claims["last_name"]
	if !ok || lastNameInterface == nil {
		return "", false
	}
	LastName, ok := lastNameInterface.(string)
	return LastName, ok
}

func GetMailFromClaims(claims map[string]any) (string, bool) {
	mailInterface, ok := claims["email"]
	if !ok || mailInterface == nil {
		return "", false
	}
	Email, ok := mailInterface.(string)
	return Email, ok
}
