package service

func checkServiceType(serviceType string) bool {
	serviceTypes := []string{"Construction", "Delivery", "Manufacture"}
	for _, sT := range serviceTypes {
		if sT == serviceType {
			return true
		}
	}

	return false
}

func checkStatus(status string) bool {
	statuses := []string{"Created", "Published", "Closed"}
	for _, s := range statuses {
		if status == s {
			return true
		}
	}
	return false
}
