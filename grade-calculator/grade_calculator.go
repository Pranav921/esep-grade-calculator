package esepunittests

type GradeCalculator struct {
	grades         []Grade
	evaluationType EvaluationType
}

type EvaluationType int

const (
	LetterGrade EvaluationType = iota
	PassFail
)

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator(evaluationType ...EvaluationType) *GradeCalculator {
	mode := LetterGrade
	if len(evaluationType) > 0 {
		mode = evaluationType[0]
	}

	return &GradeCalculator{
		grades:         make([]Grade, 0),
		evaluationType: mode,
	}
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()

	if gc.evaluationType == PassFail {
		if numericalGrade >= 70 {
			return "Pass"
		}

		return "Fail"
	}

	if numericalGrade >= 90 {
		return "A"
	} else if numericalGrade >= 80 {
		return "B"
	} else if numericalGrade >= 70 {
		return "C"
	} else if numericalGrade >= 60 {
		return "D"
	}

	return "F"
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	gc.grades = append(gc.grades, Grade{
		Name:  name,
		Grade: grade,
		Type:  gradeType,
	})
}

func (gc *GradeCalculator) gradesOfType(gradeType GradeType) []Grade {
	filtered := make([]Grade, 0)

	for _, grade := range gc.grades {
		if grade.Type == gradeType {
			filtered = append(filtered, grade)
		}
	}

	return filtered
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	assignment_average := computeAverage(gc.gradesOfType(Assignment))
	exam_average := computeAverage(gc.gradesOfType(Exam))
	essay_average := computeAverage(gc.gradesOfType(Essay))

	weighted_grade := float64(assignment_average)*.5 + float64(exam_average)*.35 + float64(essay_average)*.15

	return int(weighted_grade)
}

func computeAverage(grades []Grade) int {
	sum := 0

	for _, grade := range grades {
		sum += grade.Grade
	}

	return sum / len(grades)
}
