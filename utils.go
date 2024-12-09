package shikimori

import (
	"errors"
	"io"
	"net/http"
)

func getErrorFromBadResponse(resp *http.Response) error {
	switch resp.StatusCode {
	case 400:
		return ErrBadRequest
	case 401:
		return ErrUnauthorized
	case 403:
		return ErrForbidden
	case 404:
		return ErrNotFound
	case 422:
		return ErrUnprocessableEntity
	case 429:
		return ErrTooManyRequests
	case 500:
		return ErrInternalServer
	default:
		data, err := io.ReadAll(resp.Body)
		if err != nil && len(data) != 0 {
			resp.Body.Close()
			return errors.New("bad status " + resp.Request.Method + " " + resp.Request.URL.String() + " -> " + resp.Status + "\n" + string(data))
		} else {
			return errors.New("bad status " + resp.Request.Method + " " + resp.Request.URL.String() + " -> " + resp.Status)
		}
	}
}
