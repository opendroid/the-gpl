# the-gpl
[![Build Status](https://travis-ci.org/opendroid/the-gpl.svg?branch=master)](https://travis-ci.org/opendroid/the-gpl)

[The Go Programming 
Language](https://www.amazon.com/Programming-Language-Addison-Wesley-Professional-Computing/dp/0134190440) 
by _Alan A. A. Donovan_ and _Brian W. Kernighan_ is a classic Go book. This git repo is an attempt to share my learning from 
this book in terms of solving problems posed in the book and then some. The source code by authors is on github 
in repo [gopl.io](https://github.com/adonovan/gopl.io/).

You can access the deployed artifacts as:
 1. GPC Cloud Run `the-gpl-book` service.
   - [Post data](https://the-gpl-book-vs6xxfdoxa-uc.a.run.app/post?q="hello"&l="TheGOGPL"&a="Pike+Donovan")
   - [See Lissajous Graph](https://the-gpl-book-vs6xxfdoxa-uc.a.run.app/lis)
   - [Sinc Surface](https://the-gpl-book-vs6xxfdoxa-uc.a.run.app/sinc)
   - [Eggs Surface](https://the-gpl-book-vs6xxfdoxa-uc.a.run.app/egg)
   - [Valley Surface](https://the-gpl-book-vs6xxfdoxa-uc.a.run.app/valley)
   - [Sq Surface](https://the-gpl-book-vs6xxfdoxa-uc.a.run.app/sq)
   - [Mandel](https://the-gpl-book-vs6xxfdoxa-uc.a.run.app/mandel)
   - [Mandel BW](https://the-gpl-book-vs6xxfdoxa-uc.a.run.app/mandelbw) 
   - [Echo](https://the-gpl-book-vs6xxfdoxa-uc.a.run.app/echo?q=Hello%20%F0%9F%8C%8E%F0%9F%8C%8E%F0%9F%8C%8E)

 2. As [docker container image](https://hub.docker.com/repository/docker/uopendocker/the-gpl).
 
## Running from CLI
Assuming, the program is installed locally as `the-gpl` you can access several methods using a CLI. Some examples are:

```shell script
$ the-gpl # Prints the help of all modules
$ the-gpl bot -project=gcp-project-id # Will do a short conversation with a bot. 
$ the-gpl bot -project=gcp-project-id -chat=true # Can send messages from stdin
$ the-gpl bot -chat=true -project=gcp-project-id -lang=en-US # Chat with a bot in en-US
$ the-gpl stt -port=9999 # Will listen to RTP stream on port 9999 for 2 minutes and transcribe in real time
$ the-gpl bits -n=0xBAD0FACE # will count 1 bits in n

# Output a Lissajous graph to -file of size 1024 pixels 20 frames and 10 cycles
$ the-gpl lissajous -file ~/Downloads/lis.gif -size=1024 -frames=20 -cycles=10

# Parse various HTML content of sites URL
$ the-gpl parse -type=outline -site=https://www.airbnb.com # Creates a summary outline of a page
$ the-gpl parse -type=links -site=https://images.google.com #  Prints all links on a webpage
$ the-gpl parse -type=images -site=https://www.yahoo.com # Fetches image URLs in a site
$ the-gpl parse -type=pretty -site=https://www.google.com
$ the-gpl parse -type=crawl -site=https://www.google.com  -dir=/Users/guest/Downloads # Crawl pages to /Users/guest/Downloads/www.google.com 

# Tests array
$ the-gpl mas -fn=array
$ the-gpl mas -fn=comp -n1=123 -n2=345

# Temperature utilities
$ the-gpl temp -c=12 -f=12 -k=12 # Converts 12°C to °C/°F/°K
$ the-gpl degrees -c=12°F -f=12°K -k=12°C # Converts 12°C to °C/°F/°K

# Server-client 
$ the-gpl service -sp="clock:9999" # -sp servicePort start a service for service clock on port 9999
$ the-gpl client -cp="clock:9999"  # -cp clientPort start a client for service clock on port 9999
```

## Running from Web
You can access various example outputs using Google cloud run.

```shell script
$ the-gpl server -port=8081 # start a web server at port 8081.

```

## Command go
If you are starting out with Go, highly recommend reading [How to Write Go Code](https://golang.org/doc/code.html#ImportingLocal) first. 
Here are some sample __go commands__. 

| **Command** | **Description** | **Example** |
|:--------|:-----------|:---------|
| go doc  | Shows documentation | go doc module name |
| godoc   | Local web-page help | godoc -http=localhost:6060/pkg |
| go run  | Run a main program | run main.go -func=callMas |
| go build | Build program | go build -o out/the-gpl |
| go test | Run unit tests  | go test ./... -v |
| go test | Run benchmark tests  | go test -bench -benchmem -v . ./... |
| go get | Downloads a package | go get github.com/stretchr/testify/assert |
| go list | List packages | go list ./... |

### Managing Modules

The go modules will be installed in `$GOPATH/pkg/mod` directory when we run `go install` or `go test ./...`
commands. Some handy commands are (__run these in top go module directory__): 

```shell script
go mod init # Default
go mod init github.com/opendroid/the-gpl # Alternate way
go mod tidy # do before release
go mod vendor # Optional creates a vendor directory
go clean -modcache # Clean up packages cache

# Manually getting go modules:
go get github.com/stretchr/testify/assert
go get cloud.google.com/go/dialogflow/apiv2
go get google.golang.org/genproto/googleapis/cloud/dialogflow/v2
go get google.golang.org/api/option
```
A module is a collection of related Go packages that are versioned together as a single unit.

### Other go commands

These shell commands come handy:
```shell script
go list -m all # List all your packages
go list -u -m all
go get -u ./... # Update your packages
go get -u=patch ./...
go build ./...
go test ./...
go install # installs executable the-gpl in $GOPATH/bin/
```

## Building Docker Images

This projects builds docker images on alpine OS. The images are on `gcr.io/the-gpl` in Google 
[Container Registries](https://console.cloud.google.com/gcr/images/the-gpl/GLOBAL). 
The Alpine from Google it is lightest and fastest golang docker environment.

These are build it two ways:
1. Using GCP cloud build as part of CICD pipeline. 
2. From Mac command line

### GCP Cloud Build
The git repo is configured to trigger Cloud Build when `git push origin master` is executed. The build steps
are outlines in [cloudbuild.yaml](cloudbuild.yaml). It works alongs with [Dockerfile](Dockerfile).

The `cloudbuild.yaml` and `Dockerfile` is used by Google Cloud Build to create container 
image tag and stores it to `gcr.io/the-gpl/book`.

### Command Line
The Mac commands to create docker image on local laptop are:
1. GCloud images
2. Docker image

```shell script
# Build GCP  container image using Cloud Build
gcloud builds submit --tag gcr.io/the-gpl/gpl:v1 
gcloud container images list --repository=gcr.io/the-gpl # Show your container registries

# Build on local docker 
docker build -t the-gpl:v1 . # Build version v1
docker image ls  # See docker images
docker history the-gpl:v1  # See image layers
docker image inspect the-gpl:v1
```

## Deploying Docker Images
The deployment is not part of CICD and relies on manual steps. 

1. Deploy on Google Cloud Run
2. Local docker

### Google Cloud Run
To deploy as Cloud Run service use the [GCP console](https://console.cloud.google.com) to create service for docker image.
```shell script
 gcloud run deploy --image gcr.io/the-gpl/gpl:v4 --platform managed  --allow-unauthenticated # Deploy

 gcloud auth login # Ensure you are logged to your GCP cloud project account
 gcloud config set project the-gpl # Make sure you are in right project
```
 
### Your docker repository 
The image is also on [docker repository](https://hub.docker.com/repository/docker/uopendocker/the-gpl). 
In order to push images to docker repo commit it, add tag and push. e.g:

```shell script
docker tag the-gpl-image:v1 uopendocker/the-gpl:v1
docker push uopendocker/the-gpl:v1
```


### Your local docker
You can build a image with specific tag and run it in local docker container.
```shell script
docker build -t the-gpl:v1 . # Build version v1
docker image ls  # See docker images
```

Other handy docker commands are:
```shell script
# Sample cloud run examples, remove container after run
docker run --rm the-gpl:v1 ./the-gpl mas -fn=array
docker run --rm the-gpl:v1 ./the-gpl mas -fn=slice
docker run --rm the-gpl:v1 ./the-gpl mas -fn=comp -n1=123 -n2=-1234
docker run --rm the-gpl:v1 ./the-gpl fetch -site=http://www.google.com -site=http=//www.facebook.com
docker run -d -p 8080:8080 the-gpl:v1 ./the-gpl server -port=8080

docker container ls  # See all your containers
docker container ps -a # Your docker process
docker stop 58a639eecb9a # Container id of image "the-gpl"
dcoker rm 58a639eecb9a # Remove your container
docker image rm image:version # Remove unused images
docker container prune # Remove stopped containers
docker image prune # Remove images
``` 

You can create your own local log file by mounting it to docker container.
```shell script
# If you like to log to a file.
export DEV_LOGS="~/Logs/DevLogs/"
docker run -d -p 8080:8080 -v $DEV_LOGS:/app/logs the-gpl
```

Make sure to prune your containers and images:
```shell script
docker container prune
docker image prune
```

## Quick Links 

 - [How to Write Go Code](https://golang.org/doc/code.html#ImportingLocal) explains go modules.
 - [Command go](https://golang.org/cmd/go/)
 - [Go Modules](https://github.com/golang/go/wiki/Modules)
 - [Package Management With Go Modules: The Pragmatic Guide](https://medium.com/@adiach3nko/package-management-with-go-modules-the-pragmatic-guide-c831b4eaaf31)
 - [Golang Setup](https://www.callicoder.com/golang-installation-setup-gopat**h-workspace/)
 - [GCP cloud containers](https://cloud.google.com/run/docs/quickstarts/build-and-deploy?_ga=2.91290522.-1679093051.1593441137).
 - [Tool builder: gcr.io/cloud-builders/go](https://github.com/GoogleCloudPlatform/cloud-builders/tree/master/go)
 - [Building Docker Containers for Go Applications](https://www.callicoder.com/docker-golang-image-container-example/)
 - [The GPL Solutions](https://xingdl2007.gitbooks.io/gopl-soljutions/content/chapter-1-tutorial.html)
 
 ## Mandelbrot
 
![Color](content/media/mandel-color-256.png?raw=true "Color Mandelbrot Graph")
![B&W](content/media/mandel-bw-256.png?raw=true "Color Mandelbrot Graph")
