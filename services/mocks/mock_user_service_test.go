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

	require.Equal(t, expectedUser.ID, "2")
	require.Equal(t, expectedUser.Name, u.Name)
	require.Equal(t, expectedUser.Email, u.Email)
}
