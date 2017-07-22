package Assert

import "github.com/mirandae/g-unit/gunit"

func True(value bool) {
	if !value {
		gunit.FailTest()
		return
	}
	gunit.PassTest()
}

func False(value bool) {
	if !value {
		gunit.PassTest()
		return
	}
	gunit.FailTest()
}

func Not(value interface{}) {
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
