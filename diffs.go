package main

type diffs struct {
	lip lip
}

func newDiffs() (d diffs) {
	d.lip = newLip()
	return
}

func (d diffs) toMap() map[string]string {
	m := make(map[string]string)
	m["Lip"] = lipToSVG(d)
	return m
}
