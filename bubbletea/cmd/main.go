package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

type model struct {
	cursor   [2]int              // Cursor represents a position (row, col)
	board    [][]string          // 2D board of strings
	selected map[[2]int]struct{} // Selected positions on the board
}

func initialModel() model {
	return model{
		board: [][]string{
			{"♜", "♞", "♝", "♛", "♚", "♝", "♞", "♜", "♜", "♞", "♝", "♛", "♚", "♝", "♞", "♜"},
			{"♟", "♟", "♟", "♟", "♟", "♟", "♟", "♟", "♟", "♟", "♟", "♟", "♟", "♟", "♟", "♟"},
			{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0"},
			{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0"},
			{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0"},
			{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0"},
			{"♙", "♙", "♙", "♙", "♙", "♙", "♙", "♙", "♙", "♙", "♙", "♙", "♙", "♙", "♙", "♙"},
			{"♖", "♘", "♗", "♕", "♔", "♗", "♘", "♖", "♖", "♘", "♗", "♕", "♔", "♗", "♘", "♖"},
			{"♜", "♞", "♝", "♛", "♚", "♝", "♞", "♜", "♜", "♞", "♝", "♛", "♚", "♝", "♞", "♜"},
			{"♟", "♟", "♟", "♟", "♟", "♟", "♟", "♟", "♟", "♟", "♟", "♟", "♟", "♟", "♟", "♟"},
			{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0"},
			{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0"},
			{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0"},
			{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0"},
			{"♙", "♙", "♙", "♙", "♙", "♙", "♙", "♙", "♙", "♙", "♙", "♙", "♙", "♙", "♙", "♙"},
			{"♖", "♘", "♗", "♕", "♔", "♗", "♘", "♖", "♖", "♘", "♗", "♕", "♔", "♗", "♘", "♖"},
		},
		selected: make(map[[2]int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return tea.SetWindowTitle("Chess Board")
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor[0] > 0 {
				m.cursor[0]--
			}
		case "down", "j":
			if m.cursor[0] < len(m.board)-1 {
				m.cursor[0]++
			}
		case "left", "h":
			if m.cursor[1] > 0 {
				m.cursor[1]--
			}
		case "right", "l":
			if m.cursor[1] < len(m.board[0])-1 {
				m.cursor[1]++
			}
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	// Estilos
	labelStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	cursorStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("226"))

	// Construção do tabuleiro com a peça destacada pelo cursor
	var boardRows [][]string
	for i, row := range m.board {
		var newRow []string
		for j, piece := range row {
			if m.cursor[0] == i && m.cursor[1] == j {
				piece = cursorStyle.Render(piece)
			}
			newRow = append(newRow, piece)
		}
		boardRows = append(boardRows, newRow)
	}

	// Construção da tabela com bordas
	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderRow(true).
		BorderColumn(true).
		Rows(boardRows...).
		StyleFunc(func(row, col int) lipgloss.Style {
			return lipgloss.NewStyle().Padding(0, 1)
		})

	// Ranks e Files
	ranks := labelStyle.Render(strings.Join([]string{" A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P  "}, "   "))
	files := labelStyle.Render(strings.Join([]string{" 1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16"}, "\n\n "))

	// Combinação de ranks, files e a tabela em uma única string
	output := lipgloss.JoinVertical(lipgloss.Right, lipgloss.JoinHorizontal(lipgloss.Center, files, t.Render()), ranks)

	return output + "\nPress q to quit.\n"
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
