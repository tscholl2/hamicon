package main

import "math/rand"

//go:generate embd -n basicCSS static/basic.1.css

type diffs struct {
	rand  *rand.Rand
	body  body
	mouth mouth
	nose  nose
	eyes  eyes
	legs  legs
}

func newDiffs(r *rand.Rand) (d diffs) {
	d.rand = r
	d.body = newBody(d.rand)
	d.mouth = newMouth(d.rand)
	d.nose = newNose(d.rand)
	d.eyes = newEyes(d.rand)
	d.legs = newLegs(d.rand)
	return
}

func (d diffs) toMap() map[string]string {
	m := make(map[string]string)
	m["Body"] = bodyToSVG(d)
	m["Mouth"] = mouthToSVG(d)
	m["Nose"] = noseToSVG(d)
	m["Eyes"] = eyesToSVG(d)
	m["Legs"] = legsToSVG(d)
	m["CSS"] = basicCSS
	return m
}
