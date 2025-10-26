 🧩 Quad Project (Go)

 📘 Overview
This project implements five different **rectangle-drawing functions** in Go — `QuadA`, `QuadB`, `QuadC`, `QuadD`, and `QuadE`.  
Each function prints a unique pattern of rectangles based on the given width `x` and height `y`.  

It’s part of a Go learning challenge that tests our  understanding of loops, conditionals, and modular programming.



 🧠 Project Structure


quad/
├── go.mod
├── main.go
└── piscine/
├── quadA.go
├── quadB.go
├── quadC.go
├── quadD.go
├── quadE.go
├── quad_test.go 


 ⚙️ Setup Instructions

 Step 1: Clone the Repository
```bash
git clone https://acad.learn2earn.ng/git/tifeanae/quad
cd quad

Step 2: Initialize Go Module

go mod init quad
go mod tidy

step 3: Verify Folder Structure

Ensure your files are organized like this:

🚀 Running the Program
Example main.go

package main

import "quad/piscine"

func main() {
    piscine.QuadA(5, 3)
    println()
    piscine.QuadB(5, 3)
    println()
    piscine.QuadC(5, 3)
    println()
    piscine.QuadD(5, 3)
    println()
    piscine.QuadE(5, 3)
}

Run this using:
go run .


🧰 Utility Function — drawLine()

This shared function handles printing each line of the rectangle.

package piscine

import "fmt"

func drawLine(left, middle, right byte, width int) {
    if width <= 0 {
        return
    }
    if width == 1 {
        fmt.Printf("%c\n", left)
        return
    }
    fmt.Printf("%c", left)
    for i := 0; i < width-2; i++ {
        fmt.Printf("%c", middle)
    }
    fmt.Printf("%c\n", right)
}

✅ Example:

drawLine('o', '-', 'o', 5)
Output:

o---o

🧪 Running Tests

If you renamed your utils.go file to quad_test.go, you can run:

go test ./piscine -v

👥 Team Workflow

Each teammate should:

1.Work on one Quad function (QuadA to QuadE).

2.Use drawLine() — no repeated logic.

3.Test using provided examples.

4.Commit and push regularly.

Git Workflow Example

git add .
git commit -m "Implement QuadA"
git push origin main


🧾 Example Outputs


QuadA(5,3)
o---o
|   |
o---o

QuadB(5,3)
/***\
*   *
\***/
QuadC(5,3)

ABBBA
B   B
CBBBC

QuadD(5,3)

ABBBC
B   B
ABBBC

QuadE(5,3)

ABBBC
B   B
CBBBA
