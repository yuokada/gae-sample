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
)

type RequestInfo struct {
	UserAgent     string
	RemoteAddress string
	Cookie        string `datastore:",noindex"`
	Host          string
	Method        string `datastore:",noindex"`
	Proto         string `datastore:",noindex"`
	IsHttps       bool   `datastore:",noindex"`

	RequestDate   time.Time
	// 	Key         *datastore.Key
}

const base64GifPixel = "R0lGODlhAQABAID/AMDAwAAAACH5BAEAAAAALAAAAAABAAEAAAICRAEAOw=="

var output []byte

func BeaconHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	StoreUserInfo(ctx, r)

	output, _ = base64.StdEncoding.DecodeString(base64GifPixel)
	w.Header().Set("Content-Type", "image/gif")
	w.Header().Set("Cache-Control", "no-cache, no-store, private")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Del("Expires")
	io.WriteString(w, string(output))
}

func isSecure(r *http.Request) bool {
	t := r.TLS
	if t == nil  {
		return false
	}
	return  true
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

	key := datastore.NewIncompleteKey(ctx, "UserRequest", nil)

	if _, err := datastore.Put(ctx, key, rInfo); err != nil {
		return false, err
	}
	return true, nil
}

// func main() {
// 	//logger := log.New(os.Stdout, "", 0)
// 	logger := apachelog.NewApacheLog(os.Stderr,"...")
//
// 	//godaemon.MakeDaemon(&godaemon.DaemonAttr{})
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/hey", HeyHandler)
// 	mux.HandleFunc("/t", respHandler)
//
// 	fmt.Println("Access: http://localhost:8080/*")
// 	err := http.ListenAndServe(":8080", apachelog.WrapLoggingWriter(mux, logger))
// 	//err := http.ListenAndServe(":8080", mux)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// }
