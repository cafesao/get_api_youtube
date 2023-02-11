package getYoutube

import (
	"encoding/json"
	"fmt"
	Struct "getYoutube/Struct"
	"io"
	"log"
	"net/http"
	"os"
)

func GetApi(id string) Struct.SApiYoutube {
	url := fmt.Sprintf("https://www.googleapis.com/youtube/v3/videos?id=%v&key=%v&part=snippet,statistics,contentDetails&fields=items(snippet(title,channelTitle),statistics,contentDetails(duration))", id, GetEnv("API_KEY"))

	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Struct.SApiYoutube

	json.Unmarshal(responseData, &responseObject)

	if len(responseObject.Items) == 0 {
		log.Fatal("Not found video")
	}

	return responseObject
}
