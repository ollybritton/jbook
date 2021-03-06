package jbook

import (
	"bytes"
	"text/template"

	"github.com/gobuffalo/packr/v2"
)

// Config returns the string representation of the `book.toml` file.
func Config(journal *Journal) string {
	box := packr.New("data", "./data")

	f, err := box.FindString("book.toml")
	if err != nil {
		logger.Fatalf("Cannot get embedded 'book.toml' asset: %v", err)
	}

	tmpl, err := template.New("book.toml").Parse(f)
	if err != nil {
		logger.Fatalf("Cannot get parse template for embedded 'book.toml' asset: %v", err)
	}

	var out bytes.Buffer
	tmpl.Execute(&out, journal)

	return out.String()
}

// Summary returns the string representation of the `SUMMARY.md` file.
func Summary(journal *Journal) string {
	box := packr.New("data", "./data")

	f, err := box.FindString("SUMMARY.md")
	if err != nil {
		logger.Fatalf("Cannot get embedded 'SUMMARY.md' asset: %v", err)
	}

	tmpl, err := template.New("SUMMARY.md").Parse(f)
	if err != nil {
		logger.Fatalf("Cannot get parsed template for embedded 'SUMMARY.md' asset: %v", err)
	}

	type group struct {
		Month   string
		Link    string
		Entries []*Entry
	}

	entries := []group{}

	var curr string
	for _, entry := range journal.Entries {
		if len(entries) == 0 || curr != entry.Date.Format("January 2006") {
			entries = append(entries, group{
				entry.Date.Format("January 2006"),
				"entries/" + entry.FileDate() + ".md",
				[]*Entry{},
			})

			curr = entry.Date.Format("January 2006")
		}

		entries[len(entries)-1].Entries = append(entries[len(entries)-1].Entries, entry)
	}

	var out bytes.Buffer
	tmpl.Execute(&out, entries)

	return out.String()
}

// EntriesReadme creates the entries the string representation of the entries/README.md file.
func EntriesReadme() string {
	box := packr.New("data", "./data")

	f, err := box.FindString("entries/README.md")
	if err != nil {
		logger.Fatalf("Cannot get embedded 'entries/README.md' asset: %v", err)
	}

	return f
}

// FormatEntry creates the string representation of an entry.
func FormatEntry(entry *Entry) string {
	box := packr.New("data", "./data")

	f, err := box.FindString("entries/entry.md")
	if err != nil {
		logger.Fatalf("Cannot get embedded 'entries/README.md' asset: %v", err)
	}

	tmpl, err := template.New("entry.md").Parse(f)
	if err != nil {
		logger.Fatalf("Cannot get parse template for embedded 'entries/README.md' asset: %v", err)
	}

	var out bytes.Buffer
	tmpl.Execute(&out, entry)

	return out.String()
}
