package main

import (
	_ "embed"
	"errors"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
	"time"
)

//go:embed assets/justfile
var Justfile string

//go:embed assets/.gitignore
var Gitignore string

//go:embed assets/LICENSE
var License string

//go:embed assets/example.go
var ExampleGo string

//go:embed assets/example_test.go
var ExampleTestGo string

var ErrExists = errors.New("directory or file already exists")

func main() {
	args := os.Args
	if len(args) < 2 {
		bail("no project name provided, use `golaunch <name>`")
	}
	projectName := args[1]

	// check if target dir exists
	err := dirExists(projectName)
	catch(err)
	println("initializing project: " + projectName)

	// create project dir
	err = os.Mkdir(projectName, os.ModePerm)
	catch(err)
	err = os.Chdir(projectName)
	catch(err)

	// git init
	err = exec.Command("git", "init").Run()
	catch(err)

	// go mod init
	err = exec.Command("go", "mod", "init", projectName).Run()
	catch(err)

	// write justfile
	err = os.WriteFile("justfile", []byte(Justfile), os.ModePerm)
	catch(err)

	// write .gitignore
	err = os.WriteFile(".gitignore", []byte(Gitignore), os.ModePerm)
	catch(err)

	// write README
	err = os.WriteFile("README.md", []byte("# "+projectName+"\n"), os.ModePerm)
	catch(err)

	// write LICENSE
	err = os.WriteFile("LICENSE", []byte(license()), os.ModePerm)
	catch(err)

	// create app dir
	err = os.Mkdir("app", os.ModePerm)
	catch(err)
	err = os.Chdir("app")
	catch(err)

	// write main.go
	err = os.WriteFile("main.go", []byte(ExampleGo), os.ModePerm)
	catch(err)

	// write main_test.go
	err = os.WriteFile("main_test.go", []byte(ExampleTestGo), os.ModePerm)
	catch(err)

	println("project created!")
}

func dirExists(name string) error {
	_, err := os.Stat(name)
	if err == nil {
		return ErrExists
	}
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	return err
}

func username() string {
	user, err := user.Current()
	if err != nil {
		return "unknown-username"
	}
	return user.Username
}

func license() string {
	year := strconv.Itoa(time.Now().Year())
	s := strings.ReplaceAll(License, "{{YEAR}}", year)
	return strings.ReplaceAll(s, "{{USER}}", username())
}

func catch(err error) {
	if err != nil {
		bail(err.Error())
	}
}

func bail(msg string) {
	println(msg)
	os.Exit(1)
}
