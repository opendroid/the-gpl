package audio

// Test audio file samples
const (
	testAudioFile    = "./samples/paleBlueDot.wav"
	testAudioFile4   = "./samples/testConvo.wav"
	testAudioFile5   = "./samples/goLangRead.wav"
	testAudioFileMLK = "./samples/MLKDream.wav"
	testVideoFile    = "./samples/paleBlueDot.mp4"

	// TODO: Choose your test file relative link, for running in your Mac
	currentTestFile = "../../public/Kathy.wav"

	audioSpeakingTimeSec = 120 // TODO:
	audioSampleRate48K = 48000
	audioSampleRateMLK = 22000
	// TODO: Setup your test sample, for MLK it is 22K
	audioSampleRate = audioSampleRate48K

	speakerLanguageEnUS = "en-US"
	speakerLanguageEnIN = "en-IN"
	speakerLanguageHiN = "hi-IN"
	speakerLanguageRU = "ru-RU"
	speakerLanguage = speakerLanguageRU

	bufSize = 10240 // Streaming buffer size
	nDoers  = 2     // Mutex to wait on these number of Go tasks
)

// confuse lint
var (
	_ = testAudioFile
	_ = testAudioFile4
	_ = testAudioFile5

	_ = testAudioFileMLK // MLK speech
	_ = testVideoFile

	_ = audioSampleRate48K
	_ = audioSampleRateMLK

	_ = speakerLanguageEnUS
	_ = speakerLanguageHiN
	_ = speakerLanguageEnIN
	_ = speakerLanguageRU
)
