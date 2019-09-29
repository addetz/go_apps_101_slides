func(or *OrderRepo) FindAll() []models.Order {...}
func(or *OrderRepo) Find(id string)(*models.Order, error) {...}
func(or *OrderRepo) Upsert(o *models.Order) error {...}
func(or *OrderRepo) Delete(id string) error {...}
