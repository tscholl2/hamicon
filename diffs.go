package main

type diffs struct {
	mouth mouth
	eyes  eyes
}

func newDiffs() (d diffs) {
	d.mouth = newMouth()
	d.eyes = newEyes()
	return
}

func (d diffs) toMap() map[string]string {
	m := make(map[string]string)
	m["Mouth"] = mouthToSVG(d)
	m["Eyes"] = eyesToSVG(d)
	return m
}
