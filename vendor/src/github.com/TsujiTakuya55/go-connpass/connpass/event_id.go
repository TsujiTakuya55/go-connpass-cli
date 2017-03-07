package connpass

import (
	"fmt"
	"net/http"
	"errors"
)

type eventIdService service

func (s *eventIdService) Get(eventId string) (*Connpass, *http.Response, error) {
	var u string
	if eventId != "" {
		u = fmt.Sprintf("?event_id=%v", eventId)
	} else {
		return nil, nil, errors.New("please enter event_id")
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