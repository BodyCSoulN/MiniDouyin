package service

import (
	"errors"
	"mime/multipart"
	"os/exec"
	"path"

	"github.com/MiniDouyin/model"
	"github.com/MiniDouyin/storage"
	"github.com/MiniDouyin/utils"
)

func GetPublishList(userID int64) ([]storage.Video, error) {
	tx := model.GetVideoList(userID)
	var videoList []storage.Video
	err := tx.Find(&videoList).Error
	return videoList, err
}

func PublishVideo(file *multipart.FileHeader, title string, userID int64) error {
	// static 文件夹的目录
	dstPath := "/home/ubuntu/go/src/MiniDouyin/static/"
	fileName := file.Filename
	fileSuffix := path.Ext(fileName)
	if !utils.IsAllowedSuffix(fileSuffix) {
		return errors.New("unsupported Video Format")
	}

	uuid := utils.GenerateUUid()

	videoPath := dstPath + "video/" + uuid + fileSuffix
	if err := utils.SaveUploadedFile(file, videoPath); err != nil {
		return errors.New("save Upload Video Error: " + err.Error())
	}

	coverPath := dstPath + "cover/" + uuid + ".jpeg"
	// ffmpeg -i ./bear.mp4 -ss 1 -frames:v 1 -f image2 ./test.jpg
	// ffmpeg目录
	// ffmpegPath := "D:/ChromeDownload/ffmpeg-2022-05-26-git-0dcbe1c1aa-full_build/bin/"
	cmd := exec.Command("ffmpeg", "-i", videoPath, "-ss", "1", "-frames:v", "1", "-f", "image2", coverPath)
	err := cmd.Run()
	//err := utils.ReadFrameAsJpeg(videoPath, coverPath, 1)

	if err != nil {
		return err
	}
	// 服务器地址及端口号
	coverUrl := "1.15.235.227:8080/static/cover/" + uuid + ".jpeg"
	videoUrl := "1.15.235.227:8080/static/video/" + uuid + fileSuffix

	return model.PublishVideo(videoUrl, coverUrl, title, userID)
}
