package tabs

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/truncate"
)

const (
	Height   = 1
	ellipsis = "…"
	splitter = " | "
)

type Model struct {
	tabTitles  []string
	tabModels  []tea.Model
	totalTabs  int
	currentTab int
	Width      int
	Height     int
	// TODO: add/move styles
	TitleStyle lipgloss.Style
}

func New(totalTabs int) Model {
	var titleStyle = lipgloss.NewStyle().Align(lipgloss.Center)
	return Model{
		currentTab: -1,
		totalTabs:  totalTabs,
		Height:     Height,
		TitleStyle: titleStyle,
	}
}

func (m Model) Init() tea.Cmd {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	for _, model := range m.tabModels {
		cmd = model.Init()
		cmds = append(cmds, cmd)
	}
	return tea.Batch(cmds...)
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd

	// TODO: add keybinds
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		// just update width
		m.TitleStyle = m.TitleStyle.
			Width(msg.Width).
			MaxWidth(msg.Width)
	}

	// TODO: should i print a warning or something like that here?
	if m.currentTab >= 0 {
		m.tabModels[m.currentTab], cmd = m.tabModels[m.currentTab].Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m Model) View() string {
	if m.currentTab < 0 {
		return ""
	}

	var tabs []string
	for i, tabTitle := range m.tabTitles {
		if m.currentTab == i {
			tabs = append(tabs, activeTab.Render(tabTitle))
		} else {
			tabs = append(tabs, inactiveTab.Render(tabTitle))
		}
	}
	renderedTabs := truncate.StringWithTail(strings.Join(tabs, splitter), uint(m.Width), ellipsis)
	return lipgloss.JoinVertical(lipgloss.Top,
		m.TitleStyle.Render(renderedTabs),
		m.tabModels[m.currentTab].View())
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

func (m *Model) CurrentTab() int {
	return m.currentTab
}

func (m *Model) SetCurrentTab(tab int) {
	m.currentTab = tab
}

func (m *Model) SetWidth(width int) {
	m.Width = width
}

func (m *Model) SetTitleStyle(titleStyle lipgloss.Style) {
	m.TitleStyle = titleStyle
}
