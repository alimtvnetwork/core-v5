package reqtype

import "github.com/alimtvnetwork/core/constants"

func start(
	reqs []Request,
) *Request {
	if len(reqs) == 0 {
		return nil
	}

	r := reqs[constants.Zero]
	return &r
}
