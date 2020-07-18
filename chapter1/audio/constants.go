package audio

// Test audio file samples
const (
	testAudioFile    = "../../public/paleBlueDot.wav"
	testAudioFile4   = "../../public/testConvo.wav"
	testAudioFile5   = "../../public/goLangRead.wav"
	testAudioFileMLK = "../../public/MLKDream.wav"
	testVideoFile    = "../../public/paleBlueDot.mp4"

	// TODO: Choose your test file relative link, for running in your Mac
	currentTestFile = "../../public/Kathy.wav"

	audioSpeakingTimeSec = 2 // TODO:
	audioSampleRate48K = 48000
	audioSampleRateMLK = 22000
	// TODO: Setup your test sample, for MLK it is 22K
	audioSampleRate = audioSampleRate48K

	speakerLanguageEnUS = "en-US"
	speakerLanguageEnIN = "en-IN"
	speakerLanguageHiN = "hi-IN"
	speakerLanguage = speakerLanguageEnIN

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
)
