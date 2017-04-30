package main

import(
  "fmt"
  "bufio"
  "os"
)

func main() {
  lines := read("./large.ces")
  // lines := read("./test.txt")
  for {
    origin := lines[0]
    lines = remove(lines, 0)
    anagramCandidateList := anagramList(origin)
    anagrams := []string {origin}
    for _, object := range anagramCandidateList {
      index := search(lines, object)
      if index != -1 {
        anagrams = append(anagrams, lines[index])
        lines = remove(lines, index)
      }
    }
    if len(anagrams) > 1 {
      fmt.Println(anagrams)
    }
    if len(lines) == 0 {
      break
    }
  }
}

func search(words []string, object string) int {
  for i, word := range words {
    if object == word {
      return i
    }
  }
  return -1
}

func remove(texts []string, index int) []string {
    result := []string {}
    i := 0
    for _, text := range texts {
        if i != index {
            result = append(result, text)
        }
        i += 1
    }
    return result
}

func anagramList(word string) []string {
  if len(word) == 1 {
    return []string {word}
  }

  ret := []string {}
  for i := 0; i < len(word); i++ {
    words := anagramList(word[0:i] + word[i+1:len(word)])
    for _, w := range(words) {
      ret = append(ret, word[i:i+1] + w)
    }
  }
  return ret
}

func read(filePath string) []string  {
  f, err := os.Open(filePath)
  if err != nil {
    fmt.Fprintf(os.Stderr, "File %s could not read: %v\n", filePath, err)
    os.Exit(1)
  }

  defer f.Close()

  lines := []string {}
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  if serr := scanner.Err(); serr != nil {
    fmt.Fprintf(os.Stderr, "File %s scan error: %v\n", filePath, err)
  }

  return lines
}