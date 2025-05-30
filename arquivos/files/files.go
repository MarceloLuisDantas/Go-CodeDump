package files

import (
	"bufio"
	"log"
	"os"
)

type File struct {
	Linhas []string
	len    uint
}

func (f *File) Append(s string) {
	f.Linhas = append(f.Linhas, s)
	f.len += 1
}

func (f *File) Set(i uint, v string) {
	f.Linhas[i] = v
}

// func (f *File) Remove(i uint) (string, error) {
// 	if i >= f.len {
// 		return "", errors.New("Empty file")
// 	}

// 	f.Linhas.
// }

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Load(path string) File {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	arquivo := File{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		arquivo.Append(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return arquivo
}
