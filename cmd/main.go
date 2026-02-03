package main

import (
    "strings"

    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
)

type model struct {
    width, height int
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.WindowSizeMsg:
        m.width = msg.Width
        m.height = msg.Height
    case tea.KeyMsg:
        if msg.String() == "ctrl+q" {
            return m, tea.Quit
        }
    }
    return m, nil
}

func (m model) View() string {
    // define the splits
    leftWidth := int(float64(m.width) * 0.3)
    rightWidth := m.width - leftWidth - 2 // -2 for borders

    // left side
    leftBox := lipgloss.NewStyle().
        Width(leftWidth).
        Height(m.height).
        BorderRight(false).
        BorderStyle(lipgloss.NormalBorder()).
        Render("\n  PORTFOLIO VALUE\n\n  $178,375.58")

    // right side
    // hard coded example data
    list := strings.Join([]string{
        "SPY $478.20",
        "QQQ $409.10",
        "IWM $198.50",
        "AAPL $185.90",
        "BTC $42,100",
    }, "\n")

    // render right with align to right
    rightBox := lipgloss.NewStyle().
        Width(rightWidth).
        Align(lipgloss.Right). // here is the align
        Render(list)

    // join together
    return lipgloss.JoinHorizontal(lipgloss.Top, leftBox, rightBox)
}

func main() {
    p := tea.NewProgram(model{}, tea.WithAltScreen())
    p.Run()
}