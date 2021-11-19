package services

import (
	"flip/api"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type DBT struct {
	Port  string
	Datax *api.DataExchangeCredsResp
}

func NewDBT(datax *api.DataExchangeCredsResp, port string) (*DBT, error) {
	return &DBT{Datax: datax, Port: port}, nil
}

func (d *DBT) Console() error {
	fmt.Println("⌨️ Spinning up an interactive DBT console")
	path, _ := os.Getwd()
	mount := filepath.Join(path, "sql_models")

	docker_args := []string{
		"run",
		"-it",
		"--env", "FLIP_USERNAME=" + d.Datax.Username,
		"--env", "FLIP_PASSWORD=" + d.Datax.Password,
		"--env", "FLIP_REGION=" + d.Datax.Region,
		"--env", "FLIP_DATABASE=" + d.Datax.Database,
		"--env", "FLIP_WAREHOUSE=" + d.Datax.Warehouse,
		"--env", "FLIP_ROLE=" + d.Datax.Role,
		"-v", mount + ":/sql_models",
		d.Datax.DockerImage, "/support/dbt_console.sh",
	}
	cmdh := exec.Command("docker", docker_args...)
	// stdout, err := cmdh.Output()
	cmdh.Stdout = os.Stdout
	cmdh.Stdin = os.Stdin
	cmdh.Stderr = os.Stderr
	cmdh.Run()
	return nil
}

func (d *DBT) Docs() error {
	path, _ := os.Getwd()
	mount := filepath.Join(path, "sql_models")
	docker_args := []string{
		"run",
		"-it",
		"-p", d.Port + ":" + d.Port,
		"--env", "FLIP_USERNAME=" + d.Datax.Username,
		"--env", "FLIP_PASSWORD=" + d.Datax.Password,
		"--env", "FLIP_REGION=" + d.Datax.Region,
		"--env", "FLIP_DATABASE=" + d.Datax.Database,
		"--env", "FLIP_WAREHOUSE=" + d.Datax.Warehouse,
		"--env", "FLIP_ROLE=" + d.Datax.Role,
		"--env", "DBT_DOCS_PORT=" + d.Port,
		"-v", mount + ":/sql_models",
		d.Datax.DockerImage, "/support/dbt_docs.sh",
	}
	cmdh := exec.Command("docker", docker_args...)
	// stdout, err := cmdh.Output()
	cmdh.Stdout = os.Stdout
	cmdh.Stdin = os.Stdin
	cmdh.Stderr = os.Stderr
	cmdh.Run()
	return nil
}

func (d *DBT) ResetEnv() error {
	cmdh := exec.Command("docker", "rm", d.Datax.DockerImage, "--force")
	cmdh.Stdout = os.Stdout
	cmdh.Stdin = os.Stdin
	cmdh.Stderr = os.Stderr
	cmdh.Run()

	cmdh = exec.Command("docker", "pull", d.Datax.DockerImage)
	cmdh.Stdout = os.Stdout
	cmdh.Stdin = os.Stdin
	cmdh.Stderr = os.Stderr
	cmdh.Run()
	return nil
}
