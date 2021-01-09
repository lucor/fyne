package widget_test

import (
	"testing"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/test"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"

	"github.com/stretchr/testify/assert"
)

func TestAccordion(t *testing.T) {
	ai := widget.NewAccordionItem("foo", widget.NewLabel("foobar"))
	t.Run("Initializer", func(t *testing.T) {
		ac := &widget.Accordion{Items: []*widget.AccordionItem{ai}}
		assert.Equal(t, 1, len(ac.Items))
	})
	t.Run("Constructor", func(t *testing.T) {
		ac := widget.NewAccordion(ai)
		assert.Equal(t, 1, len(ac.Items))
	})
}

func TestAccordion_Append(t *testing.T) {
	ac := widget.NewAccordion()
	ac.Append(widget.NewAccordionItem("foo", widget.NewLabel("foobar")))
	assert.Equal(t, 1, len(ac.Items))
}

func TestAccordion_ChangeTheme(t *testing.T) {
	test.NewApp()
	defer test.NewApp()

	ac := widget.NewAccordion()
	ac.Append(widget.NewAccordionItem("foo0", widget.NewLabel("foobar0")))
	ac.Append(widget.NewAccordionItem("foo1", widget.NewLabel("foobar1")))

	w := test.NewWindow(ac)
	defer w.Close()
	w.Resize(ac.MinSize().Add(fyne.NewSize(theme.Padding()*2, theme.Padding()*2)))

	test.AssertImageMatches(t, "accordion/theme_initial.png", w.Canvas().Capture())

	test.WithTestTheme(t, func() {
		w.Resize(ac.MinSize().Add(fyne.NewSize(theme.Padding()*2, theme.Padding()*2)))
		ac.Refresh()
		time.Sleep(100 * time.Millisecond)
		test.AssertImageMatches(t, "accordion/theme_changed.png", w.Canvas().Capture())
	})
}

func TestAccordion_Close(t *testing.T) {
	t.Run("Exists", func(t *testing.T) {
		ac := widget.NewAccordion()
		ac.Append(&widget.AccordionItem{
			Title:  "foo",
			Detail: widget.NewLabel("foobar"),
			Open:   true,
		})
		ac.Close(0)
		assert.False(t, ac.Items[0].Open)
		assert.False(t, ac.Items[0].Detail.Visible())
	})
	t.Run("BelowBounds", func(t *testing.T) {
		ac := widget.NewAccordion()
		ac.Append(&widget.AccordionItem{
			Title:  "foo",
			Detail: widget.NewLabel("foobar"),
			Open:   true,
		})
		ac.Close(-1)
		assert.True(t, ac.Items[0].Open)
	})
	t.Run("AboveBounds", func(t *testing.T) {
		ac := widget.NewAccordion()
		ac.Append(&widget.AccordionItem{
			Title:  "foo",
			Detail: widget.NewLabel("foobar"),
			Open:   true,
		})
		ac.Close(1)
		assert.True(t, ac.Items[0].Open)
	})
}

func TestAccordion_CloseAll(t *testing.T) {
	ac := widget.NewAccordion()
	ac.Append(widget.NewAccordionItem("foo0", widget.NewLabel("foobar0")))
	ac.Append(widget.NewAccordionItem("foo1", widget.NewLabel("foobar1")))
	ac.Append(widget.NewAccordionItem("foo2", widget.NewLabel("foobar2")))

	ac.CloseAll()
	assert.False(t, ac.Items[0].Open)
	assert.False(t, ac.Items[1].Open)
	assert.False(t, ac.Items[2].Open)
}

func TestAccordion_Layout(t *testing.T) {
	test.NewApp()

	for name, tt := range map[string]struct {
		multiOpen bool
		items     []*widget.AccordionItem
		opened    []int
		want      string
	}{
		"single_open_one_item": {
			items: []*widget.AccordionItem{
				{
					Title:  "A",
					Detail: widget.NewLabel("11111"),
				},
			},
			want: `
				<canvas padded size="150x200">
					<content>
						<container pos="4,4" size="142x192">
							<widget pos="45,77" size="52x37" type="*widget.Accordion">
								<widget pos="0,4" size="52x29" type="*widget.Button">
									<rectangle fillColor="button" pos="2,2" size="48x25"/>
									<text bold pos="28,4" size="11x21">A</text>
									<image fillMode="contain" pos="4,4" rsc="menuDropDownIcon" size="20x21"/>
								</widget>
							</widget>
						</container>
					</content>
				</canvas>
			`,
		},
		"single_open_one_item_opened": {
			items: []*widget.AccordionItem{
				{
					Title:  "A",
					Detail: widget.NewLabel("11111"),
				},
			},
			opened: []int{0},
			want: `
				<canvas padded size="150x200">
					<content>
						<container pos="4,4" size="142x192">
							<widget pos="45,61" size="52x70" type="*widget.Accordion">
								<widget pos="0,4" size="52x29" type="*widget.Button">
									<rectangle fillColor="button" pos="2,2" size="48x25"/>
									<text bold pos="28,4" size="11x21">A</text>
									<image fillMode="contain" pos="4,4" rsc="menuDropUpIcon" size="20x21"/>
								</widget>
								<widget pos="0,37" size="52x29" type="*widget.Label">
									<text pos="4,4" size="44x21">11111</text>
								</widget>
							</widget>
						</container>
					</content>
				</canvas>
			`,
		},
		"single_open_multiple_items": {
			items: []*widget.AccordionItem{
				{
					Title:  "A",
					Detail: widget.NewLabel("11111"),
				},
				{
					Title:  "B",
					Detail: widget.NewLabel("2222222222"),
				},
			},
			want: `
				<canvas padded size="150x200">
					<content>
						<container pos="4,4" size="142x192">
							<widget pos="23,58" size="95x75" type="*widget.Accordion">
								<widget pos="0,4" size="95x29" type="*widget.Button">
									<rectangle fillColor="button" pos="2,2" size="91x25"/>
									<text bold pos="28,4" size="11x21">A</text>
									<image fillMode="contain" pos="4,4" rsc="menuDropDownIcon" size="20x21"/>
								</widget>
								<widget pos="0,42" size="95x29" type="*widget.Button">
									<rectangle fillColor="button" pos="2,2" size="91x25"/>
									<text bold pos="28,4" size="11x21">B</text>
									<image fillMode="contain" pos="4,4" rsc="menuDropDownIcon" size="20x21"/>
								</widget>
								<widget pos="0,37" size="95x1" type="*widget.Separator">
									<rectangle fillColor="disabled" size="95x1"/>
								</widget>
							</widget>
						</container>
					</content>
				</canvas>
			`,
		},
		"single_open_multiple_items_opened": {
			items: []*widget.AccordionItem{
				{
					Title:  "A",
					Detail: widget.NewLabel("11111"),
				},
				{
					Title:  "B",
					Detail: widget.NewLabel("2222222222"),
				},
			},
			opened: []int{0, 1},
			want: `
				<canvas padded size="150x200">
					<content>
						<container pos="4,4" size="142x192">
							<widget pos="23,42" size="95x108" type="*widget.Accordion">
								<widget pos="0,4" size="95x29" type="*widget.Button">
									<rectangle fillColor="button" pos="2,2" size="91x25"/>
									<text bold pos="28,4" size="11x21">A</text>
									<image fillMode="contain" pos="4,4" rsc="menuDropDownIcon" size="20x21"/>
								</widget>
								<widget pos="0,42" size="95x29" type="*widget.Button">
									<rectangle fillColor="button" pos="2,2" size="91x25"/>
									<text bold pos="28,4" size="11x21">B</text>
									<image fillMode="contain" pos="4,4" rsc="menuDropUpIcon" size="20x21"/>
								</widget>
								<widget pos="0,75" size="95x29" type="*widget.Label">
									<text pos="4,4" size="87x21">2222222222</text>
								</widget>
								<widget pos="0,37" size="95x1" type="*widget.Separator">
									<rectangle fillColor="disabled" size="95x1"/>
								</widget>
							</widget>
						</container>
					</content>
				</canvas>
			`,
		},
		"multiple_open_one_item": {
			multiOpen: true,
			items: []*widget.AccordionItem{
				{
					Title:  "A",
					Detail: widget.NewLabel("11111"),
				},
			},
			want: `
				<canvas padded size="150x200">
					<content>
						<container pos="4,4" size="142x192">
							<widget pos="45,77" size="52x37" type="*widget.Accordion">
								<widget pos="0,4" size="52x29" type="*widget.Button">
									<rectangle fillColor="button" pos="2,2" size="48x25"/>
									<text bold pos="28,4" size="11x21">A</text>
									<image fillMode="contain" pos="4,4" rsc="menuDropDownIcon" size="20x21"/>
								</widget>
							</widget>
						</container>
					</content>
				</canvas>
			`,
		},
		"multiple_open_one_item_opened": {
			multiOpen: true,
			items: []*widget.AccordionItem{
				{
					Title:  "A",
					Detail: widget.NewLabel("11111"),
				},
			},
			opened: []int{0},
			want: `
				<canvas padded size="150x200">
					<content>
						<container pos="4,4" size="142x192">
							<widget pos="45,61" size="52x70" type="*widget.Accordion">
								<widget pos="0,4" size="52x29" type="*widget.Button">
									<rectangle fillColor="button" pos="2,2" size="48x25"/>
									<text bold pos="28,4" size="11x21">A</text>
									<image fillMode="contain" pos="4,4" rsc="menuDropUpIcon" size="20x21"/>
								</widget>
								<widget pos="0,37" size="52x29" type="*widget.Label">
									<text pos="4,4" size="44x21">11111</text>
								</widget>
							</widget>
						</container>
					</content>
				</canvas>
			`,
		},
		"multiple_open_multiple_items": {
			multiOpen: true,
			items: []*widget.AccordionItem{
				{
					Title:  "A",
					Detail: widget.NewLabel("11111"),
				},
				{
					Title:  "B",
					Detail: widget.NewLabel("2222222222"),
				},
			},
			want: `
				<canvas padded size="150x200">
					<content>
						<container pos="4,4" size="142x192">
							<widget pos="23,58" size="95x75" type="*widget.Accordion">
								<widget pos="0,4" size="95x29" type="*widget.Button">
									<rectangle fillColor="button" pos="2,2" size="91x25"/>
									<text bold pos="28,4" size="11x21">A</text>
									<image fillMode="contain" pos="4,4" rsc="menuDropDownIcon" size="20x21"/>
								</widget>
								<widget pos="0,42" size="95x29" type="*widget.Button">
									<rectangle fillColor="button" pos="2,2" size="91x25"/>
									<text bold pos="28,4" size="11x21">B</text>
									<image fillMode="contain" pos="4,4" rsc="menuDropDownIcon" size="20x21"/>
								</widget>
								<widget pos="0,37" size="95x1" type="*widget.Separator">
									<rectangle fillColor="disabled" size="95x1"/>
								</widget>
							</widget>
						</container>
					</content>
				</canvas>
			`,
		},
		"multiple_open_multiple_items_opened": {
			multiOpen: true,
			items: []*widget.AccordionItem{
				{
					Title:  "A",
					Detail: widget.NewLabel("11111"),
				},
				{
					Title:  "B",
					Detail: widget.NewLabel("2222222222"),
				},
			},
			opened: []int{0, 1},
			want: `
				<canvas padded size="150x200">
					<content>
						<container pos="4,4" size="142x192">
							<widget pos="23,25" size="95x141" type="*widget.Accordion">
								<widget pos="0,4" size="95x29" type="*widget.Button">
									<rectangle fillColor="button" pos="2,2" size="91x25"/>
									<text bold pos="28,4" size="11x21">A</text>
									<image fillMode="contain" pos="4,4" rsc="menuDropUpIcon" size="20x21"/>
								</widget>
								<widget pos="0,75" size="95x29" type="*widget.Button">
									<rectangle fillColor="button" pos="2,2" size="91x25"/>
									<text bold pos="28,4" size="11x21">B</text>
									<image fillMode="contain" pos="4,4" rsc="menuDropUpIcon" size="20x21"/>
								</widget>
								<widget pos="0,37" size="95x29" type="*widget.Label">
									<text pos="4,4" size="87x21">11111</text>
								</widget>
								<widget pos="0,108" size="95x29" type="*widget.Label">
									<text pos="4,4" size="87x21">2222222222</text>
								</widget>
								<widget pos="0,70" size="95x1" type="*widget.Separator">
									<rectangle fillColor="disabled" size="95x1"/>
								</widget>
							</widget>
						</container>
					</content>
				</canvas>
			`,
		},
	} {
		t.Run(name, func(t *testing.T) {
			accordion := &widget.Accordion{
				MultiOpen: tt.multiOpen,
			}
			for _, ai := range tt.items {
				accordion.Append(ai)
			}
			for _, o := range tt.opened {
				accordion.Open(o)
			}

			window := test.NewWindow(fyne.NewContainerWithLayout(layout.NewCenterLayout(), accordion))
			window.Resize(accordion.MinSize().Max(fyne.NewSize(150, 200)))

			test.AssertRendersToMarkup(t, tt.want, window.Canvas())

			window.Close()
		})
	}
}

func TestAccordion_MinSize(t *testing.T) {
	minSizeA := fyne.MeasureText("A", theme.TextSize(), fyne.TextStyle{})
	minSizeA.Width += theme.IconInlineSize() + theme.Padding()*5
	minSizeA.Height = fyne.Max(minSizeA.Height, theme.IconInlineSize()) + theme.Padding()*2
	minSizeB := fyne.MeasureText("B", theme.TextSize(), fyne.TextStyle{})
	minSizeB.Width += theme.IconInlineSize() + theme.Padding()*5
	minSizeB.Height = fyne.Max(minSizeB.Height, theme.IconInlineSize()) + theme.Padding()*2
	minSize1 := fyne.MeasureText("111111", theme.TextSize(), fyne.TextStyle{})
	minSize1.Width += theme.Padding() * 2
	minSize1.Height += theme.Padding() * 2
	minSize2 := fyne.MeasureText("2222222222", theme.TextSize(), fyne.TextStyle{})
	minSize2.Width += theme.Padding() * 2
	minSize2.Height += theme.Padding() * 2

	minWidthA1 := fyne.Max(minSizeA.Width, minSize1.Width)
	minWidthB2 := fyne.Max(minSizeB.Width, minSize2.Width)
	minWidthA1B2 := fyne.Max(minWidthA1, minWidthB2)

	minHeightA1 := minSizeA.Height + minSize1.Height + theme.Padding()

	for name, tt := range map[string]struct {
		multiOpen bool
		items     []*widget.AccordionItem
		opened    []int
		want      fyne.Size
	}{
		"single_open_one_item": {
			items: []*widget.AccordionItem{
				{
					Title:  "A",
					Detail: widget.NewLabel("111111"),
				},
			},
			want: fyne.NewSize(minWidthA1, minSizeA.Height+theme.Padding()*2),
		},
		"single_open_one_item_opened": {
			items: []*widget.AccordionItem{
				{
					Title:  "A",
					Detail: widget.NewLabel("111111"),
				},
			},
			opened: []int{0},
			want:   fyne.NewSize(minWidthA1, minHeightA1+theme.Padding()*2),
		},
		"single_open_multiple_items": {
			items: []*widget.AccordionItem{
				{
					Title:  "A",
					Detail: widget.NewLabel("111111"),
				},
				{
					Title:  "B",
					Detail: widget.NewLabel("2222222222"),
				},
			},
			want: fyne.NewSize(minWidthA1B2, minSizeA.Height+minSizeB.Height+theme.Padding()*4+1),
		},
		"single_open_multiple_items_opened": {
			items: []*widget.AccordionItem{
				{
					Title:  "A",
					Detail: widget.NewLabel("111111"),
				},
				{
					Title:  "B",
					Detail: widget.NewLabel("2222222222"),
				},
			},
			opened: []int{0, 1},
			want:   fyne.NewSize(minWidthA1B2, minSizeA.Height+minSizeB.Height+minSize2.Height+theme.Padding()*5+1),
		},
		"multiple_open_one_item": {
			multiOpen: true,
			items: []*widget.AccordionItem{
				{
					Title:  "A",
					Detail: widget.NewLabel("111111"),
				},
			},
			want: fyne.NewSize(minWidthA1, minSizeA.Height+theme.Padding()*2),
		},
		"multiple_open_one_item_opened": {
			multiOpen: true,
			items: []*widget.AccordionItem{
				{
					Title:  "A",
					Detail: widget.NewLabel("111111"),
				},
			},
			opened: []int{0},
			want:   fyne.NewSize(minWidthA1, minHeightA1+theme.Padding()*2),
		},
		"multiple_open_multiple_items": {
			multiOpen: true,
			items: []*widget.AccordionItem{
				{
					Title:  "A",
					Detail: widget.NewLabel("111111"),
				},
				{
					Title:  "B",
					Detail: widget.NewLabel("2222222222"),
				},
			},
			want: fyne.NewSize(minWidthA1B2, minSizeA.Height+minSizeB.Height+theme.Padding()*4+1),
		},
		"multiple_open_multiple_items_opened": {
			multiOpen: true,
			items: []*widget.AccordionItem{
				{
					Title:  "A",
					Detail: widget.NewLabel("111111"),
				},
				{
					Title:  "B",
					Detail: widget.NewLabel("2222222222"),
				},
			},
			opened: []int{0, 1},
			want:   fyne.NewSize(minWidthA1B2, minSizeA.Height+minSizeB.Height+minSize1.Height+minSize2.Height+theme.Padding()*6+1),
		},
	} {
		t.Run(name, func(t *testing.T) {
			accordion := &widget.Accordion{
				MultiOpen: tt.multiOpen,
			}
			for _, ai := range tt.items {
				accordion.Append(ai)
			}
			for _, o := range tt.opened {
				accordion.Open(o)
			}

			assert.Equal(t, tt.want, accordion.MinSize())
		})
	}
}

func TestAccordion_Open(t *testing.T) {
	t.Run("Exists", func(t *testing.T) {
		ac := widget.NewAccordion()
		ac.Append(widget.NewAccordionItem("foo0", widget.NewLabel("foobar0")))
		ac.Append(widget.NewAccordionItem("foo1", widget.NewLabel("foobar1")))
		ac.Append(widget.NewAccordionItem("foo2", widget.NewLabel("foobar2")))
		assert.False(t, ac.Items[0].Open)
		assert.False(t, ac.Items[1].Open)
		assert.False(t, ac.Items[2].Open)

		ac.Open(0)
		assert.True(t, ac.Items[0].Open)
		assert.False(t, ac.Items[1].Open)
		assert.False(t, ac.Items[2].Open)

		// Opening index 1 should close index 0
		ac.Open(1)
		assert.False(t, ac.Items[0].Open)
		assert.True(t, ac.Items[1].Open)
		assert.False(t, ac.Items[2].Open)

		ac.MultiOpen = true
		ac.Open(2)
		// Opening index 2 should not close index 1
		assert.False(t, ac.Items[0].Open)
		assert.True(t, ac.Items[1].Open)
		assert.True(t, ac.Items[2].Open)
	})
	t.Run("BelowBounds", func(t *testing.T) {
		ac := widget.NewAccordion()
		ac.Append(widget.NewAccordionItem("foo", widget.NewLabel("foobar")))
		assert.False(t, ac.Items[0].Open)
		ac.Open(-1)
		assert.False(t, ac.Items[0].Open)
	})
	t.Run("AboveBounds", func(t *testing.T) {
		ac := widget.NewAccordion()
		ac.Append(widget.NewAccordionItem("foo", widget.NewLabel("foobar")))
		assert.False(t, ac.Items[0].Open)
		ac.Open(1)
		assert.False(t, ac.Items[0].Open)
	})
}

func TestAccordion_OpenAll(t *testing.T) {
	ac := widget.NewAccordion()
	ac.Append(widget.NewAccordionItem("foo0", widget.NewLabel("foobar0")))
	ac.Append(widget.NewAccordionItem("foo1", widget.NewLabel("foobar1")))
	ac.Append(widget.NewAccordionItem("foo2", widget.NewLabel("foobar2")))

	ac.OpenAll()
	// Cannot open all items if !accordion.MultiOpen
	assert.False(t, ac.Items[0].Open)
	assert.False(t, ac.Items[1].Open)
	assert.False(t, ac.Items[2].Open)

	ac.MultiOpen = true
	ac.OpenAll()
	// All items should be open
	assert.True(t, ac.Items[0].Open)
	assert.True(t, ac.Items[1].Open)
	assert.True(t, ac.Items[2].Open)
}

func TestAccordion_Remove(t *testing.T) {
	ai := widget.NewAccordionItem("foo", widget.NewLabel("foobar"))
	ac := widget.NewAccordion(ai)
	ac.Remove(ai)
	assert.Equal(t, 0, len(ac.Items))
}

func TestAccordion_RemoveIndex(t *testing.T) {
	for name, tt := range map[string]struct {
		index  int
		length int
	}{
		"Exists":      {index: 0, length: 0},
		"BelowBounds": {index: -1, length: 1},
		"AboveBounds": {index: 1, length: 1},
	} {
		t.Run(name, func(t *testing.T) {
			ac := widget.NewAccordion()
			ac.Append(widget.NewAccordionItem("foo", widget.NewLabel("foobar")))
			ac.RemoveIndex(tt.index)
			assert.Equal(t, tt.length, len(ac.Items))
		})
	}
}
