package memsample

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
)

func init() {
	r := mux.NewRouter()
	r.StrictSlash(false)

	r.HandleFunc("/hello", HeyHandler)
	r.HandleFunc("/count", CoutHandler)
	r.HandleFunc("/", HomeHandler)

	http.Handle("/", r)
}

func HeyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hey World")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// https://gist.github.com/hSATAC/5343225
	http.Redirect(w, r, "http://www.yahoo.co.jp", 301)
}

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
			fmt.Fprintln(w, err)
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
		fmt.Fprintln(w, err)
	} else {
		fmt.Fprintf(w, "%d <= ", current)
		fmt.Fprintln(w, "Count up!")
	}
	// http://localhost:8000/memcache?key=Counters
}
