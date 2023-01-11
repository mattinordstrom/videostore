package pdfhandler

import (
	"fmt"
	"log"

	"github.com/signintech/gopdf"
)

func CreatePDF(finishedPDF chan bool) {
	fmt.Println("Creating PDF..")

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()

	err := pdf.AddTTFFont("roboto", "./roboto-v19-latin-500.ttf")
	if err != nil {
		log.Print(err.Error())
		//TODO return status int to client
		finishedPDF <- true
		return
	}

	err = pdf.SetFont("roboto", "", 14)
	if err != nil {
		log.Print(err.Error())
		//TODO return status int to client
		finishedPDF <- true
		return
	}
	pdf.Cell(nil, "receipt test 123")

	pdf.SetLineWidth(1)
	pdf.Oval(100, 200, 500, 500)

	pdf.WritePdf("receipt.pdf")

	//Use to test concurrency:
	//time.Sleep(5 * time.Second)

	fmt.Println("PDF created")

	finishedPDF <- true
}
