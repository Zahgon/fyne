package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/cmd/fyne_demo/data"
	"fyne.io/fyne/v2/cmd/fyne_demo/tutorials"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

const preferenceCurrentTutorial = "currentTutorial"

var topWindow fyne.Window

func main() {
	a := app.NewWithID("io.fyne.demo")
	a.SetIcon(data.FyneLogo)
	w := a.NewWindow("Fyne Demo")
	makeTray(a, w)
	logLifecycle(a)
	topWindow = w

	w.SetMainMenu(makeMenu(a, w))
	w.SetMaster()

	content := container.NewStack()
	title := widget.NewLabel("Component name")
	intro := widget.NewLabel("An introduction would probably go\nhere, as well as a")
	intro.Wrapping = fyne.TextWrapWord
	setTutorial := func(t tutorials.Tutorial) {
		if fyne.CurrentDevice().IsMobile() {
			child := a.NewWindow(t.Title)
			topWindow = child
			child.SetContent(t.View(topWindow))
			child.Show()
			child.SetOnClosed(func() {
				topWindow = w
			})
			return
		}

		title.SetText(t.Title)
		intro.SetText(t.Intro)

		if t.Title == "Welcome" {
			title.Hide()
			intro.Hide()
		} else {
			title.Show()
			intro.Show()
		}

		content.Objects = []fyne.CanvasObject{t.View(w)}
		content.Refresh()
	}

	tutorial := container.NewBorder(
		container.NewVBox(title, widget.NewSeparator(), intro), nil, nil, nil, content,
	)
	if fyne.CurrentDevice().IsMobile() {
		w.SetContent(makeNav(setTutorial, false))
	} else {
		split := container.NewHSplit(makeNav(setTutorial, true), tutorial)
		split.Offset = 0.2
		w.SetContent(split)
	}
	w.Resize(fyne.NewSize(640, 460))

	notice := widget.NewRichTextFromMarkdown(
		"This demo has been moved to a new repository.\n\n[Fyne demo on GitHub](https://github.com/fyne-io/demo)",
	)
	notice.Segments[2].(*widget.HyperlinkSegment).Alignment = fyne.TextAlignCenter
	dialog.ShowCustom("Fyne Demo Moved", "OK", notice, w)

	w.ShowAndRun()
}

func logLifecycle(a fyne.App) { _ = "STUB: not implemented"; return }

func makeMenu(a fyne.App, w fyne.Window) *fyne.MainMenu { _ = "STUB: not implemented"; return nil }

func makeTray(a fyne.App, w fyne.Window) { _ = "STUB: not implemented"; return }

func makeNav(setTutorial func(tutorial tutorials.Tutorial), loadPrevious bool) fyne.CanvasObject {
	_ = "STUB: not implemented"
	return *new(fyne.CanvasObject)
}

func shortcutFocused(s fyne.Shortcut, cb fyne.Clipboard, f fyne.Focusable) {
	_ = "STUB: not implemented"
	return
}
