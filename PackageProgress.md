## Package Mapping

#### Legend to make sense of goapp-mdc package / status and Notes columns

* TODO - There is no work done yet and there definitely needs to be.
* investigate - There might be a goapp-mdc component that need to be created.
* LIB - This is a mdc library / styling package that is used by other mdc packages.

| [material-components-web](https://github.com/material-components/material-components-web) | [goapp-mdc](https://github.com/mlctrez/goapp-mdc/) package / status| Notes |
| --- | --- | --- |
| [material-components-web](https://github.com/material-components/material-components-web/tree/master/packages/material-components-web) | - | Top level library package, it's where root "mdc" comes from. |
| [mdc-animation](https://github.com/material-components/material-components-web/tree/master/packages/mdc-animation) | LIB | - |
| [mdc-auto-init](https://github.com/material-components/material-components-web/tree/master/packages/mdc-auto-init) | investigate | Compare with using `new mdc.<pkg>.MDCXxxx(elem)`during OnMount(). More future proof? |
| [mdc-banner](https://github.com/material-components/material-components-web/tree/master/packages/mdc-banner) | [banner](pkg/banner) | - |
| [mdc-base](https://github.com/material-components/material-components-web/tree/master/packages/mdc-base) | [base](pkg/base) | The goapp-mdc base package contains a common base struct for accessing javascript vars and a test framework, both not related to the mdc-base package. |
| [mdc-button](https://github.com/material-components/material-components-web/tree/master/packages/mdc-button) | [button](pkg/button) | - |
| [mdc-card](https://github.com/material-components/material-components-web/tree/master/packages/mdc-card) | [card](pkg/card) | - |
| [mdc-checkbox](https://github.com/material-components/material-components-web/tree/master/packages/mdc-checkbox) | [checkbox](pkg/checkbox) | - |
| [mdc-chips](https://github.com/material-components/material-components-web/tree/master/packages/mdc-chips) | pkg/chips | TODO |
| [mdc-circular-progress](https://github.com/material-components/material-components-web/tree/master/packages/mdc-circular-progress) | [pkg/progress](pkg/progress) | circular.go |
| [mdc-data-table](https://github.com/material-components/material-components-web/tree/master/packages/mdc-data-table) | pkg/table ? | TODO |
| [mdc-density](https://github.com/material-components/material-components-web/tree/master/packages/mdc-density) | LIB | - |
| [mdc-dialog](https://github.com/material-components/material-components-web/tree/master/packages/mdc-dialog) | [dialog](pkg/dialog) | - |
| [mdc-dom](https://github.com/material-components/material-components-web/tree/master/packages/mdc-dom) | LIB | - |
| [mdc-drawer](https://github.com/material-components/material-components-web/tree/master/packages/mdc-drawer) | [drawer](pkg/drawer) | TODO |
| [mdc-elevation](https://github.com/material-components/material-components-web/tree/master/packages/mdc-elevation) | LIB | - |
| [mdc-fab](https://github.com/material-components/material-components-web/tree/master/packages/mdc-fab) | [fab](pkg/fab) | - |
| [mdc-feature-targeting](https://github.com/material-components/material-components-web/tree/master/packages/mdc-feature-targeting) | LIB | - |
| [mdc-floating-label](https://github.com/material-components/material-components-web/tree/master/packages/mdc-floating-label) | pkg/floatlabel ? | - |
| [mdc-form-field](https://github.com/material-components/material-components-web/tree/master/packages/mdc-form-field) | pkg/form | TODO and move textarea/textfield to pkg/form ? |
| [mdc-icon-button](https://github.com/material-components/material-components-web/tree/master/packages/mdc-icon-button) | [pkg/icon](pkg/icon) | Currently in [pkg/iconbutton](pkg/iconbutton) - TODO for moving it |
| [mdc-image-list](https://github.com/material-components/material-components-web/tree/master/packages/mdc-image-list) | pkg/image ? | TODO |
| [mdc-layout-grid](https://github.com/material-components/material-components-web/tree/master/packages/mdc-layout-grid) | [pkg/layout](pkg/layout) | - |
| [mdc-line-ripple](https://github.com/material-components/material-components-web/tree/master/packages/mdc-line-ripple) | investigate | - |
| [mdc-linear-progress](https://github.com/material-components/material-components-web/tree/master/packages/mdc-linear-progress) | [pkg/progress](pkg/progress) | linear.go |
| [mdc-list](https://github.com/material-components/material-components-web/tree/master/packages/mdc-list) | pkg/list | TODO |
| [mdc-menu](https://github.com/material-components/material-components-web/tree/master/packages/mdc-menu) | pkg/menu | TODO |
| [mdc-menu-surface](https://github.com/material-components/material-components-web/tree/master/packages/mdc-menu-surface) | investigate | - |
| [mdc-notched-outline](https://github.com/material-components/material-components-web/tree/master/packages/mdc-notched-outline) | investigate | - |
| [mdc-progress-indicator](https://github.com/material-components/material-components-web/tree/master/packages/mdc-progress-indicator) | LIB | - |
| [mdc-radio](https://github.com/material-components/material-components-web/tree/master/packages/mdc-radio) | pkg/radio | TODO |
| [mdc-ripple](https://github.com/material-components/material-components-web/tree/master/packages/mdc-ripple) | investigate | - |
| [mdc-rtl](https://github.com/material-components/material-components-web/tree/master/packages/mdc-rtl) | investigate | - |
| [mdc-segmented-button](https://github.com/material-components/material-components-web/tree/master/packages/mdc-segmented-button) | investigate | - |
| [mdc-select](https://github.com/material-components/material-components-web/tree/master/packages/mdc-select) | pgk/select | TODO |
| [mdc-shape](https://github.com/material-components/material-components-web/tree/master/packages/mdc-shape) | investigate | - |
| [mdc-slider](https://github.com/material-components/material-components-web/tree/master/packages/mdc-slider) | [slider](pkg/slider) | Mostly done, Needs testing and api events |
| [mdc-snackbar](https://github.com/material-components/material-components-web/tree/master/packages/mdc-snackbar) | [snackbar](pkg/snackbar) | Mostly done |
| [mdc-switch](https://github.com/material-components/material-components-web/tree/master/packages/mdc-switch) | pkg/switch | TODO |
| [mdc-tab](https://github.com/material-components/material-components-web/tree/master/packages/mdc-tab) | [tab](pkg/tab) | tab.go |
| [mdc-tab-bar](https://github.com/material-components/material-components-web/tree/master/packages/mdc-tab-bar) | [tab](pkg/tab) | bar.go |
| [mdc-tab-indicator](https://github.com/material-components/material-components-web/tree/master/packages/mdc-tab-indicator) | investigate | - |
| [mdc-tab-scroller](https://github.com/material-components/material-components-web/tree/master/packages/mdc-tab-scroller) | investigate | - |
| [mdc-textfield](https://github.com/material-components/material-components-web/tree/master/packages/mdc-textfield) | [textfield](pkg/textfield) | - |
| [mdc-theme](https://github.com/material-components/material-components-web/tree/master/packages/mdc-theme) | investigate | - |
| [mdc-tokens](https://github.com/material-components/material-components-web/tree/master/packages/mdc-tokens) | investigate | - |
| [mdc-tooltip](https://github.com/material-components/material-components-web/tree/master/packages/mdc-tooltip) | pkg/tooltip | TODO |
| [mdc-top-app-bar](https://github.com/material-components/material-components-web/tree/master/packages/mdc-top-app-bar) | [bar](pkg/bar) | In Progress |
| [mdc-touch-target](https://github.com/material-components/material-components-web/tree/master/packages/mdc-touch-target) | investigate | - |
| [mdc-typography](https://github.com/material-components/material-components-web/tree/master/packages/mdc-typography) | investigate | - |
