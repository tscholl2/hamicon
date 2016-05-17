package main

import "math/rand"

type diffs struct {
	rand  *rand.Rand
	body  body
	mouth mouth
	eyes  eyes
	legs  legs
}

func newDiffs(r *rand.Rand) (d diffs) {
	d.rand = r
	d.body = newBody(d.rand)
	d.mouth = newMouth(d.rand)
	d.eyes = newEyes(d.rand)
	d.legs = newLegs(d.rand)
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
