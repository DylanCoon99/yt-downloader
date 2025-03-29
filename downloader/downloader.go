package downloader

import (
	"log"
	"context"
	"github.com/lrstanley/go-ytdlp"
)


func Download(url string) string {

	ytdlp.MustInstall(context.TODO(), nil)


	dl := ytdlp.New().
		FormatSort("res,ext:mp4:m4a").
		RecodeVideo("mp4").
		Output("%(extractor)s - %(title)s.%(ext)s")

	_, err := dl.Run(context.TODO(), url)
	if err != nil {
		panic(err)
	}

	msg := "Successfully downloaded your video!"

	log.Println(msg)
	return msg
}