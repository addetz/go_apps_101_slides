// writeResponse is a helper method that allows to write and HTTP status & response
func writeResponse(w http.ResponseWriter, status int, resp interface{}) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(status)
    if err := json.NewEncoder(w).Encode(resp); err != nil {
        fmt.Fprintf(w, fmt.Sprintf("error encoding resp %v:%s", resp, err))
    }
}