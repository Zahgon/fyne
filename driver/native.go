package driver

type NativeWindow interface {
	RunNative(func(context any))
}

type AndroidContext struct {
	VM, Env, Ctx uintptr
}

type AndroidWindowContext struct {
	AndroidContext
	NativeWindow uintptr
}

type IOSWindowContext struct{}

type UnknownContext struct{}

type WindowsWindowContext struct {
	HWND uintptr
}

type MacWindowContext struct {
	NSWindow uintptr
}

type X11WindowContext struct {
	WindowHandle uintptr
}

type WaylandWindowContext struct {
	WaylandSurface uintptr
}
