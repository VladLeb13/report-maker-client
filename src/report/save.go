package report

import (
	"io/ioutil"
	"log"
	"os"
	"tools"
)

func Save(ctx *tools.AppContex) {
	report := ctx.Context.Value("HTMLReport").(string)
	cnf := ctx.Context.Value("configuration").(*tools.Config)

	ok := checkPath(cnf.SavingFolder)
	if !ok {
		path := createTempFolder()
		saveFile(path, report)
		return
	}

	saveFile(cnf.SavingFolder, report)
	return

}

func checkPath(path string) (ok bool) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Println("Folder ", path, "not found")
		return false
	}
	return true
}

func createTempFolder() (path string) {
	dir, err := ioutil.TempDir("", "RM")
	if err != nil {
		log.Fatal(err)
	}

	return dir
}

func saveFile(path string, report string) (ok bool) {
	file, err := os.Create(path + string(os.PathSeparator) + "report.htm")
	if err != nil {
		log.Println("IO error in write report file")
		return false
	}
	defer file.Close()
	_, err = file.WriteString(report)
	if err != nil {
		log.Println("IO error in write report file")
		return false
	}

	return true
}
