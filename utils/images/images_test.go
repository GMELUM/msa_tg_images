package images

import (
	"os"
	"testing"
)

func TestMap(t *testing.T) {

	token := "***" // tg bot token

	user := "xuserz"

	file, err := GetUserImage(user, token)
	if err != nil {
		return
	}

	out, err := os.Create("photo.jpg")
	if err != nil {
		println(err)
		return
	}
	defer out.Close()

	_, err = out.Write(file)
	if err != nil {
		println(err)
		return
	}

}
