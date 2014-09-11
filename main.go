package main

import (
  "code.google.com/p/go.tools/go/vcs"
  "encoding/json"
  "flag"
  "fmt"
  "os"
)

var (
  importPathFlag = flag.String("import-path", "", "detect the repo root for the given import path")
)

func main() {
  flag.Parse()

  if *importPathFlag != "" {
    repoRoot, err := vcs.RepoRootForImportPath(*importPathFlag, false)
    if err == nil {
      fmt.Fprintln(os.Stdout, toJson(repoRoot))
      os.Exit(0)
    } else {
      fmt.Fprintln(os.Stderr, err.Error())
      os.Exit(1)
    }
  }
}

func toJson(v interface{}) string {
  out, _ := json.MarshalIndent(v, "", "  ")
  return string(out)
}
