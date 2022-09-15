package main

import "observer/Observer"

func main() {
	hhKz := Observer.NewJobSite()
	bob := Observer.NewPerson("Bob")
	hhKz.AddVacancies("Senior HTML Developer")
	hhKz.Subscribe(bob)
	hhKz.AddVacancies("Java Junior Developer")
	hhKz.RemoveVacancy("Senior HTML Developer")
}
