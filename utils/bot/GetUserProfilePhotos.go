package bot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/NicoNex/echotron/v3"
)

func GetUserProfilePhotos(user, token string) (*echotron.APIResponseUserProfile, error) {

	endpoint := fmt.Sprintf(
		"%v%v/getUserProfilePhotos?user_id=%v&limit=1&offset=0",
		"https://api.telegram.org/bot",
		token,
		user,
	)

	response, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var data echotron.APIResponseUserProfile
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil

}
