package main

import (
	"fmt"
	"sort"
	"strings"
)

func FindAnagram(str1, str2 string) bool {
	str1Slice := strings.Split(str1, "")
	str2Slice := strings.Split(str2, "")

	sort.Slice(str1Slice, func(i, j int) bool {
		return str1Slice[i] < str1Slice[j]
	})
	sort.Slice(str2Slice, func(i, j int) bool {
		return str2Slice[i] < str2Slice[j]
	})

	sortstr1 := strings.Join(str1Slice, "")
	sortstr2 := strings.Join(str2Slice, "")

	if sortstr1 == sortstr2 {
		return true
	}
	return false
}

func GroupAnagram(Str []string) [][]string {
	Arr_2D := [][]string{}
	for len(Str) > 0 {
		Str_After := []string{}
		Arr_1D := []string{}
		for j := 0; j < len(Str); j++ {
			hasil := FindAnagram(Str[0], Str[j])
			if hasil == true {
				Arr_1D = append(Arr_1D, Str[j])
			} else {
				Str_After = append(Str_After, Str[j])
			}
		}
		Str = Str_After

		Group_Angrm := [][]string{Arr_1D}
		Arr_2D = append(Arr_2D, Group_Angrm...)
	}

	return Arr_2D
}

func main() {
	fmt.Println(GroupAnagram([]string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}))
}
