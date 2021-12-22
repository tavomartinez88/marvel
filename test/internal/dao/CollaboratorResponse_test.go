package dao

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tavomartinez88/marvel/internal/dao"
)

var _ = Describe("Collaborator Response Dao tests", func() {
	var (
		response    dao.CollaboratorsDao
	)

	Context("Collaborator response dao tests", func() {
			It("test get id", func() {
				response = dao.CollaboratorsDao{
					ColID: "iron man",
				}
				field, val := response.ID()
				Expect(field).Should(Equal("colid"))
				Expect(val).Should(Equal("iron man"))
			})
	})
})
