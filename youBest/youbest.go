package youBest

import "github.com/g-unit/gunit"

// Be - assert true
func Be(value bool) {
	if !value {
		gunit.FailTest()
		return
	}
	gunit.PassTest()
}

// NotBe asserts false
func NotBe(value bool) {
	if !value {
		gunit.PassTest()
		return
	}
	gunit.FailTest()
}

// BeInTheHouse asserts existence
func BeInTheHouse(value interface{}) {
	if value != nil {
		gunit.PassTest()
		return
	}
	gunit.FailTest()
}

// MoreThanThat
func BeMoreThanThis(expected, value int) {
	if expected >= value {
		gunit.FailTest()
		return
	}
	gunit.PassTest()
}

func BeLessThanThis(expected, value int) {
	if expected >= value {
		gunit.FailTest()
		return
	}
	gunit.PassTest()
}

func BeReal(value int) {}

// Freak implements panic
func Freak() {}

func NotFreak() {}
