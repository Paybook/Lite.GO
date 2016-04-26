# Lite.GO
A very simple and light interface to demonstrate how to take advantage of the Paybook Financial API (Sync) to pull information from Mexican Banks and Tax Authority.

## Requirements
1. [GO](https://golang.org/dl/) Stable version 1.6.1, follow this guide to install [GO](https://golang.org/doc/install)
2. [NodeJS](https://nodejs.org/en/) Stable version 5.10.1
3. [Bower](http://bower.io)
4. Paybook Sync API key

## Install (cli / terminal)
1. In your $GOPATH/src create **paybook.com** directory 
2. cd $GOPATH/src/paybook.com/
3. git clone https://github.com/Paybook/Lite.GO/ lite
4. go get github.com/beego/bee
5. go install paybook.com/lite
6. cd $GOPATH/src/paybook.com/lite
7. npm install

## Configure
1. Create conf/app.conf file with the following content
```
appname = lite
httpport = 9090
runmode = dev
autorender = true
SessionOn = true
copyrequestbody = true
TemplateLeft = "<<<"
TemplateRight = ">>>"

pbsync_api_key = "YOUR-API-KEY-HERE"
pbsync_base_url = "https://sync.paybook.com/v1/"
pbsync_files_url = "https://s.paybook.com"
```

## Execute (cli / terminal)
1. cd **$GOPATH/src/paybook.com/lite**
2. Type: **bee run**
3. Test it [http://localhost:9090/](http://localhost:9090/)
4. Create a new user on Signup [http://localhost:9090/signup](http://localhost:9090/signup)
5. Login with the new user account [http://localhost:9090/login](http://localhost:9090/login)
6. Add a site account