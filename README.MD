# ESEP Grade Calculator

This library is a simple utility to calculate your grades.

You add grades for assignments, exams, and essays. The calculator averages each
category, applies the category weights, and gives back a final letter grade.

## Requirements

Go 1.24.4 or newer (this comes from `go.mod`)

## Project Files

All of the Go code lives in the `grade-calculator/` folder.

| File | What it contains |
| --- | --- |
| `go.mod` | The Go module name (`esep/grade-calculator`) and the Go version |
| `grade_calculator.go` | The `GradeCalculator` type, the `Grade` and `GradeType` types, and the grading logic |
| `grade_calculator_test.go` | The unit tests for the grade calculator |

## Grade Weights

| Category | Weight |
| --- | --- |
| Assignments | 50% |
| Exams | 35% |
| Essays | 15% |

## Letter Grade Scale

| Numerical grade | Letter grade |
| --- | --- |
| 90 and above | A |
| 80 to 89 | B |
| 70 to 79 | C |
| 60 to 69 | D |
| Below 60 | F |

A C or better is passing.

## Example

```go
calculator := NewGradeCalculator()

calculator.AddGrade("open source assignment", 100, Assignment)
calculator.AddGrade("exam 1", 95, Exam)
calculator.AddGrade("essay on ai ethics", 91, Essay)

finalGrade := calculator.GetFinalGrade() // "A"
```

## Running the Tests

All commands below should be run from inside the `grade-calculator/` folder.

```
cd grade-calculator/
```

Run the unit tests

```
go test ./...
```

Run the unit tests with verbose output

```
go test -v ./...
```

## Code Coverage

Generate a coverage profile

```
go test -coverprofile=coverage.txt ./...
```

View coverage for each function

```
go tool cover -func=coverage.txt
```

Open the HTML coverage report in your browser

```
go tool cover -html=coverage.txt
```

The `coverage.txt` file is generated, so you do not need to commit it.
