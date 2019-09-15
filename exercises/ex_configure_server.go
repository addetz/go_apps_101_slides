// configureServer configures the routes of this server and binds handler functions to them
func configureServer(handler *handlers.Handler) *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    router.Methods("GET").Path("/").Handler(http.HandlerFunc(handler.Index))
    router.Methods("POST").Path("/orders/").Handler(http.HandlerFunc(handler.OrderUpsert))
    router.Methods("DELETE").Path("/orders/{orderId}").Handler(http.HandlerFunc(handler.OrderDelete))

    return router
}