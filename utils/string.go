package utils

import (
	"strconv"
	"strings"
)

func RemoveColor(ss string) string {
	for i := 0; i <= 9; i++ {
		ss = strings.ReplaceAll(ss, "§"+strconv.Itoa(i), "")
	}
	ss = strings.ReplaceAll(ss, "§a", "")
	ss = strings.ReplaceAll(ss, "§b", "")
	ss = strings.ReplaceAll(ss, "§c", "")
	ss = strings.ReplaceAll(ss, "§d", "")
	ss = strings.ReplaceAll(ss, "§e", "")
	ss = strings.ReplaceAll(ss, "§f", "")
	ss = strings.ReplaceAll(ss, "§g", "")
	ss = strings.ReplaceAll(ss, "§k", "")
	ss = strings.ReplaceAll(ss, "§o", "")
	ss = strings.ReplaceAll(ss, "§r", "")
	ss = strings.ReplaceAll(ss, "§l", "")
	//fmt.Println(ss)
	return ss
}

func StrpadLeft(str, pad string, count int) string {
	if len(str) >= count {
		return str
	}
	count = count - len(str)
	str = strings.Repeat(pad, count) + str
	return str
}
