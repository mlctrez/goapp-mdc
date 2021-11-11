# goapp-mdc

[go-app](https://go-app.dev/) implementations of Material Components for the web

<p align="center">
    <a href="https://goreportcard.com/report/github.com/mlctrez/goapp-mdc"><img src="https://goreportcard.com/badge/github.com/mlctrez/goapp-mdc" alt="Go Report Card"></a>
	<a href="https://GitHub.com/mlctrez/goapp-mdc/releases/"><img src="https://img.shields.io/github/release/mlctrez/goapp-mdc.svg" alt="GitHub release"></a>
	<a href="https://pkg.go.dev/github.com/mlctrez/goapp-mdc"><img src="https://pkg.go.dev/badge/github.com/mlctrez/goapp-mdc.svg" alt="Go Reference"></a>
    <a href="https://github.com/mlctrez/goapp-mdc/blob/master/go.mod"><img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/mlctrez/goapp-mdc"></a>
    <a href="https://github.com/mlctrez/goapp-mdc/blob/master/LICENSE"><img alt="GitHub" src="https://img.shields.io/github/license/mlctrez/goapp-mdc"></a>
</p>

## Goal

Provide a set of components for [go-app](https://github.com/maxence-charriere/go-app) that make use of 
[material-components-web](https://github.com/material-components/material-components-web/) for 
styling, animations, and interaction.


## Dependencies

### Material Components Web

The [material-components-web](https://github.com/material-components/material-components-web/) project are 
what components in this library depend on for styling, animation, and interaction. 
Usage of this library assumes that the stock css and js from material-components-web are referenced via CDN 
or included in application in the `/web` folder. See [Getting Started](https://github.com/material-components/material-components-web/blob/master/docs/getting-started.md)
in material-components-web. 

A custom styled theme may be used, but this library depends on the
css and js objects being in the right spots. i.e. `new mdc.ripple.MDCRipple(element)` and 
`class="mdc-ripple-surface"` need to work in order for this library to render components correctly.

### Material Icons

This project also depends on the [Material Icons](https://fonts.google.com/icons?selected=Material+Icons) font 
being pulled in correctly via [Material Icons Guide](https://developers.google.com/fonts/docs/material_icons).
This project's [icon](pkg/icon) package contains a generated constants file for the list of current 
(as of Oct 2021) icons and custom names not found in these constants can be used by passing in a `MaterialIcon` 
variable with a custom icon name. The supported icon style class is currently `material-icons`. Other styles 
like `material-icons-outlined` or `material-icons-sharp` could be added with a bit of work. 

## Progress

First a word on stability - "This is a work in progress. Expect breaking changes before the v1 release."

### Packages 

This library's package structure loosely follows the naming of packages under
[material-components-web/packages](https://github.com/material-components/material-components-web/tree/master/packages/) 
but some library packages like mdc-animation won't have a `goapp-mdc` component.

[PackageProgress](PackageProgress.md) shows this mapping and also the development progress, if any.


