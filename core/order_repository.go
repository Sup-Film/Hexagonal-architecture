package core

// กำหนด Spec เพื่อให้ Adapter Secondary มาเรียกใช้งาน
// ถ้าฝั่ง Adapter จะมาเรียกใช้งานจะต้องมี Method อะไรบ้าง ให้เขียนไว้ใน Interface
type OrderRepository interface {
	// Method Save รับค่า Order struct และ คืนค่าเป็น error
	Save(order Order) error
}