package job

import (
	"hello_gin/db"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"hello_gin/entity"
	"fmt"
	"time"
)

type ContentfulJob struct {
	// filtered
}

type Item struct {
	Total int `json:"total"`
	Items []struct {
		Metadata struct {
			Tags []interface{} `json:"tags"`
		} `json:"metadata"`
		Sys struct {
			Space struct {
				Sys struct {
					Type     string `json:"type"`
					LinkType string `json:"linkType"`
					ID       string `json:"id"`
				} `json:"sys"`
			} `json:"space"`
			ID          string    `json:"id"`
			Type        string    `json:"type"`
			CreatedAt   time.Time `json:"createdAt"`
			UpdatedAt   time.Time `json:"updatedAt"`
			Environment struct {
				Sys struct {
					ID       string `json:"id"`
					Type     string `json:"type"`
					LinkType string `json:"linkType"`
				} `json:"sys"`
			} `json:"environment"`
			Revision    int `json:"revision"`
			ContentType struct {
				Sys struct {
					Type     string `json:"type"`
					LinkType string `json:"linkType"`
					ID       string `json:"id"`
				} `json:"sys"`
			} `json:"contentType"`
			Locale string `json:"locale"`
		} `json:"sys"`
		Fields struct {
			Title string `json:"title"`
			Body  struct {
				Data struct {
				} `json:"data"`
				Content []struct {
					Data struct {
					} `json:"data"`
					Content []struct {
						Data struct {
						} `json:"data"`
						Marks    []interface{} `json:"marks"`
						Value    string        `json:"value"`
						NodeType string        `json:"nodeType"`
					} `json:"content"`
					NodeType string `json:"nodeType"`
				} `json:"content"`
				NodeType string `json:"nodeType"`
			} 
		} `json:"fields"`
	} `json:"items"`
}

func (e ContentfulJob) Run() {
	fmt.Println("[Start] Run ContentfulJob!")
	resp, err := http.Get("https://cdn.contentful.com")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
			log.Fatal(err)
	}

	var data Item

	if err := json.Unmarshal(body, &data); err != nil {
			log.Fatal(err)
	}

	a := db.GetDB()

	for i := 0; i < data.Total; i++ {
		article := entity.Article{Title: data.Items[i].Fields.Title}
		// article := entity.Article{Title: data.Items[i].Fields.Title, Body: data.Items[i].Fields.Body}
		a.Create(&article)
	}
	fmt.Println("[End] Run ContentfulJob!")
}
