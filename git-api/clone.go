package gitapi

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/go-git/go-git/v6"
)

// 实现一个自定义 Writer 用于回调输出
type progressWriter struct {
	callback func(string)
}

func (p *progressWriter) Write(data []byte) (n int, err error) {
	// 将流数据按行读出
	reader := bufio.NewReader(strings.NewReader(string(data)))
	for {
		line, err := reader.ReadString('\n')
		if line != "" {
			p.callback(strings.TrimRight(line, "\n"))
		}
		if err == io.EOF {
			break
		}
	}
	return len(data), nil
}

func (g *GitApiService) GitClone(url string, path string, progress ...func(msg string)) error {
	if g.Repo != nil {
		return fmt.Errorf("仓库已存在")
	}

	var pw io.Writer
	if len(progress) > 0 && progress[0] != nil {
		pw = &progressWriter{callback: progress[0]}
	}

	r, err := git.PlainClone(path, &git.CloneOptions{
		URL:      url,
		Progress: pw, // ⭐ 将进度回调写入这里！
	})
	if err != nil {
		return err
	}

	g.Repo = r
	return nil
}
