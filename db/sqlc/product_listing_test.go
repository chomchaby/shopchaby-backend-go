package db

import (
	"context"
	"testing"
	"time"

	"github.com/chomchaby/shopchaby-backend-go/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func createRandomProductListingTx(t *testing.T) []ProductListingTxResult {

	storeTx := NewStoreTx(testDB)

	store := createRandomStore(t)
	var listing []ProductListingTxResult

	errs := make(chan error)
	results := make(chan ProductListingTxResult)

	go func() {
		product := util.RandomProductDetail()
		subproduct1 := util.RandomSubproductDetail(util.RandomBoolean())
		subproduct2 := util.RandomSubproductDetail(util.RandomBoolean())
		subproduct3 := util.RandomSubproductDetail(util.RandomBoolean())
		productimage1 := util.RandomProductImageDetail()
		productimage2 := util.RandomProductImageDetail()
		productimage3 := util.RandomProductImageDetail()

		result, err := storeTx.CreateProductListingTx(context.Background(), CreateProductListingTxParams{
			Category:    product.Category,
			Name:        product.Name,
			Description: product.Description,
			StoreID:     store.ID,
			Suspend:     product.Suspend,
			Subproducts: []struct {
				Variation   string `json:"variation" binding:"required"`
				StockAmount int32  `json:"stock_amount" binding:"required"`
				Price       int32  `json:"price" binding:"required"`
				SalePrice   int32  `json:"sale_price"`
			}{
				{
					Variation:   subproduct1.Variation,
					StockAmount: subproduct1.StockAmount,
					Price:       subproduct1.Price,
					SalePrice:   subproduct1.SalePrice.Int32,
				},
				{
					Variation:   subproduct2.Variation,
					StockAmount: subproduct2.StockAmount,
					Price:       subproduct2.Price,
					SalePrice:   subproduct2.SalePrice.Int32,
				},
				{
					Variation:   subproduct3.Variation,
					StockAmount: subproduct3.StockAmount,
					Price:       subproduct3.Price,
				},
				// You can add more subproducts here if needed
			},
			ProductImages: []struct {
				Title     string `json:"title" binding:"required"`
				PicUrl    string `json:"pic_url" binding:"required"`
				IsDefault bool   `json:"is_default" binding:"required"`
			}{
				{
					Title:     productimage1.Title,
					PicUrl:    productimage1.PicUrl,
					IsDefault: true,
				},
				{
					Title:     productimage2.Title,
					PicUrl:    productimage2.PicUrl,
					IsDefault: productimage2.IsDefault,
				},
				{
					Title:     productimage3.Title,
					PicUrl:    productimage3.PicUrl,
					IsDefault: productimage3.IsDefault,
				},
			},
		})
		listing = append(listing, result)
		errs <- err
		results <- result

	}()
	go func() {
		product := util.RandomProductDetail()
		subproduct1 := util.RandomSubproductDetail(util.RandomBoolean())
		subproduct2 := util.RandomSubproductDetail(util.RandomBoolean())
		productimage1 := util.RandomProductImageDetail()
		productimage2 := util.RandomProductImageDetail()
		result, err := storeTx.CreateProductListingTx(context.Background(), CreateProductListingTxParams{
			Category:    product.Category,
			Name:        product.Name,
			Description: product.Description,
			StoreID:     store.ID,
			Suspend:     product.Suspend,
			Subproducts: []struct {
				Variation   string `json:"variation" binding:"required"`
				StockAmount int32  `json:"stock_amount" binding:"required"`
				Price       int32  `json:"price" binding:"required"`
				SalePrice   int32  `json:"sale_price"`
			}{
				{
					Variation:   subproduct1.Variation,
					StockAmount: subproduct1.StockAmount,
					Price:       subproduct1.Price,
					SalePrice:   subproduct1.SalePrice.Int32,
				},
				{
					Variation:   subproduct2.Variation,
					StockAmount: subproduct2.StockAmount,
					Price:       subproduct2.Price,
					SalePrice:   subproduct2.SalePrice.Int32,
				},
				// You can add more subproducts here if needed
			},
			ProductImages: []struct {
				Title     string `json:"title" binding:"required"`
				PicUrl    string `json:"pic_url" binding:"required"`
				IsDefault bool   `json:"is_default" binding:"required"`
			}{
				{
					Title:     productimage1.Title,
					PicUrl:    productimage1.PicUrl,
					IsDefault: true,
				},
				{
					Title:     productimage2.Title,
					PicUrl:    productimage2.PicUrl,
					IsDefault: productimage2.IsDefault,
				},
			},
		})
		listing = append(listing, result)
		errs <- err
		results <- result
	}()
	go func() {
		product := util.RandomProductDetail()
		subproduct1 := util.RandomSubproductDetail(util.RandomBoolean())
		productimage1 := util.RandomProductImageDetail()
		productimage2 := util.RandomProductImageDetail()
		result, err := storeTx.CreateProductListingTx(context.Background(), CreateProductListingTxParams{
			Category:    product.Category,
			Name:        product.Name,
			Description: product.Description,
			StoreID:     store.ID,
			Suspend:     product.Suspend,
			Subproducts: []struct {
				Variation   string `json:"variation" binding:"required"`
				StockAmount int32  `json:"stock_amount" binding:"required"`
				Price       int32  `json:"price" binding:"required"`
				SalePrice   int32  `json:"sale_price"`
			}{
				{
					Variation:   subproduct1.Variation,
					StockAmount: subproduct1.StockAmount,
					Price:       subproduct1.Price,
					SalePrice:   subproduct1.SalePrice.Int32,
				},
				// You can add more subproducts here if needed
			},
			ProductImages: []struct {
				Title     string `json:"title" binding:"required"`
				PicUrl    string `json:"pic_url" binding:"required"`
				IsDefault bool   `json:"is_default" binding:"required"`
			}{
				{
					Title:     productimage1.Title,
					PicUrl:    productimage1.PicUrl,
					IsDefault: true,
				},
				{
					Title:     productimage2.Title,
					PicUrl:    productimage2.PicUrl,
					IsDefault: productimage2.IsDefault,
				},
			},
		})
		listing = append(listing, result)
		errs <- err
		results <- result
	}()

	// check result
	for i := 0; i < 3; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		pro := result.Product
		sub := result.Subproducts
		img := result.ProductImages

		// check product
		require.NotEmpty(t, pro)
		require.Equal(t, store.ID, pro.StoreID)
		require.NotZero(t, pro.ID)
		require.NotZero(t, pro.CreatedAt)
		require.NotZero(t, pro.Category)
		require.NotZero(t, pro.Name)
		require.NotZero(t, pro.Description)
		_, err = storeTx.GetProduct(context.Background(), pro.ID)
		require.NoError(t, err)

		// check min price, max price
		if len(sub) > 0 {
			require.True(t, pro.MinPrice.Valid)
			require.True(t, pro.MaxPrice.Valid)
			for j := 0; j < len(sub); j++ {
				// Ensure that the product's MinPrice is less than or equal to the subproduct's price
				require.True(t, pro.MinPrice.Int32 <= sub[j].Price)
				// Ensure that the product's MaxPrice is greater than or equal to the subproduct's price
				require.GreaterOrEqual(t, pro.MaxPrice.Int32, sub[j].Price)
			}
		}

		// check onsale
		onsale := false
		for j := 0; j < len(sub); j++ {
			if sub[j].SalePrice.Valid == true {
				onsale = true
				break
			}
		}
		require.Equal(t, pro.Onsale, onsale)

		// check vendible
		require.Equal(t, pro.Vendible, len(sub) > 0)

		// check subproduct
		require.NotEmpty(t, sub)
		require.NotZero(t, sub[0].ID)
		require.NotZero(t, sub[0].Variation)
		require.NotZero(t, sub[0].Price)
		require.NotZero(t, sub[0].CreatedAt)
		_, err = storeTx.GetSubproduct(context.Background(), sub[0].ID)
		require.NoError(t, err)

		// check product_image
		require.NotEmpty(t, img)
		require.NotZero(t, img[0].ID)
		require.NotZero(t, img[0].Title)
		require.NotZero(t, img[0].CreatedAt)
		_, err = storeTx.GetProductImage(context.Background(), img[0].ID)
		require.NoError(t, err)
	}

	return listing
}

func TestCreateProductListingTx(t *testing.T) {
	createRandomProductListingTx(t)

}

func updateRandomProductListingTx(t *testing.T) {

	storeTx := NewStoreTx(testDB)

	listing := createRandomProductListingTx(t)

	var productIDs []uuid.UUID
	for _, item := range listing {
		productIDs = append(productIDs, item.Product.ID)
	}

	errs := make(chan error)
	results := make(chan ProductListingTxResult)

	go func() {
		// have at least subproduct that is on sale
		product := util.RandomProductDetail()
		subproduct1 := util.RandomSubproductDetail(util.RandomBoolean())
		subproduct2 := util.RandomSubproductDetail(util.RandomBoolean())
		productimage1 := util.RandomProductImageDetail()
		productimage2 := util.RandomProductImageDetail()
		result, err := storeTx.UpdateProductListingTx(context.Background(), UpdateProductListingTxParams{
			ProductID:   listing[0].Product.ID,
			Category:    product.Category,
			Name:        product.Name,
			Description: product.Description,
			Suspend:     product.Suspend,
			Subproducts: []struct {
				Variation   string `json:"variation" binding:"required"`
				StockAmount int32  `json:"stock_amount" binding:"required"`
				Price       int32  `json:"price" binding:"required"`
				SalePrice   int32  `json:"sale_price"`
			}{
				{
					Variation:   subproduct1.Variation,
					StockAmount: subproduct1.StockAmount,
					Price:       subproduct1.Price,
					SalePrice:   subproduct1.SalePrice.Int32,
				},
				{
					Variation:   subproduct2.Variation,
					StockAmount: subproduct2.StockAmount,
					Price:       subproduct2.Price,
					SalePrice:   subproduct2.SalePrice.Int32,
				},
			},
			ProductImages: []struct {
				Title     string `json:"title" binding:"required"`
				PicUrl    string `json:"pic_url" binding:"required"`
				IsDefault bool   `json:"is_default" binding:"required"`
			}{
				{
					Title:     productimage1.Title,
					PicUrl:    productimage1.PicUrl,
					IsDefault: true,
				},
				{
					Title:     productimage2.Title,
					PicUrl:    productimage2.PicUrl,
					IsDefault: productimage2.IsDefault,
				},
			},
		})
		errs <- err
		results <- result
	}()

	// have no subproduct that is on sale
	go func() {
		product := util.RandomProductDetail()
		subproduct := util.RandomSubproductDetail(util.RandomBoolean())
		productimage1 := util.RandomProductImageDetail()
		productimage2 := util.RandomProductImageDetail()
		result, err := storeTx.UpdateProductListingTx(context.Background(), UpdateProductListingTxParams{
			ProductID:   listing[1].Product.ID,
			Category:    product.Category,
			Name:        product.Name,
			Description: product.Description,
			Suspend:     product.Suspend,
			Subproducts: []struct {
				Variation   string `json:"variation" binding:"required"`
				StockAmount int32  `json:"stock_amount" binding:"required"`
				Price       int32  `json:"price" binding:"required"`
				SalePrice   int32  `json:"sale_price"`
			}{
				{
					Variation:   subproduct.Variation,
					StockAmount: subproduct.StockAmount,
					Price:       subproduct.Price,
					SalePrice:   subproduct.SalePrice.Int32,
				},
			},
			ProductImages: []struct {
				Title     string `json:"title" binding:"required"`
				PicUrl    string `json:"pic_url" binding:"required"`
				IsDefault bool   `json:"is_default" binding:"required"`
			}{
				{
					Title:     productimage1.Title,
					PicUrl:    productimage1.PicUrl,
					IsDefault: true,
				},
				{
					Title:     productimage2.Title,
					PicUrl:    productimage2.PicUrl,
					IsDefault: productimage2.IsDefault,
				},
			},
		})
		errs <- err
		results <- result
	}()
	// has no subproduct (not vendible)
	go func() {
		product := util.RandomProductDetail()
		productimage1 := util.RandomProductImageDetail()
		productimage2 := util.RandomProductImageDetail()
		result, err := storeTx.UpdateProductListingTx(context.Background(), UpdateProductListingTxParams{
			ProductID:   listing[2].Product.ID,
			Category:    product.Category,
			Name:        product.Name,
			Description: product.Description,
			Suspend:     product.Suspend,
			Subproducts: []struct {
				Variation   string `json:"variation" binding:"required"`
				StockAmount int32  `json:"stock_amount" binding:"required"`
				Price       int32  `json:"price" binding:"required"`
				SalePrice   int32  `json:"sale_price"`
			}{},
			ProductImages: []struct {
				Title     string `json:"title" binding:"required"`
				PicUrl    string `json:"pic_url" binding:"required"`
				IsDefault bool   `json:"is_default" binding:"required"`
			}{
				{
					Title:     productimage1.Title,
					PicUrl:    productimage1.PicUrl,
					IsDefault: true,
				},
				{
					Title:     productimage2.Title,
					PicUrl:    productimage2.PicUrl,
					IsDefault: productimage2.IsDefault,
				},
			},
		})
		errs <- err
		results <- result
	}()

	// check result
	for i := 0; i < 3; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		pro := result.Product
		sub := result.Subproducts
		img := result.ProductImages

		// check product
		require.NotEmpty(t, pro)
		require.NotZero(t, pro.ID)
		require.NotZero(t, pro.Category)
		require.NotZero(t, pro.Name)
		require.NotZero(t, pro.Description)

		require.Contains(t, productIDs, pro.ID)
		require.NotEqual(t, pro.CreatedAt, pro.UpdatedAt)
		_, err = storeTx.GetProduct(context.Background(), pro.ID)
		require.NoError(t, err)

		// check min price, max price
		if len(sub) > 0 {
			require.True(t, pro.MinPrice.Valid)
			require.True(t, pro.MaxPrice.Valid)
			for j := 0; j < len(sub); j++ {
				// Ensure that the product's MinPrice is less than or equal to the subproduct's price
				require.True(t, pro.MinPrice.Int32 <= sub[j].Price)
				// Ensure that the product's MaxPrice is greater than or equal to the subproduct's price
				require.GreaterOrEqual(t, pro.MaxPrice.Int32, sub[j].Price)
			}
		}

		// check onsale
		onsale := false
		for j := 0; j < len(sub); j++ {
			if sub[j].SalePrice.Valid == true {
				onsale = true
				break
			}
		}
		require.Equal(t, pro.Onsale, onsale)

		// check vendible
		require.Equal(t, pro.Vendible, len(sub) > 0)

		// check subproduct
		if len(sub) > 0 {
			require.NotEmpty(t, sub)
			require.NotZero(t, sub[0].ID)
			require.NotZero(t, sub[0].Variation)
			require.NotZero(t, sub[0].Price)
			require.NotZero(t, sub[0].CreatedAt)

			// now we edit "subproduct" by deleting all and creating all
			require.WithinDuration(t, sub[0].CreatedAt, sub[0].UpdatedAt, time.Second)

			_, err = storeTx.GetSubproduct(context.Background(), sub[0].ID)
			require.NoError(t, err)
		}

		// check product_image
		require.NotEmpty(t, img)
		require.NotZero(t, img[0].ID)
		require.NotZero(t, img[0].Title)
		require.NotZero(t, img[0].CreatedAt)

		// now we edit "product_image" by deleting all and creating all
		require.WithinDuration(t, img[0].CreatedAt, img[0].UpdatedAt, time.Second)

		_, err = storeTx.GetProductImage(context.Background(), img[0].ID)
		require.NoError(t, err)

	}
}
func TestUpdateProductListingTx(t *testing.T) {
	updateRandomProductListingTx(t)
}
