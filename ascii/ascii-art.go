package ascii

import (
	"fmt"
	"os"
	"strings"
)

func Generate(input, banner string) (string, error) {
	var asciiArt strings.Builder

	bytes, err := os.ReadFile(fmt.Sprintf("ascii/fonts/%s.txt", banner))
	if err != nil {
		return "", fmt.Errorf("error reading file: %v", err)
	}

	var lines []string
	if banner == "thinkertoy" {
		lines = strings.Split(string(bytes), "\r\n")
	} else {
		lines = strings.Split(string(bytes), "\n")
	}

	words := strings.Split(input, "\n")

	if banner == "standard" || banner == "shadow" || banner == "thinkertoy" || banner == "colossal" {
		for i, word := range words {
			if word == "" {
				if i < len(words)-1 {
					asciiArt.WriteString("\n")
				}
				continue
			}
			for h := 1; h < 9; h++ {
				for _, l := range word {
					for lineIndex, line := range lines {
						if lineIndex == (int(l)-32)*9+h {
							asciiArt.WriteString(line)
						}
					}
				}
				asciiArt.WriteString("\n")
			}
		}
	} else {
		var lineCount int
		var offset int

		bannerDetails, err := GetBannerDetails(banner)
		if err != nil {
			return "", fmt.Errorf("unknown banner: %v", err)
		}

		lineCount = bannerDetails.LineCount
		offset = bannerDetails.Offset

		for _, word := range words {
			arr := []rune(word)
			printArt(&asciiArt, arr, lines, lineCount, offset)
		}
	}

	return asciiArt.String(), nil
}

func printArt(asciiArt *strings.Builder, arr []rune, lines []string, lineCount int, offset int) {
	for line := 1; line <= lineCount; line++ {
		for _, r := range arr {
			skip := (r * rune(lineCount)) - rune(offset)
			asciiArt.WriteString(lines[line+int(skip)-1])
		}
		asciiArt.WriteRune('\n')
	}
}
