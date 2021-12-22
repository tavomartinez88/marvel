package handlers

import (
	"errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"github.com/tavomartinez88/marvel/pkg/handlers"
	"github.com/tavomartinez88/marvel/pkg/models"
)

var _ = Describe("Handler Collaborator Tests", func() {
	var (
		handler handlers.ICollaboratorHandler
		serviceMock ServiceCollaboratorMock
	)

	Context("Handler collaborator tests", func() {
		It("Get not-empty list", func() {
			serviceMock.On("GetCollaborators", mock.Anything).Return(GetMockCollaborator(), nil)
			handler = &handlers.CollaboratorHandler{
				Service: &serviceMock,
			}

			response, err := handler.GetCollaborators("iron man")

			Expect(err).To(BeNil())
			Expect(response).NotTo(BeNil())
			Expect(response.LastSync).To(Equal("01-01-2021"))
			Expect(len(response.Colorists)>0).Should(BeTrue())
			Expect(len(response.Editors)>0).Should(BeTrue())
			Expect(len(response.Writers)>0).Should(BeTrue())
		})

		It("Get empty list", func() {
			serviceMock.On("GetCollaborators", mock.Anything).Return(models.CollaboratorsResponse{}, nil)
			handler = &handlers.CollaboratorHandler{
				Service: &serviceMock,
			}

			response, err := handler.GetCollaborators("iron man")

			Expect(err).To(BeNil())
			Expect(response).NotTo(BeNil())
			Expect(len(response.Colorists)==0).Should(BeTrue())
			Expect(len(response.Editors)==0).Should(BeTrue())
			Expect(len(response.Writers)==0).Should(BeTrue())
		})

		It("Get with error", func() {
			serviceMock.On("GetCollaborators", mock.Anything).Return(models.CollaboratorsResponse{}, errors.New("error"))
			handler = &handlers.CollaboratorHandler{
				Service: &serviceMock,
			}

			response, err := handler.GetCollaborators("iron man")

			Expect(err).NotTo(BeNil())
			Expect(response).NotTo(BeNil())
		})

		It("Get NewHandler", func() {
			handler = handlers.NewCollaboratorHandler()
			Expect(handler).NotTo(BeNil())
		})
	})

	BeforeEach(func() {
		serviceMock = ServiceCollaboratorMock{}
	})

})

type ServiceCollaboratorMock struct {
	mock.Mock
}

func (scm *ServiceCollaboratorMock) GetCollaborators(name string) (models.CollaboratorsResponse, error) {
	args := scm.Called()
	return args.Get(0).(models.CollaboratorsResponse), args.Error(1)
}

func GetMockCollaborator() models.CollaboratorsResponse {
	return models.CollaboratorsResponse{
		LastSync: "01-01-2021",
		Editors: []string{"a","b","c"},
		Writers: []string{"a","b","c"},
		Colorists: []string{"a","b","c"},
	}
}