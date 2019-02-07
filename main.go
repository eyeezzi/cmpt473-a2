package main

import ( 
	"fmt" 
	"os"
	"path/filepath"
	"log"
	"strings"
	"os/exec"
	"errors"
)

func getFiles(dir string) []string {
	var list []string

	_ = filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if !(path == dir) && !f.IsDir() { list = append(list, path) }
		return nil
	})

	return list
}

func runSUT(csvFile, jsonDir, msgDir string) (jsonFile string, msgFile string) {
	// setup the destination files
	ext := filepath.Ext(csvFile)
	filename := strings.TrimSuffix(filepath.Base(csvFile), ext)

	jsonFile = fmt.Sprintf("%s.json", filepath.Join(jsonDir, filename))
	msgFile = fmt.Sprintf("%s.log", filepath.Join(msgDir, filename))

	// run the software-under-test
	cmd := fmt.Sprintf("cat %s | ./csv2json > %s 2> %s", csvFile, jsonFile, msgFile)

	c := exec.Command("bash", "-c", cmd)
	if err := c.Run(); err != nil {
		// don't care
	}

	return
}

func diff(file1, file2 string) error {
	cmd := exec.Command("diff", file1, file2)
	out, err := cmd.Output()

	if err != nil {
		return errors.New(string(out))
	}

	return nil
}

func clean(dirs ...string) {
	for _, dir := range dirs {
		s := fmt.Sprintf("rm -rf ./%s/*", dir)
		_, err := exec.Command("bash", "-c", s).Output()

		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	verbose := true

	// sut := "csv2json"
	testFilesDir := "TestData/TestFiles"

	expOutDir := "TestData/ExpectedOutput"
	expMsgDir := "TestData/ExpectedMessages"

	outDir := "TestOutput/Files"
	msgDir := "TestOutput/Messages"

	testFiles := getFiles(testFilesDir)
	expOutFiles := getFiles(expOutDir)
	expMsgFiles := getFiles(expMsgDir)

	// ensure every testfile has a corresponding expected output and message file
	if !(len(testFiles) == len(expOutFiles) && len(testFiles) == len(expMsgFiles)) {
		log.Fatal("Each TestFile requires an expectedOutput and expectedMessage file.")
	}

	// clean the destination directories
	clean(outDir, msgDir)

	for i, file := range testFiles {
		
		// run the testfile
		outFile, msgFile := runSUT(file, outDir, msgDir)

		// Assert output == expectedOutput && msg = expectedMessage
		outErr := diff(outFile, expOutFiles[i])
		msgErr := diff(msgFile, expMsgFiles[i])

		if outErr != nil || msgErr != nil { 
			fmt.Printf("Test %v: failed ---------------------\n", i)
		} else {
			fmt.Printf("Test %v: passed ---------------------\n", i)
		}

		if verbose && (outErr != nil) {	
			fmt.Printf("\tOutput Diff: \n%s", outErr.Error())		
		}
		if verbose && (msgErr != nil) {
			fmt.Printf("\tMessage Diff: \n%s", msgErr.Error())			
		}
	}
}

/* dumpster

*/