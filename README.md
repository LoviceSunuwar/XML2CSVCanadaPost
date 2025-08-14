# XML → CSV Extractor

A simple GUI application to convert Canada Post XML shipment files into CSV format for easy readability and analysis.

---

## What the project does

This application allows users to:

- Select a Canada Post XML file containing shipment data.  
- Extract relevant fields like Contact Name, Company, Address, Tracking Number, etc.  
- Save the extracted data as a **CSV file** that can be opened in Excel or any spreadsheet software.  

The tool is designed to simplify the process of converting XML to CSV without requiring knowledge of command-line tools.

---

## Third-party packages used

1. **[Fyne](https://fyne.io/)**  
   - Used to build the **GUI** for the application.  
   - Provides cross-platform windows, buttons, file dialogs, and more.  
   - Chosen because it is lightweight and easy to use for desktop GUI apps in Go.  

2. **[xmlquery](https://github.com/antchfx/xmlquery)**  
   - Used to **parse XML** files and perform XPath queries.  
   - Makes it easy to extract specific elements from structured XML documents.  

---

## Installation

1. **Clone the repository**

```bash
git clone https://github.com/yourusername/XML2CSVCanadaPost.git
cd XML2CSVCanadaPost


# XML → CSV Extractor

A simple GUI application to convert Canada Post XML shipment files into CSV format for easy readability and analysis.

---

## What the project does

This application allows users to:

- Select a Canada Post XML file containing shipment data.  
- Extract relevant fields like Contact Name, Company, Address, Tracking Number, etc.  
- Save the extracted data as a **CSV file** that can be opened in Excel or any spreadsheet software.  

The tool is designed to simplify the process of converting XML to CSV without requiring knowledge of command-line tools.

---

## Third-party packages used

1. **[Fyne](https://fyne.io/)**  
   - Used to build the **GUI** for the application.  
   - Provides cross-platform windows, buttons, file dialogs, and more.  
   - Chosen because it is lightweight and easy to use for desktop GUI apps in Go.  

2. **[xmlquery](https://github.com/antchfx/xmlquery)**  
   - Used to **parse XML** files and perform XPath queries.  
   - Makes it easy to extract specific elements from structured XML documents.  

---

## Installation

1. **Clone the repository**

```bash
git clone https://github.com/yourusername/XML2CSVCanadaPost.git
cd XML2CSVCanadaPost

2. Install Go dependencies
- go mod download
3. Run the application
- On macOS or Linux:
- go run .
- On Windows (if Go is installed):
- go run .
- Build a standalone executable
- macOS/Linux:
- go build -o XmlToCsv
- Windows:
- go build -o XmlToCsv.exe
- The .exe file can be shared with others and run without installing Go.

## How to use it
- Launch the application (XmlToCsv or XmlToCsv.exe).
- Click “Choose XML…” and select your Canada Post XML file.
- Click “Generate CSV…” and choose where to save the CSV file.
- The app will extract the relevant fields and save them in the CSV format.

## Story time
One of my mates shared a story about how they were handling documenting shipments every day.
While UPS had an export CSV functionality, Canada Post didn’t.
The only export option was XML, which was unreadable for him.
That’s when I thought: “This just needs a simple XML → CSV converter.”
My friend has a different interest and doesn’t like command-line applications, so I decided he needed a GUI.
This is where Fyne came in — a lightweight GUI framework for Go that made it easy to turn a simple converter into a friendly desktop application.