package main

import (
	. "github.com/ricallinson/simplebdd"
	"testing"
)

func TestServer(t *testing.T) {

	Describe("a", func() {
		It("should return [true]", func() {
			AssertEqual(true, false)
		})
	})

	Report(t)
}
