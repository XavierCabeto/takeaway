package application_test

import (
	"testing"

	"github.com/XavierCabeto/takeaway/application"
	mock_application "github.com/XavierCabeto/takeaway/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()
	service := application.ProductService{Persistence: persistence}

	result, err := service.Get("abc")
	require.Nil(t, err)
	require.Equal(t, product, result)

}

func TestProductService_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	product := mock_application.NewMockProductInterface(ctrl)
	products := []application.ProductInterface{}
	products = append(products, product)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().GetAll().Return(products, nil).AnyTimes()
	service := application.ProductService{Persistence: persistence}

	result, err := service.GetAll()
	require.Nil(t, err)
	require.Equal(t, products, result)

}

func TestProductService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()
	service := application.ProductService{Persistence: persistence}

	result, err := service.Create("Product 1", 10)
	require.Nil(t, err)
	require.Equal(t, product, result)
}
