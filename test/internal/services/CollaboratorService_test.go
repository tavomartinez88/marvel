package services

import (
	"errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"github.com/tavomartinez88/marvel/internal/dao"
	errorClient "github.com/tavomartinez88/marvel/internal/error"
	"github.com/tavomartinez88/marvel/internal/services"
	"github.com/tavomartinez88/marvel/internal/utils"
)

var _ = Describe("Collaborator Service tests", func() {
	var (
		service    services.ICollaboratorService
		clientMock ClientMock
		daoMock    DaoMock
	)

	Context("Collaborator service tests", func() {

		It("Invalid parameter input", func() {
			service = &services.CollaboratorService{
				Client: &clientMock,
				Dao:    &daoMock,
			}

			errExpect := errorClient.ClientError{
				HttpStatus: 400,
				Message:    "Name is invalid",
			}

			_, err := service.GetCollaborators("1ronman")
			Expect(err).NotTo(BeNil())
			Expect(errExpect).Should(Equal(err))

		})

		It("Get Collaborators there is a problem with db", func() {

			daoMock.On("GetCollaborators", mock.Anything).Return(dao.CollaboratorsDao{}, errors.New("error db"))
			service = &services.CollaboratorService{
				Client: &clientMock,
				Dao: &daoMock,
			}

			errExpect := errorClient.ClientError{
				HttpStatus: 400,
				Message: "error db",
			}

			_, err := service.GetCollaborators("ironman")
			Expect(err).NotTo(BeNil())
			Expect(errExpect).Should(Equal(err))
		})

		It("Get Collaborators and get error on GetCollaboratorsByHeroId", func() {
			var n int64 = 123
			daoMock.On("GetCollaborators", mock.Anything).Return(dao.CollaboratorsDao{}, errors.New("record not found"))
			clientMock.On("GetHeroId", mock.Anything).Return(n, nil)
			clientMock.On("GetCollaboratorsByHeroId", mock.Anything).Return([]utils.Collaborator{}, errors.New("error client"))

			service = &services.CollaboratorService{
				Client: &clientMock,
				Dao: &daoMock,
			}

			_, err := service.GetCollaborators("ironman")
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal("error client"))
		})

		It("Get Collaborators and get error on GetHeroId", func() {
			var n int64 = 0
			daoMock.On("GetCollaborators", mock.Anything).Return(dao.CollaboratorsDao{}, errors.New("record not found"))
			clientMock.On("GetHeroId", mock.Anything).Return(n, errors.New("error on db"))
			clientMock.On("GetCollaboratorsByHeroId", mock.Anything).Return([]utils.Collaborator{}, errors.New("error client"))

			service = &services.CollaboratorService{
				Client: &clientMock,
				Dao: &daoMock,
			}

			_, err := service.GetCollaborators("ironman")
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal("error on db"))
		})

		It("Get Collaborators and fail parse date", func() {
			var n int64 = 123
			daoMock.On("GetCollaborators", mock.Anything).Return(dao.CollaboratorsDao{
				LastSync: "any-date",
			}, nil)
			daoMock.On("SaveOrUpdateCollaborators", mock.Anything).Return(nil)
			clientMock.On("GetHeroId", mock.Anything).Return(n, nil)
			clientMock.On("GetCollaboratorsByHeroId", mock.Anything).Return([]utils.Collaborator{
				{
					Name: "james",
					Role: "editor",
				},
			}, nil)

			service = &services.CollaboratorService{
				Client: &clientMock,
				Dao: &daoMock,
			}

			errExpect := errorClient.ClientError{
				HttpStatus: 500,
				Message: "parsing time \"any-date\" as \"02-01-2006 15:04:05\": cannot parse \"any-date\" as \"02\"",
			}

			_, err := service.GetCollaborators("ironman")
			Expect(err).NotTo(BeNil())
			Expect(err).To(Equal(errExpect))
		})

		It("Get Characters and get error to save", func() {
			var n int64 = 123
			daoMock.On("GetCollaborators", mock.Anything).Return(dao.CollaboratorsDao{
				LastSync: "01-01-2021 00:00:00",
				Colorists: []string{"juan"},
				Editors: []string{"juan"},
				Writers: []string{"juan"},
			}, nil)
			daoMock.On("SaveOrUpdateCollaborators", mock.Anything).Return(errors.New("error to save"))
			clientMock.On("GetHeroId", mock.Anything).Return(n, nil)
			clientMock.On("GetCollaboratorsByHeroId", mock.Anything).Return([]utils.Collaborator{
				{
					Name: "james",
					Role: "editor",
				},
			}, nil)

			service = &services.CollaboratorService{
				Client: &clientMock,
				Dao: &daoMock,
			}

			errExpect := errorClient.ClientError{
				HttpStatus: 500,
				Message: "error to save",
			}

			_, err := service.GetCollaborators("ironman")
			Expect(err).NotTo(BeNil())
			Expect(err).To(Equal(errExpect))
		})

		It("Get Characters on happy path", func() {
			var n int64 = 123
			daoMock.On("GetCollaborators", mock.Anything).Return(dao.CollaboratorsDao{
				LastSync: "01-01-2021 00:00:00",
				Colorists: []string{"juan"},
				Editors: []string{"juan"},
				Writers: []string{"juan"},
			}, nil)
			daoMock.On("SaveOrUpdateCollaborators", mock.Anything).Return(nil)
			clientMock.On("GetHeroId", mock.Anything).Return(n, nil)
			clientMock.On("GetCollaboratorsByHeroId", mock.Anything).Return([]utils.Collaborator{
				{
					Name: "james",
					Role: "editor",
				},
				{
					Name: "james",
					Role: "colorist",
				},
				{
					Name: "james",
					Role: "writer",
				},
			}, nil)

			service = &services.CollaboratorService{
				Client: &clientMock,
				Dao: &daoMock,
			}

			response, err := service.GetCollaborators("ironman")
			Expect(err).To(BeNil())
			Expect(len(response.Colorists)>0).To(Equal(true))
			Expect(len(response.Writers)>0).To(Equal(true))
			Expect(len(response.Editors)>0).To(Equal(true))
		})

		BeforeEach(func() {
			clientMock = ClientMock{}
			daoMock = DaoMock{}
		})
	})
})