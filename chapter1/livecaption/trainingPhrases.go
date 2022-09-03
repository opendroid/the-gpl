package livecaption

var (
	// trainingPhrases contains [supported class tokens ].
	//
	// [supported class tokens ]: https://cloud.google.com/speech-to-text/docs/class-tokens
	trainingPhrases = []string{"Hello", "$TIME", "$PERCENT"}

	_ = trainingPhrases
)
