package module

import (
	"encoding/json"
	"fmt"
	"github.com/brentp/xopen"
	"io/ioutil"
	"log"
	"os"
)

const (
	DestFolder = "/Users/lynam/dev/go-builder/deploy"
)

type ListFile struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type downloader struct {

}

func NewDownloader() *downloader {
	return &downloader{}
}

func (s *downloader) Run() {
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

	s.createFolderIfNotExist(DestFolder)

	for _, fUrl := range list {
		itemData, err := s.downloadFile(fUrl.Path)
		if err != nil {
			log.Fatalln(err)
		}
		fileDest := fmt.Sprintf(`%s/%s`, DestFolder, fUrl.Name)

		err = ioutil.WriteFile(fileDest, itemData, os.FileMode(0644))
		if err != nil {
			log.Fatalln(err)
		}
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
