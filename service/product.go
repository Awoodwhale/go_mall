package service

import (
	"context"
	"go_mall/dao"
	"go_mall/model"
	"go_mall/pkg/e"
	"go_mall/pkg/utils"
	"go_mall/serializer"
	"mime/multipart"
	"sync"
)

// ProductService
// @Description: product service
type ProductService struct {
	ID            uint   `json:"id" form:"id"`
	Name          string `json:"name" form:"name"`
	CategoryId    uint   `json:"category_id" form:"category_id"`
	Title         string `json:"title" form:"title"`
	Info          string `json:"info" form:"info"`
	ImagePath     string `json:"image_path" form:"image_path"`
	Price         string `json:"price" form:"price"`
	DiscountPrice string `json:"discount_price" form:"discount_price"`
	OnSale        bool   `json:"on_sale" form:"on_sale"`
	Num           int    `json:"num" form:"num"`
	model.BasePage
}

// CreateProduct
// @Description: 创建product
// @receiver service *ProductService
// @param c context.Context
// @param uid uint
// @param files []*multipart.FileHeader
// @return serializer.Response
func (service *ProductService) CreateProduct(c context.Context, uid uint, files []*multipart.FileHeader) serializer.Response {
	var (
		code             = e.Success
		msg              string
		err              error
		boss             *model.User
		productImagePath string
	)
	// 过基本验证
	switch {
	case service.Name == "":
		code, msg = e.Error, "商品名称不可为空"
	case service.CategoryId <= 0:
		code, msg = e.Error, "商品分类不可为空"
	case service.Title == "":
		code, msg = e.Error, "商品标题不可为空"
	case service.Price == "":
		code, msg = e.Error, "商品价格不可为空"
	}
	if code != e.Success {
		return serializer.Response{
			Code:    code,
			Message: msg,
		}
	}
	// 没有打折那么打折价就是原价
	if service.DiscountPrice == "" {
		service.DiscountPrice = service.Price
	}

	// 获取boss信息
	userDao := dao.NewUserDao(c)
	boss, err = userDao.GetUserById(uid)
	if err != nil {
		utils.Logger.Errorln("service product GetUserById,", err)
		code = e.ErrorWithSQL
		return serializer.Response{
			Code:    code,
			Message: e.GetMessageByCode(code),
			Error:   err.Error(),
		}
	}

	// 如果存在商品照片,上传的第一张图片作为product的img
	if len(files) > 0 {
		file, err := files[0].Open()
		if err != nil {
			utils.Logger.Errorln("service product open product image,", err)
			code = e.ErrorWithFileOpen
			return serializer.Response{
				Code:    code,
				Message: e.GetMessageByCode(code),
				Error:   err.Error(),
			}
		}
		productImagePath, err = UploadProductImage(file, files[0].Filename, uid)
		if err != nil {
			utils.Logger.Errorln("service product upload product image,", err)
			code = e.ErrorWithUploadProduct
			return serializer.Response{
				Code:    code,
				Message: e.GetMessageByCode(code),
				Error:   err.Error(),
			}
		}
	}
	product := &model.Product{
		Name:          service.Name,
		CategoryID:    service.CategoryId,
		Title:         service.Title,
		Info:          service.Info,
		ImagePath:     productImagePath,
		Price:         service.Price,
		DiscountPrice: service.DiscountPrice,
		OnSale:        true,
		Num:           service.Num,
		BossID:        uid,
		BossName:      boss.UserName,
		BossAvatar:    boss.Avatar,
	}

	// 将product写入MySQL
	productDao := dao.NewProductDaoByDB(userDao.DB) // 复用
	if err = productDao.CreateProduct(product); err != nil {
		utils.Logger.Errorln("service product CreateProduct to MySQL,", err)
		code = e.ErrorWithSQL
		return serializer.Response{
			Code:    code,
			Message: e.GetMessageByCode(code),
			Error:   err.Error(),
		}
	}

	// 读取数据库中刚刚存入的product的数据（更新product的id）
	product, err = productDao.GetProductByName(product.Name)
	if err != nil {
		utils.Logger.Errorln("service product GetProductByName,", err)
		code = e.ErrorWithSQL
		return serializer.Response{
			Code:    code,
			Message: e.GetMessageByCode(code),
			Error:   err.Error(),
		}
	}

	// 将product image写入mysql
	productImageDao := dao.NewProductImageDaoByDB(userDao.DB)
	if err = productImageDao.CreateProductImage(&model.ProductImage{
		ProductID: product.ID,
		ImagePath: product.ImagePath,
	}); err != nil {
		code = e.ErrorWithSQL
		utils.Logger.Errorf("service product write to mysql product image [%v], %v\n", files[0].Filename, err)
		return serializer.Response{
			Code:    code,
			Message: e.GetMessageByCode(code),
			Error:   err.Error(),
		}
	}

	// 还有多的文件需要上传,开协程去上传
	if len(files) > 1 {
		wg := new(sync.WaitGroup)
		wg.Add(len(files) - 1)
		for _, file := range files[1:] {
			_file := file // cp一份file
			go func() {   // 开协程跑上传文件
				defer wg.Done() // 函数执行完一定要去done
				openFile, err := _file.Open()
				if err != nil {
					// 文件打开错误就不上传
					code = e.ErrorWithUploadProduct
					utils.Logger.Errorf("service product open product image _file [%v], %v\n", _file.Filename, err)
					return
				}
				productImagePath, err = UploadProductImage(openFile, _file.Filename, uid)
				if err != nil {
					// 将product image上传到本地
					code = e.ErrorWithUploadProduct
					utils.Logger.Errorf("service product upload product image _file [%v], %v\n", _file.Filename, err)
					return
				}
				// 写入product image的dao
				productImageDao = dao.NewProductImageDaoByDB(productImageDao.DB)
				if err = productImageDao.CreateProductImage(&model.ProductImage{
					ProductID: product.ID,
					ImagePath: productImagePath,
				}); err != nil {
					code = e.ErrorWithSQL
					utils.Logger.Errorf("service product write to mysql product image [%v], %v\n", _file.Filename, err)
					return
				}
			}()

		}
		wg.Wait()
	}

	// 添加成功返回
	return serializer.Response{
		Code:    code,
		Message: e.GetMessageByCode(code),
		Data:    serializer.BuildProduct(product),
	}
}

// ListProduct
// @Description: 获取product的列表
// @receiver service *ProductService
// @param c context.Context
// @return serializer.Response
func (service *ProductService) ListProduct(c context.Context) serializer.Response {
	var (
		code     = e.Success
		err      error
		products []*model.Product
	)

	if service.PageSize == 0 {
		service.PageSize = 15
	}

	condition := map[string]any{}
	if service.CategoryId != 0 {
		condition["category_id"] = service.CategoryId
	}

	productDao := dao.NewProductDao(c)
	total, err := productDao.CountProductByCondition(condition)
	if err != nil {
		code = e.Error
		utils.Logger.Errorln("service product ListProduct CountProductByCondition, ", err)
		return serializer.Response{
			Code:    code,
			Message: e.GetMessageByCode(code),
			Error:   err.Error(),
		}
	}

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		productDao = dao.NewProductDaoByDB(productDao.DB)
		products, _ = productDao.ListProductByCondition(condition, &service.BasePage)
		wg.Done()
	}()
	wg.Wait()

	return serializer.BuildListResponse(serializer.BuildProducts(products), uint(total))
}
