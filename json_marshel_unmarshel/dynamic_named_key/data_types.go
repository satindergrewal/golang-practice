package main

// Currencies type
type Currencies struct {
	Reservein           float64 `json:"reservein"`
	Nativein            float64 `json:"nativein"`
	Reserveout          float64 `json:"reserveout"`
	Lastconversionprice float64 `json:"lastconversionprice"`
	Viaconversionprice  float64 `json:"viaconversionprice"`
	Fees                float64 `json:"fees"`
	Conversionfees      float64 `json:"conversionfees"`
}

// CurrencyInfo type
type CurrencyInfo struct {
	Name                 string        `json:"name"`
	Version              int           `json:"version"`
	Options              int           `json:"options"`
	Parent               string        `json:"parent"`
	Systemid             string        `json:"systemid"`
	Currencyid           string        `json:"currencyid"`
	Notarizationprotocol int           `json:"notarizationprotocol"`
	Proofprotocol        int           `json:"proofprotocol"`
	Idregistrationprice  int           `json:"idregistrationprice"`
	Idreferrallevels     int           `json:"idreferrallevels"`
	Minnotariesconfirm   int           `json:"minnotariesconfirm"`
	Billingperiod        int           `json:"billingperiod"`
	Notarizationreward   int           `json:"notarizationreward"`
	Startblock           int           `json:"startblock"`
	Endblock             int           `json:"endblock"`
	Currencies           []string      `json:"currencies"`
	Weights              []float64     `json:"weights"`
	Conversions          []float64     `json:"conversions"`
	Initialsupply        float64       `json:"initialsupply"`
	Prelaunchcarveout    float64       `json:"prelaunchcarveout"`
	Initialcontributions []float64     `json:"initialcontributions"`
	Preconversions       []float64     `json:"preconversions"`
	Eras                 []interface{} `json:"eras"`
}

// GetCurrencyConverter type
type GetCurrencyConverter struct {
	CurrencyInfo     map[string]CurrencyInfo `json:"-"`
	Lastnotarization struct {
		Version             int    `json:"version"`
		Currencyid          string `json:"currencyid"`
		Notaryaddress       string `json:"notaryaddress"`
		Notarizationheight  int    `json:"notarizationheight"`
		Mmrroot             string `json:"mmrroot"`
		Notarizationprehash string `json:"notarizationprehash"`
		Work                string `json:"work"`
		Stake               string `json:"stake"`
		Currencystate       struct {
			Flags             int    `json:"flags"`
			Currencyid        string `json:"currencyid"`
			Reservecurrencies []struct {
				Currencyid     string  `json:"currencyid"`
				Weight         float64 `json:"weight"`
				Reserves       float64 `json:"reserves"`
				Priceinreserve float64 `json:"priceinreserve"`
			} `json:"reservecurrencies"`
			Initialsupply        float64               `json:"initialsupply"`
			Emitted              float64               `json:"emitted"`
			Supply               float64               `json:"supply"`
			Currencies           map[string]Currencies `json:"currencies"`
			Nativefees           int64                 `json:"nativefees"`
			Nativeconversionfees int64                 `json:"nativeconversionfees"`
		} `json:"currencystate"`
		Prevnotarization  string        `json:"prevnotarization"`
		Prevheight        int           `json:"prevheight"`
		Crossnotarization string        `json:"crossnotarization"`
		Crossheight       int           `json:"crossheight"`
		Nodes             []interface{} `json:"nodes"`
	} `json:"lastnotarization"`
	Multifractional struct {
		Name                 string        `json:"name"`
		Version              int           `json:"version"`
		Options              int           `json:"options"`
		Parent               string        `json:"parent"`
		Systemid             string        `json:"systemid"`
		Currencyid           string        `json:"currencyid"`
		Notarizationprotocol int           `json:"notarizationprotocol"`
		Proofprotocol        int           `json:"proofprotocol"`
		Idregistrationprice  int           `json:"idregistrationprice"`
		Idreferrallevels     int           `json:"idreferrallevels"`
		Minnotariesconfirm   int           `json:"minnotariesconfirm"`
		Billingperiod        int           `json:"billingperiod"`
		Notarizationreward   int           `json:"notarizationreward"`
		Startblock           int           `json:"startblock"`
		Endblock             int           `json:"endblock"`
		Currencies           []string      `json:"currencies"`
		Weights              []float64     `json:"weights"`
		Conversions          []float64     `json:"conversions"`
		Initialsupply        float64       `json:"initialsupply"`
		Prelaunchcarveout    float64       `json:"prelaunchcarveout"`
		Initialcontributions []float64     `json:"initialcontributions"`
		Preconversions       []float64     `json:"preconversions"`
		Eras                 []interface{} `json:"eras"`
	} `json:"multifractional,omitempty"`
}
