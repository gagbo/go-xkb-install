package rules

import (
	"bufio"
	"fmt"
	"github.com/gagbo/go-xkb-install/internal/utils"
	"io/ioutil"
	"strings"
)

func AddLstVariant(layout, variantName, desc string) error {
	lstPath := "/usr/share/X11/xkb/rules/evdev.lst"
	content, err := ioutil.ReadFile(lstPath)
	if err != nil {
		return fmt.Errorf("could not read %v: %s", lstPath, err)
	}
	contentStr := string(content)

	variantLine := fmt.Sprintf("%s\t%s: %s\n", variantName, layout, desc)
	variantHeader := "! variant"

	if strings.Index(contentStr, variantLine) != -1 {
		return fmt.Errorf("the variant %s seems to already be listed", variantLine)
	}

	header := strings.Index(contentStr, variantHeader)
	if header == -1 {
		return fmt.Errorf("the variant header is missing")
	}

	scanner := bufio.NewScanner(strings.NewReader(contentStr))
	fileContent := ""
	for scanner.Scan() {
		line := scanner.Text()
		fileContent += line
		if line == variantHeader {
			fmt.Println("Found variant header!")
			fileContent += variantLine
		}
	}

	_, err = utils.BackupAndWrite(fileContent, lstPath)

	if err != nil {
		return fmt.Errorf("couldn't write rules files: %s", err)
	}
	return nil
}
