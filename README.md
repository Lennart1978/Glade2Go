# Glade2Go
Convert .glade file into .go file !
Usage: programname source.glade target.go

The .go file will look like this:

package ui

const ui = "... data from .glade file here ..."

func GetUI() string {
  return ui
}

You can load your .glade string with : builder, err := gtk.BuilderNewFromString(ui.GetUI) ...
