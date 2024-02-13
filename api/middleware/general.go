package middleware

import (
	"net/http"
	"sessionmanagement/error"
	"sessionmanagement/utils"

	"github.com/alexedwards/scs/v2"
)

func VerifyidentityAndSession(flag int, sessionManager *scs.SessionManager) func(handler http.Handler) http.Handler {
	//flag == 0 => normal user, flag == 1 => admin
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// authenticated := sessionManager.Pop(r.Context(), "authenticated")
			// fmt.Println(authenticated)
			// sessionManager.Remove(r.Context(), "authenticated")
			// fmt.Println(sessionManager.Exists(r.Context(), "authenticated"))

			if sessionManager.GetBool(r.Context(), "authenticated") {
				if flag == 1 && !(sessionManager.GetBool(r.Context(), "isadmin")) {
					utils.ErrorGenerator(w, errorhandling.UnauthorizedError)
					return
				}
				handler.ServeHTTP(w, r)
			} else {
				utils.ErrorGenerator(w, errorhandling.SessionExpired)
				return
			}
		})
	}
}
