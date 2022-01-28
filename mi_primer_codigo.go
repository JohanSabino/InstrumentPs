package main

import (
	_ "bytes"
	"encoding/json"
	"fmt"
	"os"
)

var affirmations = map[string]string{
	"i1":"Tristeza",
	"i2":"Pesimismo ",
	"i3":"Fracaso",
	"i4":"Pérdida de Placer ",
	"i5":"Sentimientos de Culpa",
	"i6":"Sentimientos de Castigo",
	"i7":"Dissconformidad con uno mismo",
	"i8":"Pensamientos o Deseos Suicidas",
	"i9":"Pensamientos o Deseos Suicidas",
	"i10":"Llanto",
	"i11":"Agitación",
	"i12":"Pérdida de Interés",
	"i13":"Indecisión",
	"i14":"Desvalorización",
	"i15":"Pérdida de Energía",
	"i16":"Cambios en los Hábitos de Sueño",
	"i17":"Irritabilidad",
	"i18":"Cambios en el Apetito",
	"i19":"Dificultad de Concentración",
	"i20":"Cansancio o Fatiga",
	"i21":"Pérdida de Interés en el Sexo",
}

type Affirmation struct {
	Affirmations string `json:"affirmations"`
	Score        int `json:"score"`
}

func (a Affirmation) ShowAnswer() {
	fmt.Println("Has respondido la afirmación 1:", a.Affirmations, ", con el valor:", a.Score)
}

func main() {
	p1 := Affirmation{
		Affirmations: affirmations["i1"],
		Score:        2,
	}
	os.Environ()

	s1 := body()
	fmt.Println(string(s1))
	p1.ShowAnswer()
	Scores(&p1.Score)
}
func body() []byte {
	bs, err := json.Marshal(affirmations)
	if err != nil {
		fmt.Println("Fall", err)
	}
	return bs
}
func Scores(Score *int) int {
		for *Score = 0; *Score != 1; *Score++ {
		fmt.Println(Score)
	}
	return *Score
}