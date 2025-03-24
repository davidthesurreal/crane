package main 


import (
  "github.com/enescakir/emoji"
  "os"
  "errors"
  "fmt"
  "github.com/fatih/color"
)

func main() {

  Text("Hello World!", "error")
}

func StringInput(prompt string, emojiIsEnabled ...bool) string {
  var emojiEnabled bool

  if len(emojiIsEnabled) > 0 {
    emojiEnabled = emojiIsEnabled[0]
  }
  var input string

  if emojiEnabled {
    fmt.Printf("%v %v: ", emoji.Pencil, prompt)
  }else {
  fmt.Printf("%v:", prompt)
  }
  fmt.Scan(&input)

  return input
}

func GetArg(index int) (string, error){
  args := os.Args

  if (len(args) > index) {
    return args[index], nil
  }

  return "", errors.New("Not enough arguments!")
}

func Text(message string, mtype string) error {
  switch mtype {
  case "success":
    color.Green(message)
    return nil

  case "error":
    color.Red(message)
    return nil

  case "info":
    color.Blue(message)
    return nil

  case "warning": 
    color.Yellow(message)
    return nil

  default:
    return errors.New("Invalid message type, message types are only [success, error, info and warning].")
  }

}
