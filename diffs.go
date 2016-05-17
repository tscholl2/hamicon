package main

type diffs struct {
	body  body
	mouth mouth
	eyes  eyes
	legs  legs
}

func newDiffs() (d diffs) {
	d.body = newBody()
	d.mouth = newMouth()
	d.eyes = newEyes()
	d.legs = newLegs()
	return
}

func (d diffs) toMap() map[string]string {
	m := make(map[string]string)
	m["Body"] = bodyToSVG(d)
	m["Mouth"] = mouthToSVG(d)
	m["Eyes"] = eyesToSVG(d)
	m["Legs"] = legsToSVG(d)
	return m
}
