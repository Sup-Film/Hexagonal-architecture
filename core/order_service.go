package core

import "errors"

// * กำหนด Spec Primary port
type OrderService interface {
	CreateOrder(order Order) error
}

// * สร้าง Struct ของการรับข้อมูลของ Business Logic
// * โดยมี Field ที่เก็บค่าของ OrderRepository ที่เป็น Secondary port
type orderServiceImpl struct {
	repo OrderRepository
}

// * Func ที่ทำการสร้าง instance ของ OrderService
func newOrderService(repo OrderRepository) OrderService {
	// เป็นการนำ OrderRepository มาเก็บไว้ใน Struct ของ orderServiceImpl
	return &orderServiceImpl{repo: repo}
}

/*
* Implment Method ของ OrderService เพื่อทำหน้าที่เป็น Primary port
* สามารถคุยผ่าน Primary port ได้โดยการรับ CreateOrder เป็นช่องทางในการรับข้อมูลมา
* และคุยผ่าน Secondary port ได้โดยผ่านตัว Struct orderServiceImpl ที่รับค่าจาก OrderRepository เข้ามา
 */
func (s *orderServiceImpl) CreateOrder(order Order) error {
	// Business Logic function
	if order.Total <= 0 {
		return errors.New("total must be greater than 0")
	}

	if err := s.repo.Save(order); err != nil {
		return err
	}

	return nil
}
