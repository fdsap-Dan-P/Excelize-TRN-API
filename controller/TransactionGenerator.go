package controller

import (
	"APITransactionGenerator/middleware/go-utils/database"
	"APITransactionGenerator/struct/request"
	"APITransactionGenerator/struct/response"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
)

// Developer			Roldan
// @summary 	  		API TRANSACTION GENERATION
// @Description	  		Provides Excel File Report of Transaction in a certain range of date
// @Tags		  		JANUS REPORT GENERATION
// @Produce		  		json
// @Success		  		200 {object} response.TransactionResponse
// @Failure		  		400 {object} response.ResponseModel
// @Router				/public/v1/transaction/fetch_transaction [post]
func TransactionCount(c *fiber.Ctx) error {
	getTransactionResult := []response.TransactionResponse{}
	getUserInput := request.TransactionRequest{}

	if parsErr := c.BodyParser(&getUserInput); parsErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Bad request",
			Data:    parsErr.Error(),
		})
	}

	if fetchErr := database.DBConn.Debug().Raw("SELECT DISTINCT * FROM func_transactionreport(?, ?)", getUserInput.Start_Date, getUserInput.End_Date).Scan(&getTransactionResult).Error; fetchErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "can't fetch table",
			Data:    fetchErr,
		})
	}

	///////////////////////////////////////////////////////////////////////////////////////////
	///                         E X C E L I Z E   -   S H E E T   1                        ///
	/////////////////////////////////////////////////////////////////////////////////////////
	file := excelize.NewFile()

	StreamWritter, err := file.NewStreamWriter("Sheet1")
	if err != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Excelize: Stream Writter Error",
			Data:    err.Error(),
		})
	}

	styleID, err := file.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
		Fill:      excelize.Fill{Type: "pattern", Color: []string{"#DFEBF6"}, Pattern: 1},
	})
	if err != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Excelize: Style ID Error",
			Data:    err.Error(),
		})
	}

	if err := StreamWritter.SetRow("A1", []interface{}{
		excelize.Cell{Value: "TRANSACTION REPORT", StyleID: styleID},
	}, excelize.RowOpts{Height: 30, Hidden: false}); err != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Excelize: Top Row Error",
			Data:    err.Error(),
		})
	}

	header := []interface{}{}
	for _, cell := range []string{
		"AMOUNT", "REFERENCE", "DATE", "TIME", "ACCOUNT NUMBER", "CID", "PARTICULAR", "SOURCE", "TARGET",
		"TOTAL FEE", "AGENT FEE", "BANK FEE", "STANDING",
	} {
		header = append(header, cell)
	}

	if err := StreamWritter.SetRow("A2", header); err != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Excelize: Header Error",
			Data:    err.Error(),
		})
	}

	for rowID := 3; rowID < len(getTransactionResult)+3; rowID++ {
		row := make([]interface{}, 10000)
		// count := 0
		for colID := 0; colID < 13; colID++ {
			row[colID] = getTransactionResult[rowID-3].Got_amount
			colID++
			row[colID] = getTransactionResult[rowID-3].Got_reference
			colID++
			row[colID] = getTransactionResult[rowID-3].Got_date
			colID++
			row[colID] = getTransactionResult[rowID-3].Got_time
			colID++
			row[colID] = getTransactionResult[rowID-3].Got_iid
			colID++
			row[colID] = getTransactionResult[rowID-3].Got_cid
			colID++
			row[colID] = getTransactionResult[rowID-3].Got_particular
			colID++
			row[colID] = getTransactionResult[rowID-3].Got_source_account
			colID++
			row[colID] = getTransactionResult[rowID-3].Got_target
			colID++
			row[colID] = getTransactionResult[rowID-3].Got_total_fee
			colID++
			row[colID] = getTransactionResult[rowID-3].Got_agent_fee
			colID++
			row[colID] = getTransactionResult[rowID-3].Got_bank_fee
			colID++
			row[colID] = getTransactionResult[rowID-3].Got_standing
			// count++
			// colID++

			// fmt.Println(getTransactionResult[colID].Got_reference)
			// fmt.Println()
		}
		cell, _ := excelize.CoordinatesToCellName(1, rowID)
		if err := StreamWritter.SetRow(cell, row); err != nil {
			fmt.Println("RetCode: 400")
			fmt.Println("Message: cannot coordinate to cell name")
			fmt.Println("------------------------------------")
			fmt.Println(err)
			fmt.Println("------------------------------------")
		}
	}

	if err := StreamWritter.MergeCell("A1", "M1"); err != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Excelize: Merging Cells Error",
			Data:    err.Error(),
		})
	}

	if err := StreamWritter.Flush(); err != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Excelize: Flush Error",
			Data:    err.Error(),
		})
	}

	file.SaveAs("./files/" + getUserInput.File_Name)

	return c.JSON(response.ResponseModel{
		RetCode: "200",
		Message: "success",
		Data:    getTransactionResult,
	})
}

/////////////////////////////////////////////////////////////////////////////////////////////////
///                    G  E  T    P  A  T  H    F  U  N  C  T  I  O  N                       ///
///////////////////////////////////////////////////////////////////////////////////////////////

// Developer			Roldan
// @summary 	  		API TRANSACTION GENERATION FILE PATH GETTER
// @Description	  		Excel File Path Getter To Save into the Database when it was Downloaded
// @Tags		  		JANUS REPORT GENERATION
// @Produce		  		json
// @Success		  		200 {object} response.GetCellValuePath
// @Failure		  		400 {object} response.ResponseModel
// @Router				/public/v1/transaction/download_file [post]
func GetPathFunc(c *fiber.Ctx) error {
	fileName := request.FileToDownload{}

	f, err := excelize.OpenFile("./files/" + fileName.FileName)
	if err != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Excelize: Opening File Error",
			Data:    err.Error(),
		})
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println("RetCode: 400")
			fmt.Println("Message: cannot close file: ./files")
			fmt.Println("------------------------------------")
			fmt.Println(err)
			fmt.Println("------------------------------------")
		}
	}()

	// sample link: "http://www.africau.edu/images/default/sample.pdf
	// fmt.Println("DOWNLOAD: ", dl)

	//////////////////////////////////
	///       D O W N L O A D     ///
	////////////////////////////////

	response, err := http.Get("/files/")
	if err != nil {
		// return c.JSON(err)
		fmt.Println("RetCode: 400")
		fmt.Println("Message: cannot get file")
		fmt.Println("------------------------------------")
		fmt.Println(err)
		fmt.Println("------------------------------------")
	}

	defer response.Body.Close()

	file, err := os.Create("newXLSX.xlsx")

	if err != nil {
		// return c.JSON(err)
		fmt.Println("RetCode: 400")
		fmt.Println("Message: cannot download the file")
		fmt.Println("------------------------------------")
		fmt.Println(err)
		fmt.Println("------------------------------------")
	}

	defer file.Close()

	_, err = io.Copy(file, response.Body)

	if err != nil {
		// return c.JSON(err)
		fmt.Println("RetCode: 400")
		fmt.Println("Message: cannot get file data")
		fmt.Println("------------------------------------")
		fmt.Println(err)
		fmt.Println("------------------------------------")
	}

	return c.JSON(fileName)
}
