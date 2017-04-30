package main

import(
  "fmt"
  "bufio"
  "os"
  "strings"
  "sort"
)

func main() {
  lines := read("./10196pla.ces")
  m := map[string]string {}
  for _, line := range lines {
    m[line] = sortChars(line)
  }

  a := List{}
  for k, v := range m {
    e := Entry{k, v}
    a = append(a, e)
  }
  sort.Sort(a)

  current := ""
  previous := ""
  anagrams := []string {}
  for _, entry := range a {
    current = entry.value
    if current == previous {
      anagrams = append(anagrams, entry.name)
    } else {
      if len(anagrams) > 1 {
        fmt.Println(anagrams)
      }
      anagrams = []string {}
      anagrams = append(anagrams, entry.name)
    }
    previous = current
  }
}

type Entry struct {
    name  string
    value string
}
type List []Entry

func (l List) Len() int {
    return len(l)
}

func (l List) Swap(i, j int) {
    l[i], l[j] = l[j], l[i]
}

func (l List) Less(i, j int) bool {
    if l[i].value == l[j].value {
        return (l[i].name < l[j].name)
    } else {
        return (l[i].value < l[j].value)
    }
}

func sortChars(word string) string {
  chars := strings.Split(word, "")
  sort.Strings(chars)
  return strings.Join(chars, "")
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