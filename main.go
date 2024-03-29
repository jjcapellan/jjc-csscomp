package csscomp

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func minColor(hexCol string) string {
	if hexCol[1] == hexCol[2] && hexCol[3] == hexCol[4] && hexCol[5] == hexCol[6] {
		return "#" + string(hexCol[1]) + string(hexCol[3]) + string(hexCol[5])
	}
	return hexCol
}

func quitQuotes(str string) string {
	str = strings.ReplaceAll(str, "\"", "")
	return str
}

func quitQuoteIfNoSpace(str string) string {
	str = strings.TrimSpace(str)
	if strings.Index(str, " ") == -1 {
		return strings.ReplaceAll(str, "\"", "")
	}
	return str
}

func rmffQuotes(str string) string {
	rgx := regexp.MustCompile(`\s*\"[^\"]*\"\s*[\,;]`)
	str = rgx.ReplaceAllStringFunc(str, quitQuoteIfNoSpace)
	return str
}

func intToHex(str string) string {
	number, _ := strconv.Atoi(str)
	strHex := fmt.Sprintf("%x", number)
	if len(strHex) == 1 {
		strHex = "0" + strHex
	}
	return strHex
}

func rgbToHex(str string) string {
	str = strings.ReplaceAll(str, " ", "")
	str = strings.Split(str, "(")[1]
	str = strings.TrimRight(str, ")")
	strArr := strings.Split(str, ",")
	str = ""
	for i := range strArr {
		str += intToHex(strArr[i])
	}
	return "#" + str
}

// Removes zero units sufix
func rmZeroSufix(str string) string {
	rgx := regexp.MustCompile(`\s0[^0-9\s\;\,\.]{1,2}`)
	str = rgx.ReplaceAllString(str, " 0")

	rgx = regexp.MustCompile(`:0[^0-9\s\;\,\.]{1,2}`)
	str = rgx.ReplaceAllString(str, ":0")

	return str
}

func rmInnerSpaces(str string) string {
	return strings.ReplaceAll(str, " ", "")
}

// Deletes spaces for all text
func rmSpaces(str string) string {
	rgx := regexp.MustCompile(`\s*[\{\}\,;\=:\)\(\]\[\+\>\~]\s*`)
	str = rgx.ReplaceAllStringFunc(str, strings.TrimSpace)

	rgx = regexp.MustCompile(`\s\s+`)
	str = rgx.ReplaceAllString(str, " ")

	rgx = regexp.MustCompile(`\-\s+\.`)
	str = rgx.ReplaceAllStringFunc(str, rmInnerSpaces)

	return str
}

// Deletes all comments
func rmComments(str string) string {
	rgx := regexp.MustCompile(`\/\*.*\*\/`)
	str = rgx.ReplaceAllString(str, "")
	return str
}

// Deletes all line jumps
func rmJumps(str string) string {
	rgx := regexp.MustCompile(`[\n]`)
	str = rgx.ReplaceAllString(str, "")
	return str
}

// Deletes begin line spaces
func rmBeginSpaces(str string) string {
	rgx := regexp.MustCompile(`^\s*`)
	str = rgx.ReplaceAllString(str, "")
	return str
}

// Preserve first char
func rmLeadZeroFunc(str string) string {
	return string(str[0]) + " ."
}

// Deletes leading zeros
func rmLeadingZeros(str string) string {
	rgx := regexp.MustCompile(`[^0-9]0\.`)
	str = rgx.ReplaceAllStringFunc(str, rmLeadZeroFunc)
	return str
}

// Deletes quotes when possible
func rmQuotes(str string) string {
	rgx := regexp.MustCompile(`\[[^\[\]]+\=\"[^\"]*\"\]`)
	str = rgx.ReplaceAllStringFunc(str, quitQuoteIfNoSpace)

	rgx = regexp.MustCompile(`font-family:[^;]*;`)
	str = rgx.ReplaceAllStringFunc(str, rmffQuotes)

	rgx = regexp.MustCompile(`url\s*\(\s*\"[^\"]*\"\s*\)`)
	str = rgx.ReplaceAllStringFunc(str, quitQuotes)

	return str
}

// Converts rgb values to hexadecimal notation
func rmRgb(str string) string {
	rgx := regexp.MustCompile(`rgb\s*\([^,]+,[^,]+,[^,]+\)`)
	str = rgx.ReplaceAllStringFunc(str, rgbToHex)
	return str
}

// Converts hex colors to its shorthand
func minColors(str string) string {
	rgx := regexp.MustCompile(`\#[0-9a-fA-F]{6}`)
	str = rgx.ReplaceAllStringFunc(str, minColor)
	return str
}

// Deletes semicolon
func rmSemicolon(str string) string {
	rgx := regexp.MustCompile(`;\s*\}`)
	str = rgx.ReplaceAllString(str, "}")
	return str
}

func msToSecFunc(str string) string {
	numstr := strings.TrimRight(str, "ms")
	ms, _ := strconv.Atoi(numstr)
	secs := float32(ms) / float32(1000)
	secstr := fmt.Sprint(secs)
	return secstr + "s"
}

// Converts ms to s
func msToSecond(str string) string {
	rgx := regexp.MustCompile(`[0-9]+ms`)
	str = rgx.ReplaceAllStringFunc(str, msToSecFunc)
	return str
}

func fontWeightToNumber(str string) string {
	rgx := regexp.MustCompile(`weight:normal`)
	str = rgx.ReplaceAllString(str, "weight:400")

	rgx = regexp.MustCompile(`weight:bold`)
	str = rgx.ReplaceAllString(str, "weight:700")
	return str
}

// Get returns the minified css string
func Get(txt string) string {
	str := rmComments(txt)

	str = rmJumps(str)

	str = msToSecond(str)

	str = rmLeadingZeros(str)

	str = rmSpaces(str)

	str = rmZeroSufix(str)

	str = rmQuotes(str)

	str = rmSemicolon(str)

	str = rmRgb(str)

	str = minColors(str)

	str = fontWeightToNumber(str)

	return str
}
