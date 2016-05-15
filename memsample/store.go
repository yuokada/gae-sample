package memsample

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/user"
)

const (
	ckey = "JRQ43DhDBqI3ZphoqVMcFNY1F"
	csec = "i5X6W7CGzDzrk0qxNPew6JUyQwXA8sRm3ZukJ9QYysxOXFjrrq"
	atok = "203129378-ACePDK8V0yFsvm4FJt9QbpRiNyhRBcq9Zfan7CvV"
	asec = "ulD1IjHJD8EFEnbqwsJGR0bi33i1PyZdlC83J7KdgVGuP"
)

type Employee struct {
	FirstName          string
	LastName           string
	HireDate           time.Time
	AttendedHRTraining bool
}

func DStoreHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	u, err := user.CurrentOAuth(ctx, "")
	if err != nil {
		http.Error(w, "OAuth Authorization header required", http.StatusUnauthorized)
		return
	}
	//if !u.Admin {
	//	err_s := fmt.Sprintf("Admin login only : %s", u.Email)
	//	//http.Error(w, "Admin login only", http.StatusUnauthorized)
	//	http.Error(w, err_s, http.StatusUnauthorized)
	//	return
	//}
	f(ctx, w)
	fmt.Fprintf(w, `Welcome, admin user %s!`, u)
}

func f(ctx context.Context, w http.ResponseWriter) {
	employee := &Employee{
		FirstName: "Antonio",
		LastName:  "Salieri",
		HireDate:  time.Now(),
	}
	employee.AttendedHRTraining = true

	key := datastore.NewIncompleteKey(ctx, "Employee", nil)
	if _, err := datastore.Put(ctx, key, employee); err != nil {
		// Handle err
		fmt.Fprintln(w, err)
	}
	fmt.Fprintln(w, "Data Putt succeed!")
	return
}
