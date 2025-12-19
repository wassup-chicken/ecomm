package clients

// Example: How to use HTTPClient in your handlers
//
// In your handler:
//
// func GetJob(w http.ResponseWriter, r *http.Request) {
//     ctx := r.Context()  // Get request context
//
//     // Optionally add timeout for external call
//     ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
//     defer cancel()
//
//     // Create client
//     client := NewHTTPClient("https://api.example.com")
//
//     // Make request with context
//     var result JobResponse
//     if err := client.Get(ctx, "/jobs/123", &result); err != nil {
//         if err == context.DeadlineExceeded {
//             http.Error(w, "External service timeout", http.StatusGatewayTimeout)
//             return
//         }
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//         return
//     }
//
//     json.NewEncoder(w).Encode(result)
// }
