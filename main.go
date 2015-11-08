package main
import (
	. "github.com/ArthurHlt/mattermost-cf-integrator/mci"
	"fmt"
	"os"
	"path"
	"os/exec"
)
func main() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	configFilePath := path.Join(wd, "config", "config.json")
	config, err := ExtractConfig(configFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	err = CloudifyConfig(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	PushConfig(config, configFilePath)
	mattermostExec := exec.Command(path.Join(wd, "bin", "platform"))
	mattermostExec.Stdout = os.Stdout
	mattermostExec.Stderr = os.Stderr
	err = mattermostExec.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

}