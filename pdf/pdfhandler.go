package pdf

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/signintech/gopdf"
)

const (
	PDF_SUCCESS      = 1
	PDF_FAIL_ADDFONT = 2
	PDF_FAIL_SETFONT = 3
)

func CreatePDF(finishedPDF chan int, rentalId uuid.UUID, videoName string, customer string) {
	fmt.Println("Creating PDF..")

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()

	err := pdf.AddTTFFont("roboto", "./roboto-v19-latin-500.ttf")
	if err != nil {
		log.Print(err.Error())
		finishedPDF <- PDF_FAIL_ADDFONT
		return
	}

	err = pdf.SetFont("roboto", "", 14)
	if err != nil {
		log.Print(err.Error())
		finishedPDF <- PDF_FAIL_SETFONT
		return
	}
	pdf.Cell(nil, "Receipt: "+rentalId.String())
	pdf.SetXY(10, 70)
	pdf.Cell(nil, "Customer: "+customer)
	pdf.SetXY(10, 100)
	pdf.Cell(nil, "Video: "+videoName)
	pdf.SetXY(10, 400)
	pdf.Cell(nil, "Date: "+time.Now().Format("2006-01-02 15:04"))

	pdf.SetLineWidth(1)
	pdf.Oval(450, 10, 550, 110)

	pdf.WritePdf(rentalId.String() + ".pdf")

	//Use to test concurrency:
	//time.Sleep(5 * time.Second)

	fmt.Println("PDF created")

	finishedPDF <- PDF_SUCCESS
}
