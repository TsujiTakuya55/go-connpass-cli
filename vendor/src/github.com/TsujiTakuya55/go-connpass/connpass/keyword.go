package connpass

import (
	"fmt"
	"net/http"
	"errors"
)

type keywordService service

func (s *keywordService) Get(keyword string) (*Connpass, *http.Response, error) {
	var u string
	if keyword != "" {
		u = fmt.Sprintf("?keyword=%v", keyword)
	} else {
		return nil, nil, errors.New("please enter keyword")
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	connpass := new(*Connpass)

	resp, err := s.client.Do(req, &connpass)
	if err != nil {
		return nil, resp, err
	}
	return *connpass, resp, err
}
