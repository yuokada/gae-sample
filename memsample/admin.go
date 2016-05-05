package memsample

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
)

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Hello, this is AdminHandler")
	ctx := appengine.NewContext(r)
	u, err := user.CurrentOAuth(ctx, "")
	if err != nil {
		http.Error(w, "OAuth Authorization header required", http.StatusUnauthorized)
		return
	}
	if !u.Admin {
		err_s := fmt.Sprintf("Admin login only : %s", u.Email)
		//http.Error(w, "Admin login only", http.StatusUnauthorized)
		http.Error(w, err_s, http.StatusUnauthorized)
		return
	}
	fmt.Fprintf(w, `Welcome, admin user %s!`, u)

}
