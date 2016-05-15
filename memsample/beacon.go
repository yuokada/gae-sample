package memsample

import (
	"encoding/base64"
	"io"
	"net/http"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/user"
	// _ "github.com/lestrrat/go-apache-logformat"
	"fmt"
)

type RequestInfo struct {
	UserAgent     string
	RemoteAddress string
	Cookie        string `datastore:",noindex"`
	Host          string
	Method        string `datastore:",noindex"`
	Proto         string `datastore:",noindex"`
	IsHttps       bool   `datastore:",noindex"`

	RequestDate time.Time
	// 	Key         *datastore.Key
	//Id            *datastore.Key
}

const base64GifPixel = "R0lGODlhAQABAID/AMDAwAAAACH5BAEAAAAALAAAAAABAAEAAAICRAEAOw=="

var output []byte

func BeaconHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	_, err := StoreUserInfo(ctx, r)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	output, _ = base64.StdEncoding.DecodeString(base64GifPixel)
	w.Header().Set("Content-Type", "image/gif")
	w.Header().Set("Cache-Control", "no-cache, no-store, private")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Del("Expires")
	io.WriteString(w, string(output))
}

func isSecure(r *http.Request) bool {
	t := r.TLS
	if t == nil {
		return false
	}
	return true
}

func StoreUserInfo(ctx context.Context, r *http.Request) (bool, error) {
	// get User Infomation
	_, err := user.CurrentOAuth(ctx, "")
	if err != nil {
		// http.Error(w, "OAuth Authorization header required", http.StatusUnauthorized)
		return false, err
	}

	rInfo := &RequestInfo{
		UserAgent:     r.UserAgent(),
		RemoteAddress: r.RemoteAddr,
		Proto:         r.Proto,
		Host:          r.Host,
		Method:        r.Method,
		IsHttps:       isSecure(r),
		RequestDate:   time.Now(),
	}

	key := datastore.NewIncompleteKey(ctx, "UserRequests", nil)
	if _, err := datastore.Put(ctx, key, rInfo); err != nil {
		return false, err
	}
	return true, nil
}

// https://godoc.org/google.golang.org/appengine/datastore
//type Counter struct {
//	Count int
//}
//
//func inc(ctx context.Context, key *datastore.Key) (int, error) {
//	var x Counter
//	if err := datastore.Get(ctx, key, &x); err != nil && err != datastore.ErrNoSuchEntity {
//		return 0, err
//	}
//	x.Count++
//	if _, err := datastore.Put(ctx, key, &x); err != nil {
//		return 0, err
//	}
//	return x.Count, nil
//}
//func handle(w http.ResponseWriter, r *http.Request) {
//	ctx := appengine.NewContext(r)
//	var count int
//	err := datastore.RunInTransaction(ctx, func(ctx context.Context) error {
//		var err1 error
//		count, err1 = inc(ctx, datastore.NewKey(ctx, "Counter", "singleton", 0, nil))
//		return err1
//	}, nil)
//	if err != nil {
//		serveError(ctx, w, err)
//		return
//	}
//	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
//	fmt.Fprintf(w, "Count=%d", count)
//}
