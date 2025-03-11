package nanoid

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("nanoid", func() {
	It("should generate a nanoid", func() {
		nanoid, err := New(21, ALPHABET_NUMBERS+ALPHABET_LOWERCASE+ALPHABET_UPPERCASE)
		Expect(err).NotTo(HaveOccurred())
		Expect(len(nanoid)).To(Equal(21))
	})

	It("should not generate a nanoid with 0 size", func() {
		nanoid, err := New(0, ALPHABET_NUMBERS+ALPHABET_LOWERCASE+ALPHABET_UPPERCASE)
		Expect(err).To(HaveOccurred())
		Expect(nanoid).To(BeEmpty())
	})

	It("should not generate a nanoid with empty alphabet", func() {
		nanoid, err := New(21, "")
		Expect(err).To(HaveOccurred())
		Expect(nanoid).To(BeEmpty())
	})

	It("should generate a nanoid without alike characters", func() {
		nanoid, err := NewWithoutAlike(21, ALPHABET_NUMBERS+ALPHABET_LOWERCASE+ALPHABET_UPPERCASE)
		Expect(err).NotTo(HaveOccurred())
		Expect(len(nanoid)).To(Equal(21))
		Expect(nanoid).NotTo(ContainSubstring("1lI0Oouv5Ss"))
	})

	It("should generate a safe nanoid", func() {
		nanoid := NewSafe()
		Expect(len(nanoid)).To(Equal(21))
	})
})
