package services

import (
	"errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"github.com/tavomartinez88/marvel/internal/dao"
	errorClient "github.com/tavomartinez88/marvel/internal/error"
	"github.com/tavomartinez88/marvel/internal/services"
	"github.com/tavomartinez88/marvel/pkg/models"
)

var _ = Describe("Character Service tests", func() {
	var (
		service services.ICharacterService
		clientMock ClientMock
		daoMock DaoMock
	)

	Context("Character service tests", func() {
		It("Invalid parameter input", func() {
			service = &services.CharacterService{
				Client: &clientMock,
				Dao: &daoMock,
			}

			errExpect := errorClient.ClientError{
				HttpStatus: 400,
				Message: "Name is invalid",
			}

			_, err := service.GetCharacters("1ronman")
			Expect(err).NotTo(BeNil())
			Expect(errExpect).Should(Equal(err))
		})

		It("Get Characters there is a problem with db", func() {

			daoMock.On("GetCharacters", mock.Anything).Return(dao.CharacterDao{}, errors.New("error db"))
			service = &services.CharacterService{
				Client: &clientMock,
				Dao: &daoMock,
			}

			errExpect := errorClient.ClientError{
				HttpStatus: 400,
				Message: "error db",
			}

			_, err := service.GetCharacters("ironman")
			Expect(err).NotTo(BeNil())
			Expect(errExpect).Should(Equal(err))
		})

		It("Get Characters and there isn't record", func() {
			var n int64 = 0
			daoMock.On("GetCharacters", mock.Anything).Return(dao.CharacterDao{}, errors.New("record not found"))
			clientMock.On("GetHeroId", mock.Anything).Return(n, errors.New("error to save"))

			service = &services.CharacterService{
				Client: &clientMock,
				Dao: &daoMock,
			}

			_, err := service.GetCharacters("ironman")
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal("error to save"))
		})

		It("Get Characters and get error on GetCaractersByHeroId", func() {
			var n int64 = 123
			daoMock.On("GetCharacters", mock.Anything).Return(dao.CharacterDao{}, errors.New("record not found"))
			clientMock.On("GetHeroId", mock.Anything).Return(n, nil)
			clientMock.On("GetCaractersByHeroId", mock.Anything).Return([]models.Character{}, errors.New("error client"))

			service = &services.CharacterService{
				Client: &clientMock,
				Dao: &daoMock,
			}

			_, err := service.GetCharacters("ironman")
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal("error client"))
		})

		It("Get Characters and get error on save in db", func() {
			var n int64 = 123
			daoMock.On("GetCharacters", mock.Anything).Return(dao.CharacterDao{}, errors.New("record not found"))
			daoMock.On("SaveOrUpdateCharacters", mock.Anything).Return(errors.New("error to save"))
			clientMock.On("GetHeroId", mock.Anything).Return(n, nil)
			clientMock.On("GetCaractersByHeroId", mock.Anything).Return([]models.Character{
				{
					Character: "Iron man",
					Comics: []string{"Avengers"},
				},
			}, nil)

			service = &services.CharacterService{
				Client: &clientMock,
				Dao: &daoMock,
			}

			errExpect := errorClient.ClientError{
				HttpStatus: 500,
				Message: "error to save",
			}

			_, err := service.GetCharacters("ironman")
			Expect(err).NotTo(BeNil())
			Expect(err).To(Equal(errExpect))
		})

		It("Get Characters and fail parse date", func() {
			var n int64 = 123
			daoMock.On("GetCharacters", mock.Anything).Return(dao.CharacterDao{
				CharID: "iron man",
				LastSync: "any-date",
				Characters: []models.Character{
					{
						Character: "Iron man",
						Comics: []string{"Avengers"},
					},
				},
			}, nil)
			daoMock.On("SaveOrUpdateCharacters", mock.Anything).Return(nil)
			clientMock.On("GetHeroId", mock.Anything).Return(n, nil)
			clientMock.On("GetCaractersByHeroId", mock.Anything).Return([]models.Character{
				{
					Character: "Iron man",
					Comics: []string{"Avengers"},
				},
			}, nil)

			service = &services.CharacterService{
				Client: &clientMock,
				Dao: &daoMock,
			}

			errExpect := errorClient.ClientError{
				HttpStatus: 500,
				Message: "parsing time \"any-date\" as \"02-01-2006 15:04:05\": cannot parse \"any-date\" as \"02\"",
			}

			_, err := service.GetCharacters("ironman")
			Expect(err).NotTo(BeNil())
			Expect(err).To(Equal(errExpect))
		})

		It("Get Characters on happy path", func() {
			var n int64 = 123
			daoMock.On("GetCharacters", mock.Anything).Return(dao.CharacterDao{
				CharID: "iron man",
				LastSync: "01-01-2021 00:00:00",
				Characters: []models.Character{
					{
						Character: "Iron man",
						Comics: []string{"Avengers"},
					},
				},
			}, nil)
			daoMock.On("SaveOrUpdateCharacters", mock.Anything).Return(nil)
			clientMock.On("GetHeroId", mock.Anything).Return(n, nil)
			clientMock.On("GetCaractersByHeroId", mock.Anything).Return([]models.Character{
				{
					Character: "Iron man",
					Comics: []string{"Avengers"},
				},
			}, nil)

			service = &services.CharacterService{
				Client: &clientMock,
				Dao: &daoMock,
			}

			response, err := service.GetCharacters("ironman")
			Expect(err).To(BeNil())
			Expect(len(response.Characters)>0).To(Equal(true))
			Expect(response.Characters[0].Character).To(Equal("Iron man"))
			Expect(len(response.Characters[0].Comics)>0).To(Equal(true))
			Expect(response.Characters[0].Comics[0]).To(Equal("Avengers"))
		})

		BeforeEach(func() {
			clientMock = ClientMock{}
			daoMock = DaoMock{}
		})
	})

})

