package common

import "strings"

func ExtractFilePathFromURL(url string) string {
	index := strings.Index(url, "/uploads/")
	if index == -1 {
		return "" // Không tìm thấy "/uploads/"
	}

	// Trích xuất phần đường dẫn file từ URL
	filePath := "./" + url[index+1:]

	return filePath
}
