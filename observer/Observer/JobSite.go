package Observer

type JobSite struct {
	subscribers []Observer
	vacancies   []string
}

func NewJobSite() *JobSite {
	j := new(JobSite)
	j.subscribers = make([]Observer, 0, 5)
	j.vacancies = make([]string, 0, 5)
	return j
}

func (j *JobSite) AddVacancies(vacancy string) {
	j.vacancies = append(j.vacancies, vacancy)
	j.SendAll()
}

func (j *JobSite) RemoveVacancy(vacancy string) {
	for i := range j.vacancies {
		if j.vacancies[i] == vacancy {
			j.vacancies = remove(j.vacancies, i)
			break
		}
	}
	j.SendAll()
}

func (j *JobSite) Subscribe(observer Observer) {
	j.subscribers = append(j.subscribers, observer)
}

func (j *JobSite) Unsubscribe(observer Observer) {
	for i := range j.subscribers {
		if j.subscribers[i] == observer {
			j.subscribers = remove(j.subscribers, i)
			break
		}
	}
}

func (j *JobSite) SendAll() {
	for _, observer := range j.subscribers {
		observer.handleEvent(j.vacancies)
	}
}

func remove[T any](slice []T, i int) []T {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
