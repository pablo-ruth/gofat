package profiler

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/bradfitz/slice"
)

// Dependency infos
type Dependency struct {
	Name  string
	Size  uint64
	cache string
}

// Profile binary size of each dependency
func Profile() ([]Dependency, error) {
	fmt.Println("Compiling...")

	cmd := exec.Command("go", "build", "-o", "tmp-build-output", "-work", "-a")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return []Dependency{}, err
	}

	fmt.Println("Parsing output...")
	tmpDir, err := parseCmdOutput(output)
	if err != nil {
		return []Dependency{}, err
	}
	defer func() {
		fmt.Println("Removing temp folder and temp binary")

		err = os.Remove("tmp-build-output")
		if err != nil {
			fmt.Printf("Failed to remove temp binary ./tmp-build-output")
		}

		err = os.RemoveAll(tmpDir)
		if err != nil {
			fmt.Printf("Failed to remove temp folder %s\n", tmpDir)
		}
	}()

	fmt.Println("Reading packagefile definitions...")
	file, err := os.Open(fmt.Sprintf("%s/b001/importcfg.link", tmpDir))
	if err != nil {
		return []Dependency{}, err
	}
	defer file.Close()

	dependencies := []Dependency{}
	filescanner := bufio.NewScanner(file)
	for filescanner.Scan() {
		line := filescanner.Text()

		if !strings.HasPrefix(line, "packagefile") {
			continue
		}

		splitStr := strings.Split(line, " ")
		if len(splitStr) <= 1 {
			continue
		}

		splitImport := strings.Split(splitStr[1], "=")
		if len(splitImport) <= 1 {
			continue
		}

		stat, err := os.Stat(splitImport[1])
		if err != nil {
			fmt.Printf("Failed to stat file %s: %s", splitImport[1], err)
			continue
		}

		dependencies = append(dependencies, Dependency{Size: uint64(stat.Size()), Name: splitImport[0], cache: splitImport[1]})
	}

	slice.Sort(dependencies[:], func(i, j int) bool {
		return dependencies[i].Size < dependencies[j].Size
	})

	return dependencies, nil
}

func parseCmdOutput(output []byte) (string, error) {
	outputStr := fmt.Sprintf("%s", output)
	if !strings.HasPrefix(outputStr, "WORK=") {
		return "", errors.New("No WORK output from build command")
	}

	tmpDir := strings.TrimSpace(outputStr[5:])
	if len(tmpDir) == 0 {
		return "", errors.New("Parse WORK folder is an empty path")
	}

	return tmpDir, nil
}
