package main

import (
	. "github.com/jroimartin/gocui"
	"unicode/utf8"
	"fmt"
)

type UIManager struct {
	inputView, outputView *View
	glyph                 rune
}

func startGui() {
	gui, err := NewGui(OutputNormal)
	if err != nil {
		panic(err)
	}

	glyph, _ := utf8.DecodeRuneInString(*glyph)
	gui.SetManager(&UIManager{glyph: glyph})
	SetKeyBindings(gui)

	gui.Cursor = true
	gui.Mouse = false
	gui.Highlight = true

	defer gui.Close()
	if err := gui.MainLoop(); err != nil && err != ErrQuit {
		panic(err)
	}
}

func (mngr *UIManager) Layout(gui *Gui) error {
	w, h := gui.Size()

	//Input view
	if view, err := gui.SetView("input", 0, 0, w*2/3-1, 4); err != nil {
		if err != ErrUnknownView {
			return err
		}

		view.Title = "Input"
		view.Editor = EditorFunc(mngr.InputViewEditor)
		view.Editable = true
		view.Wrap = false
		view.Highlight = false

		gui.SetCurrentView(view.Name())
		gui.SetViewOnTop(view.Name())

		mngr.inputView = view
	}

	//Output view
	if view, err := gui.SetView("output", 0, 5, w-1, h-1); err != nil {
		if err != ErrUnknownView {
			return err
		}

		view.Title = "Output"
		view.Editable = false
		view.Wrap = false

		mngr.outputView = view
	}

	//Shortcuts view
	if view, err := gui.SetView("shortcuts", w * 2/3, 0, w-1, 4); err != nil {
		if err != ErrUnknownView {
			return err
		}

		view.Title = "Shortcuts"
		view.Editable = false

		fmt.Fprintln(view, "^C exit\n^R change glyph")
	}

	return nil
}

func (mngr *UIManager) InputViewEditor(v *View, key Key, ch rune, mod Modifier) {
	DefaultEditor.Edit(v, key, ch, mod)

	mngr.outputView.Clear()
	generateASCIIArt(mngr.outputView, mngr.inputView.Buffer(), mngr.glyph)
}

func SetKeyBindings(gui *Gui) {
	gui.SetKeybinding("", KeyCtrlC, ModNone, func(gui *Gui, view *View) error {
		return ErrQuit
	})
}
