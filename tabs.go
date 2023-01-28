package tabs

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/truncate"
)

const (
	TabHeight = 1
	ellipsis  = "…"
	splitter  = " | "
)

type Model struct {
	tabTitles  []string
	tabModels  []tea.Model
	totalTabs  int
	currentTab int
	width      int
	height     int
	// TODO: add/move styles
	TitleStyle lipgloss.Style
}

func New(totalTabs int) Model {
	titleStyle := lipgloss.NewStyle().Align(lipgloss.Center)
	return Model{
		currentTab: -1,
		totalTabs:  totalTabs,
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

	// TODO: call Init method on page change for the first time
	// TODO: add keybinds
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.SetSize(msg.Width, msg.Height)
	case tea.KeyMsg:
		if msg.Type == tea.KeyRight || msg.String() == "l" {
			m.currentTab = (m.currentTab + 1) % m.totalTabs
			return m, nil
		} else if msg.Type == tea.KeyLeft || msg.String() == "h" {
			m.currentTab = m.currentTab - 1
			if m.currentTab <= -1 {
				m.currentTab = m.totalTabs - 1
			}
			return m, nil
		}
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
	renderedTabs := truncate.StringWithTail(strings.Join(tabs, splitter), uint(m.width), ellipsis)
	content := lipgloss.JoinVertical(lipgloss.Top,
		m.TitleStyle.Render(renderedTabs),
		m.tabModels[m.currentTab].View())
	return lipgloss.NewStyle().Height(m.height).Render(content)
}

func (m *Model) TabTitles() []string {
	return m.tabTitles
}

func (m *Model) SetTabTitles(titles []string) {
	// TODO: error out, instead of just returning
	// TODO: do not let the user change the number of tabs once it's already created?
	if len(titles) != m.totalTabs {
		return
	}
	m.tabTitles = titles
}

func (m *Model) TabModels() []tea.Model {
	return m.tabModels
}

func (m *Model) SetTabModels(models []tea.Model) {
	// TODO: error out, instead of just returning
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

func (m *Model) SetTitleStyle(titleStyle lipgloss.Style) {
	m.TitleStyle = titleStyle
}

func (m *Model) SetWidth(width int) {
	m.SetSize(width, m.height)
}

func (m *Model) SetHeight(height int) {
	m.SetSize(m.width, height)
}

func (m *Model) SetSize(width, height int) {
	m.width = width
	m.height = height
	m.TitleStyle = m.TitleStyle.
		Width(m.width).
		MaxWidth(m.width)
}

func (m *Model) TotalTabs() int {
	return m.totalTabs
}
