package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

type Feld struct {
	s    [][]bool
	w, h int
}

func NewFeld(w, h int) *Feld { // erstellt ein 2d array mit bool werten
	s := make([][]bool, h)
	for i := range s {
		s[i] = make([]bool, w)
	}
	return &Feld{s: s, w: w, h: h}
}

func (f *Feld) Set(x, y int, b bool) { // Set setzt die Zelle auf den entsprechenden Wert
	f.s[y][x] = b
}

func (f *Feld) Amleben(x, y int) bool { // Amleben überprüft die Zelle am leben ist
	x += f.w
	x %= f.w
	y += f.h
	y %= f.h
	return f.s[y][x]
}

func (f *Feld) Next(x, y int) bool { // Next überprüft ob die Zelle im nächsten schritt noch lebt
	Amleben := 0 // Zähle die benachbarten Zellen die am Leben sind
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if (j != 0 || i != 0) && f.Amleben(x+i, y+j) {
				Amleben++
			}
		}
	}
	// Regeln des Spiels
	return Amleben == 3 || Amleben == 2 && f.Amleben(x, y)
}

type Leben struct {
	a, b *Feld
	w, h int
}

func NewLeben(w, h int) *Leben { //NewLeben startet das Spiel mit einem "random seed".
	a := NewFeld(w, h)
	for i := 0; i < (w * h / 4); i++ {
		a.Set(rand.Intn(w), rand.Intn(h), true)
	}
	return &Leben{
		a: a, b: NewFeld(w, h),
		w: w, h: h,
	}
}

func (l *Leben) Update() {
	// tauscht das Feld mit dem nächsten Schritt
	for y := 0; y < l.h; y++ {
		for x := 0; x < l.w; x++ {
			l.b.Set(x, y, l.a.Next(x, y))
		}
	}
	l.a, l.b = l.b, l.a
}

func (l *Leben) String() string { // gibt das Feld als String wieder
	var buf bytes.Buffer
	for y := 0; y < l.h; y++ {
		for x := 0; x < l.w; x++ {
			b := byte(' ')
			if l.a.Amleben(x, y) {
				b = '0'
			}
			buf.WriteByte(b)
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}

func main() {
	l := NewLeben(40, 15)
	for i := 0; i < 50; i++ { // Gibt die ersten 50 Generationen wieder 
		l.Update()
		fmt.Print("\x0c %s", l) // printed das Feld 
		time.Sleep(time.Second / 30)
	}
}
