package images

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"sort"
	"unicode"

	"github.com/gmelum/msa_tg_images/utils/bot"
)

func GetUserImage(user, token string) ([]byte, error) {

	var url string = ""

	if isNumber(user) {
		uri, err := GetPhotoOnID(user, token)
		if err != nil {
			return nil, err
		}
		url = uri
	} else {
		uri, err := GetPhotoOnShortName(user, token)
		if err != nil {
			return nil, err
		}
		url = uri
	}

	resp, err := http.Get(url)
	if err != nil {
		println(err)
		return nil, err
	}
	defer resp.Body.Close()

	file, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return file, nil

}

func GetChatImage(user, token string) ([]byte, error) {

	if !isNumber(user) && user[0:1] != "@" {
		user = "@" + user
	}

	photos, err := bot.GetChat(user, token)

	if err != nil {
		return nil, err
	}

	if !photos.Ok || photos.Result == nil || photos.Result.Photo.SmallFileID == "" {
		return nil, errors.New("image is null")
	}

	photo := photos.Result.Photo.SmallFileID

	file, err := bot.GetFile(photo, token)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.telegram.org/file/bot%v/%v", token, file.Result.FilePath)

	resp, err := http.Get(url)
	if err != nil {
		println(err)
		return nil, err
	}
	defer resp.Body.Close()

	image, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return image, nil

}

func GetPhotoOnID(user, token string) (string, error) {

	photos, err := bot.GetUserProfilePhotos(user, token)
	if err != nil {
		return "", err
	}
	if !photos.Ok || photos.Result == nil || photos.Result.TotalCount == 0 {
		return "", errors.New("image is null")
	}

	photo := photos.Result.Photos[0]
	sort.SliceStable(photo, func(i, j int) bool {
		return photo[i].Width < photo[j].Width
	})

	if photo[0].FileID == "" {
		return "", err
	}

	file, err := bot.GetFile(photo[0].FileID, token)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("https://api.telegram.org/file/bot%v/%v", token, file.Result.FilePath), nil
}

func GetPhotoOnShortName(user, token string) (string, error) {

	if len(user) > 30 {
		return "", errors.New("very long user")
	}

	resp, err := http.Get("https://t.me/" + user)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	html, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	re := regexp.MustCompile(`<img.*?src="(.+?)">`)

	match := re.FindStringSubmatch(string(html))
	if len(match) > 0 {
		return fmt.Sprintf(match[1]), nil
	}

	return "", errors.New("image is not defined")

}

func isNumber(s string) bool {
	if s[0:1] == "-" {
		s = s[1:]
	}
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}
