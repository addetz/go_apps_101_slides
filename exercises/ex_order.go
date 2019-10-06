type Order struct {
    ID string
    // Status as OrderStatus, Items as slice of item pointers to avoid copying them
}
