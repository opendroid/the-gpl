# Audio Stream
This exercise is about streaming. The private method `sendStreamToGCP` sends an audio stream
of bytes, read from a test audio `.wav` file in `public/` directory. The `readStreamFromGCP`
it then receives the speech-to-text stream and prints it. To run this please
ensure that you have set `GOOGLE_APPLICATION_CREDENTIALS` with appropriate permissions 
for your GCP Project.

The go routines run in parallel:
1. sendStreamToGCP: audio file => Reader() stream => StreamingRecognizeClient Send()
2. readStreamFromGCP:  StreamingRecognizeClient Recv() => stream Writer() => os.Stdout

The API limits  are specified in [Quotas & limits](https://cloud.google.com/speech-to-text/quotas).
There is no limit on this streaming version. Below is the transcription of `I Have a Dream, Martin Luther King Jr.`
speech by this API. The [speech audio](../../public/MLKDream.wav) is 16 minutes long (at 22K bit rate).

## Mac Audio
On a Mac you can create a sample .wav file using `ffmpeg` pre-installed program.

### Audio Samples on Mac

As example epilogue of [A Pale Blue Dot](https://www.planetary.org/explore/space-topics/earth/pale-blue-dot.html) recorded 
in `public/paleBlueDot.wav`:

```text
There is perhaps no better demonstration of the folly of human 
conceits than this distant image of our tiny world. To me, it underscores our 
responsibility to deal more kindly with one another, and to preserve and cherish 
the pale blue dot, the only home we've ever known.
```

The [Google ML](https://cloud.google.com/sdk/gcloud/reference/ml/speech/recognize) command line shows the reference output:
` gcloud ml speech recognize ../../public/paleBlueDot.wav  --language-code=en_US`

Output:
```json
{
  "results": [
    {
      "alternatives": [
        {
          "confidence": 0.9397845,
          "transcript": "there is perhaps no better demonstration of the Folly of human conceit than this image of a world to me and of course a responsibility to deal more kindly with one another and to preserve and cherish the pale blue dot"
        }
      ]
    }
  ]
}
```
You can check out yourself what stream client returned.

## Handy Mac Shell Commands
You can use `ffmpeg` Macc command line program to record a .wav file. Some examples are:
 
```shell script
# List audio devices, Mac
ffmpeg -f avfoundation -list_devices true -i ""
# Record 20 seconds of audio from the built in microphone and save it in playBlueDot.mp3
ffmpeg -f avfoundation -i ":1" -t 20 ../../public/playBlueDot.wav
# Stream to a RTP port
ffmpeg -f avfoundation -i ":1" -acodec libmp3lame -ab 32k -ac 1 -f rtp rtp://0.0.0.0:12345
# Meta data
mdls chapter1/audio/playBlueDot.wav
# Test playback
afplay chapter1/audio/playBlueDot.wav
```

The same command can be used to recording Video files:
```shell script
# Record from video device 0 and audio device 0:
ffmpeg -r 30 -f avfoundation -i "0:1" ../../public/paleBlueDot.mp4
ffmpeg -f avfoundation -framerate 30 -video_size 640x480 -i "0:1" ../../public/paleBlueDot.mp4
```

Formats flag in the command are:
```text
ffmpeg flags:
  -f = "force format". In this case we're forcing the use of AVFoundation
  -i = input source. Typically it's a file, but you can use devices.
        "0:1" = Record both audio and video from FaceTime camera and built-in mic
        "0" = Record just video from FaceTime camera
        ":1" = Record just audio from built-in mic
  -t = time in seconds. If you want it to run indefinitely until you stop it  (ControlC)
```

## Material:
You may find following links handy if you like to check more audio related stuff in Golang.

 - [FFMPEG ](https://ffmpeg.org/), a complete, cross-platform solution to record, convert and stream audio and video.
 - [Creating Multiple Outputs](https://trac.ffmpeg.org/wiki/Creating%20multiple%20outputs)
 - [How to stream audio using FFMPEG?](https://apple.stackexchange.com/questions/326419/how-to-stream-audio-using-ffmpeg)
 - [PortAudioAPI Overview](http://portaudio.com/docs/v19-doxydocs/api_overview.html)
 - [Golang Package PortAudio](https://pkg.go.dev/github.com/gordonklaus/portaudio?tab=doc)
 - [How to build an audio streaming server in Go](https://medium.com/@valentijnnieman_79984/how-to-build-an-audio-streaming-server-in-go-part-1-1676eed93021)
 - [Live Caption Sample Code](https://github.com/GoogleCloudPlatform/golang-samples/blob/master/speech/livecaption/livecaption.go)
 - [Transcribing audio from streaming input](https://cloud.google.com/speech-to-text/docs/streaming-recognize)
 - [Endless streaming tutorials](https://cloud.google.com/speech-to-text/docs/endless-streaming-tutorial)
 - [I Have a Dream, Martin Luther King Jr.](https://archive.org/details/MLKDream)

 ## MLK Speech Transcribed
 Words transcribed: 529 v/s actual words in speech 881.
 ```text
  I say to you today my friend
  so even though we Face the difficulties of today and tomorrow
  I still have a dream
  it is a dream deeply rooted in the American dream
  I have a dream
  one day
  this nation will rise up
  live out the true meaning of its trees
  we hold these truths to be self-evident that all men are created equal
  I have a dream
  that one day on the Red Hills of Georgia
  sons of former slaves and the sons of former slave owners
  will baby be able to sit down together at the table of Brotherhood I have a dream
  the one thing
  even the state of Mississippi a state sweltering with the keto Zone Injustice
  sweltering with the heat of Oppression
  be transformed into an oasis of freedom and Justice I have a trees
  my four little children
  one day live in a nation where they will not be judged by the color of their skin but by the content of a character I have a dream today
  I have a dream that one day
  in Alabama with its vicious racists
  what's its Governor having his lips dripping with the words of interposition and nullification one day right there in Alabama little black boys and black girls little join hands with little white balls and white girls as sisters and brothers I have a dream today
  I have a dream that one day I shall be exalted never healed in Mountain shall be made Low Places would be made friends and the Crooked places will be made at All Flesh shall see it together and I hope this is a piece that I go back to the Southwest River space we will be able to shoot out of the Mountain of Despair a stone of hope we will be able to transform the jangling call Javon nation into a beautiful Symphony of Brotherhood with this face we will be able to work together to pray together to struggle together to go to jail for Freedom together knowing that we will be free one day
  this will be the day
  this will be the day when all of God's children
  be able to sing with new meaning my country tears would be
  sweet land of liberty of Beyonce the Pilgrim's Pride from every Mountainside Let Freedom Ring Americans to be a great nation this must be come true and so Let Freedom Ring
  from the mighty mountains of New York
  Let Freedom Ring from the highwomen alligators of Pennsylvania Let Freedom Ring from the smoke
  Let Freedom Ring from the probation. California but not only that
  Let Freedom Ring from Stone Mountain of Georgia
  Let Freedom Ring from Lookout Mountain of Tennessee and Mississippi State play tomorrow in Japanese
  turn wheel of freedom
  when we let it ring from every finish whatever Hamlet from every state and their Parsippany
  we will be able to speed up that they put on Jews and Gentiles Protestants and Catholics will be able to tell her I'm sending the words of the old Negro spiritual free at last free at last thank God Almighty we are free
```