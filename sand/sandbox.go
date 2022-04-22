/*package main

import (
	"fmt"
)


func main() {

	var width, height = 50, 58
	num := 10 //короткое объявление
	fmt.Printf("Hello world %d!\nWxH : %dx%d\n", num, width, height)

	height, width = 10, 20
	num = 100040
	fmt.Printf("Изменение переменных | %d!\nWxH : %dx%d\n", num, width, height)
	width, height, weight := 30, 20, 50

	fmt.Printf("Добавление новой переменной +  изменение\nWxH : %dx%d\n%d g\n", width, height, weight)
	//в го нету префиксного инкримента

	var utf8 rune = 2374 //rune символы из unicode/ utf-8
	fmt.Printf("%c\n", utf8)

	a1 := [...]int{1, 4, 2, 5} //определение размера массива при объявлении
	fmt.Println(a1)

	//слайсы
	buf5 := make([]int, 5, 7) //len = 5, capacity = 7
	fmt.Println(buf5, "; len =", len(buf5), "; capacity =", cap(buf5))
	buf5 = append(buf5, 6, 7) //можно добавить еще два элемента без проблем
	fmt.Println(buf5, "; len =", len(buf5), "; capacity =", cap(buf5))
	buf5 = append(buf5, 8, 10, 39) // если len = capacity, то есть аллоцированная память закончится, то будет создан новый слайс в два раза большей capacity
	fmt.Println(buf5, "; len =", len(buf5), "; capacity =", cap(buf5))

	//добавление другого слайса
	otherBuf := make([]int, 8)
	buf5 = append(buf5, otherBuf...) //... - оператор подставляет элементы другого слайса
	fmt.Println(buf5, "; len =", len(buf5), "; capacity =", cap(buf5))

	//можно скопировать слайс, при изменении элементов одного слайса в другом тоже поменяется
	//если слайс переполнится при добавлении, то создастя новый и связь с тем, скопированным слайсом, будет утеряна

	ints := []int{1, 2, 3, 4}
	copy(ints[1:3], []int{5, 6}) //копирование в существующий слайс
	fmt.Println(ints)

	//хеш-таблица
	profile := make(map[string]string, 10) //с нужной ёмкостью
	fmt.Printf("%d\n", len(profile))

	var user map[string]string = map[string]string{
		"name":     "Vasily",
		"lastName": "Romanov",
	}
	mapLenght := len(user)
	fmt.Printf("%d %v\n", mapLenght, user)

	//проверка на существование ключа
	_, mNameExist := user["middleName"]
	fmt.Println("mNameExist", mNameExist)

	delete(user, "lastName")

	//в условных операторах есть блок инициализации
	if keyValue, keyExist := user["name"]; keyExist {
		fmt.Println("name =", keyValue)
	}

	//range в циклах
	sl := []int{8, 53, 4}
	for idx, val := range sl {
		fmt.Println("range slice by idx-value", idx, val)
	}

	person := map[int]string{1: "Vasily", 2: "Romanov"}
	for key, val := range person {
		fmt.Println("range map by key-val", key, val)
	}
}
*/

package main

import "fmt"

func test() int {
	x := 0
	defer func() {
		fmt.Println(x, "bef")
		x++
		fmt.Println(x, "aft")
	}()

	x++
	return x
}

func main() {
	fmt.Println(test())
}
