package recognizer

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func LoadHashSelect(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")

		if len(parts) != 4 {
			continue
		}

		hash, prompt, _x, _y := parts[0], parts[1], parts[2], parts[3]

		x, _ := strconv.Atoi(_x)
		y, _ := strconv.Atoi(_y)

		Selectlist[prompt] = append(Selectlist[prompt], HashData{
			X:    x,
			Y:    y,
			Hash: hash,
		})
		count++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return count, nil
}

func LoadHash(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")

		if len(parts) != 2 {
			continue
		}

		hash, prompt := parts[0], parts[1]

		if _, exists := Hashlist[prompt]; !exists {
			Hashlist[prompt] = []string{}
		}

		Hashlist[prompt] = append(Hashlist[prompt], hash)
		count++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return count, nil
}
