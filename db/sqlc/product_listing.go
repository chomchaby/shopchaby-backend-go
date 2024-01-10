package db

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

// AddProductListingTxParams contains the input parameters of the transfer transaction
type CreateProductListingTxParams struct {
	Category    string    `json:"category" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	StoreID     uuid.UUID `json:"store_id" binding:"required"`
	Suspend     bool      `json:"suspend" binding:"required"`
	Subproducts []struct {
		Variation   string `json:"variation" binding:"required"`
		StockAmount int32  `json:"stock_amount" binding:"required"`
		Price       int32  `json:"price" binding:"required"`
		SalePrice   int32  `json:"sale_price"`
	} `json:"subproducts" binding:"required"`
	ProductImages []struct {
		Title     string `json:"title" binding:"required"`
		PicUrl    string `json:"pic_url" binding:"required"`
		IsDefault bool   `json:"is_default" binding:"required"`
	} `json:"product_images" binding:"required"`
}

type ProductListingTxResult struct {
	Product       Product        `json:"product"`
	Subproducts   []Subproduct   `json:"subproducts"`
	ProductImages []ProductImage `json:"product_images"`
}

// AddProductToStoreTx performs adding product and subproduct into store
func (storeTx *SQLStoreTx) CreateProductListingTx(ctx context.Context, arg CreateProductListingTxParams) (ProductListingTxResult, error) {
	var result ProductListingTxResult

	err := storeTx.execTx(ctx, func(q *Queries) error {

		var err error
		product, err := q.CreateProduct(ctx, CreateProductParams{
			Category:    arg.Category,
			Name:        arg.Name,
			Description: arg.Description,
			StoreID:     arg.StoreID,
			Suspend:     arg.Suspend,
		})

		if err != nil {
			return err
		}

		productID := product.ID
		result.Subproducts = make([]Subproduct, len(arg.Subproducts)) // Initialize the result slice

		for i, subprod := range arg.Subproducts {

			var sale_price sql.NullInt32
			if subprod.SalePrice == 0 {
				sale_price = sql.NullInt32{Int32: 0, Valid: false}
			} else if subprod.SalePrice < subprod.Price {
				sale_price = sql.NullInt32{Int32: subprod.SalePrice, Valid: true}
			} else {
				return errors.New("error: Invalid Sale Price")

			}

			subproduct, err := q.CreateSubproduct(ctx, CreateSubproductParams{
				ProductID:   productID,
				Variation:   subprod.Variation,
				StockAmount: subprod.StockAmount,
				Price:       subprod.Price,
				SalePrice:   sale_price,
				// Fill in other fields as needed for CreateSubproductParams
			})
			if err != nil {
				return err
			}
			result.Subproducts[i] = subproduct // Store the created subproduct in the result slice
		}

		result.ProductImages = make([]ProductImage, len(arg.ProductImages)) // Initialize the result slice

		for i, img := range arg.ProductImages {
			image, err := q.CreateProductImage(ctx, CreateProductImageParams{
				ProductID: productID,
				Title:     img.Title,
				PicUrl:    img.PicUrl,
				IsDefault: img.IsDefault,
				// Fill in other fields as needed
			})
			if err != nil {
				return err
			}
			result.ProductImages[i] = image // Store the created subproduct in the result slice
		}

		product, err = q.GetProduct(ctx, productID)
		if err != nil {
			return err
		}

		result.Product = product

		return nil
	})

	return result, err
}

type UpdateProductListingTxParams struct {
	ProductID   uuid.UUID `json:"product_id" binding:"required"`
	Category    string    `json:"category" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Suspend     bool      `json:"suspend" binding:"required"`
	Subproducts []struct {
		Variation   string `json:"variation" binding:"required"`
		StockAmount int32  `json:"stock_amount" binding:"required"`
		Price       int32  `json:"price" binding:"required"`
		SalePrice   int32  `json:"sale_price"`
	} `json:"subproducts" binding:"required"`
	ProductImages []struct {
		Title     string `json:"title" binding:"required"`
		PicUrl    string `json:"pic_url" binding:"required"`
		IsDefault bool   `json:"is_default" binding:"required"`
	} `json:"product_images" binding:"required"`
}

func (storeTx *SQLStoreTx) UpdateProductListingTx(ctx context.Context, arg UpdateProductListingTxParams) (ProductListingTxResult, error) {
	var result ProductListingTxResult

	err := storeTx.execTx(ctx, func(q *Queries) error {

		product, err := q.UpdateProduct(ctx, UpdateProductParams{
			ID:          arg.ProductID,
			Category:    arg.Category,
			Name:        arg.Name,
			Description: arg.Description,
			Suspend:     arg.Suspend,
		})

		if err != nil {
			return err
		}

		productID := product.ID

		// delete old subproduct
		oldSub, err := q.ListSubproductsByProductID(ctx, productID)
		if err != nil {
			return err
		}

		for _, subprod := range oldSub {
			err := q.DeleteSubproduct(ctx, subprod.ID)
			if err != nil {
				return err
			}
		}

		// create new subproduct
		result.Subproducts = make([]Subproduct, len(arg.Subproducts)) // Initialize the result slice
		for i, subprod := range arg.Subproducts {

			var sale_price sql.NullInt32
			if subprod.SalePrice == 0 {
				sale_price = sql.NullInt32{Int32: 0, Valid: false}
			} else if subprod.SalePrice < subprod.Price {
				sale_price = sql.NullInt32{Int32: subprod.SalePrice, Valid: true}
			} else {
				return errors.New("error: Invalid Sale Price")

			}

			subproduct, err := q.CreateSubproduct(ctx, CreateSubproductParams{
				ProductID:   productID,
				Variation:   subprod.Variation,
				StockAmount: subprod.StockAmount,
				Price:       subprod.Price,
				SalePrice:   sale_price,
				// Fill in other fields as needed for CreateSubproductParams
			})
			if err != nil {
				return err
			}
			result.Subproducts[i] = subproduct // Store the created subproduct in the result slice
		}

		// delete old product image
		oldImage, err := q.ListProductImagesByProductID(ctx, productID)
		if err != nil {
			return err
		}
		for _, img := range oldImage {
			err := q.DeleteProductImage(ctx, img.ID)
			if err != nil {
				return err
			}
		}

		// create new product image
		result.ProductImages = make([]ProductImage, len(arg.ProductImages)) // Initialize the result slice
		for i, img := range arg.ProductImages {
			image, err := q.CreateProductImage(ctx, CreateProductImageParams{
				ProductID: productID,
				Title:     img.Title,
				PicUrl:    img.PicUrl,
				IsDefault: img.IsDefault,
				// Fill in other fields as needed for CreateSubproductParams
			})
			if err != nil {
				return err
			}
			result.ProductImages[i] = image // Store the created subproduct in the result slice
		}

		product, err = q.GetProduct(ctx, productID)
		if err != nil {
			return err
		}

		result.Product = product

		return nil
	})
	return result, err
}
