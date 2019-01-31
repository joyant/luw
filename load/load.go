package load

import (
    "encoding/xml"
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"
)

type Item struct {
    Word     string `xml:"word"`
    Trans    string `xml:"trans"`
    Phonetic string `xml:"phonetic"`
}

type WordBook struct {
    XMLName xml.Name `xml:"wordbook"`
    Items   []Item   `xml:"item"`
}

func Load() []Item {
    c, err := CurrentDir()
    if err != nil {
        panic(err)
    }

    dir := filepath.Join(c, "resource")
    fs, err := ioutil.ReadDir(dir)
    if err != nil {
        panic(err)
    }

    ret := make([]Item, 0)
    for _, info := range fs {
        if ! info.IsDir() {
            f := filepath.Join(dir, info.Name())
            b, err := ioutil.ReadFile(f)
            if err != nil {
                panic(err)
            }
            var wordBook WordBook
            err = xml.Unmarshal(b, &wordBook)
            if err != nil {
                panic(err)
            }
            for _,item := range wordBook.Items {
                if isPureWord(item.Word) {
                    ret = append(ret,item)
                }
            }
        }
    }

    return ret
}

func isPureWord(s string) bool {
    for _,w := range s {
        if w < 'a' || w > 'z' {
            return false
        }
    }
    return true
}

func CurrentDir() (string, error) {
    dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
        return "", err
    }
    dir = strings.Replace(dir, `\\`, "/", -1)
    return dir, nil
}
