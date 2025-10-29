package bleveapi

import (
	"fmt"

	ttmlSearchCore "github.com/xiaowumin-mark/ttml-search-core"
)

type BleveApiSearchKV struct {
	Key string `json:"key"`

	Value string `json:"value"`
}

type BleveApiSearchResult struct {
	Meta      ttmlSearchCore.LyricIndex
	Lyrics    ttmlSearchCore.LyricItem
	Highlight map[string][]string
}

func (b *BleveApiService) BleveApiSearch(args []BleveApiSearchKV, highlight, isLyrics bool) ([]*BleveApiSearchResult, error) {

	if len(args) == 0 { // 搜索全部
		return nil, fmt.Errorf("请输入搜索条件")
	}

	if b.MetaIndex == nil || b.LyricIndex == nil {
		return nil, fmt.Errorf("索引未初始化")
	}
	var sf []ttmlSearchCore.LyricsSearchFiled
	for _, item := range args {
		sf = append(sf, ttmlSearchCore.LyricsSearchFiled{
			FieldName:  item.Key,
			FieldValue: item.Value,
		})
	}
	var resu []*BleveApiSearchResult
	if isLyrics {

		request := b.LyricIndex.BuildSearchFiled(
			sf,
		)
		_, res, err := b.LyricIndex.FindList(request, highlight)
		if err != nil {
			return resu, err
		}

		for _, item := range res.Hits {
			songid := item.Fields["songId"].(string)
			// 二次查询 获取元数据
			request := b.MetaIndex.BuildSearchFiled([]ttmlSearchCore.LyricsSearchFiled{
				{"id", songid},
			})
			_, metaRes, err := b.MetaIndex.FindOne(request, false)
			if err != nil {
				return resu, err
			}
			var metaFields map[string]interface{}
			if len(metaRes.Hits) > 0 {
				//fmt.Println("Meta:", metaRes.Hits[0].Fields)
				//fmt.Println("tiele:", metaRes.Hits[0].Fields["title"])
				metaFields = metaRes.Hits[0].Fields

			} else {
				return resu, fmt.Errorf("没有找到歌曲")
			}
			resu = append(resu, &BleveApiSearchResult{
				Meta: *BuildMeta(metaFields),
				Lyrics: ttmlSearchCore.LyricItem{
					ID:        item.Fields["id"].(string),
					SongID:    item.Fields["songId"].(string),
					Text:      item.Fields["text"].(string),
					Trans:     item.Fields["trans"].(string),
					Roman:     item.Fields["roman"].(string),
					StartTime: int64(item.Fields["startTime"].(float64)),
					EndTime:   int64(item.Fields["endTime"].(float64)),
				},
				Highlight: item.Fragments,
			})
		}
	} else {
		request := b.MetaIndex.BuildSearchFiled(
			sf,
		)
		_, res, err := b.MetaIndex.FindList(request, highlight)
		if err != nil {
			return resu, err
		}

		for _, item := range res.Hits {
			resu = append(resu, &BleveApiSearchResult{
				Meta:      *BuildMeta(item.Fields),
				Lyrics:    ttmlSearchCore.LyricItem{},
				Highlight: item.Fragments,
			})
		}
	}
	return resu, nil

}

func BuildMeta(metaFields map[string]interface{}) *ttmlSearchCore.LyricIndex {
	var li ttmlSearchCore.LyricIndex
	if title, ok := metaFields["title"].([]any); ok {
		// 转换为[]string
		li.Title = make([]string, len(title))
		for i, v := range title {
			li.Title[i] = v.(string)
		}
	} else {
		li.Title = []string{metaFields["title"].(string)}
	}
	if artist, ok := metaFields["artist"].([]any); ok {
		li.Artist = make([]string, len(artist))
		for i, v := range artist {
			li.Artist[i] = v.(string)
		}
	} else {
		li.Artist = []string{metaFields["artist"].(string)}
	}
	if album, ok := metaFields["album"].([]any); ok {
		li.Album = make([]string, len(album))
		for i, v := range album {
			li.Album[i] = v.(string)
		}
	} else {
		li.Album = []string{metaFields["album"].(string)}
	}
	li.LyricsCount = int(metaFields["lyricsCount"].(float64))
	li.ID = metaFields["id"].(string)
	return &li
}
