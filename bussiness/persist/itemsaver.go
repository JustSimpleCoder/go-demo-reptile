package persist

import (
	"context"
	"github.com/olivere/elastic/v7"
	"log"
	"strings"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {

		itemCount := 0
		for {
			item := <-out
			itemCount++
			log.Printf("Item saver %d : %v", itemCount, item)
			save(item)
		}

	}()
	return out
}

func save(item interface{}) {

	client, err := elastic.Client{elastic.SetSniff(false)}
	if err != nil {
		panic(err)
	}
	response, err := client.Index().Index("reptile").Type("paper").BodyJson(item).Do(context.Background())

	if err != nil {

		panic(err)
	}

	log.Printf("%v", response)
}
