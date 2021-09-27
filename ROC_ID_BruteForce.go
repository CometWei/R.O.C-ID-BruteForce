package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {

	var Id string = ""
	print("Enter the R.O.C ID：")
	fmt.Scanf("%s", &Id) //Test Data:A190355428
	Id = strings.ToUpper(Id)
	id_len := len([]rune(Id))

	r, _ := regexp.Compile("[^A-Za-z0-9]")
	var id_range int = 1

	if r.MatchString(Id) {
		fmt.Printf("It's strings error.\n")
	} else if id_len < 10 {

		if id_len == 0 {

			id_city_sex := GetCityRandomString(1) + GetSexRandomString(1)
			for id_range < 100000000 {
				Id = id_city_sex + fmt.Sprintf("%08d", id_range)
				ID_check(Id)
				id_range++
			}

		} else if id_len == 1 {

			id_city_sex := Id + GetSexRandomString(1)
			for id_range < 100000000 {

				Id = id_city_sex + fmt.Sprintf("%08d", id_range)
				ID_check(Id)
				id_range++

			}

		} else {
			id_len = 10 - id_len
			id_city_sex := Id

			switch id_len {
			case 8:
				for id_range < 100000000 {

					Id = id_city_sex + fmt.Sprintf("%08d", id_range)
					ID_check(Id)
					id_range++

				}
			case 7:
				for id_range < 10000000 {

					Id = id_city_sex + fmt.Sprintf("%07d", id_range)
					ID_check(Id)
					id_range++

				}
			case 6:
				for id_range < 1000000 {

					Id = id_city_sex + fmt.Sprintf("%06d", id_range)
					ID_check(Id)
					id_range++

				}
			case 5:
				for id_range < 100000 {

					Id = id_city_sex + fmt.Sprintf("%05d", id_range)
					ID_check(Id)
					id_range++

				}
			case 4:
				for id_range < 10000 {

					Id = id_city_sex + fmt.Sprintf("%04d", id_range)
					ID_check(Id)
					id_range++

				}
			case 3:
				for id_range < 1000 {

					Id = id_city_sex + fmt.Sprintf("%03d", id_range)
					ID_check(Id)
					id_range++

				}
			case 2:
				for id_range < 100 {

					Id = id_city_sex + fmt.Sprintf("%02d", id_range)
					ID_check(Id)
					id_range++

				}
			case 1:
				for id_range < 10 {

					Id = id_city_sex + fmt.Sprintf("%d", id_range)
					ID_check(Id)
					id_range++

				}
			}
		}

	} else if id_len == 10 {
		ID_check(Id)
	} else if id_len > 10 {
		fmt.Printf("%s Is so long.\n", Id)
	}

}

func GetSexRandomString(l int) string {
	str := "12"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func GetCityRandomString(l int) string {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func ID_check(id string) {
	id_check := ROC_ID_check(id)
	if id_check == true {
		fmt.Printf("%s,%v\n", id, id_check)
	}
}

func ROC_ID_check(id string) bool {

	city_match, _ := regexp.Compile("[A-Za-z]")
	city := city_match.FindString(id)

	id_code_match, _ := regexp.Compile("([0-9]+)")
	id_code_string := id_code_match.FindString(id)
	id_code, _ := strconv.Atoi(id_code_string)

	city_code := city_num(city)

	var code [11]int

	code[0] = city_code / 10
	code[1] = city_code % 10
	code[2] = id_code / 100000000
	code[3] = (id_code / 10000000) % 10
	code[4] = (id_code / 1000000) % 10
	code[5] = (id_code / 100000) % 10
	code[6] = (id_code / 10000) % 10
	code[7] = (id_code / 1000) % 10
	code[8] = (id_code / 100) % 10
	code[9] = (id_code / 10) % 10
	code[10] = id_code % 10

	if ((code[2] == 1) || (code[2] == 2)) == false {
		fmt.Printf("Sex is error.\n")
		return false
	}

	var math_count int = 9
	var code_count int = 1
	var code_math [10]int

	for math_count != 0 {
		code_math[code_count] = (code[code_count] * math_count)
		code_count++
		math_count--
	}

	code_count = 0
	var Check_num int = 0
	for code_count != 10 {
		Check_num = Check_num + int(code_math[code_count])
		code_count++
	}

	Check_num = Check_num + code[0] + code[10]

	Check_num = Check_num % 10

	if Check_num == 0 {
		return true
	} else {
		return false
	}

}

func city_num(city string) int { //ID City Convert

	switch city {
	case "A":
		return 10
	case "B":
		return 11
	case "C":
		return 12
	case "D":
		return 13
	case "E":
		return 14
	case "F":
		return 15
	case "G":
		return 16
	case "H":
		return 17
	case "I":
		return 34
	case "J":
		return 18
	case "K":
		return 19
	case "L":
		return 20
	case "M":
		return 21
	case "N":
		return 22
	case "O":
		return 35
	case "P":
		return 23
	case "Q":
		return 24
	case "R":
		return 25
	case "S":
		return 26
	case "T":
		return 27
	case "U":
		return 28
	case "V":
		return 29
	case "W":
		return 32
	case "X":
		return 30
	case "Y":
		return 31
	case "Z":
		return 33
	}
	return 0
}
