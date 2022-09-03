package tabs

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	tabTitles  []string
	tabModels  []tea.Model
	totalTabs  uint
	currentTab uint
	width      int
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

	for _, model := range m.tabModels {
		cmd = model.Init()
		cmds = append(cmds, cmd)
	}
	return tea.Batch(cmds...)
}

// TODO: set size function and handle `tea.WindowSizeMsg` in Update loop
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
        // TODO: how to go about this?
        // should i componsate for the height taken up by the header?
        // or should i define `Height` as a constant and let the user handle it?
		_, v := docStyle.GetFrameSize()
        msg.Height -= v
	}

	m.tabModels[m.currentTab], cmd = m.tabModels[m.currentTab].Update(msg)
	return m, cmd
}

func (m Model) View() string {
	var tabs []string
	for i, tabTitle := range m.tabTitles {
		if m.currentTab == uint(i) {
			tabs = append(tabs, activeTab.Render(tabTitle))
		} else {
			tabs = append(tabs, inactiveTab.Render(tabTitle))
		}
	}

	// TODO: take width into account, ellipsize extra text
	renderedTabs := lipgloss.NewStyle().
		Width(m.width).
		MaxWidth(m.width).
		Render(lipgloss.JoinHorizontal(lipgloss.Top, strings.Join(tabs, "|")))

	return lipgloss.JoinVertical(lipgloss.Top,
		renderedTabs, docStyle.Render(m.tabModels[m.currentTab].View()))
}

func (m *Model) TabTitles() []string {
	return m.tabTitles
}

func (m *Model) SetTabTitles(titles []string) {
	if len(titles) != int(m.totalTabs) {
		return
	}
	m.tabTitles = titles
}

func (m *Model) TabModels() []tea.Model {
	return m.tabModels
}

func (m *Model) SetTabModels(models []tea.Model) {
	if len(models) != int(m.totalTabs) {
		return
	}
	m.tabModels = models
}

func (m *Model) Width() int {
	return m.width
}

func (m *Model) SetWidth(w int) {
	m.width = w
}
