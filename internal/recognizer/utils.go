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

		currentValue, _ := Selectlist.Load(prompt)

		// Create a new HashData entry.
		newHashData := HashData{
			Hash: hash,
			X:    x,
			Y:    y,
		}

		// Append the new HashData entry to the existing slice (or create a new slice if it doesn't exist).
		var updatedValue []HashData
		if currentValue != nil {
			updatedValue = append(currentValue.([]HashData), newHashData)
		} else {
			updatedValue = []HashData{newHashData}
		}

		// Store the updated value back into the 'Selectlist'.
		Selectlist.Store(prompt, updatedValue)

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

		// Load the current value associated with 'prompt'.
		currentValue, _ := Hashlist.Load(prompt)

		// Append the new hash to the existing slice (or create a new slice if it doesn't exist).
		var updatedValue []string
		if currentValue != nil {
			updatedValue = append(currentValue.([]string), hash)
		} else {
			updatedValue = []string{hash}
		}

		// Store the updated value back into the 'Hashlist'.
		Hashlist.Store(prompt, updatedValue)
		count++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return count, nil
}

func LoadAnswer(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "|")

		if len(parts) != 2 {
			continue
		}

		Answerlist.Store(parts[0], parts[1])
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return count, nil
}
