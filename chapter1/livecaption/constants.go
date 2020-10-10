package livecaption

// Test livecaption file samples
const (
	testAudioFile    = "./samples/paleBlueDot.wav"
	testAudioFile4   = "./samples/testConvo.wav"
	testAudioFile5   = "./samples/goLangRead.wav"
	testAudioFileMLK = "./samples/MLKDream.wav"

	// TODO: Choose your test file relative link, for running in your Mac
	currentTestFile = testAudioFile

	audioSpeakingTimeSec = 120 // TODO:
	audioSampleRate48K   = 48000
	audioSampleRateMLK   = 22000
	// TODO: Setup your test sample, for MLK it is 22K
	audioSampleRate = audioSampleRate48K

	speakerLanguageEnUS = "en-US"
	speakerLanguageEnIN = "en-IN"
	speakerLanguageHiN  = "hi-IN"
	speakerLanguageRU   = "ru-RU"
	speakerLanguage     = speakerLanguageEnUS
	speakerShowIntermediate = true

	// defaultRTPPort where RTP livecaption is being streamed
	defaultRTPPort = 9999

	bufSize = 10240 // Streaming buffer size
	nDoers  = 2     // Mutex to wait on these number of Go tasks
)

// Reference: https://twinnation.org/articles/35/how-to-add-colors-to-your-console-terminal-output-in-go
type LineColor string
const (
	Reset  LineColor = "\033[0m"
	Red    LineColor = "\033[31m"
	Green  LineColor = "\033[32m"
	Yellow LineColor = "\033[33m"
	Blue   LineColor = "\033[34m"
	Purple LineColor = "\033[35m"
	Cyan   LineColor = "\033[36m"
	Gray   LineColor = "\033[37m"
	White  LineColor = "\033[97m"
	termWidth = 80
)

// confuse lint
var (
	_ = testAudioFile
	_ = testAudioFile4
	_ = testAudioFile5

	_ = testAudioFileMLK // MLK speech

	_ = audioSampleRate48K // Bit Rates
	_ = audioSampleRateMLK

	_ = speakerLanguageEnUS // Locales
	_ = speakerLanguageHiN
	_ = speakerLanguageEnIN
	_ = speakerLanguageRU

	_ = Reset // Line Colors
	_ = Red
	_ = Green
	_ = Yellow
	_ = Blue
	_ = Purple
	_ = Cyan
	_ = Gray
	_ = White
)
