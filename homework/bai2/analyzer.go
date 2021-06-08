package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

func GroupPeopleByCity(p []Person) (result map[string][]Person) {
	result = make(map[string][]Person)
	for _, person := range p {
		result[person.City] = append(result[person.City], person)
	}
	return result
}

func GroupPeopleByJob(p []Person) map[string]int {
	result := make(map[string]int)
	for _, person := range p {
		result[person.Job]++
	}
	return result
}

func Top5JobsByNumber(p []Person) []string {
	peopleByJob := GroupPeopleByJob((p))
	var jobNames []string
	for job, _ := range peopleByJob {
		jobNames = append(jobNames, job)
	}

	sort.Slice(jobNames, func(i, j int) bool {
		return peopleByJob[jobNames[i]] > peopleByJob[jobNames[j]]
	})
	if len(jobNames) < 5 {
		return jobNames[0:]
	} else {
		return jobNames[0:5]
	}
}

func Top5CitiesByNumber(p []Person) []string {
	peopleByCity := GroupPeopleByCity(p)
	peopleByCityNumber := make(map[string]int)
	for city, _ := range peopleByCity {
		peopleByCityNumber[city] = len(peopleByCity[city])
	}
	var cityNames []string
	for city, _ := range peopleByCityNumber {
		cityNames = append(cityNames, city)
	}
	sort.Slice(cityNames, func(i, j int) bool {
		return peopleByCityNumber[cityNames[i]] > peopleByCityNumber[cityNames[j]]
	})
	if len(cityNames) < 5 {
		return cityNames[0:]
	} else {
		return cityNames[0:5]
	}
}

func TopJobByNumberInEachCity(p []Person) {
	peopleByCity := GroupPeopleByCity(p)
	countJobInCity := make(map[string]map[string]int)
	for city, _ := range peopleByCity {
		countJob := GroupPeopleByJob(peopleByCity[city])
		countJobInCity[city] = countJob
	}
	for city, objJob := range countJobInCity {
		var jobNames []string
		for key, _ := range objJob {
			jobNames = append(jobNames, key)
		}
		sort.Slice(jobNames, func(i, j int) bool {
			return objJob[jobNames[i]] > objJob[jobNames[j]]
		})
		fmt.Println(city + " - " + jobNames[0])
	}
}

func AverageSalaryByJob(p []Person) []map[string]int {
	listPeopleByJob := make(map[string][]Person)
	listAverageSalaryOfJob := make([]map[string]int, 0)
	for _, person := range p {
		listPeopleByJob[person.Job] = append(listPeopleByJob[person.Job], person)
	}
	for jobName, jobValue := range listPeopleByJob {
		objJob := make(map[string]int)
		sum := 0
		for _, people := range jobValue {
			sum += people.Salary
		}
		objJob[jobName] = sum / len(jobValue)
		listAverageSalaryOfJob = append(listAverageSalaryOfJob, objJob)
	}
	return listAverageSalaryOfJob
}

func FiveCitiesHasTopAverageSalary(p []Person) []string {
	listPeopleByCity := GroupPeopleByCity(p)
	listAverageSalaryOfCity := make(map[string]int)
	for city, listPerson := range listPeopleByCity {
		sum := 0
		for _, person := range listPerson {
			sum += person.Salary
		}
		listAverageSalaryOfCity[city] = sum / len(listPerson)
	}
	listCityName := make([]string, 0)
	for cityName, _ := range listAverageSalaryOfCity {
		listCityName = append(listCityName, cityName)
	}
	sort.Slice(listCityName, func(i, j int) bool {
		return listAverageSalaryOfCity[listCityName[i]] > listAverageSalaryOfCity[listCityName[j]]
	})
	if len(listCityName) < 5 {
		return listCityName[0:]
	} else {
		return listCityName[0:5]
	}
}

func FiveCitiesHasTopSalaryForDeveloper(p []Person) []string {
	listPeopleByCity := GroupPeopleByCity(p)
	listAverageSalaryOfDeveloperGroupCity := make(map[string]int)
	for city, listPerson := range listPeopleByCity {
		sum := 0
		count := 0
		for _, person := range listPerson {
			if person.Job == "developer" {
				sum += person.Salary
				count++
			}
		}
		if count != 0 {
			listAverageSalaryOfDeveloperGroupCity[city] = sum / count
		}
	}
	listCityName := make([]string, 0)
	for cityName, _ := range listAverageSalaryOfDeveloperGroupCity {
		listCityName = append(listCityName, cityName)
	}
	sort.Slice(listCityName, func(i, j int) bool {
		return listAverageSalaryOfDeveloperGroupCity[listCityName[i]] > listAverageSalaryOfDeveloperGroupCity[listCityName[j]]
	})
	if len(listCityName) < 5 {
		return listCityName[0:]
	} else {
		return listCityName[0:5]
	}
}

func AverageAgePerJob(p []Person) map[string]string {
	groupPeopleByJob := make(map[string][]Person)
	listAverageAgeOfJob := make(map[string]string)
	for _, person := range p {
		groupPeopleByJob[person.Job] = append(groupPeopleByJob[person.Job], person)
	}
	dt := time.Now().String()[0:4]
	yearCurent, err := strconv.ParseInt(dt, 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	for job, listPeople := range groupPeopleByJob {
		var sumAge float64 = 0
		for _, people := range listPeople {
			yearPeople, err := strconv.ParseInt(strings.Split(people.Birthday, "-")[0], 10, 64)
			if err != nil {
				fmt.Println(err)
			}
			sumAge += float64(yearCurent - yearPeople)
		}
		listAverageAgeOfJob[job] = fmt.Sprintf("%.1f", sumAge/float64(len(listPeople)))
	}
	return listAverageAgeOfJob
}

func AverageAgePerCity(p []Person) map[string]string {
	listPersonByCity := GroupPeopleByCity(p)
	listAverageAgeOfCity := make(map[string]string)

	dt := time.Now().String()[0:4]
	yearCurent, err := strconv.ParseInt(dt, 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	for city, listPeople := range listPersonByCity {
		var sumAge float64 = 0
		for _, people := range listPeople {
			yearPeople, err := strconv.ParseInt(strings.Split(people.Birthday, "-")[0], 10, 64)
			if err != nil {
				fmt.Println(err)
			}
			sumAge += float64(yearCurent - yearPeople)
		}
		listAverageAgeOfCity[city] = fmt.Sprintf("%.1f", sumAge/float64(len(listPeople)))
	}
	return listAverageAgeOfCity
}
