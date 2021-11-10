package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(bytes.NewBufferString(src))
	for scanner.Scan() {
		txt := strings.TrimSpace(scanner.Text())
		if txt == "" {
			continue
		}
		parts := strings.SplitN(txt, "'", 3)
		fmt.Printf("const %s MDCName = %q\n", parts[1], parts[1])
	}

}

const src = `
autoInit.register('MDCBanner', banner.MDCBanner);
autoInit.register('MDCCheckbox', checkbox.MDCCheckbox);
autoInit.register('MDCChip', chips.MDCChip);
autoInit.register('MDCChipSet', chips.MDCChipSet);
autoInit.register('MDCCircularProgress', circularProgress.MDCCircularProgress);
autoInit.register('MDCDataTable', dataTable.MDCDataTable);
autoInit.register('MDCDialog', dialog.MDCDialog);
autoInit.register('MDCDrawer', drawer.MDCDrawer);
autoInit.register('MDCFloatingLabel', floatingLabel.MDCFloatingLabel);
autoInit.register('MDCFormField', formField.MDCFormField);
autoInit.register('MDCIconButtonToggle', iconButton.MDCIconButtonToggle);
autoInit.register('MDCLineRipple', lineRipple.MDCLineRipple);
autoInit.register('MDCLinearProgress', linearProgress.MDCLinearProgress);
autoInit.register('MDCList', list.MDCList);
autoInit.register('MDCMenu', menu.MDCMenu);
autoInit.register('MDCMenuSurface', menuSurface.MDCMenuSurface);
autoInit.register('MDCNotchedOutline', notchedOutline.MDCNotchedOutline);
autoInit.register('MDCRadio', radio.MDCRadio);
autoInit.register('MDCRipple', ripple.MDCRipple);
autoInit.register('MDCSegmentedButton', segmentedButton.MDCSegmentedButton);
autoInit.register('MDCSelect', select.MDCSelect);
autoInit.register('MDCSlider', slider.MDCSlider);
autoInit.register('MDCSnackbar', snackbar.MDCSnackbar);
autoInit.register('MDCSwitch', switchControl.MDCSwitch as unknown as MDCAttachable);
autoInit.register('MDCTabBar', tabBar.MDCTabBar);
autoInit.register('MDCTextField', textField.MDCTextField);
autoInit.register('MDCTooltip', tooltip.MDCTooltip);
autoInit.register('MDCTopAppBar', topAppBar.MDCTopAppBar);
`
