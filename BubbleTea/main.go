package main

// A simple program demonstrating the paginator component from the Bubbles
// component library.

import (
	"fmt"
	"log"
	"strings"

	"charm.land/bubbles/v2/paginator"
	"charm.land/lipgloss/v2"

	tea "charm.land/bubbletea/v2"
)

type styles struct {
	activeDot   lipgloss.Style
	inactiveDot lipgloss.Style
}

func newStyles(bgIsDark bool) (s styles) {
	lightDark := lipgloss.LightDark(bgIsDark)

	s.activeDot = lipgloss.NewStyle().Foreground(lightDark(lipgloss.Color("235"), lipgloss.Color("252"))).SetString("•")
	s.inactiveDot = s.activeDot.Foreground(lightDark(lipgloss.Color("250"), lipgloss.Color("238"))).SetString("•")
	return s
}

type model struct {
	items     []string
	paginator paginator.Model
}

func newModel() model {
	var items []string
	for i := 1; i < 101; i++ {
		text := fmt.Sprintf("Item %d", i)
		items = append(items, text)
	}

	p := paginator.New()
	p.Type = paginator.Dots
	p.PerPage = 7
	p.SetTotalPages(len(items))

	m := model{
		paginator: p,
		items:     items,
	}

	m.updateStyles(true) // default to dark styles
	return m
}

func (m *model) updateStyles(isDark bool) {
	styles := newStyles(isDark)
	m.paginator.ActiveDot = styles.activeDot.String()
	m.paginator.InactiveDot = styles.inactiveDot.String()
}

func (m model) Init() tea.Cmd {
	return tea.RequestBackgroundColor
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.BackgroundColorMsg:
		m.updateStyles(msg.IsDark())
		return m, nil
	// ─── マウススクロールのハンドリングを追加 ───
    case tea.MouseWheelMsg:
		mouse := msg.Mouse()
        switch mouse.Button {
        case tea.MouseWheelUp: // 上スクロールで前のページへ
            m.paginator.PrevPage()
            return m, nil
        case tea.MouseWheelDown: // 下スクロールで次のページへ
            m.paginator.NextPage()
            return m, nil
        }
	case tea.KeyPressMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
			// ─── カスタムキーバインドの追加 ───
        case "n": // 'n' キーで次のページへ
            m.paginator.NextPage()
            return m, nil
        case "p": // 'p' キーで前のページへ
            m.paginator.PrevPage()
            return m, nil
			//─── 5ページ進む (Shift + j) ───
   		case "J":
        	targetPage := m.paginator.Page + 5
        	// 総ページ数を超えないように丸める
        	if targetPage >= m.paginator.TotalPages {
         	   targetPage = m.paginator.TotalPages - 1
    		}
        	m.paginator.Page = targetPage
        	return m, nil

   			// ─── 5ページ戻る (Shift + k) ───
    	case "K":
			targetPage := m.paginator.Page - 5
       		// 0ページ未満にならないように丸める
        	if targetPage < 0 {
       	    	targetPage = 0
        	}
        	m.paginator.Page = targetPage
        	return m, nil
    	}
	}
	m.paginator, cmd = m.paginator.Update(msg)
	return m, cmd
}

func (m model) View() tea.View {
	var b strings.Builder
	b.WriteString("\n  Paginator Example\n\n")
	start, end := m.paginator.GetSliceBounds(len(m.items))
	for _, item := range m.items[start:end] {
		b.WriteString("  • " + item + "\n\n")
	}
	b.WriteString("  " + m.paginator.View())
	b.WriteString("\n\n  h/l ←/→ page • q: quit\n")
	view := tea.NewView(b.String())
	view.MouseMode = tea.MouseModeCellMotion
	return view
}

func main() {
	p := tea.NewProgram(newModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}