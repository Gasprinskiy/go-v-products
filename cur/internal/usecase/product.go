package usecase

import (
	"jun2/cur/internal/entity/global"
	"jun2/cur/internal/entity/product"
	"jun2/cur/internal/entity/stock"
	"jun2/cur/internal/entity/variation"
	"jun2/cur/rimport"
)

type ProductUsecase struct {
	rimport.RepositoryImports
}

func NewProductUsecase(ri rimport.RepositoryImports) *ProductUsecase {
	return &ProductUsecase{
		RepositoryImports: ri,
	}
}

// CreateProduct создает продукт и вариацию к ниму
func (u *ProductUsecase) CreateProduct(param product.AddProductParams) (createRes product.ProductCreationResult, err error) {
	// NewProduct параметры для создания нового продукта
	newProduct := product.NewProduct(param)
	// создание нового продукта и получение id новой записи
	createRes.ProductID, err = u.Repository.Product.CreateProduct(newProduct)
	if err != nil {
		err = global.ErrInternalError
		return
	}

	// присвоение param.ProductID id нового продукта
	param.ProductID = createRes.ProductID

	// NewVariation параметры для создания новой вариации продукты
	newVariation := variation.NewVariation(param.ProductID, param.VariationType, param.UnitType)
	// создание новой вариации и получение id новой записи
	createRes.VariationID, err = u.Repository.Variation.CreateVariation(newVariation)
	if err != nil {
		err = global.ErrInternalError
		return
	}

	return
}

// FindProductByID получение информации о продукте по ID
func (u *ProductUsecase) FindProductByID(id int) (variation.ProductWithVariationList, error) {
	var data variation.ProductWithVariationList
	var variationInfo []variation.Variation
	// поиск продукта по ID
	product, err := u.Repository.Product.FindProductByID(id)
	if err != nil {
		return data, err
	}

	data = variation.ProductWithVariationList{Product: product}
	// поиск вариации продукта по ID
	variationInfo, err = u.Repository.Variation.FindVariationInfoByID(id)
	if err != nil && err == global.ErrNoData {
		data.VariationList = variationInfo
		return data, nil
	}
	if err != nil {
		return data, err
	}

	var stockInfo []stock.ProductStock
	// загрузка информации в каких складах находится варициации продукта
	for index, vValue := range variationInfo {
		stockInfo, err = u.Repository.Stock.FindVariationAvailability(vValue.VariationID)
		switch err {
		case nil:
			for _, sValue := range stockInfo {
				if sValue.VariationID == vValue.VariationID {
					variationInfo[index].StockAvailability = append(variationInfo[index].StockAvailability, sValue)
				}
			}

		// если вариации нет на складе, то цикл продолжает работу что бы вернуть найденную информацию
		case global.ErrNoData:
			variationInfo[index].StockAvailability = []stock.ProductStock{}
			continue
		default:
			continue
		}
	}

	data.VariationList = variationInfo
	return data, err
}

// FindProductList поиск продуктов по переданным параметрам
func (u *ProductUsecase) FindProductList(tags string, limit, offset int) (data product.ProductListWithTotalCount, err error) {
	// обработка limit
	if limit <= 0 {
		limit = global.DefaultLimit
	}

	// если параметр paramtags передат, выполняеться поиск продуктов по ключевым словам и параметрам limit / offset
	if len(tags) > 0 {
		// поиск продуктов по ключевым словам и параметрам limit / offset
		data.ProductList, err = u.Repository.Product.FindProductListByTags(limit, offset, tags)
		if err != nil {
			return
		}
		// поиск количестова записей по ключевым словам
		data.TotalCount, err = u.Repository.Product.FindProductCountByTags(tags)
		if err != nil {
			return
		}
		return
	}
	// поиск продуктов параметрам limit / offset
	data.ProductList, err = u.Repository.Product.FindProductList(limit, offset)
	if err != nil {
		return
	}

	// получение количества записей
	data.TotalCount, err = u.Repository.Product.LoadProductCount()
	if err != nil {
		return
	}

	return
}
