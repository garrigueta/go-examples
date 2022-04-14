package main

import (
	"log"
	"os"
)

// Create file
// func main() {
// 	f, err := os.Create("create.txt")
// 	defer f.Close()
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	fmt.Println(f)
// }

// Open and Close file
// the closer function
// func closer(f *os.File) error {
// 	f.Close()
// 	fmt.Println(f.Name(), "successfully closed")
// 	return nil
// }
// func main() {
// 	f, err := os.Open("file.txt")
// 	defer closer(f)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	fmt.Println("file successfully opened:", f.Name())
// }

// Remove (delete) file
// func main() {
// 	err := os.Remove("del.text")
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	fmt.Println("file successfully removed")
// }

// Copy file
// func main() {
// 	src, err := os.Open("src.txt")
// 	defer src.Close()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// the flags allow the program to create it, if the file does not exist
// 	dst, err := os.OpenFile("dst.txt", os.O_RDWR|os.O_CREATE, 0755)
// 	defer dst.Close()
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	// now we can copy
// 	b, err := io.Copy(dst, src)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	fmt.Println(b)
// }

// Rename (move) a file
// func main() {
// 	oldPath := "file.txt"
// 	newPath := "./new/new.txt"
// 	err := os.Rename(oldPath, newPath)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// Truncate a file
// myFile.txt originally contains "test file contents"
// func main() {
// 	err := os.Truncate("myFile.txt", 10)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// FileInfo
// Stat returns a FileInfo structure that describes a file
func main() {
	f, err := os.Stat("myFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("file name:", f.Name())
	log.Println("file size:", f.Size())
	log.Println("file permissions:", f.Mode())
}
