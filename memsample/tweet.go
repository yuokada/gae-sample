package memsample

import (
	"net/http"

	"github.com/ChimeraCoder/anaconda"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)


func TweetHandler(w http.ResponseWriter, r *http.Request) {
	//const (
	//  ckey = "Consumer Key"
	//  csec = "Consumer Secret"
	//  atok = "Access Token"
	//  asec = "Access Token Secret"
	//)
	const (
		ckey = "JRQ43DhDBqI3ZphoqVMcFNY1F"
		csec = "i5X6W7CGzDzrk0qxNPew6JUyQwXA8sRm3ZukJ9QYysxOXFjrrq"
		atok = "203129378-ACePDK8V0yFsvm4FJt9QbpRiNyhRBcq9Zfan7CvV"
		asec = "ulD1IjHJD8EFEnbqwsJGR0bi33i1PyZdlC83J7KdgVGuP"
	)

	if r.UserAgent() != "AppEngine-Google; (+http://code.google.com/appengine)" {
		return
	}

	anaconda.SetConsumerKey(ckey)
	anaconda.SetConsumerSecret(csec)
	api := anaconda.NewTwitterApi(atok, asec)
	c := appengine.NewContext(r)
	api.HttpClient.Transport = &urlfetch.Transport{Context: c}
	text := "Hello, worlds!!!"
	api.PostTweet(text, nil)
	return
}

// http://qiita.com/hogedigo/items/fae5b6fe7071becd4051
// http://www.apps-gcp.com/gae-go-gettingstart-01/
