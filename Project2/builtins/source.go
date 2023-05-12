package builtins

import (
"bufio"
"fmt"
"io"
"os"
"os/exec"
"strings"
)

// SourceFile reads and executes commands from the specified file.
func SourceFile(filename string) error {
file, err := os.Open(filename)
if err != nil {
return err
}
defer file.Close()

reader := bufio.NewReader(file)
for {
line, err := reader.ReadString('\n')
if err == io.EOF {
break
}
if err != nil {
return err
}

line = strings.TrimSpace(line)
if len(line) == 0 || strings.HasPrefix(line, "#") {
continue
}

args := strings.Split(line, " ")
name, args := args[0], args[1:]

cmd := exec.Command(name, args...)
cmd.Stderr = os.Stderr
cmd.Stdout = os.Stdout

if err := cmd.Run(); err != nil {
return fmt.Errorf("error executing command '%s': %v", line, err)
}
}

return nil
}
