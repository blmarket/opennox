package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, "Usage: %s <go_file> <function_name> <variable_name>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Example: %s audio_impl.go sub_452050 v1\n", os.Args[0])
		os.Exit(1)
	}

	filename := os.Args[1]
	functionName := os.Args[2]
	variableName := os.Args[3]

	if err := processFile(filename, functionName, variableName); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func processFile(filename, functionName, variableName string) error {
	// Read the file
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// Find the target function and process it
	modified := false
	inTargetFunction := false
	braceCount := 0

	for i, line := range lines {
		// Check if we're entering the target function
		if !inTargetFunction && strings.Contains(line, "func") && strings.Contains(line, functionName) {
			inTargetFunction = true
			braceCount = 0
			continue
		}

		if inTargetFunction {
			// Count braces to determine when the function ends
			braceCount += strings.Count(line, "{") - strings.Count(line, "}")

			// Process the line for unsafe.Add calls
			newLine := processLine(line, variableName)
			if newLine != line {
				lines[i] = newLine
				modified = true
			}

			// Check if we've reached the end of the function
			if braceCount <= 0 && strings.Contains(line, "}") {
				inTargetFunction = false
			}
		}
	}

	if !modified {
		fmt.Printf("No modifications made to function %s\n", functionName)
		return nil
	}

	// Write the modified content back to the file
	outFile, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)
	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}
	writer.Flush()

	fmt.Printf("Successfully processed %s, function %s, variable %s\n", filename, functionName, variableName)
	return nil
}

func processLine(line, variableName string) string {
	// Pattern to match unsafe.Add(unsafe.Pointer(variableName), offset)
	// This regex captures the offset part
	pattern := fmt.Sprintf(`unsafe\.Add\(unsafe\.Pointer\(%s\),\s*([^)]+)\)`, regexp.QuoteMeta(variableName))
	re := regexp.MustCompile(pattern)

	return re.ReplaceAllStringFunc(line, func(match string) string {
		// Extract the offset from the match
		submatches := re.FindStringSubmatch(match)
		if len(submatches) != 2 {
			return match // Return original if we can't parse
		}

		offsetStr := strings.TrimSpace(submatches[1])
		offset := parseOffset(offsetStr)
		if offset < 0 {
			return match // Return original if we can't parse offset
		}

		// Calculate field number (4 bytes per field for 32-bit)
		fieldNum := offset / 4

		// Return the field access
		return fmt.Sprintf("&%s.field_%d", variableName, fieldNum)
	})
}

func parseOffset(offsetStr string) int {
	// Handle simple integer
	if val, err := strconv.Atoi(offsetStr); err == nil {
		return val
	}

	// Handle multiplication like "4*26"
	parts := strings.Split(offsetStr, "*")
	if len(parts) == 2 {
		left := strings.TrimSpace(parts[0])
		right := strings.TrimSpace(parts[1])

		leftVal, err1 := strconv.Atoi(left)
		rightVal, err2 := strconv.Atoi(right)

		if err1 == nil && err2 == nil {
			return leftVal * rightVal
		}
	}

	return -1 // Cannot parse
}
