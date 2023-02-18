package book

import (
	"context"
	"encoding/json"
	"fidibo_interview/contract"
	"fidibo_interview/dto"
	"fidibo_interview/entity"
	"fmt"
	"log"
	"net/http"
)

type Interactor struct {
	bookCache contract.BookCache
}

func New(bc contract.BookCache) contract.BookInteractor {
	return Interactor{
		bookCache: bc,
	}
}

func (i Interactor) SearchBook(ctx context.Context, req dto.SearchBookRequest) (dto.SearchBookResponse, error) {
	// check the cache
	cachedBooks, err := i.bookCache.Get(req.Keyword)
	if err == nil && cachedBooks != nil {
		// cache hit
		log.Println("cache hit")
		return dto.SearchBookResponse{
			Result: *cachedBooks}, nil
	}

	// cache missed
	url := fmt.Sprintf("https://search.fidibo.com%s", getUrlParam(req.Keyword))

	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		return dto.SearchBookResponse{}, err
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Println("failed to close the response body - err: ", err.Error())
		}
	}()

	fidibo := new(entity.FidiboBook)
	err = json.NewDecoder(resp.Body).Decode(fidibo)
	if err != nil {
		return dto.SearchBookResponse{}, err
	}
	log.Println("decoded successfully")

	result := dto.SearchBookResponse{
		Result: make([]entity.Book, 0),
	}
	for _, n := range fidibo.AllBooks.AllHits.AllSources {
		result.Result = append(result.Result, n.Bk)
	}

	defer func() {
		// cache the result in redis
		err = i.bookCache.Set(req.Keyword, &result.Result)
		if err != nil {
			log.Println("failed to cache the result - err: ", err.Error())
		}
	}()

	return result, nil
}

func getUrlParam(q string) string {
	if q == "" {
		return ""
	}

	return fmt.Sprintf("?q=%s", q)
}
