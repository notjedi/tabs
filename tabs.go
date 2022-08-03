package tabs

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	tabTitles  []string
	tabModels  []tea.Model
	totalTabs  uint
	currentTab uint
	// TODO: add styles
}

func New(totalTabs uint) Model {
	return Model{
		currentTab: 0,
		totalTabs:  totalTabs,
	}
}

func (m Model) Init() tea.Cmd {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	for _, model := range tabModels {
		cmd = model.Init()
		cmds = append(cmds, cmd)
	}
	return tea.Batch(cmds...)
}

// TODO: set size function and handle `tea.WindowSizeMsg` in Update loop
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {

	var cmd tea.Cmd
	var cmds []tea.Cmd

	m.tabModels[currentTab], cmd = m.tabModels[currentTab].Update(msg)
	cmds = append(cmds, cmd)

	return tea.Batch(cmds...)

}

func (m Model) View() string {
	// TODO: style(title) + content
	return ""
}

func (m *Model) TabTitles() []string {
	return m.tabTitles
}

func (m *Model) SetTabTitles(titles []string) {
	if len(titles) != m.totalTabs {
		return
	}
	m.tabTitles = titles
}

func (m *Model) TabModels() []tea.Model {
	return m.tabModels
}

func (m *Model) SetTabModels(models []tea.Model) {
	if len(models) != m.totalTabs {
		return
	}
	m.tabModels = models
}
