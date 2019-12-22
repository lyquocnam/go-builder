package module

import "github.com/vektra/gitreader"

type downloader struct {

}

func NewDownloader() *downloader {
	return &downloader{}
}

func (*downloader) Run() {
	repo, err := gitreader.OpenRepo("/path/to/repo")
	if err != nil {
		panic(err)
	}

	blob, err := repo.CatFile("HEAD", "path/to/file")
	if err != nil {
		panic(err)
	}

	// WARNING: use Blob as an io.Reader instead if you can!
	bytes, err := blob.Bytes()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", bytes)

	repo.Close()
}
