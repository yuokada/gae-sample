package memsample

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
)

func init() {
	r := mux.NewRouter()
	r.StrictSlash(false)

	// GAE Practice Entrypoint and Handlers
	r.HandleFunc("/hello", HeyHandler)
	r.HandleFunc("/count", CoutHandler)
	r.HandleFunc("/", HomeHandler)

	// Beacon
	track := r.PathPrefix("/t").Subrouter()
	track.HandleFunc("/b", BeaconHandler)

	// Admin
	sub := r.PathPrefix("/admin").Subrouter()
	sub.HandleFunc("/", AdminHandler)
	sub.HandleFunc("/register", RegisterSeacretsHandler)
	r.HandleFunc("/register", RegisterSeacretsHandler)

	// Bot
	bot := r.PathPrefix("/bot").Subrouter()
	bot.HandleFunc("/tweet", TweetHandler)
	bot.HandleFunc("/store", DStoreHandler) // debug now

	http.Handle("/", r)
}

// Hello World Handler: sample code
func HeyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hey World")
}

// Redirect to another page: sample code
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// https://gist.github.com/hSATAC/5343225
	http.Redirect(w, r, "http://www.yahoo.co.jp", 301)
}

// Access Counter : memcache sample code
func CoutHandler(w http.ResponseWriter, r *http.Request) {
	// https://cloud.google.com/appengine/docs/go/memcache/using
	k := "Counters"
	var current int
	c := appengine.NewContext(r)

	item0, err := memcache.Get(c, k)
	if err != nil && err != memcache.ErrCacheMiss {
		return
	}
	if err == nil {
		current, err = strconv.Atoi(string(item0.Value))
		if err != nil {
			//fmt.Fprintf(w, "Here? [% s]\n", string(item0.Value[:]))
			errors.Fprint(w, err)
			return
		}
	} else {
		fmt.Fprintf(w, "memcache miss\n")
	}

	current = current + 1
	item := &memcache.Item{
		Key:   k,
		Value: []byte(fmt.Sprintf("%d", current)),
	}
	// fmt.Fprintf(w, "memcache will set: Key=%q Val=[%#v], Var=[%q]\n",
	// item.Key, item.Value, fmt.Sprintf("%d", current))
	if err := memcache.Set(c, item); err == memcache.ErrCacheMiss {
		fmt.Fprintf(w, "%s\n", memcache.ErrCacheMiss)
	} else if err != nil {
		errors.Fprint(w, err)
	} else {
		fmt.Fprintf(w, "%d <= ", current)
		fmt.Fprintln(w, "Count up!")
	}
}
