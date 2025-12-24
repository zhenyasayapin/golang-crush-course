package staff

type Employee struct {
	FirstName string
	LastName  string
	FullTime  bool
	Salary    int
}

type Office struct {
	AllStaff []Employee
}

var OverpaidThreshold = 75000
var underpaidThreshold = 50000

func (office *Office) All() []Employee {
	return office.AllStaff
}

func (office *Office) Overpaid() []Employee {
	var overpaid []Employee
	for _, emp := range office.AllStaff {
		if emp.Salary > OverpaidThreshold {
			overpaid = append(overpaid, emp)
		}
	}
	return overpaid
}

func (office *Office) Underpaid() []Employee {
	var underpaid []Employee
	for _, emp := range office.AllStaff {
		if emp.Salary < underpaidThreshold {
			underpaid = append(underpaid, emp)
		}
	}
	return underpaid
}
