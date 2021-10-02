package erratum

import "errors"

func Use(opener ResourceOpener, input string) (err error) {
	res, err := opener()
	for err != nil {
		_, ok := err.(TransientError)
		if !ok {
			return err
		}
		res, err = opener()
	}

	defer res.Close()

	defer func() {
		if r := recover(); r != nil {
			v, ok := r.(FrobError)
			if ok {
				res.Defrob(v.defrobTag)
				err = v.inner
			} else {
				err = errors.New("meh")
			}
		}
	}()

	res.Frob(input)

	return err
}
