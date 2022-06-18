package rules

import (
	"github.com/gagbo/go-xkb-install/internal/utils"

	"fmt"

	"github.com/beevik/etree"
)

// These settings prevent the re-encoding of HTML-dangerous characters like '>é&...
var settings = etree.WriteSettings{
	CanonicalText:    true,
	CanonicalEndTags: false,
	CanonicalAttrVal: false,
	UseCRLF:          false,
}

func UpdateVariant(layout, variantName, desc string) error {
	doc := etree.NewDocument()
	doc.WriteSettings = settings
	rulePath := "/usr/share/X11/xkb/rules/evdev.xml"
	if err := doc.ReadFromFile(rulePath); err != nil {
		return err
	}

	root := doc.Root()
	if root == nil {
		return fmt.Errorf("invalid xml file, no root found")
	}

	path := etree.MustCompilePath(fmt.Sprintf("/xkbConfigRegistry/layoutList/layout/configItem[name='%s']/..", layout))

	layoutConfig := root.FindElementPath(path)
	if layoutConfig == nil {
		return fmt.Errorf("layout '%s' not found", layout)
	}

	path = etree.MustCompilePath(fmt.Sprintf("./variantList/variant/configItem[name='%s']/..", variantName))
	if layoutConfig.FindElementPath(path) != nil {
		return fmt.Errorf("there is already a '%s' variant in the variantList", variantName)
	}

	variantList := layoutConfig.FindElement("./variantList")
	variantConfigItem := variantList.CreateElement("variant").CreateElement("configItem")
	variantConfigItem.CreateElement("name").CreateText(variantName)
	variantConfigItem.CreateElement("description").CreateText(desc)

	// Sanity checks
	if layoutConfig.FindElementPath(path) == nil {
		return fmt.Errorf("after writing, couldn't find '%s' variant in the layoutConfig", variantName)
	}
	if root.FindElement(
		fmt.Sprintf(
			"/xkbConfigRegistry/layoutList/layout/configItem[name='%s']/../variantList/variant/configItem[name='%s']/..",
			layout,
			variantName)) == nil {
		return fmt.Errorf("after writing, couldn't find '%s' variant in the layout %s of the root", variantName, layout)
	}

	doc.Indent(2)
	ruleStr, err := doc.WriteToString()
	if err != nil {
		return fmt.Errorf("couldn't serialize rules to XML: %s", err)
	}

	_, err = utils.BackupAndWrite(ruleStr, rulePath)

	if err != nil {
		return fmt.Errorf("couldn't write rules files: %s", err)
	}
	return nil
}
