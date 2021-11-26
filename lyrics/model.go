package lyrics

type ColorLyrics struct {
	Lyrics *LyricsInfo `json:"lyrics"`
	Colors *ColorsInfo `json:"colors"`
}

type LyricsInfo struct {
	SyncType string        `json:"syncType"`
	Lines    []*LyricsLine `json:"lines"`
}

type ColorsInfo struct {
	Background    int `json:"background"`
	Text          int `json:"text"`
	HighlightText int `json:"highlightText"`
}
type LyricsLine struct {
	StartTimeMs int64  `json:"startTimeMs,string"`
	Words       string `json:"words"`
}
