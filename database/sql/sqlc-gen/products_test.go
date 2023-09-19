package database

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/fredele20/order.ms-leta/utils"
	"github.com/stretchr/testify/require"
)

func createRandomProduct(t *testing.T) Product {
	args := CreateProductParams{
		ID:          utils.GenerateId(),
		Name:        utils.RandomName(),
		Description: utils.RandomDescription(),
		Quantity:    utils.RandomInt(1, 20),
		Price:       utils.RandomPrice(),
	}

	product, err := testQueries.CreateProduct(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, product)

	require.Equal(t, args.ID, product.ID)
	require.Equal(t, args.Name, product.Name)
	require.Equal(t, args.Description, product.Description)
	require.Equal(t, args.Quantity, product.Quantity)
	require.Equal(t, args.Price, product.Price)

	require.NotZero(t, product.CreatedAt)

	return product
}

func TestCreateProduct(t *testing.T) {
	createRandomProduct(t)
}

func TestGetProduct(t *testing.T) {
	product1 := createRandomProduct(t)

	product2, err := testQueries.GetProduct(context.Background(), product1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, product2)

	require.Equal(t, product1.ID, product2.ID)
	require.Equal(t, product1.Name, product2.Name)
	require.Equal(t, product1.Description, product2.Description)
	require.Equal(t, product1.Quantity, product2.Quantity)
	require.Equal(t, product1.Price, product2.Price)
	require.WithinDuration(t, product1.CreatedAt, product2.CreatedAt, time.Second)
}

func TestDeleteProduct(t *testing.T) {
	product1 := createRandomProduct(t)

	err := testQueries.DeleteProduct(context.Background(), product1.ID)
	require.NoError(t, err)

	product2, err := testQueries.GetProduct(context.Background(), product1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, product2)
}

func TestListProducts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomProduct(t)
	}

	args := ListProductsParams{
		Limit: 5,
		Offset: 5,
	}

	products, err := testQueries.ListProducts(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, products, 5)

	for _, product := range products {
		require.NotEmpty(t, product)
	}
}
