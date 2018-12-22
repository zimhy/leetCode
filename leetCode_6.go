package main

import "fmt"

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	runes := []rune(s)
	var splitedRunes [][]rune
	start := 0
	end := 2 * (numRows - 1)
	length := len(s)
	for {
		if end >= length {
			end = length
			splitedRunes = append(splitedRunes, runes[start:end])
			break
		} else {
			splitedRunes = append(splitedRunes, runes[start:end])
			start += 2 * (numRows - 1)
			end += 2 * (numRows - 1)
		}
	}
	//for j := 0; j < len(splitedRunes); j++ {
	//	fmt.Println(string(splitedRunes[j]))
	//}

	var result []rune
	/*
		PAYP
	    ALIS
	    HIRI
	    NG
	*/
	for col := 0; col <= numRows-1; col++ {
		//fmt.Println("---------------------")
		for line := 0; line < len(splitedRunes); line++ {
			if col == 0 || col == numRows-1 {
				if len(splitedRunes[line]) > col {
					//fmt.Println(line, col)
					result = append(result, splitedRunes[line][col])
				}
			} else {
				if len(splitedRunes[line]) > col {
					//fmt.Println(line, col)
					result = append(result, splitedRunes[line][col])
				}
				if len(splitedRunes[line]) > 2*(numRows-1)-col {
					//fmt.Println(line, 2*(numRows-1)-col-1)
					result = append(result, splitedRunes[line][2*(numRows-1)-col])
				}
			}
		}
	}
	//fmt.Printf("%v", string(result))
	return string(result)
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
}
