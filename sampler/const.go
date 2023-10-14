package sampler

const (
	EULER           = "Euler"
	EULER_A         = "Euler a"
	EULER_ASCESTRAL = "Euler a" // EULER_A alias

	LMS        = "LMS"
	LMS_KARRAS = "LMS Karras"

	HEUN = "Heun"

	DPM_FAST     = "DPM fast"
	DPM_ADAPTIVE = "DPM adaptive"
	DPM_PLUS_PLUS = "DPM++"

	DPM2          = "DPM2"
	DPM2_KARRAS   = "DPM2 Karras"
	DPM2_A        = "DPM2 a"
	DPM2_A_KARRAS = "DPM2 a Karras"

	DPM_PLUS_PLUS_2S_A        	= "DPM++ 2S a"
	DPM_PLUS_PLUS_2S_A_KARRAS 	= "DPM++ 2S a Karras"
	
	DPM_PLUS_PLUS_2M          	= "DPM++ 2M"
	DPM_PLUS_PLUS_2M_KARRAS   	= "DPM++ 2M Karras"

	DPM_PLUS_PLUS_SDE         	= "DPM++ SDE"
	DPM_PLUS_PLUS_SDE_KARRAS  	= "DPM++ SDE Karras"

	DPM_PLUS_PLUS_2M_SDE_EXP  	= "DPM++ 2M SDE Exponential"
	DPM_PLUS_PLUS_2M_SDE_HEUN 	= "DPM++ 2M SDE Heun" 
	DPM_PLUS_PLUS_2M_SDE_HEUN_KARRAS = "DPM++ 2M SDE Heun Karras"
	DPM_PLUS_PLUS_2M_SDE_HEUN_EXP = "DPM++ 2M SDE Heun Exponential" 

	DPM_PLUS_PLUS_3M_SDE 		= "DPM++ 3M SDE" 
	DPM_PLUS_PLUS_3M_SDE_KARRAS = "DPM++ 3M SDE Karras" 
	DPM_PLUS_PLUS_3M_SDE_EXP 	= "DPM++ 3M SDE Exponential"

	RESTART = "Restart"
	DDIM   = "DDIM"
	PLMS   = "PLMS"
	UNI_PC = "UniPC"
)
