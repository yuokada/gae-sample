package memsample

import (
	"net/http"

	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
	"log"
	"time"
)

// https://cloud.google.com/appengine/docs/go/datastore/creating-entities
// https://console.cloud.google.com/datastore/settings

// http://qiita.com/hogedigo/items/fae5b6fe7071becd4051
// http://www.apps-gcp.com/gae-go-gettingstart-01/

type Secrets struct {
	Type string

	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

func GetSeacrets(c context.Context) (*Secrets, error) {
	log.Printf("%#v\n", c)
	sec := &Secrets{
		Type: "twitter",
		ConsumerKey:       "JRQ43DhDBqI3ZphoqVMcFNY1F",
		ConsumerSecret:    "i5X6W7CGzDzrk0qxNPew6JUyQwXA8sRm3ZukJ9QYysxOXFjrrq",
		AccessToken:       "203129378-ACePDK8V0yFsvm4FJt9QbpRiNyhRBcq9Zfan7CvV",
		AccessTokenSecret: "ulD1IjHJD8EFEnbqwsJGR0bi33i1PyZdlC83J7KdgVGuP",
	}
	return sec, nil
}

func TweetHandler(w http.ResponseWriter, r *http.Request) {
	// Deny Request except without GAE
	if r.Header.Get("X-Appengine-Cron") != "true" {
		return
	}

	// Fetch Seacrets From Datastore(NoSQL) or else.
	c := appengine.NewContext(r)
	sec, err := GetSeacrets(c)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	anaconda.SetConsumerKey(sec.ConsumerKey)
	anaconda.SetConsumerSecret(sec.ConsumerSecret)
	api := anaconda.NewTwitterApi(sec.AccessToken, sec.AccessTokenSecret)
	api.HttpClient.Transport = &urlfetch.Transport{Context: c}
	t := time.Now()
	text := "Hello, world!! powered by GAE /" + t.String()
	api.PostTweet(text, nil)
	return
}
