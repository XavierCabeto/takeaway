package cli_test

import (
	"fmt"
	"testing"

	"github.com/XavierCabeto/takeaway/adapters/cli"
	"github.com/XavierCabeto/takeaway/application"
	mock_application "github.com/XavierCabeto/takeaway/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Product Test"
	productPrice := 25.99
	productId := "abc"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	products := []application.ProductInterface{}
	products = append(products, productMock)
	service.EXPECT().GetAll().Return(products, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Product ID %s with the name %s has been created with the price %f",
		productId, productName, productPrice)
	result, err := cli.Run(service, "create", "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f",
		productId, productName, productPrice)
	result, err = cli.Run(service, "get", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\n",
		productId, productName, productPrice)
	result, err = cli.Run(service, "getAll", "", "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
}
