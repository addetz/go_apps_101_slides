func(or *OrderRepo) FindAll() []models.Order {
    // implementation here
}
func(or *OrderRepo) Find(id string)(*models.Order, error) {
	// implementation here 
}
func(or *OrderRepo) Upsert(o *models.Order) error {
	// implementation here
}
func(or *OrderRepo) Delete(id string) error {
    // implementation here
}
