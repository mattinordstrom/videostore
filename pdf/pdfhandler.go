package pdf

import (
	"fmt"
	"log"
	"time"

	"github.com/signintech/gopdf"
)

const (
	pdfSuccess     = 1
	pdfFailAddfont = 2
	pdfFailSetfont = 3
)

func CreatePDF(finishedPDF chan int, rentalID fmt.Stringer, videoName string, customer string) {
	fmt.Println("Creating PDF..")

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()

	err := pdf.AddTTFFont("roboto", "./roboto-v19-latin-500.ttf")
	if err != nil {
		log.Print(err.Error())
		finishedPDF <- pdfFailAddfont

		return
	}

	err = pdf.SetFont("roboto", "", 14)
	if err != nil {
		log.Print(err.Error())
		finishedPDF <- pdfFailSetfont

		return
	}

	if err := pdf.Cell(nil, "Receipt: "+rentalID.String()); err != nil {
		panic(err)
	}

	pdf.SetXY(10, 70)

	if err := pdf.Cell(nil, "Customer: "+customer); err != nil {
		panic(err)
	}

	pdf.SetXY(10, 100)

	if err := pdf.Cell(nil, "Video: "+videoName); err != nil {
		panic(err)
	}

	pdf.Line(1, 370, 400, 370)

	pdf.SetXY(10, 400)

	if err := pdf.Cell(nil, "Date: "+time.Now().Format("2006-01-02 15:04")); err != nil {
		panic(err)
	}

	// LOGO ///////////////////////////////
	pdf.SetStrokeColor(255, 0, 0) // red
	pdf.SetLineWidth(2)
	pdf.SetFillColor(0, 255, 0) // green

	if err := pdf.Rectangle(410, 20, 570, 60, "DF", 3, 10); err != nil {
		panic(err)
	}

	pdf.SetFillColor(0, 0, 0)
	pdf.SetXY(455, 30)

	if err := pdf.SetFontSize(22); err != nil {
		panic(err)
	}

	if err := pdf.Cell(nil, "RETRO"); err != nil {
		panic(err)
	}
	////////////////////////////////////////

	pdfErr := pdf.WritePdf("pdf_output/" + rentalID.String() + ".pdf")
	if pdfErr != nil {
		log.Fatal("Failed to write PDF")
	}

	// Use to test concurrency:
	// time.Sleep(5 * time.Second)

	fmt.Println("PDF created")

	finishedPDF <- pdfSuccess
}
