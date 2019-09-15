func writeResponse(w http.ResponseWriter, status int, resp interface{}) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(status)
    if err := json.NewEncoder(w).Encode(resp); err != nil {
        panic(err)
    }
}