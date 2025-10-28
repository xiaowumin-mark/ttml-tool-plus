package bleveapi

import (
	"fmt"
	"strings"

	ttmlSearchCore "github.com/xiaowumin-mark/ttml-search-core"
	"github.com/xiaowumin-mark/ttml-search-core/ttml"
)

type BleveApiService struct {
	Config *map[string]any

	MetaIndex  *ttmlSearchCore.LyricsClient
	LyricIndex *ttmlSearchCore.LyricsClient
}

type Task struct {
	ID   string
	Path string
}
type LyricDoc struct {
	LyricIndex ttmlSearchCore.LyricIndex
	Lyrics     []*ttmlSearchCore.LyricItem
}

// 从 TTMLLyric 生成可索引结构
func BuildIndexData(id string, lyric ttml.TTMLLyric) (ttmlSearchCore.LyricIndex, []*ttmlSearchCore.LyricItem) {
	title := []string{}
	artist := []string{}
	album := []string{}
	for _, m := range lyric.Metadata {
		switch strings.ToLower(m.Key) {
		case "musicname":
			title = m.Value
		case "artists":
			artist = m.Value
		case "album":
			album = m.Value
		}
	}

	var lyrics []*ttmlSearchCore.LyricItem
	var textBuilder strings.Builder
	for index, line := range lyric.LyricLines {
		var lyric ttmlSearchCore.LyricItem
		for _, word := range line.Words {
			textBuilder.WriteString(word.Word)
		}
		lyric.Text = textBuilder.String()
		textBuilder.Reset()
		if line.TranslatedLyric != "" {
			lyric.Trans = line.TranslatedLyric
		}
		if line.RomanLyric != "" {
			lyric.Roman = line.RomanLyric
		}
		lyric.StartTime = int64(line.StartTime)
		lyric.EndTime = int64(line.EndTime)
		lyric.LineIndex = int64(index)
		lyric.SongID = id
		lyric.ID = fmt.Sprintf("%s-%d", id, index)
		lyrics = append(lyrics, &lyric)
	}

	return ttmlSearchCore.LyricIndex{
			ID:          id,
			Title:       title,
			Artist:      artist,
			Album:       album,
			LyricsCount: len(lyrics),
		},
		lyrics
}
