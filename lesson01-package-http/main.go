package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/demo", demoHandler)

	log.Println("Server is starting...")
	err := http.ListenAndServe(":8080", nil) // nếu dùng nil thì sẽ dùng http.DefaultServeMux
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	} else {
		log.Println("Server is running on port 8080")

	}
}

func demoHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%+v", r)
	if r.Method != http.MethodGet {
		// Check request method chỉ được dùng 1 phương thức
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	response := map[string]string{
		"message": "Chào mừng bạn đã đén thế giới của mình",
		"author":  "Dacoband",
	}

	w.Header().Set("Content-Type", "application/json") // chắc chắn kiểu dữ liệu là JSON
	w.Header().Set("X-Course", "Go Programming")

	// Cách 1 : Sử dụng json.Marshal
	// data, err := json.Marshal(response)
	// if err != nil {
	// 	http.Error(w, "Lỗi mã hóa JSON", http.StatusInternalServerError)
	// 	return
	// }

	// w.Write(data) // Ghi dữ liệu JSON vào phản hồi

	//Cách 2 : Sử dụng json.NewEncoder
	json.NewEncoder(w).Encode(response)

}
