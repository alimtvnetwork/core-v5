package reqtype

func end(
	reqs []Request,
) *Request {
	if len(reqs) == 0 {
		return nil
	}

	r := reqs[len(reqs)-1]
	return &r
}
