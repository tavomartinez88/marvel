package dao

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tavomartinez88/marvel/internal/dao"
)

var _ = Describe("Character Response Dao tests", func() {
	var (
		response    dao.CharacterDao
	)

	Context("Character response dao tests", func() {
			It("test get id", func() {
				response = dao.CharacterDao{
					CharID: "iron man",
				}
				field, val := response.ID()
				Expect(field).Should(Equal("charid"))
				Expect(val).Should(Equal("iron man"))
			})
	})
})
