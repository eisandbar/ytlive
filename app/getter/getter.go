package getter

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/eisandbar/ytlive/app/internal"
	"github.com/eisandbar/ytlive/app/store"
)

func NewGetter(ds store.Store) {
	streamChan := make(chan store.Stream, 20)
	getter := Getter{}
	getter.Store = ds

	searchDB := func() {
		log.Println("Searching DB for streams without a category")
		getter.SearchDB(streamChan)
	}

	go internal.Ticker(time.Second*50, searchDB)

	getCategory := func() {
		stream := <-streamChan
		log.Printf("Getting category for stream: %s", stream.Id)
		getter.GetCategory(stream)
	}

	go internal.Ticker(time.Second*5, getCategory)
}

type Getter struct {
	Store store.Store
}

// Searches db for streams without a category
func (getter *Getter) SearchDB(streamChan chan<- store.Stream) {
	streams := getter.Store.List(store.WithMaxResults(100))
	log.Printf("Found %v streams without a category", len(streams))
	for _, stream := range streams {

		// Filter out non gaming streams
		stream.Category = getCategoryFromId(stream)

		if stream.Category == "" {
			stream.Category = store.Pending // Change status so next SearchDB doesn't add again
			streamChan <- stream            // Send to queue
		}

		getter.Store.Update(stream) // Updates with non-gaming category or pending
	}
}

// Sends a get request to the stream to check for a category
func (getter *Getter) GetCategory(stream store.Stream) {
	url := genURL(stream)
	category := getCategory(url)
	getter.Store.SaveCategory(category)
	stream.Category = category.Category
	getter.Store.Update(stream)
}

func getCategory(url string) store.Category {
	text := getRequest(url)
	data := matchEnds("{\"richMetadataRenderer\":", "},{\"r", "", "", text)
	// data := matchEnds("playabilityStatus\":", ",\"streamingData", "", "", text)
	var rmr RichMetadataRenderer
	if len(data) == 0 {
		log.Println("No metadata renderer found when looking for category")
		return store.Category{Category: store.NoCategory}
	}
	err := json.Unmarshal(data[0], &rmr)
	if err != nil || rmr.Title.SimpleText == "" {
		return store.Category{Category: store.NoCategory}
	}
	return store.Category{Category: rmr.Title.SimpleText, Url: getThumbnail(rmr.Thumbnail)}
}

func genURL(stream store.Stream) string {
	return fmt.Sprintf("https://www.youtube.com/watch?v=%s", stream.Id)
}

func getRequest(url string) []byte {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Failed to make request, %s", err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Bad response, %s", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body, %s", err)
	}
	return body
}

// Function return the thumbnail with the best resolution
func getThumbnail(thumbnail Thumbnail) string {
	if len(thumbnail.Thumbnails) == 0 {
		return ""
	}

	url := thumbnail.Thumbnails[0].Url
	width := thumbnail.Thumbnails[0].Width

	for _, t := range thumbnail.Thumbnails {
		if width < t.Width {
			url = t.Url
		}
	}
	return url
}

func getCategoryFromId(stream store.Stream) string {
	switch stream.CategoryId {
	case "1":
		return "Film & Animation"
	case "2":
		return "Autos & Vehicles"
	case "10":
		return "Music"
	case "15":
		return "Pets & Animals"
	case "17":
		return "Sports"
	case "18":
		return "Short Movies"
	case "19":
		return "Travel & Events"
	// For gaming we want the exact game streamed
	case "20":
		return stream.Category
	case "21":
		return "Videoblogging"
	case "22":
		return "People & Blogs"
	case "23":
		return "Comedy"
	case "24":
		return "Entertainment"
	case "25":
		return "News & Politics"
	case "26":
		return "Howto & Style"
	case "27":
		return "Education"
	case "28":
		return "Science & Technology"
	case "29":
		return "Nonprofits & Activism"
	case "30":
		return "Movies"
	case "31":
		return "Anime/Animation"
	case "32":
		return "Action/Adventure"
	case "33":
		return "Classics"
	case "34":
		return "Comedy"
	case "35":
		return "Documentary"
	case "36":
		return "Drama"
	case "37":
		return "Family"
	case "38":
		return "Foreign"
	case "39":
		return "Horror"
	case "40":
		return "Sci-Fi/Fantasy"
	case "41":
		return "Thriller"
	case "42":
		return "Shorts"
	case "43":
		return "Shows"
	case "44":
		return "Trailers"
	default:
		return ""
	}
}
