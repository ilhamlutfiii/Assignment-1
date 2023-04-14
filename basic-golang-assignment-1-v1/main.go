package main

import (
	"fmt"
	"strings"

	"a21hc3NpZ25tZW50/helper"
)

var Students string = "A1234_Aditira_TI, B2131_Dito_TK, A3455_Afis_MI"
var StudentStudyPrograms string = "TI_Teknik Informatika, TK_Teknik Komputer, SI_Sistem Informasi, MI_Manajemen Informasi"

func Login(id string, name string) string {
	var log string
	if id == "" || name == "" {
		log = "ID or Name is undefined!"
	} else if len(id) != 5 {
		log = "ID must be 5 characters long!"
	} else if strings.Contains(Students, id+"_") {
		studentsData := strings.Split(Students, ", ")
		found := false
		for _, student := range studentsData {
			if strings.Contains(student, id+"_"+name+"_") {
				studentInfo := strings.Split(student, "_")
				log = fmt.Sprintf("Login berhasil: %s (%s)", studentInfo[1], studentInfo[2])
				found = true
				break
			}
		}
		if !found {
			log = "Login gagal: data mahasiswa tidak ditemukan"
		}
	} else {
		log = "Login gagal: data mahasiswa tidak ditemukan"
	}
	return log
}

func isIdRegistered(id string) bool {
	studentsData := strings.Split(Students, ", ")
	for _, student := range studentsData {
		if strings.Contains(student, id+"_") {
			return true
		}
	}
	return false
}

func Register(id string, name string, major string) string {
	var reg string
	if id == "" || name == "" || major == "" {
		reg = "ID, Name or Major is undefined!"
	} else if len(id) != 5 {
		reg = "ID must be 5 characters long!"
	} else if isIdRegistered(id) {
		reg = "Registrasi gagal: id sudah digunakan"
	} else {
		Students += ", " + id + "_" + name + "_" + major
		reg = fmt.Sprintf("Registrasi berhasil: %s (%s)", name, major)
	}
	return reg
}

func GetStudyProgram(code string) string {
	var get string
	if code == "" {
		get = "Code is undefined!"
	} else if !strings.Contains(StudentStudyPrograms, code+"_") {
		get = "Program studi tidak ditemukan"
	} else {
		studyProgramsData := strings.Split(StudentStudyPrograms, ", ")
		for _, studyProgram := range studyProgramsData {
			if strings.Contains(studyProgram, code+"_") {
				studyProgramInfo := strings.Split(studyProgram, "_")
				get = studyProgramInfo[1]
				break
			}
		}
	}

	return get
}

func main() {
	fmt.Println("Selamat datang di Student Portal!")

	for {
		helper.ClearScreen()
		fmt.Println("Students: ", Students)
		fmt.Println("Student Study Programs: ", StudentStudyPrograms)

		fmt.Println("\nPilih menu:")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Get Program Study")
		fmt.Println("4. Keluar")

		var pilihan string
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			helper.ClearScreen()
			var id, name string
			fmt.Print("Masukkan id: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan name: ")
			fmt.Scan(&name)

			fmt.Println(Login(id, name))

			helper.Delay(5)
		case "2":
			helper.ClearScreen()
			var id, name, jurusan string
			fmt.Print("Masukkan id: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan name: ")
			fmt.Scan(&name)
			fmt.Print("Masukkan jurusan: ")
			fmt.Scan(&jurusan)
			fmt.Println(Register(id, name, jurusan))

			helper.Delay(5)
		case "3":
			helper.ClearScreen()
			var kode string
			fmt.Print("Masukkan kode: ")
			fmt.Scan(&kode)

			fmt.Println(GetStudyProgram(kode))
			helper.Delay(5)
		case "4":
			fmt.Println("Terima kasih telah menggunakan Student Portal.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
