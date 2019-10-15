package main

import (
	"fmt"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func main() {
	m := pdf.NewMaroto(consts.Landscape, consts.A4)

	m.Row(20, func() {
		m.Col(func() {
			m.FileImage("internal/assets/images/frontpage.png")
		})
		m.Col(func() {
			m.Text("DevFest - SUL", props.Text{
				Top:   10,
				Size:  20,
				Style: consts.BoldItalic,
				Align: consts.Center,
			})
		})
		m.Col(func() {
			m.FileImage("internal/assets/images/frontpage.png")
		})
	})

	talk := "Uma nova forma simples, elegante e rápida de criar PDFs em Go"
	text := "Nós da organização do DevFest - SUL (GDG-Cascavel, GDG-Pelotas, GDG-Porto Alegre e GDG-Floripa), " +
		"certificamos que (%s) participou como ouvinte da palestra (%s). Que foi realizada no dia 19/10/2019 durante" +
		" o período de 11:40h e 11:55h, na ACATE - Associação Catarinense de Tecnologia - Florianópolis, SC."

	m.Row(130, func() {
		m.Col(func() {
			m.Text(fmt.Sprintf(text, "Fulano de Tal", talk), props.Text{
				Top:    55,
				Family: consts.Helvetica,
				Size:   15,
				Align:  consts.Center,
			})
		})
	})

	m.Row(20, func() {
		m.Row(20, func() {
			m.Col(func() {
				m.Signature("Ouvinte")
			})
			m.Col(func() {
				m.Signature("Palestrante")
			})
			m.Col(func() {
				m.Signature("Organização")
			})
		})
	})

	_ = m.OutputFileAndClose("maroto.pdf")
}
