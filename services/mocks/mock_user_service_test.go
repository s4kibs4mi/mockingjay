package mocks

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/s4kibs4mi/mockingjay/services"
	"github.com/stretchr/testify/require"
)

func Test_UserService_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectedID := "1"
	expectedName := "Sakib"
	expectedEmail := "sakib@liveklass.io"
	expectedUser := &services.User{ID: expectedID, Name: expectedName, Email: expectedEmail}

	mockUserService := NewMockUserService(ctrl)
	mockUserService.EXPECT().CreateUser(expectedName, expectedEmail).Return(expectedUser, nil).Times(1)

	u, err := mockUserService.CreateUser(expectedName, expectedEmail)
	require.NoError(t, err)

	require.Equal(t, expectedUser.ID, u.ID)
	require.Equal(t, expectedUser.Name, u.Name)
	require.Equal(t, expectedUser.Email, u.Email)
}

func Test_UserService_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectedID := "1"
	expectedName := "Sakib"
	expectedEmail := "sakib@liveklass.io"
	expectedUser := &services.User{ID: expectedID, Name: expectedName, Email: expectedEmail}

	mockUserService := NewMockUserService(ctrl)
	mockUserService.EXPECT().UpdateUser(expectedID, expectedName, expectedEmail).Return(expectedUser, nil)

	u, err := mockUserService.UpdateUser(expectedID, expectedName, expectedEmail)
	require.NoError(t, err)

	require.Equal(t, expectedUser.ID, u.ID)
	require.Equal(t, expectedUser.Name, u.Name)
	require.Equal(t, expectedUser.Email, u.Email)
}

func Test_UserService_GetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectedID := "1"
	expectedName := "Sakib"
	expectedEmail := "sakib@liveklass.io"
	expectedUser := &services.User{ID: expectedID, Name: expectedName, Email: expectedEmail}

	mockUserService := NewMockUserService(ctrl)
	mockUserService.EXPECT().GetUser(expectedID).Return(expectedUser, nil)

	u, err := mockUserService.GetUser(expectedID)
	require.NoError(t, err)

	require.Equal(t, expectedUser.ID, u.ID)
	require.Equal(t, expectedUser.Name, u.Name)
	require.Equal(t, expectedUser.Email, u.Email)
}
