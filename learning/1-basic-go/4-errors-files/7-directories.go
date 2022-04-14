package main

import (
	"fmt"
	"log"
	"os"
)

// MkdirAll Create directory
// func main() {
// 	// Stat() returns a FileInfo which implements methods that describe a file
// 	// func Stat(name string) (FileInfo, error)
// 	d, err := os.Stat("subdir")
// 	fmt.Println("error returned by os.Stat() is:", err)
// 	// if err is nil, then os.Stat() found a file/directory with the same name
// 	if err == nil {
// 		fmt.Println(d.Mode(), d.IsDir())
// 		log.Fatal("file/directory name already exists")
// 	}
// 	if errors.Is(err, os.ErrNotExist) {
// 		// func MkdirAll(path string, perm FileMode) error
// 		err := os.MkdirAll("subdir", 0777)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Println("subdir directory created")
// 	}
// }

// MkdirAll Nested directories
// func main() {
// 	//func Join(elem ...string) string
// 	p := filepath.Join("../test", "subdir1", "subdir2")
// 	err := os.MkdirAll(p, 0777)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(p, "nested directory created")
// }

// ReadDir List files
func main() {
	// formerly ioutil.ReadDir()
	// func ReadDir(name string) ([]DirEntry, error)
	ls, err := os.ReadDir("../")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range ls {
		fmt.Println(f.Name(), f.IsDir())
	}
}
