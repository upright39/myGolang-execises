//Package census simulates a syterm used to collect census data.
package execises16

//Resident representss a resident in this city
type Resident struct {
	name    string
	age     int
	address map[string]string
}

// NewResident registers a new resident in this city.

func NewResident(name string, age int, address map[string]string) *Resident {

	newR := &Resident{
		name:    name,
		age:     age,
		address: address,
	}
	return newR
}

//HasRequieredInfo determins if a given resident has all of the required informations.
func (r *Resident) HasRequiredInfo() bool {

	if r.name != "" && r.address["street"] != "" {
		return true
	} else {
		return false
	}
}

// Deletes deletes a resident's infomation.
func (r *Resident) Delete() {

	if r.name != "" && r.age != 0 && r.address != nil {
		r.name = ""
		r.age = 0
		r.address = nil

		// or
		//*r = Resident{}

	}
}

// Count counts all residents that have provided the required information.

func Count(residents []*Resident) int {

	Count := 0

	for _, v := range residents {

		if v.HasRequiredInfo() {
			Count++

		}

	}
	return Count
}
