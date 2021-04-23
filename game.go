package main 

import ( 
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

type Feld struct{
	brett  [][] bool
	hoehe, breite  int 
}


type Leben struct {				// Speichert den nächsten schritt des Feldes
	a, b *Feld 
	hoehe, breite int 
}


func (f *Feld) Set(w, h int, b bool) {  		// bekommt den "index" des Feldes und setzt ihn auf wahr oder falsch
	f.brett[w][h] = b
}

func (f *Feld) Next(x, y int) bool{				//ermittelt die Nachbarzellen die am leben sind 
	alive := 0 
	for i := -1; i<=1; i++{
		for j :=-1; j<=1; j++{
			if(j != 0|| i != 0) && f.AmLeben(x+i, y+j){
				alive++
			}
		}
	}
	return alive == 3 || alive == 2 && f.AmLeben(x, y)
}

func (f *Feld) AmLeben(h, w int) bool{			// 
	h += f.hoehe
	h %= f.hoehe
	w += f.breite
	w %= f.breite
	return f.brett[w][h]

}

func NeuesBrett(hoehe, breite int) *Feld {				//erzeugt ein neues Feld mit bool Werten 
	brett := make([][]bool, hoehe)
	for i := range brett {
		brett[i] = make([]bool, breite)
	}
	return &Feld{brett: brett, breite: breite, hoehe: hoehe}
}


func NeuStart(hoehe, breite int) *Leben {				//erstellt ein Brett mit einem zufälligen Muster
	a := NeuesBrett(hoehe, breite)						//ein viertel der Zellen sind am Leben 
	for i := 0; i < (hoehe * breite / 4); i++ {
		a.Set(rand.Intn(hoehe), rand.Intn(breite), true)	
	}
	return &Leben{
		a: a, b: NeuesBrett(hoehe, breite),
		hoehe: hoehe, breite: breite,
	}
}

func (l *Leben) Update(){							// tauscht das aktive Feld mit dem neu generierten 
	for y := 0; y < l.hoehe; y++{
		for x := 0; x < l.breite; x++{
			l.b.Set(x, y, l.a.Next(x,y))
		}
	}
	l.a, l.b = l.b, l.a 
}

func (l *Leben) String() string {					// gibt das Feld als String wieder 
	var buf bytes.Buffer
	for y := 0; y < l.hoehe; y++ {
		for x := 0; x < l.breite; x++ {
			b := byte(' ')
			if l.a.AmLeben(x, y) {
				b = '0'
			}
			buf.WriteByte(b)
		}
		buf.WriteByte(' ')
	}
	return buf.String()
}

func main() {
	ns := NeuStart(50, 10)
	for i := 0; i < 50; i++ {				// gibt die ersten 50. generationen wieder 
		ns.Update()
		fmt.Print("\x0c %d", ns) 			// printed das Feld 
		time.Sleep(time.Second / 30)
	}
}