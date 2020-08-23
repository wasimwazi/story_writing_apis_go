package story

import (
	"errors"
	"storyapi/utils"
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
