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

func runSUT(csvFile, jsonDir, msgDir string, noHeaderFlag string) (jsonFile string, msgFile string) {
	// setup the destination files
	ext := filepath.Ext(csvFile)
	filename := strings.TrimSuffix(filepath.Base(csvFile), ext)

	jsonFile = fmt.Sprintf("%s.json", filepath.Join(jsonDir, filename))
	msgFile = fmt.Sprintf("%s.log", filepath.Join(msgDir, filename))

	// run the software-under-test
	// cmd := fmt.Sprintf("cat %s | ./csv2json > %s 2> %s", csvFile, jsonFile, msgFile)
	cmd := fmt.Sprintf("./bin/csv2json --src %s %s > %s 2> %s", csvFile, noHeaderFlag, jsonFile, msgFile)

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
	verbose := false

	// sut := "csv2json"
	testFilesDir := "TestData/TestFiles"

	expOutDir := "TestData/ExpectedOutput"
	expMsgDir := "TestData/ExpectedMessages"

	outDir := "TestOutput/Files"
	msgDir := "TestOutput/Messages"

	testFiles := getFiles(testFilesDir)
	expOutFiles := getFiles(expOutDir)
	expMsgFiles := getFiles(expMsgDir)

	// No-header flag for each test configuration
	noHeaderFlags := map[int]string {
		1 : "",
		2 : "-no-headers",
		3 : "",
		4 : "-no-headers",
		5 : "-no-headers",
		6 : "",
		7 : "",
		8 : "-no-headers",
	}

	// ensure every testfile has a corresponding expected output and message file
	if !(len(testFiles) == len(expOutFiles) && len(testFiles) == len(expMsgFiles)) {
		log.Fatal("Each TestFile requires an expectedOutput and expectedMessage file.")
	}

	fmt.Println(`
-----------------------
Running Tests
-----------------------
	`)

	// clean the destination directories
	clean(outDir, msgDir)

	fmt.Println("Cleaning Output Directories\n")

	for i, file := range testFiles {
		
		// run the testfile
		outFile, msgFile := runSUT(file, outDir, msgDir, noHeaderFlags[i+1])

		// Assert output == expectedOutput && msg = expectedMessage
		outErr := diff(outFile, expOutFiles[i])
		msgErr := diff(msgFile, expMsgFiles[i])

		if outErr != nil || msgErr != nil { 
			fmt.Printf("Test %v: failed\n", i+1)
		} else {
			fmt.Printf("Test %v: passed\n", i+1)
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