/*
//466f0f67fe8cecc689d22b68118d40a0c3d0747b98c3020482926d01db16d6fc
 * unipdf_license_loading_offline.go:
 * Illustrates how to load an offline (perpetual) license key.
 * Offline keys can be purchased at https://www.unidoc.io
 *
 * Run as: go run unipdf_license_loading_offline.go
*/

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/unidoc/unipdf/v3/common/license"
	"github.com/unidoc/unipdf/v3/creator"
	"github.com/unidoc/unipdf/v3/model"
)

// Example of an offline perpetual license key.
const offlineLicenseKey = `
 -----BEGIN UNIDOC LICENSE KEY-----
 466f0f67fe8cecc689d22b68118d40a0c3d0747b98c3020482926d01db16d6fc
 -----END UNIDOC LICENSE KEY-----
 `

/*
func init() {
	// The customer name needs to match the entry that is embedded in the signed key.
	customerName := `CJ001`

	// Good to load the license key in `init`. Needs to be done prior to using the library, otherwise operations
	// will result in an error.
	err := license.SetLicenseKey(offlineLicenseKey, customerName)
	if err != nil {
		panic(err)
	}
}*/

func init() {
	// To get your free API key for metered license, sign up on: https://cloud.unidoc.io
	// Make sure to be using UniPDF v3.19.1 or newer for Metered API key support.
	err := license.SetMeteredKey("466f0f67fe8cecc689d22b68118d40a0c3d0747b98c3020482926d01db16d6fc")
	if err != nil {
		fmt.Printf("ERROR: Failed to set metered key: %v\n", err)
		fmt.Printf("Make sure to get a valid key from https://cloud.unidoc.io\n")
		panic(err)
	}
}

func main() {
	c := creator.New()
	c.SetPageMargins(50, 50, 100, 70)

	helvetica, _ := model.NewStandard14Font("Helvetica")
	helveticaBold, _ := model.NewStandard14Font("Helvetica-Bold")

	p := c.NewParagraph("UniDoc")
	p.SetFont(helvetica)
	p.SetFontSize(48)
	p.SetMargins(15, 0, 150, 0)
	p.SetColor(creator.ColorRGBFrom8bit(56, 68, 77))
	c.Draw(p)

	p = c.NewParagraph("Example Page")
	p.SetFont(helveticaBold)
	p.SetFontSize(30)
	p.SetMargins(15, 0, 0, 0)
	p.SetColor(creator.ColorRGBFrom8bit(45, 148, 215))
	c.Draw(p)

	t := time.Now().UTC()
	dateStr := t.Format("1 Jan, 2006 15:04")

	p = c.NewParagraph(dateStr)
	p.SetFont(helveticaBold)
	p.SetFontSize(12)
	p.SetMargins(15, 0, 5, 60)
	p.SetColor(creator.ColorRGBFrom8bit(56, 68, 77))
	c.Draw(p)

	loremTxt := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt" +
		"ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut " +
		"aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore" +
		"eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt " +
		"mollit anim id est laborum."

	p = c.NewParagraph(loremTxt)
	p.SetFontSize(16)
	p.SetColor(creator.ColorBlack)
	p.SetLineHeight(1.5)
	p.SetMargins(0, 0, 5, 0)
	p.SetTextAlignment(creator.TextAlignmentJustify)
	c.Draw(p)

	err := c.WriteToFile("report.pdf")
	if err != nil {
		log.Println("Write file error:", err)
	}
}
