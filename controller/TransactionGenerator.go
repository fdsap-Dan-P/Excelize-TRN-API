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
	pathGetter := response.FormulaFromDB{}

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

	if fetchErr := database.DBConn.Debug().Table("excel_formula").Select("formula").Where("formula_use = 'path_getter'").Find(&pathGetter).Error; fetchErr != nil {
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
		fmt.Println("RetCode: 400")
		fmt.Println("Message: Error")
		fmt.Println("------------------------------------")
		fmt.Println(err)
		fmt.Println("------------------------------------")
	}

	styleID, err := file.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
		Fill:      excelize.Fill{Type: "pattern", Color: []string{"#DFEBF6"}, Pattern: 1},
	})
	if err != nil {
		fmt.Println("RetCode: 400")
		fmt.Println("Message: cannot set style")
		fmt.Println("------------------------------------")
		fmt.Println(err)
		fmt.Println("------------------------------------")
	}

	if err := StreamWritter.SetRow("A1", []interface{}{
		excelize.Cell{Value: "TRANSACTION REPORT", StyleID: styleID},
	}, excelize.RowOpts{Height: 30, Hidden: false}); err != nil {
		fmt.Println("RetCode: 400")
		fmt.Println("Message: cannot set row")
		fmt.Println("------------------------------------")
		fmt.Println(err)
		fmt.Println("------------------------------------")
	}

	header := []interface{}{}
	for _, cell := range []string{
		"AMOUNT", "REFERENCE", "DATE", "TIME", "ACCOUNT NUMBER", "CID", "PARTICULAR", "SOURCE", "TARGET",
		"TOTAL FEE", "AGENT FEE", "BANK FEE", "STANDING",
	} {
		header = append(header, cell)
	}

	if err := StreamWritter.SetRow("A2", header); err != nil {
		fmt.Println("RetCode: 400")
		fmt.Println("Message: cannot set row")
		fmt.Println("------------------------------------")
		fmt.Println(err)
		fmt.Println("------------------------------------")
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
		fmt.Println("RetCode: 400")
		fmt.Println("Message: cannot merge cell")
		fmt.Println("------------------------------------")
		fmt.Println(err)
		fmt.Println("------------------------------------")
	}

	if err := StreamWritter.Flush(); err != nil {
		fmt.Println("RetCode: 400")
		fmt.Println("Message: Error")
		fmt.Println("------------------------------------")
		fmt.Println(err)
		fmt.Println("------------------------------------")
	}

	//////////////////////////////////////////////////////////////////////////////////////
	/////                           S H E E T    # 2                                /////
	////////////////////////////////////////////////////////////////////////////////////

	file.NewSheet("PATH")

	file.SetCellValue("PATH", "A1", "PATH")

	if err := file.SetCellFormula("PATH", "B1", pathGetter.Formula); err != nil {
		fmt.Println("RetCode: 400")
		fmt.Println("Message: cannot convert into text")
		fmt.Println("------------------------------------")
		fmt.Println(err)
		fmt.Println("------------------------------------")
	}

	if err := file.SaveAs(getUserInput.File_Path + getUserInput.File_Name + ".xlsx"); err != nil {
		fmt.Println("RetCode: 400")
		fmt.Println("Message: cannot save the  file")
		fmt.Println("------------------------------------")
		fmt.Println(err)
		fmt.Println("------------------------------------")
	}

	dl := c.Download("Excelize.xlsx", "Excelize.xlsx")
	fmt.Println("DOWNLOAD: ", dl)
	cPath := c.Path("Excelize.xlsx", "Excelize.xlsx")
	fmt.Println("PATH: ", cPath)

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
// @Router				/public/v1/transaction/get_path [get]
func GetPathFunc(c *fiber.Ctx) error {
	gotPath := response.GetCellValuePath{}

	// if parsErr := c.BodyParser(&gotPath); parsErr != nil {
	// 	return c.JSON(response.ResponseModel{
	// 		RetCode: "400",
	// 		Message: "Bad request",
	// 		Data:    parsErr.Error(),
	// 	})
	// }

	f, err := excelize.OpenFile("Excelize.xlsx")
	if err != nil {
		fmt.Println("RetCode: 400")
		fmt.Println("Message: cannot open file: ", gotPath)
		fmt.Println("------------------------------------")
		fmt.Println(err)
		fmt.Println("------------------------------------")
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println("RetCode: 400")
			fmt.Println("Message: cannot close file: ", gotPath)
			fmt.Println("------------------------------------")
			fmt.Println(err)
			fmt.Println("------------------------------------")
		}
	}()

	getVal, err := f.GetCellValue("PATH", "B1")
	if err != nil {
		fmt.Println("RetCode: 400")
		fmt.Println("Message: cannot get the cell value: PATH")
		fmt.Println("------------------------------------")
		fmt.Println(err)
		fmt.Println("------------------------------------")
	}

	fmt.Println("CELL VALUE: ", getVal)

	if fetchErr := database.DBConn.Debug().Raw("INSERT INTO public.file_path (file_name) VALUES (?)", getVal).Scan(&gotPath).Error; fetchErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "can't inser into table",
			Data:    fetchErr,
		})
	}

	c.SendFile("Excelize.xlsx", true)
	// dl := c.Download("http://www.africau.edu/images/default/sample.pdf", "download_Excelize.xlsx")
	// fmt.Println("DOWNLOAD: ", dl)
	path := c.Path("Excelize.xlsx")
	fmt.Println(path)

	///////////////////////////////////
	///      F I L E   P A T H     ///
	/////////////////////////////////

	fmt.Println("TRY 2nd Commit")
	// filepath.Dir()

	//////////////////////////////////
	///       D O W N L O A D     ///
	////////////////////////////////

	response, err := http.Get("www.africau.edu/images/default/sample.pdf")

	if err != nil {
		return err
	}

	defer response.Body.Close()

	file, err := os.Create("newCsv.csv")

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, response.Body)

	if err != nil {
		return err
	}

	return c.JSON(gotPath)
}
