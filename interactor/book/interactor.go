package interactor

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

type Interactor struct{}

func New() contract.BookInteractor {
	return Interactor{}
}

func (i Interactor) SearchBook(ctx context.Context, req dto.SearchBookRequest) (dto.SearchBookResponse, error) {
	url := fmt.Sprintf("https://search.fidibo.com%s", getUrlParam(req.Keyword))

	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		return dto.SearchBookResponse{}, err
	}
	defer func() {
		err = resp.Body.Close()
		log.Println("failed to close the response body - err: ", err.Error())
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
	return result, nil
}

func getUrlParam(q string) string {
	if q == "" {
		return ""
	}

	return fmt.Sprintf("?q=%s", q)
}
