package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Person struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Job      string `json:"job"`
	City     string `json:"city"`
	Salary   int    `json:"salary"`
	Birthday string `json:"birthdate"`
}

func (p *Person) String() string {
	return fmt.Sprintf("name: %s, email: %s, job: %s, city: %s, salary: %d, birthday: %s",
		p.Name, p.Email, p.Job, p.City, p.Salary, p.Birthday)
}

func main() {
	// Open our jsonFile
	jsonFile, err := os.Open("data.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened person.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var people []Person
	json.Unmarshal(byteValue, &people)

	// 2.1 Gom tất cả những người trong cùng một thành phố lại
	fmt.Println("_____Bai 2.1_______Group people by city_____________")
	peopleByCity := GroupPeopleByCity(people)
	for key, value := range peopleByCity {
		fmt.Println(key)
		for _, person := range value {
			fmt.Println("  ", (&person).Name)
		}
	}

	// 2.2 Nhóm các nghề nghiệp và đếm số người làm
	fmt.Println("_____Bai 2.2________Group people by job______________")
	peopleByJob := GroupPeopleByJob(people)
	for key, value := range peopleByJob {
		fmt.Println(key, "-", value)
	}

	// 2.3 Tìm 5 nghề có nhiều người làm nhất, đếm từ cao xuống thấp
	fmt.Println("______Bai 2.3______Top 5 job has many people work______________")
	topJobs := Top5JobsByNumber(people)
	for i := 0; i < len(topJobs); i++ {
		fmt.Println(topJobs[i], "-", peopleByJob[topJobs[i]])
	}

	// 2.4 Tìm 5 thành phố có nhiều người trong danh sách ở nhất, đếm từ cao xuống thấp
	fmt.Println("______Bai 2.4______Top 5 city has many people______________")
	Top5CitiesByNumber := Top5CitiesByNumber(people)
	fmt.Println(Top5CitiesByNumber)

	// 2.5 Trong mỗi thành phố, hãy tìm ra nghề nào được làm nhiều nhất
	fmt.Println("_____Bai 2.5_______Top Job By Number In Each City______________")
	TopJobByNumberInEachCity(people)

	// 2.6 Ứng với một nghề, hãy tính mức lương trung bình
	fmt.Println("____Bai 2.6________Average Salary By Job______________")
	listAverageSalaryOfJob := AverageSalaryByJob(people)
	for _, job := range listAverageSalaryOfJob {
		for key, value := range job {
			fmt.Print(key + " - ")
			fmt.Println(value)
		}
	}

	//  2.7 Năm thành phố có mức lương trung bình cao nhất
	fmt.Println("_____Bai 2.7_______Five Cities Has Top Average Salary______________")
	fmt.Println(FiveCitiesHasTopAverageSalary(people))

	//  2.8 Năm thành phố có mức lương trung bình của developer cao nhất
	fmt.Println("_____Bai 2.8_______Five Cities Has Top Salary For Developer______________")
	fmt.Println(FiveCitiesHasTopSalaryForDeveloper(people))

	// 2.9 Tuổi trung bình từng nghề nghiệp
	fmt.Println("_____Bai 2.9_______Average Age Per Job______________")
	objAverageAgePerJob := AverageAgePerJob(people)
	for key, value := range objAverageAgePerJob {
		fmt.Println(key + " - " + value)
	}

	// 2.10 Tuổi trung bình ở từng thành phố
	fmt.Println("____Bai 2.10________Average Age Per City______________")
	objAverageAgePerCity := AverageAgePerCity(people)
	for key, value := range objAverageAgePerCity {
		fmt.Println(key + " - " + value)
	}
}
