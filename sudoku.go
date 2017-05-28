/*
 *  A simple sudoku solver written in Go
 *  by Peter Cunha (May 28, 2017)
 */

package main

import (
  "fmt"
)

var found int = 0
var nums      = []int   { 0, 0, 0, 0, 0, 0, 0, 0, 0 }
var board     = [][]int {
        {0,0,0, 5,1,2, 0,6,0},
        {2,0,0, 0,0,0, 9,1,0},
        {1,7,0, 0,0,6, 0,5,8},

        {0,0,0, 0,8,0, 1,0,0},
        {0,1,0, 3,0,4, 0,9,6},
        {0,0,2, 0,6,0, 0,0,0},

        {9,2,0, 4,0,0, 0,8,0},
        {0,0,3, 0,0,0, 0,0,1},
        {0,6,0, 0,7,1, 0,0,0},
}

func main() {
  iterations := 0
  for !isBoardFinished() {

    for i := range board {
      for j := range board[i] {
        board[i][j] = calcCell(i, j)
        resetNums()
      }
    }
    iterations++

    if found == 0 {
      fmt.Println("Board appears to be unsolvable!")
      return
    } else {
      found = 0
    }
  }

  fmt.Println("Solved in", iterations, "iterations.")
  fmt.Println()
  printBoard()
}

// Given the coordinates for a cell in the array, this function tries to return
// the solved cell. Returns 0 if the answer can't be found.
func calcCell(i int, j int) int {
  cell := board[i][j]

  if cell != 0 {
    return cell
  }

  checkRow(i)
  checkCol(j)
  checkBox(i, j)

  unknowns := 0
  lastFound := 0

  for index := 0; index < 9; index++ {
    if nums[index] == 0 {
      unknowns++
      lastFound = index+1
    }
  }

  if unknowns == 1 {
    found++
    return lastFound
  } else {
    return 0
  }
}

// Check row for numbers and add it to nums[] array
func checkRow(i int)  {
  for col := 0; col < 9; col++ {
    if board[i][col] != 0 {
      // Mark the number as found in row
      nums[board[i][col] - 1] = 1
    }
  }
}

// Check col for numbers and add it to nums[] array
func checkCol(j int)  {
  for row := 0; row < 9; row++ {
    if board[row][j] != 0 {
      // Mark the number as found in col
      nums[board[row][j] - 1] = 1
    }
  }
}

// Check the 3x3 box and add it to nums[] array
func checkBox(ih int, jh int)  {
  var boxi int = ih / 3
  var boxj int = jh / 3

  for i := 3*boxi; i < 3*boxi + 3; i++ {
    for j := 3*boxj; j < 3*boxj + 3; j++ {
      if board[i][j] != 0 {
        // Mark the number as found in box
        nums[board[i][j] - 1] = 1
      }
    }
  }
}

// Resets the found numbers array. This is run after every calcCell()
func resetNums()  {
  for index := 0; index < len(nums); index++ {
    nums[index] = 0
  }
}

// Check if board is solved
func isBoardFinished() bool  {
  for i := range board {
    for j := range board[i] {
      if board[i][j] == 0 {
        return false
      }
    }
  }
  return true
}

// Print the board 
func printBoard() {
  for i := range board {
    for j := range board[i] {
      fmt.Print(board[i][j])
      if j != 8 {
        fmt.Print(", ")
      }
    }
    fmt.Println()
  }
}
