package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sessionmanagement/api/route"
	"sessionmanagement/config"
	"sessionmanagement/constants"
	"sessionmanagement/db"
	"time"

	"github.com/alexedwards/scs/v2"
)

var sessionManager *scs.SessionManager

func main() {
	sessionManager = scs.New()
	sessionManager.Lifetime = 30 * time.Minute
	//Lifetime refers to the maximum duration that a session can exist regardless of activity.

	sessionManager.IdleTimeout = 10 * time.Minute
	//Idle timeout refers to the duration of inactivity within a session before the session is considered expired.

	//Put() to add data to session.
	//Get() to retrieve those data from session like GetInt64, GetString, GetBool, GetFloat.
	//Pop() to retrieve data and remove key from session. (for reference: api/middleware/general.go)
	//Keys() to returns a sorted slice of keys in the session data.
	//Remove() to remove key from session data. (same as Pop() but pop first return value corresponding to that key and then remove key from session data.)
	//Destroy() to remove all session data.

	config.LoadEnv()
	conn, err := db.DBConnection()

	if err != nil {
		return
	}
	defer conn.Close(context.Background())

	r := route.UsersRoutes(conn, sessionManager)

	fmt.Println("Server started on port no. " + constants.PORT_NO)
	log.Fatal(http.ListenAndServe(constants.PORT_NO, r))
}
