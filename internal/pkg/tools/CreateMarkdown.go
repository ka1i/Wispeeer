package tools

import (
	"bufio"
	"os"
	"time"

	"github.com/wispeeer/wispeeer/pkg/logger"
)

func CreateMarkdown(fileName string, title string, tags string) error {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	articlePostedTime := time.Now().UTC().Format("2006-01-02 15:04:05")

	fileWrite := bufio.NewWriter(file)
	//Markdown header
	fileWrite.WriteString("title: " + title + "\n")
	fileWrite.WriteString("posted: " + articlePostedTime + "\n")
	fileWrite.WriteString("tags: " + tags + "\n")
	fileWrite.WriteString("categories: " + tags + "\n")
	fileWrite.WriteString("------\n")
	fileWrite.WriteString("\n\n")
	fileWrite.WriteString("# Absract\n")
	fileWrite.WriteString("<!--more-->\n\n")
	fileWrite.WriteString("# Reference\n\n")

	//Flush buffer
	fileWrite.Flush()

	logger.Task("new").Infof("Created: %s\n", fileName)
	logger.Task("new").Infof("Posted : %s\n", articlePostedTime)
	return nil
}
