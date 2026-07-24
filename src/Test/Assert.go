

func AssertImpl(message string, success bool) func() {
	return func() {
		if !success {
			panic(message)
		}
	}
}

func CheckThrows(fn func(any) any) func() bool {
	return func() bool {
		var success bool
		func() {
			defer func() {
				if r := recover(); r != nil {
					success = true
				}
			}()
			fn(nil)
			success = false
		}()
		return success
	}
}
