package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

type singletonDatabase struct {
	capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

// sync.Once / init func - thread safety
// sync.Once - lazy loading

func readData(path string) (map[string]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := map[string]int{}

	for scanner.Scan() {
		k := scanner.Text()
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		result[k] = v
	}

	return result, nil
}

var once sync.Once
var instance *singletonDatabase

func GetSingletonDatabase() *singletonDatabase {
	once.Do(func() {
		caps, e := readData("./capitals.txt")
		db := singletonDatabase{caps}
		if e == nil {
			db.capitals = caps
		}
		instance = &db
	})

	return instance
}

func main() {
	db1 := GetSingletonDatabase()
	fmt.Println(db1.GetPopulation("ABC"))

	db2 := GetSingletonDatabase()
	fmt.Println(db2.GetPopulation("ABC"))

	fmt.Println(db1)
	fmt.Println(db2)
}
