package main
import "fmt"
func fullJustify(words []string, maxWidth int) []string {
    lineWidthUsed := 0
	linesArray := [][]string {{}}
	lengthsArray := []int {}
	spaces := 0
	for _ , word := range words {
		if lineWidthUsed + len(word) + spaces <= maxWidth  {
			linesArray[len(linesArray) - 1] = append(linesArray[len(linesArray) - 1], word)
			lineWidthUsed += len(word)
			spaces ++
		} else {
			lengthsArray = append(lengthsArray, lineWidthUsed)
			linesArray = append(linesArray, []string{word})
			lineWidthUsed = len(word)
			spaces =1
		}
		
	}
	lengthsArray = append(lengthsArray, lineWidthUsed)

	return linesFill(linesArray, lengthsArray, maxWidth)
}

func linesFill(linesArray [][]string , lengthsArray []int, maxWidth int) []string {
	lines := make([]string, len(linesArray))

	for i := 0; i < len(linesArray) - 1; i++ {

		line := linesArray[i]
		totalWords := len(line)
		diff := maxWidth - lengthsArray[i]

		innerSpaces , extraSpaces := 0 , 0
		if totalWords >1 {
			innerSpaces = diff / (totalWords - 1)
			extraSpaces = diff % (totalWords - 1)
		} else {
			innerSpaces = diff
			extraSpaces = 0
		}

		
		//lines[i] = ""
		if len(line) > 1 { 
			for j := 0; j < len(line)-1; j++ {
			
				lines[i] += line[j]
				if extraSpaces > 0  {
					lines[i] += " "
					extraSpaces --
				}
	
				for k := 0; k < innerSpaces; k++ {
					lines[i] += " "
				}
			}
			lines[i] += line[len(line) - 1]
		} else {
			lines[i] += line[0]
			for k := 0; k < innerSpaces; k++ {
				lines[i] += " "
			}
		}
		
		
	}
	// last line
	line := linesArray[len(linesArray) - 1]
	for _ , word := range line {
		lines[len(linesArray) - 1] +=  word
		if len(lines[len(linesArray) - 1]) < maxWidth { 
			lines[len(linesArray) - 1] += " "
		}
	}

	rem := maxWidth - len(lines[len(linesArray) - 1])
	for i := 0; i < rem; i++ {

		lines[len(linesArray) - 1] += " "
	}

	return lines
}

func main() {
    //words := []string{"What", "must", "be", "acknowledgment", "shall", "be"}
    words := []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}
    //words := []string{"This", "is", "an", "example", "of", "text", "justification."}
    maxWidth := 20

    lines := fullJustify(words, maxWidth)

    for _, line := range lines {
        fmt.Print(line)
		fmt.Println("|")
    }
}