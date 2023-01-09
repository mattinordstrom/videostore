package pdfhandler

import (
	"fmt"

	"github.com/signintech/gopdf"
)

func CreatePDF(finishedPDF chan bool) {
	fmt.Println("Creating PDF..")

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()

	pdf.SetLineWidth(1)
	pdf.Oval(100, 200, 500, 500)

	//TODO Set add font and text
	//pdf.Cell(nil, "Test pdf output 123")

	pdf.WritePdf("receipt.pdf")

	//Use to test concurrency:
	//time.Sleep(5 * time.Second)

	fmt.Println("PDF created")

	finishedPDF <- true
}
