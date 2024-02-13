package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"sessionmanagement/api/model/request"
	"sessionmanagement/api/model/response"
	"sessionmanagement/api/service"
	"sessionmanagement/api/validation"
	"sessionmanagement/error"
	"sessionmanagement/utils"

	"github.com/alexedwards/scs/v2"
)

type AuthController interface {
	UserLogin(w http.ResponseWriter, r *http.Request)
}

type authController struct {
	authService    service.AuthService
	sessionManager *scs.SessionManager
}

func NewAuthController(s service.AuthService, sessionManager *scs.SessionManager) AuthController {
	return authController{
		authService:    s,
		sessionManager: sessionManager,
	}
}

func (a authController) UserLogin(w http.ResponseWriter, r *http.Request) {
	var userLoginRequest request.User

	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.ErrorGenerator(w, errorhandling.ReadBodyError)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &userLoginRequest)
	if err != nil {
		utils.ErrorGenerator(w, errorhandling.ReadDataError)
		return
	}

	isEmail := validation.EmailValidation(userLoginRequest.Email)
	if !isEmail {
		utils.ErrorGenerator(w, errorhandling.EmailvalidationError)
		return
	}

	var user response.User
	user, err = a.authService.UserLogin(userLoginRequest)

	if err != nil {
		utils.ErrorGenerator(w, errorhandling.LoginFailedError)
		return
	}

	a.sessionManager.Put(r.Context(), "authenticated", true)
	a.sessionManager.Put(r.Context(), "userid", user.ID)
	a.sessionManager.Put(r.Context(), "isadmin", user.IsAdmin)

	utils.ResponseGenerator(w, http.StatusOK, user)
	return
}
