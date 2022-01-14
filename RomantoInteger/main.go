package main

func main() {
	println(romanToInt("III"))
	println(romanToInt("IV"))
	println(romanToInt("IX"))
	println(romanToInt("LVIII"))
	println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		c := string(s[i])
		switch c {
		case "I":
			if i+1 < len(s) && string(s[i+1]) == "V" {
				res += 4
				i++
			} else if i+1 < len(s) && string(s[i+1]) == "X" {
				res += 9
				i++
			} else {
				res += 1
			}
		case "X":
			if i+1 < len(s) && string(s[i+1]) == "L" {
				res += 40
				i++
			} else if i+1 < len(s) && string(s[i+1]) == "C" {
				res += 90
				i++
			} else {
				res += 10
			}
		case "C":
			if i+1 < len(s) && string(s[i+1]) == "D" {
				res += 400
				i++
			} else if i+1 < len(s) && string(s[i+1]) == "M" {
				res += 900
				i++
			} else {
				res += 100
			}
		case "V":
			res += 5
		case "L":
			res += 50

		case "D":
			res += 500
		case "M":
			res += 1000
		}
	}
	return res
}
