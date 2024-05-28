package main

import (
	"context"
	"fmt"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type model struct {
	table table.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return baseStyle.Render(m.table.View())
}

type containerStat struct {
	id      string
	image   string
	command string
	ports   []types.Port
	state   string
	status  string
}

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	containers, err := cli.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		panic(err)
	}

	containerStats := []containerStat{}

	for _, ctr := range containers {
		containerStats = append(containerStats, containerStat{
			id:      ctr.ID,
			image:   ctr.Image,
			command: ctr.Command,
			ports:   ctr.Ports,
			state:   ctr.State,
			status:  ctr.Status,
		})
	}

	columns := []table.Column{
		{Title: "ID", Width: 10},
		{Title: "Image", Width: 20},
		{Title: "Status", Width: 10},
		{Title: "State", Width: 10},
	}

	rows := []table.Row{}
	for _, ctrStats := range containerStats {
		rows = append(rows, table.Row{
			trim(ctrStats.id, 10),
			trim(ctrStats.image, 20),
			ctrStats.status,
			ctrStats.state,
		})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(min(len(rows), 10)),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	m := model{t}
	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program: ", err)
		os.Exit(1)
	}
}

func trim(content string, maxChars int) string {
	if len(content) < maxChars {
		return content
	}
	return content[0:maxChars]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
