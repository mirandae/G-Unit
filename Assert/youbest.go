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

func NonNil(value interface{}) {
	if value != nil {
		gunit.PassTest()
		return
	}
	gunit.FailTest()
}
