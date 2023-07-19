package todo

import (
	"fmt"
	"github.com/alexeyco/simpletable"
	"time"
)

// List all to-dos
func (t *TodoList) List() {
	// create new table
	table := simpletable.New()

	// set table header
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Completed?"},
			{Align: simpletable.AlignCenter, Text: "CreatedAt"},
			{Align: simpletable.AlignCenter, Text: "CompletedAt"},
		},
	}

	// define table body cells
	var cells [][]*simpletable.Cell

	// add table body items
	for index, item := range *t {
		index++

		task := item.Task
		done := "Yes"
		createdAt := item.CreatedAt.Format(time.RFC822)
		completedAt := item.CompletedAt.Format(time.RFC822)

		if !item.Done {
			done = "No"
			completedAt = item.CompletedAt.Format("###################")
		}

		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", index)},
			{Text: task},
			{Text: done},
			{Text: createdAt},
			{Text: completedAt},
		})
	}

	// set table body
	table.Body = &simpletable.Body{Cells: cells}

	// set table footer
	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 5, Text: fmt.Sprintf("You have %d pending to-do(s)", t.CountPending())},
	}}

	// set table style
	table.SetStyle(simpletable.StyleUnicode)

	// print table
	table.Println()
}

// List completed to-do's
func (t *TodoList) ListCompleted() {
	// create new table
	table := simpletable.New()

	// set table header
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "CreatedAt"},
			{Align: simpletable.AlignCenter, Text: "CompletedAt"},
		},
	}

	// define table body cells
	var cells [][]*simpletable.Cell

	// add table body items
	for index, item := range *t {
		if item.Done {
			index++

			task := item.Task
			createdAt := item.CreatedAt.Format(time.RFC822)
			completedAt := item.CompletedAt.Format(time.RFC822)

			cells = append(cells, *&[]*simpletable.Cell{
				{Text: fmt.Sprintf("%d", index)},
				{Text: task},
				{Text: createdAt},
				{Text: completedAt},
			})
		} else {
			continue
		}
	}

	// set table body
	table.Body = &simpletable.Body{Cells: cells}

	// set table style
	table.SetStyle(simpletable.StyleUnicode)

	// print table
	table.Println()
}
