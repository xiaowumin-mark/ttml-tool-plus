package bleveapi

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	gitapi "ttml-tool-plus/git-api"

	"github.com/blevesearch/bleve/v2"
	ttmlSearchCore "github.com/xiaowumin-mark/ttml-search-core"
	"github.com/xiaowumin-mark/ttml-search-core/ttml"
)

func (b *BleveApiService) BleveApiInit(onProgress ...func(string, string, float64)) error {

	var progress func(string, string, float64)
	if len(onProgress) > 0 && onProgress[0] != nil {
		progress = onProgress[0]
	} else {
		progress = func(string, string, float64) {} // 空实现，免判断
	}
	var initProgress float64 = 0

	var db gitapi.GitApiService
	db.Config = b.Config
	// 检查是否存在git仓库
	progress("Checking", "检查仓库是否存在...", initProgress)
	err := db.GitOpen()
	if err != nil {
		// git没有仓库

		return err
	}

	progress("Init", "初始化索引...", initProgress)
	b.MetaIndex = ttmlSearchCore.NewLyrics("lyrics-meta", "gse", filepath.Join((*b.Config)["bleveIndexPath"].(string), "lyrics-meta.bleve"))
	docMapping := bleve.NewDocumentMapping()

	fieldMapping := bleve.NewTextFieldMapping()
	fieldMapping.Store = true
	fieldMapping.Analyzer = "simple"

	// 正确地将字段映射添加到文档映射
	docMapping.AddFieldMappingsAt("Title", fieldMapping)
	docMapping.AddFieldMappingsAt("Artist", fieldMapping)
	docMapping.AddFieldMappingsAt("Album", fieldMapping)
	nMpig := bleve.NewTextFieldMapping()
	nMpig.Store = true
	nMpig.Index = false

	docMapping.AddFieldMappingsAt("LyricsCount", nMpig)
	// main.go - lyrics-meta 索引修正
	idMapping := bleve.NewTextFieldMapping()
	idMapping.Store = true
	idMapping.Index = true // 确保可以搜索和关联
	docMapping.AddFieldMappingsAt("ID", idMapping)

	b.MetaIndex.AddDocMapping(docMapping, "lyricsMetadata")

	b.LyricIndex = ttmlSearchCore.NewLyrics("lyrics-lines", "gse", filepath.Join((*b.Config)["bleveIndexPath"].(string), "lyrics-lines.bleve"))
	lyricDoc := bleve.NewDocumentMapping()

	// 1. 歌词文本映射 (Text)
	lyricsTextMapping := bleve.NewTextFieldMapping()
	lyricsTextMapping.Store = true
	lyricsTextMapping.Analyzer = "simple"
	lyricsTextMapping.IncludeTermVectors = true // ⭐️ 必须设置，用于高亮！

	// 2. 翻译和罗马音映射 (Trans, Roman)
	lyricsStoredMapping := bleve.NewTextFieldMapping()
	lyricsStoredMapping.Store = true
	lyricsStoredMapping.Index = false // 节省空间

	// 3. 数值/关联字段映射 (SongId, LineIndex)
	numericIndexMapping := bleve.NewTextFieldMapping()
	numericIndexMapping.Store = true // 存储
	numericIndexMapping.Index = true // ⭐️ 必须索引，用于关联查询！
	numericIndexMapping.IncludeTermVectors = false

	lyricDoc.AddFieldMappingsAt("Text", lyricsTextMapping)
	lyricDoc.AddFieldMappingsAt("Trans", lyricsStoredMapping)
	lyricDoc.AddFieldMappingsAt("Roman", lyricsStoredMapping)

	lyricDoc.AddFieldMappingsAt("StartTime", lyricsStoredMapping) // StartTime Index=true, Store=true 是可以的
	lyricDoc.AddFieldMappingsAt("EndTime", lyricsStoredMapping)
	lyricDoc.AddFieldMappingsAt("LineIndex", lyricsStoredMapping)
	lyricDoc.AddFieldMappingsAt("SongID", numericIndexMapping) // ⭐️ 核心关联字段
	lyInd := bleve.NewTextFieldMapping()
	lyInd.Store = true
	lyInd.Index = true
	lyricDoc.AddFieldMappingsAt("ID", lyInd)

	b.LyricIndex.AddDocMapping(lyricDoc, "lyric_line") // 建议使用更清晰的 Type 字符串
	progress("Init", "初始化索引完成...", initProgress)

	isonce, err := b.MetaIndex.CreateIndexMapping().CreateIndex()
	if err != nil {
		return err
	}
	islinesonce, err := b.LyricIndex.CreateIndexMapping().CreateIndex()
	if err != nil {
		return err
	}

	if islinesonce || isonce {
		progress("Parse", "解析歌词文件...", initProgress)

		// 读取文件数量 filepath.Join((*b.Config)["dbPath"].(string), "raw-lyrics")
		files, err := os.ReadDir(filepath.Join((*b.Config)["dbPath"].(string), "raw-lyrics"))
		if err != nil {
			return err
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		results := ttmlSearchCore.RunFileWorkers(
			ctx,
			filepath.Join((*b.Config)["dbPath"].(string), "raw-lyrics"),
			".ttml",
			10,
			func(fileName string) (string, error) {
				parts := strings.Split(fileName, "-")
				if len(parts) < 2 {
					return "", fmt.Errorf("invalid file name: %s", fileName)
				}
				return parts[0], nil
			},
			func(task ttmlSearchCore.Task) (LyricDoc, error) {
				data, err := os.ReadFile(task.Path)
				if err != nil {
					return LyricDoc{}, err
				}
				lyric, err := ttml.ParseTTML(string(data))
				if err != nil {
					return LyricDoc{}, err
				}
				m, l := BuildIndexData(task.ID, lyric)
				return LyricDoc{
					LyricIndex: m,
					Lyrics:     l,
				}, nil
			},
		)
		batch := b.MetaIndex.GetIndex().NewBatch()
		var index int
		for results := range results {
			index++
			initProgress = float64(index) / float64(len(files)) * 100
			progress("Parse", fmt.Sprint(results.LyricIndex.ID, results.LyricIndex.Title), initProgress)
			err := batch.Index(results.LyricIndex.ID, results.LyricIndex)
			if err != nil {
				progress("Error", fmt.Sprint("编制文件出错 ", err), initProgress)

			}
			lbatch := b.LyricIndex.GetIndex().NewBatch()
			for _, l := range results.Lyrics {
				err := lbatch.Index(l.ID, l)
				if err != nil {
					progress("Error", fmt.Sprint("编制文件出错 ", err), initProgress)

				}
			}
			err = b.LyricIndex.AddDocs(lbatch)
			if err != nil {
				progress("Error", fmt.Sprint("编制文件出错 ", err), initProgress)

			}

			progress("Index", fmt.Sprint("索引文件", results.LyricIndex.Title, "完成..."), initProgress)

		}
		err = b.MetaIndex.AddDocs(batch)
		if err != nil {
			progress("Error", fmt.Sprint("编制文件出错 ", err), initProgress)
		}
		initProgress = 100
		progress("Index", "索引完成", initProgress)
	} else {
		initProgress = 100
		progress("Index", "索引文件完成", initProgress)
	}

	return nil
}
