package binding

type not struct {
	Bool
}

var _ Bool = (*not)(nil)

func Not(data Bool) Bool { _ = "STUB: not implemented"; return *new(Bool) }

func (n *not) Get() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (n *not) Set(value bool) error { _ = "STUB: not implemented"; return nil }

type and struct {
	booleans
}

var _ Bool = (*and)(nil)

func And(data ...Bool) Bool { _ = "STUB: not implemented"; return *new(Bool) }

func (a *and) Get() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (a *and) Set(value bool) error { _ = "STUB: not implemented"; return nil }

type or struct {
	booleans
}

var _ Bool = (*or)(nil)

func Or(data ...Bool) Bool { _ = "STUB: not implemented"; return *new(Bool) }

func (o *or) Get() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (o *or) Set(value bool) error { _ = "STUB: not implemented"; return nil }

type booleans struct {
	data []Bool
}

func (g *booleans) AddListener(listener DataListener) { _ = "STUB: not implemented"; return }

func (g *booleans) RemoveListener(listener DataListener) { _ = "STUB: not implemented"; return }
