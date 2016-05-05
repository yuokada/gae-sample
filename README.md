# OverView 
GAE Goの練習

## Prepare

環境変数をセットしてPATHを通しておく。

    $ export   PATH=~/go_appengine:$PATH


## Development

ローカルで開発用のサーバーを起動する。

```
% goapp serve .
INFO     2016-05-05 05:18:09,640 devappserver2.py:706] Skipping SDK update check.
WARNING  2016-05-05 05:18:09,645 api_server.py:378] Could not initialize images API; you are likely missing the Python "PIL" module.
WARNING  2016-05-05 05:18:09,648 simple_search_stub.py:1090] Could not read search indexes from /var/folders/d9/8bnn6sfj32x6pv14t1k5hflw0000gq/T/appengine.yama-stage.yuokada/search_indexes
INFO     2016-05-05 05:18:09,651 api_server.py:171] Starting API server at: http://localhost:59747
INFO     2016-05-05 05:18:09,655 dispatcher.py:182] Starting module "default" running at: http://localhost:8080
INFO     2016-05-05 05:18:09,657 admin_server.py:117] Starting admin server at: http://localhost:8000
```

サーバーが起動したらこれらのポートにアクセスして動作を確認する
- http://localhost:8000/
- http://localhost:8080/

## Deploy

Example:

    % goapp serve .
    INFO     2016-05-04 19:37:06,923 devappserver2.py:706] Skipping SDK update check.
    WARNING  2016-05-04 19:37:06,927 api_server.py:378] Could not initialize images API; you are likely missing the Python "PIL" module.
    INFO     2016-05-04 19:37:06,932 api_server.py:171] Starting API server at: http://localhost:51704
    INFO     2016-05-04 19:37:06,936 dispatcher.py:182] Starting module "default" running at: http://localhost:8080
    INFO     2016-05-04 19:37:06,939 admin_server.py:117] Starting admin server at: http://localhost:8000
    ^Cgoapp: caught SIGINT, waiting for dev_appserver.py to shut down
    INFO     2016-05-04 19:37:11,426 shutdown.py:44] Shutting down.
    INFO     2016-05-04 19:37:11,427 api_server.py:578] Applying all pending transactions and saving the datastore
    INFO     2016-05-04 19:37:11,427 api_server.py:581] Saving search indexes



## Link
later
- [Y.A.M の 雑記帳: GAE Go で static なサイトをホストする](http://y-anz-m.blogspot.jp/2014/10/gae-go-static.html "Y.A.M の 雑記帳: GAE Go で static なサイトをホストする")

### twitter bot
- [ChimeraCoder/anaconda: A Go client library for the Twitter 1.1 API][anaconda]
- [Twitter APIの使い方まとめ](https://syncer.jp/twitter-api-matome "Twitter APIの使い方まとめ")




[anaconda](https://github.com/ChimeraCoder/anaconda "ChimeraCoder/anaconda: A Go client library for the Twitter 1.1 API")