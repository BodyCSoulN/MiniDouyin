package utils

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io"
	"mime/multipart"
	"os"
)

type UserClaims struct {
	ID int64
	jwt.RegisteredClaims
}

var myKey = ""

func ParseToken(tokenString string) (int64, error) {

	userClaim := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		return -1, err
	}

	if !claims.Valid {
		return -1, fmt.Errorf("parse Token Error:%v", err)
	}
	return userClaim.ID, nil
}

func GenerateUUid() string {
	return uuid.NewV4().String()
}

var videoFileExt = []string{".mp4", ".flv"}

func IsAllowedSuffix(fileSuffix string) bool {
	for _, AllowExt := range videoFileExt {
		if fileSuffix == AllowExt {
			return true
		}
	}
	return false
}

// ReadFrameAsJpeg
// generate cover image from the first frame of video by ffmpeg
func ReadFrameAsJpeg(inFileName, coverPath string, frameNum int) error {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(inFileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()

	if err != nil {
		return errors.New("could not generate cover, FFmpeg Error: " + err.Error())
	}
	img, err := imaging.Decode(buf)
	if err != nil {
		return errors.New("could not generate cover, Decode Error: " + err.Error())
	}
	err = imaging.Save(img, coverPath)
	if err != nil {
		return errors.New("could not generate cover, Save Error: " + err.Error())
	}
	return nil
}

func SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	//创建 dst 文件
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	// 拷贝文件
	_, err = io.Copy(out, src)
	return err
}
