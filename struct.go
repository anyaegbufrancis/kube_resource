package main

type AutoGenerated struct {
	APIVersion string `json:"apiVersion"`
	Items      []struct {
		APIVersion string `json:"apiVersion"`
		Kind       string `json:"kind"`
		Metadata   struct {
			Name      string `json:"name"`
			Namespace string `json:"namespace"`
		} `json:"metadata,omitempty"`
		Spec struct {
			Containers []struct {
				Image     string `json:"image"`
				Name      string `json:"name"`
				Resources struct {
					Limits struct {
						CPU    string `json:"cpu"`
						Memory string `json:"memory"`
					} `json:"limits"`
					Requests struct {
						CPU    string `json:"cpu"`
						Memory string `json:"memory"`
					} `json:"requests"`
				} `json:"resources"`
			} `json:"containers"`
		} `json:"spec"`
	} `json:"items"`
}

type Output struct {
	Namespace     string  `json:"namespace"`
	CPURequest    float64 `json:"cpuRequest"`
	CPULimit      float64 `json:"cpuLimit"`
	MemoryRequest float64 `json:"memoryRequest"`
	MemoryLimit   float64 `json:"memoryLimit"`
}
