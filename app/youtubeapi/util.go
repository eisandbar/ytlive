package youtubeapi

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/eisandbar/ytlive/app/internal"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func NewService() *youtube.Service {
	ctx := context.Background()

	// Creating new youtube service
	api_key, err := internal.LoadToken("key.txt")
	if err != nil {
		log.Fatalf("Unable to create YouTube service: %v", err)
	}
	service, err := youtube.NewService(ctx, option.WithAPIKey(api_key))
	if err != nil {
		log.Fatalf("Unable to create YouTube service: %v", err)
	}
	return service
}

func SaveListResponse(file string, listResponse any) {
	data, err := json.Marshal(listResponse)
	if err != nil {
		log.Fatalf("%+v\nerror: %s", data, err)
	}
	err = os.WriteFile(file, data, 0644)
	if err != nil {
		fmt.Println("Failed to write file")
	}
}

func LoadListResponse(file string, listResponse any) {

	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatalln("Failed to read file")
	}

	err = json.Unmarshal(data, listResponse)
	if err != nil {
		log.Fatalf("Error unmarshalling listResponse, %v", err)
	}

}

func GetToken(ctx context.Context, config *oauth2.Config) *oauth2.Token {
	ch := make(chan string)
	randState := fmt.Sprintf("st%d", time.Now().UnixNano())

	ts := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.Path == "/favicon.ico" {
			http.Error(rw, "", 404)
			return
		}
		if req.FormValue("state") != randState {
			log.Printf("State doesn't match: req = %#v", req)
			http.Error(rw, "", 500)
			return
		}
		if code := req.FormValue("code"); code != "" {
			fmt.Fprintf(rw, "<h1>Success</h1>Authorized.")
			rw.(http.Flusher).Flush()
			ch <- code
			return
		}
		log.Printf("no code")
		http.Error(rw, "", 500)
	}))
	defer ts.Close()

	config.RedirectURL = ts.URL
	authURL := config.AuthCodeURL(randState)
	fmt.Printf("Visit the URL for the auth dialog: %v", authURL)
	code := <-ch

	token, err := config.Exchange(ctx, code)
	if err != nil {
		log.Fatalf("Failed to generate token")
	}
	return token
}
