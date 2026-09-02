package desktop

type Window interface {
	RequestFullScreenSecondary()

	RequestAlwaysOnTop()

	RequestPosition(x, y int)
}
