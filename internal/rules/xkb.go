package rules

import (
	"github.com/gagbo/go-xkb-install/internal/utils"

	"fmt"
)

func AddLayout(layout, pathToXkb string) error {
	rulePath := fmt.Sprintf("/usr/share/X11/xkb/symbols/%s", layout)

	if _, err := utils.BackupAndAppend(pathToXkb, rulePath); err != nil {
		return fmt.Errorf("could not update symbols file: %w", err)
	}

	return nil
}
