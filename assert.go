package main

func AssertTrue(value bool) {
	if !value {
		FailTest()
		return
	}
	PassTest()
}

func AssertFalse(value bool) {
	if !value {
		PassTest()
		return
	}
	FailTest()
}
