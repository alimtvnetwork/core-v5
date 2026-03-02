package reqtype

func end(
	reqs []Request,
) any {
	if len(reqs) == 0 {
		return nil
	}

	return reqs[len(reqs)-1]
}
