package main

import (
	"encoding/json"
	"fmt"
)

const data = `{
    "name": "FUP-0244",
    "owner": "abdulrahman.muhammad@ebda.earth",
    "creation": "2026-01-13 10:21:24.125258",
    "modified": "2026-01-14 18:23:21.299814",
    "modified_by": "abdulrahman.muhammad@ebda.earth",
    "docstatus": 0,
    "idx": 0,
    "farm": "Nasreddin Mohammed Karkoub",
    "farm_name": "Nasreddin Mohammed Karkoub",
    "farm_code": "EG/4180",
    "region": "Luxor",
    "area_feddan": 5.0,
    "farm_owner": "\u0646\u0635\u0631 \u0627\u0644\u062f\u064a\u0646 \u0645\u062d\u0645\u062f \u0643\u0631\u0643\u0648\u0628",
    "phone": "100361410",
    "latitude": "25.3470457",
    "longitude": "32.4648413",
    "upscaling_project": "Phase 3",
    "subupscalingproject": "Phase 3 3K",
    "farm_group": "First Group",
    "gps": "25.35135135135135,32.479117411511",
    "visit_date": "2026-01-13",
    "follower_type": "\u0645\u0647\u0646\u062f\u0633 \u0632\u0631\u0627\u0639\u0649",
    "follower_name": "\u0639\u0628\u062f\u0627\u0644\u0631\u062d\u0645\u0646 \u0645\u062d\u0645\u062f ",
    "picture_of_follower": "/private/files/IMG_\u0662\u0660\u0662\u0666\u0660\u0661\u0661\u0663_\u0660\u0669\u0661\u0662\u0662\u0661.jpg",
    "farmers_count": 1,
    "last_ebda_visit": "\u0627\u0644\u0634\u0647\u0631 \u0627\u0644\u0645\u0627\u0636\u064a",
    "has_health_insurance_card": 1,
    "storage_exist": "\u0646\u0639\u0645",
    "warehouses_notes": "\u062d\u0628\u0648\u0628 \u0648\u063a\u0644\u0627\u0644",
    "records_farm_book": 1,
    "record_image": "/private/files/17682925488332940548414778308900.jpg",
    "intercropping_or_green_manure": "\u0644\u0627",
    "intercropping_percent": 0.0,
    "planted_trees_or_hedge": "\u0646\u0639\u0645",
    "trees_count": 20,
    "has_animals": "\u0646\u0639\u0645",
    "animals_type_count": "2\u0628\u0642\u0631\u0629",
    "fertilization_source": "\u0627\u0644\u0623\u0633\u0645\u062f\u0629 \u0627\u0644\u062d\u064a\u0648\u064a\u0629 / \u0627\u0644\u0645\u0646\u062a\u062c\u0627\u062a \u0627\u0644\u0628\u064a\u0648\u0644\u0648\u062c\u064a\u0629 \u0648 \u0627\u0644\u0633\u0645\u0627\u062f \u0627\u0644\u0639\u0636\u0648\u064a",
    "compost_source": "\u0645\u0646 \u062f\u0627\u062e\u0644 \u0627\u0644\u0645\u0632\u0631\u0639\u0629",
    "compost_production": 50.0,
    "compost_qtys": 50.0,
    "uses_bio_products": "\u0646\u0639\u0645",
    "comp_qty": 50000.0,
    "qurts": 1.0,
    "horns": 20.0,
    "\u0627\u0644\u0643\u0645\u064a\u0627\u062a_\u0627\u0644\u0645\u0633\u062a\u062e\u062f\u0645\u0629": "20\u062c\u0631\u0627\u0645 \u0642\u0631\u0648\u0646",
    "pests_or_diseases_last_season": "\u0644\u0627",
    "pest_control_method": "\u0645\u0643\u0627\u0641\u062d\u0629 \u062d\u064a\u0648\u064a\u0629",
    "bio_control_details": "\u0635\u0627\u0628\u0648\u0646 \u0628\u0648\u062a\u0627\u0633\u064a",
    "uses_natural_enemies": 1,
    "natural_enemies_details": "\u0627\u0628\u0648 \u0627\u0644\u0639\u064a\u062f",
    "weed_disposal": "\u0646\u0632\u0639 \u064a\u062f\u0648\u064a",
    "irrigation_method": "\u063a\u0645\u0631",
    "water_source": "\u0646\u0647\u0631 \u0627\u0644\u0646\u064a\u0644",
    "water_shortage_or_salinity": 1,
    "energy_type": "\u062f\u064a\u0632\u0644",
    "climate_challenges_this_season": 1,
    "yield_vs_last_season": "\u0632\u064a\u0627\u062f\u0629",
    "need_cert_support": 1,
    "current_challenges": "\u0644\u0627 \u062a\u0648\u062c\u062f",
    "support_needed": "\u0645\u0633\u062a\u062d\u0636\u0631\u0627\u062a \u0648\u0623\u0633\u0645\u062f\u0629 \u0648\u0645\u0648\u0627\u062f \u0645\u0643\u0627\u0641\u062d\u0629 \u062d\u064a\u0648\u064a\u0629",
    "follower_assessment": "\u062c\u064a\u062f\u0629",
    "follower_recommendations": "\u0627\u0644\u0627\u0633\u062a\u0645\u0631\u0627\u0631 \u0641\u064a \u0627\u0644\u0632\u0631\u0627\u0639\u0629 \u0627\u0644\u062d\u064a\u0648\u064a\u0629 \u0644\u0632\u064a\u0627\u062f\u0629 \u0627\u0644\u0627\u0646\u062a\u0627\u062c",
    "doctype": "Farm FollowUp",
    "services_used": [
      {
        "name": "64dau3l8it",
        "owner": "abdulrahman.muhammad@ebda.earth",
        "creation": "2026-01-13 10:21:24.125258",
        "modified": "2026-01-14 18:23:21.299814",
        "modified_by": "abdulrahman.muhammad@ebda.earth",
        "docstatus": 0,
        "idx": 1,
        "service": "\u0623\u062f\u0648\u064a\u0629",
        "parent": "FUP-0244",
        "parentfield": "services_used",
        "parenttype": "Farm FollowUp",
        "doctype": "Healthe Services Table"
      }
    ],
    "farmers_names": [
      {
        "name": "199l6bfmp2",
        "owner": "abdulrahman.muhammad@ebda.earth",
        "creation": "2026-01-13 10:21:24.125258",
        "modified": "2026-01-14 18:23:21.299814",
        "modified_by": "abdulrahman.muhammad@ebda.earth",
        "docstatus": 0,
        "idx": 1,
        "farmer": "\u0646\u0635\u0631 \u0627\u0644\u062f\u064a\u0646 \u0645\u062d\u0645\u062f \u0643\u0631\u0643\u0648\u0628",
        "status": "",
        "parent": "FUP-0244",
        "parentfield": "farmers_names",
        "parenttype": "Farm FollowUp",
        "doctype": "Farmer FollowUp"
      }
    ],
    "bios_products_details": [
      {
        "name": "64dpelhqle",
        "owner": "abdulrahman.muhammad@ebda.earth",
        "creation": "2026-01-13 10:21:24.125258",
        "modified": "2026-01-14 18:23:21.299814",
        "modified_by": "abdulrahman.muhammad@ebda.earth",
        "docstatus": 0,
        "idx": 1,
        "fertilizers": "Horn",
        "parent": "FUP-0244",
        "parentfield": "bios_products_details",
        "parenttype": "Farm FollowUp",
        "doctype": "Fertilizers Followup"
      },
      {
        "name": "64d405qlha",
        "owner": "abdulrahman.muhammad@ebda.earth",
        "creation": "2026-01-13 10:21:24.125258",
        "modified": "2026-01-14 18:23:21.299814",
        "modified_by": "abdulrahman.muhammad@ebda.earth",
        "docstatus": 0,
        "idx": 2,
        "fertilizers": "Compost",
        "parent": "FUP-0244",
        "parentfield": "bios_products_details",
        "parenttype": "Farm FollowUp",
        "doctype": "Fertilizers Followup"
      },
      {
        "name": "64da9tq78q",
        "owner": "abdulrahman.muhammad@ebda.earth",
        "creation": "2026-01-13 10:21:24.125258",
        "modified": "2026-01-14 18:23:21.299814",
        "modified_by": "abdulrahman.muhammad@ebda.earth",
        "docstatus": 0,
        "idx": 3,
        "fertilizers": "Biofert",
        "parent": "FUP-0244",
        "parentfield": "bios_products_details",
        "parenttype": "Farm FollowUp",
        "doctype": "Fertilizers Followup"
      },
      {
        "name": "64dmrg00bg",
        "owner": "abdulrahman.muhammad@ebda.earth",
        "creation": "2026-01-13 10:21:24.125258",
        "modified": "2026-01-14 18:23:21.299814",
        "modified_by": "abdulrahman.muhammad@ebda.earth",
        "docstatus": 0,
        "idx": 4,
        "fertilizers": "Quartz",
        "parent": "FUP-0244",
        "parentfield": "bios_products_details",
        "parenttype": "Farm FollowUp",
        "doctype": "Fertilizers Followup"
      }
    ],
    "curent_crops": [
      {
        "name": "64d6mcbt2u",
        "owner": "abdulrahman.muhammad@ebda.earth",
        "creation": "2026-01-13 10:21:24.125258",
        "modified": "2026-01-14 18:23:21.299814",
        "modified_by": "abdulrahman.muhammad@ebda.earth",
        "docstatus": 0,
        "idx": 1,
        "crops": "Wheat",
        "parent": "FUP-0244",
        "parentfield": "curent_crops",
        "parenttype": "Farm FollowUp",
        "doctype": "Crops Followup"
      },
      {
        "name": "7dt354al11",
        "owner": "abdulrahman.muhammad@ebda.earth",
        "creation": "2026-01-13 10:21:24.125258",
        "modified": "2026-01-14 18:23:21.299814",
        "modified_by": "abdulrahman.muhammad@ebda.earth",
        "docstatus": 0,
        "idx": 2,
        "crops": "Clover",
        "parent": "FUP-0244",
        "parentfield": "curent_crops",
        "parenttype": "Farm FollowUp",
        "doctype": "Crops Followup"
      }
    ]
  }`

type BioProductFollowUp struct {
	Name       string `json:"name"`
	Owner      string `json:"owner"`
	Creation   string `json:"creation"`
	Modified   string `json:"modified"`
	ModifiedBy string `json:"modified_by"`
	DocStatus  int    `json:"docstatus"`
	Idx        int    `json:"idx"`

	Fertilizers string `json:"fertilizers"`

	Parent      string `json:"parent"`
	ParentField string `json:"parentfield"`
	ParentType  string `json:"parenttype"`
	Doctype     string `json:"doctype"`
}

type CropFollowUp struct {
	Name       string `json:"name"`
	Owner      string `json:"owner"`
	Creation   string `json:"creation"`
	Modified   string `json:"modified"`
	ModifiedBy string `json:"modified_by"`
	DocStatus  int    `json:"docstatus"`
	Idx        int    `json:"idx"`

	Crops string `json:"crops"`

	Parent      string `json:"parent"`
	ParentField string `json:"parentfield"`
	ParentType  string `json:"parenttype"`
	Doctype     string `json:"doctype"`
}

type FarmerFollowUp struct {
	Name       string `json:"name"`
	Owner      string `json:"owner"`
	Creation   string `json:"creation"`
	Modified   string `json:"modified"`
	ModifiedBy string `json:"modified_by"`
	DocStatus  int    `json:"docstatus"`
	Idx        int    `json:"idx"`

	Farmer string `json:"farmer"`
	Status string `json:"status"`

	Parent      string `json:"parent"`
	ParentField string `json:"parentfield"`
	ParentType  string `json:"parenttype"`
	Doctype     string `json:"doctype"`
}

type FarmFollowUp struct {
	Name       string `json:"name"`
	Owner      string `json:"owner"`
	Creation   string `json:"creation"`
	Modified   string `json:"modified"`
	ModifiedBy string `json:"modified_by"`
	DocStatus  int    `json:"docstatus"`
	Idx        int    `json:"idx"`

	Farm                string  `json:"farm"`
	FarmName            string  `json:"farm_name"`
	FarmCode            string  `json:"farm_code"`
	Region              string  `json:"region"`
	AreaFeddan          float64 `json:"area_feddan"`
	Latitude            string  `json:"latitude"`
	Longitude           string  `json:"longitude"`
	UpscalingProject    string  `json:"upscaling_project"`
	SubUpscalingProject string  `json:"subupscalingproject"`
	FarmGroup           string  `json:"farm_group"`

	FarmOwner              string `json:"farm_owner"`
	Phone                  string `json:"phone"`
	PictureOfFollower      string `json:"picture_of_follower"`
	WarehousesNotes        string `json:"warehouses_notes"`
	RecordImage            string `json:"record_image"`
	AnimalsTypeCount       string `json:"animals_type_count"`
	UsedQuantities         string `json:"الكميات_المستخدمة"`
	BioControlDetails      string `json:"bio_control_details"`
	ChemicalControlDetails string `json:"chemical_control_details"`
	NaturalEnemiesDetails  string `json:"natural_enemies_details"`
	WeedDisposalOther      string `json:"weed_disposal_other"`
	IrrigationMethodOther  string `json:"irrigation_method_other"`
	WaterSourceOther       string `json:"water_source_other"`
	EnergyDetails          string `json:"energy_details"`
	SoilAnalysisDetails    string `json:"soil_analysis_details"`
	CurrentChallenges      string `json:"current_challenges"`
	SupportNeededOther     string `json:"support_needed_other"`
	FollowerAssessment     string `json:"follower_assessment"`
	Recommendations        string `json:"follower_recommendations"`

	GPS       string `json:"gps"`
	VisitDate string `json:"visit_date"`

	FollowerType  string `json:"follower_type"`
	FollowerName  string `json:"follower_name"`
	FarmersCount  int    `json:"farmers_count"`
	LastEBDAVisit string `json:"last_ebda_visit"`

	HasHealthInsuranceCard int    `json:"has_health_insurance_card"`
	StorageExist           string `json:"storage_exist"`
	RecordsFarmBook        int    `json:"records_farm_book"`

	IntercroppingOrGreenManure string  `json:"intercropping_or_green_manure"`
	IntercroppingPercent       float64 `json:"intercropping_percent"`

	PlantedTreesOrHedge string `json:"planted_trees_or_hedge"`
	TreesCount          int    `json:"trees_count"`
	HasAnimals          string `json:"has_animals"`

	FertilizationSource string  `json:"fertilization_source"`
	CompostSource       string  `json:"compost_source"`
	CompostProduction   float64 `json:"compost_production"`
	CompostQtys         float64 `json:"compost_qtys"`

	UsesBioProducts string  `json:"uses_bio_products"`
	CompQty         float64 `json:"comp_qty"`
	Qurts           float64 `json:"qurts"`
	Horns           float64 `json:"horns"`

	PestsOrDiseasesLastSeason string `json:"pests_or_diseases_last_season"`
	PestControlMethod         string `json:"pest_control_method"`
	UsesNaturalEnemies        int    `json:"uses_natural_enemies"`

	WeedDisposal            string `json:"weed_disposal"`
	IrrigationMethod        string `json:"irrigation_method"`
	WaterSource             string `json:"water_source"`
	WaterShortageOrSalinity int    `json:"water_shortage_or_salinity"`

	EnergyType                  string `json:"energy_type"`
	ClimateChallengesThisSeason int    `json:"climate_challenges_this_season"`
	YieldVsLastSeason           string `json:"yield_vs_last_season"`

	NeedCertSupport int    `json:"need_cert_support"`
	SupportNeeded   string `json:"support_needed"`

	Doctype      string `json:"doctype"`
	ServicesUsed []any  `json:"services_used"`

	FarmersNames        []FarmerFollowUp     `json:"farmers_names"`
	BiosProductsDetails []BioProductFollowUp `json:"bios_products_details"`
	CurrentCrops        []CropFollowUp       `json:"curent_crops"`
}

func main() {
	f := FarmFollowUp{}
	err := json.Unmarshal([]byte(data), &f)
	if err != nil {
		panic(err)
	}

	fmt.Println(f.CurrentCrops)
}
