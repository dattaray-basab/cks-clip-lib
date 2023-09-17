package transform


import (
   "io/fs"
   "path/filepath"
)


func GetFileFromDir() {
   filepath.WalkDir(".", func(s string, d fs.DirEntry, e error) error {
    if e != nil { return e }
    if ! d.IsDir() {
       println(s)
    }
    return nil
 })
}