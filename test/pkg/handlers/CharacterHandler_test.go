package handlers

import (
	"errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"github.com/tavomartinez88/marvel/pkg/handlers"
	"github.com/tavomartinez88/marvel/pkg/models"
)

var _ = Describe("Handler Character Tests", func() {
	var (
		handler handlers.ICharacterHandler
		serviceMock ServiceCharacterMock
	)

	Context("Handler character tests", func() {
		It("Get not-empty list", func() {
			serviceMock.On("GetCharacters", mock.Anything).Return(GetMockCharacter(), nil)
			handler = &handlers.CharacterHandler{
				Service: &serviceMock,
			}

			response, err := handler.GetCharacters("iron man")

			Expect(err).To(BeNil())
			Expect(response).NotTo(BeNil())
			Expect(response.LastSync).To(Equal("01-01-2021"))
			Expect(len(response.Characters)>0).Should(BeTrue())
		})

		It("Get empty list", func() {
			serviceMock.On("GetCharacters", mock.Anything).Return(models.CharacterResponse{}, nil)
			handler = &handlers.CharacterHandler{
				Service: &serviceMock,
			}

			response, err := handler.GetCharacters("iron man")

			Expect(err).To(BeNil())
			Expect(response).NotTo(BeNil())
			Expect(len(response.Characters)==0).Should(BeTrue())
		})

		It("Get with error", func() {
			serviceMock.On("GetCharacters", mock.Anything).Return(models.CharacterResponse{}, errors.New("error"))
			handler = &handlers.CharacterHandler{
				Service: &serviceMock,
			}

			response, err := handler.GetCharacters("iron man")

			Expect(err).NotTo(BeNil())
			Expect(response).NotTo(BeNil())
		})

		It("Get NewHandler", func() {
			handler = handlers.NewCharacterHandler()
			Expect(handler).NotTo(BeNil())
		})
	})

	BeforeEach(func() {
		serviceMock = ServiceCharacterMock{}
	})

})

type ServiceCharacterMock struct {
	mock.Mock
}

func (scm *ServiceCharacterMock) GetCharacters(name string) (models.CharacterResponse, error) {
	args := scm.Called()
	return args.Get(0).(models.CharacterResponse), args.Error(1)
}

func GetMockCharacter() models.CharacterResponse {
	characters := []models.Character{
		{
			Character: "iron man",
			Comics:    []string{"im1", "im2", "im3"},
		},
	}

	return models.CharacterResponse{
		LastSync: "01-01-2021",
		Characters: characters,
	}
}