package main

import "fmt"

// GlobalLimit const
const GlobalLimit = 100

// MaxCacheSize const
const MaxCacheSize int = 10 * GlobalLimit

// CacheKeyBook const
const CacheKeyBook = "book_"

// CacheKeyCD const
const CacheKeyCD = "cd_"

var cache map[string]string

func main() {

	cache = make(map[string]string)

	SetBook("1234-5678", "Get Ready To GO")
	SetCD("1234-5678", "Get Ready to Go Audio Book")

	fmt.Println("Book: ", GetBook("1234-5678"))
	fmt.Println("CD: ", GetCD("1234-5678"))
}

func cacheGet(key string) string {
	return cache[key]
}

func cacheSet(key, value string) {
	if len(cache)+1 >= MaxCacheSize {
		return
	}
	cache[key] = value
}

// GetBook function
func GetBook(isbn string) string {
	return cacheGet(CacheKeyBook + isbn)
}

// SetBook function
func SetBook(isbn string, name string) {
	cacheSet(CacheKeyBook+isbn, name)
}

// GetCD function
func GetCD(sku string) string {
	return cacheGet(CacheKeyCD + sku)
}

// SetCD function
func SetCD(sku string, title string) {
	cacheSet(CacheKeyCD+sku, title)
}
