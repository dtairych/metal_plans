package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type EquinixMetalResponse struct {
	Plans []Plan `json:"plans"`
}

type Plan struct {
	Slug            string   `json:"slug"`
	Legacy          bool     `json:"legacy"`
	DeploymentTypes []string `json:"deployment_types"`
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	Specs           Specs    `json:"specs"`
	Pricing         Pricing  `json:"pricing"`
}

type Pricing struct {
	Hour float32 `json:"hour"`
}

type Specs struct {
	CPUs   []CPU   `json:"cpus"`
	Memory Memory  `json:"memory"`
	Drives []Drive `json:"drives"`
	NICs   []NIC   `json:"nics"`
}

type CPU struct {
	Count        int         `json:"count"`
	Type         string      `json:"type"`
	Manufacturer string      `json:"manufacturer,omitempty"`
	Model        string      `json:"model,omitempty"`
	Cores        json.Number `json:"cores,omitempty"`
	Speed        string      `json:"speed,omitempty"`
}

type Memory struct {
	Total string `json:"total"`
}

type Drive struct {
	Count    int    `json:"count"`
	Size     string `json:"size"`
	Type     string `json:"type"`
	Category string `json:"category,omitempty"`
}

type NIC struct {
	Count int    `json:"count"`
	Type  string `json:"type"`
}

func main() {
	apiToken := os.Getenv("EQUINIX_API_TOKEN")
	if apiToken == "" {
		log.Fatal("EQUINIX_API_TOKEN environment variable is not set")
	}

	url := "https://api.equinix.com/metal/v1/plans"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("X-Auth-Token", apiToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(body)
	var response EquinixMetalResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal(err)
	}

	for _, plan := range response.Plans {
		fmt.Println("Server:", plan.Name)
		fmt.Println("Description:", plan.Description)
		fmt.Println("Legacy: ", plan.Legacy)
		fmt.Println("Deployment Options: ", plan.DeploymentTypes)
		fmt.Println("CPUs:")
		for _, cpu := range plan.Specs.CPUs {
			fmt.Printf("  %d x %s\n", cpu.Count, cpu.Type)
			fmt.Printf("    Manufacturer: %s\n", cpu.Manufacturer)
			fmt.Printf("    Model: %s\n", cpu.Model)
			fmt.Printf("    Cores: %d\n", cpu.Cores)
			fmt.Printf("    Speed: %s\n", cpu.Speed)
		}
		fmt.Println("Memory:")
		fmt.Printf("  Total: %s\n", plan.Specs.Memory.Total)
		fmt.Println("Drives:")
		for _, drive := range plan.Specs.Drives {
			fmt.Printf("  %d x %s %s %s\n", drive.Count, drive.Size, drive.Type, drive.Category)
		}
		fmt.Println("NICs:")
		for _, nic := range plan.Specs.NICs {
			fmt.Printf("  %d x %s\n", nic.Count, nic.Type)
		}
		fmt.Println("------------------------")
	}
}
