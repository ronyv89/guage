package tools_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "guage/internal/tools"
)

var _ = Describe("Shell", func() {
	Describe("ExecCommand", func() {
		Context("For a valid shell command", func() {
			It("Executes a shell command", func() {
				_, err := ExecCommand(".", "go", "help")
				Expect(err).To(BeNil())
			})
		})
		Context("For an invalid shell command", func() {
			It("Returns error", func() {
				_, err := ExecCommand(".", "invalidcommand", "help")
				Expect(err).NotTo(BeNil())
			})
		})
	})
})
