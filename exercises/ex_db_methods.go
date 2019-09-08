func InitDB() *DB {...}
func(db *DB) UpsertOrder(o models.Order) error {...}
func(db *DB) FindOrder(id string) (models.Order, error) {...}
func(db *DB) FindAllOrders() models.Orders {...}
func(db *DB) DeleteOrder(id string) error {...}




