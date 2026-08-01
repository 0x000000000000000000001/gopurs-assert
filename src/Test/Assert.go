package Test_Assert

func AssertImpl(message string, success bool, _ interface{}) interface{} {
	if !success {
		panic(message)
	}
	return nil
}

func CheckThrows(fn func(interface{}) interface{}, _ interface{}) bool {
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
