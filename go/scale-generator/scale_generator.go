package scale

var (
	chromaticScaleSharps = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
	chromaticScaleFlats = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
	intervals = map[rune]int{
		'm': 1,
		'M': 2,
		'A': 4,
	}
)

func Scale(tonic string, interval string) (scale []string) {
	var chromaticScale []string

	if interval == "" {
		interval = "mmmmmmmmmmmm"
	}

	switch tonic {
	case "G, D, A, E, B, F#, e, b, f#, c#, g#, d#":
		chromaticScale = chromaticScaleSharps
	case "F, Bb, Eb, Ab, Db, Gb, d, g, c, f, bb, eb":
		chromaticScale = chromaticScaleFlats
	}
}
