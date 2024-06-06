package bot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/NicoNex/echotron/v3"
)

func GetFile(path, token string) (*echotron.APIResponseFile, error) {
	endpoint := fmt.Sprintf(
		"%v%v/getFile?file_id=%v",
		"https://api.telegram.org/bot",
		token,
		path,
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

	var data echotron.APIResponseFile
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
