package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/wispeeer/wispeeer/internal/pkg/tools"
	"github.com/wispeeer/wispeeer/pkg/info"
	"github.com/wispeeer/wispeeer/pkg/utils"
)

func (c *CMD) NewPost(title string) error {
	title = utils.SafeFormat(title, " ", "", "")
	var safeName = utils.SafeFormat(title, "-", "md", ".")
	var filePath = filepath.Join(c.Options.SourceDir, c.Options.ArticleDir, safeName)
	if utils.IsExist(filePath) {
		return fmt.Errorf("article %v is exist", safeName)
	}
	// 创建文章文件
	err := tools.CreateMarkdown(filePath, title, "[void]")
	if err != nil {
		return fmt.Errorf("create article %s is failed", safeName)
	}
	return nil
}

func (c *CMD) NewPage(title string) error {
	// 检查发布文件夹状态
	title = utils.SafeFormat(title, " ", "", "")

	utils.SafeMkdir(c.Options.SourceDir)
	utils.SafeMkdir(filepath.Join(c.Options.SourceDir, title))

	var safeName = info.IndexMarkdownTitleStr
	var filePath = filepath.Join(c.Options.SourceDir, title, safeName)
	if utils.IsExist(filePath) {
		return fmt.Errorf("page %v is exist", filePath)
	}
	// 创建文件
	err := tools.CreateMarkdown(filePath, title, "["+title+"]")
	if err != nil {
		return fmt.Errorf("create page %s is failed", safeName)
	}
	return nil
}
