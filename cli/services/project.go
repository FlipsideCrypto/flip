package services

import (
	"fmt"
	"os"
	"os/exec"
)

type Project struct {
	Name string
}

func NewProject(name string) (*Project, error) {
	return &Project{Name: name}, nil
}

const DBT_PROJECT_TEMPLATE_URL = "https://github.com/jfmyers/flip-starter-project"
const DBT_PROJECT_TEMPLATE_BRANCH = "main"

func (p *Project) Init() error {
	projectName := p.Name
	if _, err := os.Stat("./" + projectName); !os.IsNotExist(err) {
		fmt.Println(fmt.Sprintf("🙅 Project `%s` already exists in this directory!", projectName))
		return nil
	}

	git_clone_args := []string{"clone", "--depth=1", "--branch", DBT_PROJECT_TEMPLATE_BRANCH, DBT_PROJECT_TEMPLATE_URL, projectName}
	cmdh := exec.Command("git", git_clone_args...)
	cloneStdout, cloneErr := cmdh.Output()

	if cloneErr != nil {
		fmt.Println("❌ Error cloning default project: ")
		fmt.Println(cloneStdout)
		fmt.Println(cloneErr.Error())
		return nil
	}

	rmCmd := exec.Command("rm", "-rf", "./"+projectName+"/.git")
	_, rmErr := rmCmd.Output()
	if rmErr != nil {
		fmt.Println("❌ Error removing .git config", rmErr.Error())
		return nil
	}
	return nil
}
