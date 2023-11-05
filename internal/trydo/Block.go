package trydo

type Block struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}

func (it Block) Do() {
	if it.Finally != nil {
		defer it.Finally()
	}

	if it.Catch != nil {
		defer func() {
			r := recover()
			it.Catch(r)
		}()
	}

	it.Try()
}
