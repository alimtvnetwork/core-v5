package reqtype

import "gitlab.com/auk-go/core/constants"

func start(
	reqs []Request,
) *Request {
	if len(reqs) == 0 {
		return nil
	}

	r := reqs[constants.Zero]
	return &r
}
