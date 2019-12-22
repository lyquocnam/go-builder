package module

import (
	"encoding/json"
	"fmt"
	"github.com/brentp/xopen"
	"github.com/labstack/gommon/color"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type ListFile struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type downloader struct {
	logger *color.Color
}

func NewDownloader() *downloader {
	logger := color.New()
	logger.SetOutput(log.Writer())

	return &downloader{
		logger: logger,
	}
}

func (s *downloader) Infof(format string, params ...interface{}) {
	msg := fmt.Sprintf(format, params...)
	s.logger.Println(s.logger.Green(msg))
}

func (s *downloader) Warnf(format string, params ...interface{}) {
	msg := fmt.Sprintf(format, params...)
	s.logger.Println(s.logger.Yellow(msg))
}


func (s *downloader) Run(override bool) {
	listUrl := "https://raw.githubusercontent.com/lyquocnam/go-builder/master/template/list.json"
	data, err := s.downloadFile(listUrl)
	if err != nil {
		log.Fatalln(err)
	}

	var list []ListFile
	if err := json.Unmarshal(data, &list); err != nil {
		log.Fatalln(err)
	}

	if len(list) <= 0 {
		log.Println("No files in list")
		return
	}

	currentDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	currentDir = fmt.Sprintf(`%s/deploy`, currentDir)

	s.createFolderIfNotExist(currentDir)

	for _, fUrl := range list {
		fileDest := fmt.Sprintf(`%s/%s`, currentDir, fUrl.Name)
		// check exist
		if !override && s.checkFileExist(fileDest) {
			s.Warnf(`❌ '%s' skipped.`, fUrl.Name)
			continue
		}

		itemData, err := s.downloadFile(fUrl.Path)
		if err != nil {
			log.Fatalln(err)
		}

		err = ioutil.WriteFile(fileDest, itemData, os.FileMode(0644))
		if err != nil {
			log.Fatalln(err)
		}

		s.Infof(`🍀 Downloaded '%s'`, fileDest)
	}
}

func (*downloader) downloadFile(url string) ([]byte, error) {
	f, err := xopen.Ropen(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	return ioutil.ReadAll(f)
}

func (*downloader) saveFile(fileDest string, data []byte, ) error {
	return ioutil.WriteFile(fileDest, data, os.FileMode(0644))
}

func (*downloader) createFolderIfNotExist(folderPath string) {
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		os.Mkdir(folderPath, os.ModePerm)
	}
}

func (*downloader) checkFileExist(dest string) bool {
	if _, err := os.Stat(dest); err != nil {
		return false
	}
	return true
}
