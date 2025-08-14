package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/antchfx/xmlquery"
)

type Mapping struct {
	Header string
	Expr   string
}

func main() {
	//a := app.New()
	a := app.NewWithID("com.lovice.xml2csv")
	w := a.NewWindow("XML → CSV Extractor")
	w.Resize(fyne.NewSize(600, 400))

	selectedXML := widget.NewLabel("No XML selected")
	selectedXML.Wrapping = fyne.TextWrapWord

	chooseXMLBtn := widget.NewButton("Choose XML…", nil)
	generateBtn := widget.NewButton("Generate CSV…", nil)
	generateBtn.Disable()

	var selectedXMLURI fyne.URI
	chooseXMLBtn.OnTapped = func() {
		fd := dialog.NewFileOpen(func(rc fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			if rc == nil {
				return
			}
			defer rc.Close()
			selectedXMLURI = rc.URI()
			selectedXML.SetText(fmt.Sprintf("Selected: %s", selectedXMLURI.Path()))
			generateBtn.Enable()
		}, w)
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".xml"}))
		fd.Show()
	}

	generateBtn.OnTapped = func() {
		if selectedXMLURI == nil {
			dialog.ShowInformation("Missing XML", "Please choose an XML file first.", w)
			return
		}

		sd := dialog.NewFileSave(func(wc fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			if wc == nil {
				return
			}
			defer wc.Close()

			path := wc.URI().Path()
			if filepath.Ext(path) == "" {
				path += ".csv"
			}

			bufWriter := bufio.NewWriter(wc)
			mappings := []Mapping{
				{"ContactName", "delivery-spec/destination/recipient/contact-name"},
				{"Company", "delivery-spec/destination/recipient/company"},
				{"Status", "status"},
				{"AddressLine1", "delivery-spec/destination/recipient/address-line-1"},
				{"City", "delivery-spec/destination/recipient/city"},
				{"ProvinceState", "delivery-spec/destination/recipient/prov-state"},
				{"CustomerRef1", "delivery-spec/reference/customer-ref1"},
				{"Tracking Number", "delivery-spec/reference/item-id"},
			}
			recordXPath := "/delivery-requests/delivery-request"

			if err := convertXMLToCSV(selectedXMLURI.Path(), recordXPath, mappings, bufWriter); err != nil {
				dialog.ShowError(err, w)
				return
			}
			bufWriter.Flush()
			dialog.ShowInformation("Done", "CSV generated successfully.", w)
		}, w)
		sd.SetFileName("output.csv")
		sd.SetFilter(storage.NewExtensionFileFilter([]string{".csv"}))
		sd.Show()
	}

	form := container.NewVBox(
		chooseXMLBtn,
		selectedXML,
		widget.NewSeparator(),
		generateBtn,
	)

	w.SetContent(container.NewCenter(form))
	w.ShowAndRun()
}

func convertXMLToCSV(xmlPath, recordXPath string, mappings []Mapping, w io.Writer) error {
	reader, err := os.Open(xmlPath)
	if err != nil {
		return fmt.Errorf("failed to open XML: %w", err)
	}
	defer reader.Close()

	doc, err := xmlquery.Parse(reader)
	if err != nil {
		return fmt.Errorf("failed to parse XML: %w", err)
	}

	records := xmlquery.Find(doc, recordXPath)
	if len(records) == 0 {
		return fmt.Errorf("no nodes matched Record XPath: %s", recordXPath)
	}

	csvw := csv.NewWriter(w)
	defer csvw.Flush()

	headers := make([]string, len(mappings))
	for i, m := range mappings {
		headers[i] = m.Header
	}
	if err := csvw.Write(headers); err != nil {
		return err
	}

	for _, rec := range records {
		row := make([]string, len(mappings))
		for i, m := range mappings {
			val := extractValue(rec, m.Expr)
			if m.Header == "Tracking Number" {
				val = "'" + val
			}
			row[i] = val
		}
		if err := csvw.Write(row); err != nil {
			return err
		}
	}

	return nil
}

func extractValue(node *xmlquery.Node, expr string) string {
	if strings.HasPrefix(expr, "@") {
		attr := strings.TrimPrefix(expr, "@")
		return node.SelectAttr(attr)
	}
	if child := xmlquery.FindOne(node, expr); child != nil {
		return child.InnerText()
	}
	return ""
}
