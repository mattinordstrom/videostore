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

	pdf.Line(1, 370, 400, 370)

	pdf.SetXY(10, 400)
	pdf.Cell(nil, "Date: "+time.Now().Format("2006-01-02 15:04"))

	// LOGO ///////////////////////////////
	pdf.SetStrokeColor(255, 0, 0) // red
	pdf.SetLineWidth(2)
	pdf.SetFillColor(0, 255, 0) // green
	pdf.Rectangle(410, 20, 570, 60, "DF", 3, 10)

	pdf.SetFillColor(0, 0, 0)
	pdf.SetXY(455, 30)
	pdf.SetFontSize(22)
	pdf.Cell(nil, "RETRO")
	////////////////////////////////////////

	pdf.WritePdf("pdf_output/" + rentalId.String() + ".pdf")

	//Use to test concurrency:
	//time.Sleep(5 * time.Second)

	fmt.Println("PDF created")

	finishedPDF <- PDF_SUCCESS
}
