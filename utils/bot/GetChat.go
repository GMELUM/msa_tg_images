package bot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/NicoNex/echotron/v3"
)

type APIChatInfo struct {
	Result *echotron.ChatFullInfo `json:"result,omitempty"`
	echotron.APIResponseBase
}

func GetChat(chat, token string) (*APIChatInfo, error) {

	endpoint := fmt.Sprintf(
		"%v%v/getChat?chat_id=%v",
		"https://api.telegram.org/bot",
		token,
		chat,
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

	var data APIChatInfo
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil

}
