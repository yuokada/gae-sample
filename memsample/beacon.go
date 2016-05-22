package memsample

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	//"google.golang.org/appengine/user"
	"github.com/pkg/errors"
)

type RequestInfo struct {
	UserAgent     string
	Referrer      string
	RemoteAddress string
	Cookie        string
	Host          string
	Method        string `datastore:",noindex"`
	Proto         string `datastore:",noindex"`
	IsHttps       bool   `datastore:",noindex"`

	RequestDate time.Time
}

const base64GifPixel = "R0lGODlhAQABAID/AMDAwAAAACH5BAEAAAAALAAAAAABAAEAAAICRAEAOw=="

var output []byte

func BeaconHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	_, err := storeUserInfo(ctx, w, r)
	if err != nil {
		errors.Fprint(w, err)
		return
	}

	// Make response
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

func storeUserInfo(ctx context.Context, w http.ResponseWriter, r *http.Request) (bool, error) {
	// get User Infomation
	//_, err := user.CurrentOAuth(ctx, "")
	//if err != nil {
	//	// http.Error(w, "OAuth Authorization header required", http.StatusUnauthorized)
	//	return  false, errors.Wrap(err, "Store Failed")
	//}

	rInfo := &RequestInfo{
		UserAgent:     r.UserAgent(),
		Referrer:      r.Referer(),
		RemoteAddress: r.RemoteAddr,
		Proto:         r.Proto,
		Host:          r.Host,
		Method:        r.Method,
		IsHttps:       isSecure(r),
		RequestDate:   time.Now(),
	}

	terr := datastore.RunInTransaction(ctx, func(ctx context.Context) error {
		key := datastore.NewIncompleteKey(ctx, "UserRequests", nil)
		//key := datastore.NewKey(ctx, "UserRequests", "tweetID", 0, nil)

		if _, err := datastore.Put(ctx, key, rInfo); err != nil {
			fmt.Fprintln(w, err)
			return errors.Wrap(err, "Transaction Failed!")
		}
		return nil
	}, nil)
	if terr != nil {
		return false, errors.Wrap(terr, "Store Failed!")
	}
	return true, nil
}
