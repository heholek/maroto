package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"strings"
)

func main() {
	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.AddPage()

	info := pdf.RegisterImageOptions("internal/assets/images/frontpage.png", gofpdf.ImageOptions{
		ReadDpi:   false,
		ImageType: "",
	})

	pdf.Image("internal/assets/images/frontpage.png", 45, 8, info.Width()*0.16, info.Height()*0.16, false, "", 0, "")
	pdf.Image("internal/assets/images/frontpage.png", 235, 8, info.Width()*0.16, info.Height()*0.16, false, "", 0, "")

	pdf.SetFont("Arial", "BI", 20)
	pdf.CellFormat(270, 15, "DevFest - SUL", "", 0, "C", false, 0, "")

	fontTextSize := 15.0
	textTop := 70.0
	pdf.SetFont("Helvetica", "", fontTextSize)

	width, _ := pdf.GetPageSize()
	left, top, right, _ := pdf.GetMargins()
	maxWidth := width - right - left

	talk := "Uma nova forma simples, elegante e rápida de criar PDFs em Go"
	text := "Nós da organização do DevFest - SUL (GDG-Cascavel, GDG-Pelotas, GDG-Porto Alegre e GDG-Floripa), " +
		"certificamos que (%s) participou como ouvinte da palestra (%s). Que foi realizada no dia 19/10/2019 durante" +
		" o período de 11:40h e 11:55h, na ACATE - Associação Catarinense de Tecnologia - Florianópolis, SC."

	translator := pdf.UnicodeTranslatorFromDescriptor("")
	finalText := translator(fmt.Sprintf(text, "Fulano de Tal", talk))
	words := strings.Split(finalText, " ")

	currentlySize := 0.0
	actualLine := 0
	lines := []string{}
	lines = append(lines, "")

	for _, word := range words {
		wordWidth := pdf.GetStringWidth(word + " ")
		if wordWidth+currentlySize < maxWidth {
			lines[actualLine] = lines[actualLine] + word + " "
			currentlySize += pdf.GetStringWidth(word + " ")
		} else {
			lines = append(lines, "")
			actualLine++
			lines[actualLine] = lines[actualLine] + word + " "
			currentlySize = pdf.GetStringWidth(word + " ")
		}
	}

	for index, line := range lines {
		lineWidth := pdf.GetStringWidth(line)
		pdf.Text(left+((maxWidth-lineWidth)/2), textTop+top+float64(index)*fontTextSize/2.0, line)
	}

	pdf.SetFont("Arial", "B", 10)

	pdf.Text(30, 180, translator("Ouvinte"))
	pdf.Line(25, 170, 55, 170)

	pdf.Text(140, 180, translator("Palestrante"))
	pdf.Line(130, 170, 220, 170)

	pdf.Text(250, 180, translator("Organização"))
	pdf.Line(230, 170, 290, 170)

	_ = pdf.OutputFileAndClose("gofpdf.pdf")
}
