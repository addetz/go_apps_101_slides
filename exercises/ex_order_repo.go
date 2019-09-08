func(or *OrderRepo) FindAll() []models.Order {
    // implementation here
}
func(or *OrderRepo) FindOrder(id string)(*models.Order, error) {
	// implementation here 
}
func(or *OrderRepo) UpsertOrder(o *models.Order) error {
	// implementation here
}
func(or *OrderRepo) Delete(id string) error {
    // implementation here
}
