package story

import (
	"errors"
	"net/http"
	"storyapi/utils"
	"strconv"
	"strings"
)

//ValidatePost : Function to validate the post request
func ValidatePost(request AddStoryRequest) error {
	s := strings.Split(request.Word, " ")
	if len(s) > 1 {
		return errors.New(utils.MultipleWordError)
	}
	return nil
}

func validateGetStories(request *GetStoryRequest, r *http.Request) (*GetStoryRequest, error) {
	limit := r.URL.Query().Get("limit")
	if limit != "" {
		Limit, err := strconv.Atoi(limit)
		if err != nil {
			return nil, errors.New("invalid query parameter, limit")
		}
		request.Limit = Limit
	}
	offset := r.URL.Query().Get("offset")
	if offset != "" {
		storyOffset, err := strconv.Atoi(offset)
		if err != nil {
			return nil, errors.New("invalid query parameter, offset")
		}
		request.Offset = storyOffset
	}
	sort := r.URL.Query().Get("sort")
	if sort != "" {
		if sort != utils.CREATEDAT && sort != utils.UPDATEDAT && sort != utils.TITLE {
			return nil, errors.New("invalid query parameter, sort")
		}
		request.Sort = sort
	}
	order := r.URL.Query().Get("order")
	if order != "" {
		if order != utils.ASC && order != utils.DESC {
			return nil, errors.New("invalid query parameter, order")
		}
		request.Sort = order
	}
	return request, nil
}
