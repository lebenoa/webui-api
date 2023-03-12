package upscaler

type model struct {
	Lanczos               string
	Nearest               string
	LDSR                  string
	SwinIR4x              string
	ESRGAN4x              string
	REsrgan4xPlus         string
	REsrgan4xPlusAnime6B  string
	REsrganGeneral4xV3    string
	REsrganGeneralWdn4xV3 string
	REsrganAnimeVideo     string
	REsrgan2xPlus         string

	LANCZOS                    string
	NEAREST                    string
	SWIN_IR_4x                 string
	ESRGAN_4x                  string
	R_ESRGAN_4x_PLUS           string
	R_ESRGAN_4x_PLUS_ANIME_6B  string
	R_ESRGAN_GENERAL_4x_V3     string
	R_ESRGAN_GENERAL_WDN_4x_V3 string
	R_ESRGAN_ANIME_VIDEO       string
	R_ESRGAN_2x_PLUS           string
}

var Model = model{
	Lanczos: Lanczos,
	LANCZOS: Lanczos,

	Nearest: Nearest,
	NEAREST: Nearest,

	LDSR: LDSR,

	SwinIR4x:   SwinIR4x,
	SWIN_IR_4x: SwinIR4x,

	ESRGAN4x:  ESRGAN4x,
	ESRGAN_4x: ESRGAN4x,

	REsrgan4xPlus:    REsrgan4xPlus,
	R_ESRGAN_4x_PLUS: REsrgan4xPlus,

	REsrgan4xPlusAnime6B:      REsrgan4xPlusAnime6B,
	R_ESRGAN_4x_PLUS_ANIME_6B: REsrgan4xPlusAnime6B,

	REsrganGeneral4xV3:     REsrganGeneral4xV3,
	R_ESRGAN_GENERAL_4x_V3: REsrganGeneral4xV3,

	REsrganGeneralWdn4xV3:      REsrganGeneralWdn4xV3,
	R_ESRGAN_GENERAL_WDN_4x_V3: REsrganGeneralWdn4xV3,

	REsrganAnimeVideo:    REsrganAnimeVideo,
	R_ESRGAN_ANIME_VIDEO: REsrganAnimeVideo,

	REsrgan2xPlus:    REsrgan2xPlus,
	R_ESRGAN_2x_PLUS: REsrgan2xPlus,
}
