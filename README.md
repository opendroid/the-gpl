# the-gpl
[![Build Status](https://travis-ci.org/opendroid/the-gpl.svg?branch=master)](https://travis-ci.org/opendroid/the-gpl)

[The Go Programming 
Language](https://www.amazon.com/Programming-Language-Addison-Wesley-Professional-Computing/dp/0134190440) 
is a seminal programming language authored by Avengers: _Alan A. A. Donovan_ and _Brian W. Kernighan_. 
This git repo is an attempt to share my learning from this book to showcase problems posed in the book and then some. 
The source-code by authors is on GitHub at [gopl.io](https://github.com/adonovan/gopl.io/).

You can experience the deployed artifacts on the web as:
1. GPC Cloud Run `the-gpl-book` service.
   - [Home](https://the-gpl.com/index), [About](https://the-gpl.com/about) 
   - [Post data](https://the-gpl.com/index?q="hello"&l="TheGOGPL"&a="Pike+Donovan")
   - [Lissajous Graph](https://the-gpl.com/lis)
   - [Mandel](https://the-gpl.com/mandel), [Black & White](https://the-gpl.com/mandelbw)
   - Surfaces: [Sinc](https://the-gpl.com/sinc), [Eggs](https://the-gpl.com/egg), [Valley](https://the-gpl.com/valley), [Sq](https://the-gpl.com/sq)
2. As [docker container image](https://hub.docker.com/repository/docker/uopendocker/the-gpl).
3. Curls eg: `curl https://the-gpl.com/who`

## Running from Web
You can start a webserver and see simple web-server examples. The command is:
```shell script
$ the-gpl server -port=8080 # start a web server at port 8080.
```

## Running from CLI
Assuming, the program is installed locally as `the-gpl` you can access several methods using a CLI. 
To use autocompletion `source the-gpl-autocomplete.sh`.
Some examples are:

### Google API examples
To use the Google  _Dialogflow Agent_ and _Speech-to-text_ set up:
1. **GOOGLE_APPLICATION_CREDENTIALS** shell variable
2. Enable APIs
 
Here are commands to communicate with the Dialogflow Agent.
```shell script
$ the-gpl # Prints the help of all modules
$ the-gpl bot -project=gcp-project-id # Will do a short conversation with an agent. 
$ the-gpl bot -project=gcp-project-id -chat=true # Can send messages from stdin
$ the-gpl bot -chat=true -project=gcp-project-id -lang=en-US # Chat with an agent in en-US
```

To run live-caption speech to text first start a microphone stream on RTP port, and then use the-gpl to listen and apply STT.
```shell script
$ ffmpeg -f avfoundation -i ":1" -acodec pcm_s16le -ar 48000 -f s16le udp://localhost:9999 # macOS Start microphone streaming
$ the-gpl stt -port=9999 # Will listen to RTP stream on port 9999 for 2 minutes and transcribe in real time
```

### Simple Examples from book
Use these commands to run utilities:
1. bits: That counts number of 1 bits in a Hex input 
2. Array examples using `mas` i.e., maps, arrays and string utilities.
3. Temperature conversions among `°F, °K and °C`.
4. Measure disk usage in a directory using command `du`
5. Saving Lissajous gif to a file.

```shell script
$ the-gpl bits -n=0xBAD0FACE # will count 1 bits in n
$ the-gpl mas -fn=array # Tests array
$ the-gpl mas -fn=comp -n1=123 -n2=345 # Compare n1 and n2

# Temperature utilities
$ the-gpl temp -c=12 -f=12 -k=12 # Converts 12°C to °C/°F/°K
$ the-gpl degrees -c=12°F -f=12°K -k=12°C # Converts 12°C to °C/°F/°K

# du: Disk Usage calculates size of all files in a directory recursively, using go-routines
$ the-gpl du -dir=$HOME/gocode

# Output a Lissajous graph to -file of size 1024 pixels 20 frames and 10 cycles
$ the-gpl lissajous -file ~/Downloads/lis.gif -size=1024 -frames=20 -cycles=10
```

### Crawling Examples
These commands fetch a website and does various operations on it.
```shell script
# Parse various HTML content of sites URL
$ the-gpl parse -type=outline -site=https://www.airbnb.com # Creates a summary outline of a page
$ the-gpl parse -type=links -site=https://images.google.com #  Prints all links on a webpage
$ the-gpl parse -type=images -site=https://www.yahoo.com # Fetches image URLs in a site
$ the-gpl parse -type=pretty -site=https://www.google.com
$ the-gpl parse -type=crawl -site=https://www.google.com  -dir=/Users/guest/Downloads # Crawl pages to /Users/guest/Downloads/www.google.com 
```

### Simple Servers and Clients

```shell script
# Server-client 
$ the-gpl service -sp="clock:9999"  # -sp servicePort start clock  service on port 9999
$ the-gpl client  -cp="clock:9999"  # -cp clientPort  start clock  client  on port 9999
$ nc localhost 9999                 # use Mac netcat 'nc' client on port 9999
$ the-gpl service -sp="reverb:9998" # -sp servicePort start reverb service on port 9998
$ the-gpl client  -cp="reverb:9998" # -cp clientPort  start reverb client  on port 9998
$ the-gpl service -sp="chat:9997"   # starts a chat service. Join using:
$ nc localhost 9997                 # Joins chat session as a client
```

## Building from Local Machine

The GPL application can be built using:
1. The plain old `go install`
2. GCP Cloud run container.
3. As a docker image on [docker.com](https://hub.docker.com/r/uopendocker/the-gpl).

See [The GPL Docker](https://github.com/opendroid/the-gpl/wiki/The-GPL-Docker) wiki for docker steps,
 
 ## Mandelbrot
Here are some sample Mandelbrot fractals created.

![Color](public/images/media/mandel-color-256.png?raw=true "Color Mandelbrot Graph")
![B&W](public/images/media/mandel-bw-256.png?raw=true "Color Mandelbrot Graph")
