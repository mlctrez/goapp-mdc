package demo

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/icon"
	"github.com/mlctrez/goapp-mdc/pkg/list"
)

func addRoute(nav *list.Item, compo app.Composer) {
	nav.Type = list.ItemTypeAnchor
	navigationItems = append(navigationItems, nav)
	app.Route(nav.Href, compo)
}

func Routes() {
	addRoute(&list.Item{Text: "Home", Graphic: icon.MIHome, Href: "/"}, &Index{})
	addRoute(&list.Item{Text: "Banner", Graphic: icon.MIVoicemail, Href: "/banner"}, &BannerDemo{})
	addRoute(&list.Item{Text: "Button", Graphic: icon.MISmartButton, Href: "/button"}, &ButtonDemo{})
	addRoute(&list.Item{Text: "Card", Graphic: icon.MICreditCard, Href: "/card"}, &CardDemo{})
	addRoute(&list.Item{Text: "Checkbox", Graphic: icon.MICheckBox, Href: "/checkbox"}, &CheckboxDemo{})
	addRoute(&list.Item{Text: "Dialog", Graphic: icon.MISpeaker, Href: "/dialog"}, &DialogDemo{})
	addRoute(&list.Item{Text: "Drawer", Graphic: icon.MIDashboard, Href: "/drawer"}, &DrawerDemo{})
	addRoute(&list.Item{Text: "Fab", Graphic: icon.MIFavorite, Href: "/fab"}, &FabDemo{})
	addRoute(&list.Item{Text: "Form", Graphic: icon.MIInput, Href: "/form"}, &FormDemo{})
	addRoute(&list.Item{Text: "Icon", Graphic: icon.MIIcecream, Href: "/icon"}, &IconDemo{})
	addRoute(&list.Item{Text: "List", Graphic: icon.MIList, Href: "/list"}, &ListDemo{})
	addRoute(&list.Item{Text: "Tab", Graphic: icon.MITab, Href: "/tab"}, &TabDemo{})
	navigationItems = append(navigationItems, &list.Item{Type: list.ItemTypeDivider})
	addRoute(&list.Item{Text: "Code", Graphic: icon.MICode, Href: "/code"}, &CodeDemo{})
}
