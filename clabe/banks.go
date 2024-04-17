package clabe

// Bank represents a bank in Mexico. In Mexico, banking institutions are identified by a three-digit number
// assigned by the ABM. The following table is a Bank Catalogue provided by El Servicio de Administración Tributaria
// (SAT.gob.mx),
type Bank struct {
	Code string
	Name string
}

// Banks is a list of all banks in Mexico. [List]
//
// [List]: https://www.gob.mx/sat/documentos/catalogo-de-bancos
var Banks = []Bank{
	{Code: "001", Name: "BANXICO"},
	{Code: "002", Name: "BANAMEX"},
	{Code: "006", Name: "BANCOMEXT"},
	{Code: "009", Name: "BANOBRAS"},
	{Code: "012", Name: "BBVA MEXICO"},
	{Code: "014", Name: "SANTANDER"},
	{Code: "019", Name: "BANJERCITO"},
	{Code: "021", Name: "HSBC"},
	{Code: "030", Name: "BAJIO"},
	{Code: "036", Name: "INBURSA"},
	{Code: "042", Name: "MIFEL"},
	{Code: "044", Name: "SCOTIABANK"},
	{Code: "058", Name: "BANREGIO"},
	{Code: "059", Name: "INVEX"},
	{Code: "060", Name: "BANSI"},
	{Code: "062", Name: "AFIRME"},
	{Code: "072", Name: "BANORTE"},
	{Code: "106", Name: "BANK OF AMERICA"},
	{Code: "108", Name: "MUFG"},
	{Code: "110", Name: "JP MORGAN"},
	{Code: "112", Name: "BMONEX"},
	{Code: "113", Name: "VE POR MAS"},
	{Code: "126", Name: "CREDIT SUISSE"},
	{Code: "127", Name: "AZTECA"},
	{Code: "128", Name: "AUTOFIN"},
	{Code: "129", Name: "BARCLAYS"},
	{Code: "130", Name: "COMPARTAMOS"},
	{Code: "132", Name: "MULTIVA BANCO"},
	{Code: "133", Name: "ACTINVER"},
	{Code: "135", Name: "NAFIN"},
	{Code: "136", Name: "INTERCAM BANCO"},
	{Code: "137", Name: "BANCOPPEL"},
	{Code: "138", Name: "ABC CAPITAL"},
	{Code: "140", Name: "CONSUBANCO"},
	{Code: "141", Name: "VOLKSWAGEN"},
	{Code: "143", Name: "CIBANCO"},
	{Code: "145", Name: "BBASE"},
	{Code: "147", Name: "BANKAOOL"},
	{Code: "148", Name: "PAGATODO"},
	{Code: "150", Name: "INMOBILIARIO"},
	{Code: "151", Name: "DONDE"},
	{Code: "152", Name: "BANCREA"},
	{Code: "154", Name: "BANCO COVALTO"},
	{Code: "155", Name: "ICBC"},
	{Code: "156", Name: "SABADELL"},
	{Code: "157", Name: "SHINHAN"},
	{Code: "158", Name: "MIZUHO BANK"},
	{Code: "159", Name: "BANK OF CHINA"},
	{Code: "160", Name: "BANCO S3"},
	{Code: "166", Name: "BaBien"},
	{Code: "168", Name: "HIPOTECARIA FED"},
	{Code: "600", Name: "MONEXCB"},
	{Code: "601", Name: "GBM"},
	{Code: "602", Name: "MASARI"},
	{Code: "605", Name: "VALUE"},
	{Code: "608", Name: "VECTOR"},
	{Code: "613", Name: "MULTIVA CBOLSA"},
	{Code: "616", Name: "FINAMEX"},
	{Code: "617", Name: "VALMEX"},
	{Code: "620", Name: "PROFUTURO"},
	{Code: "630", Name: "CB INTERCAM"},
	{Code: "631", Name: "CI BOLSA"},
	{Code: "634", Name: "FINCOMUN"},
	{Code: "638", Name: "NU MEXICO"},
	{Code: "646", Name: "STP"},
	{Code: "648", Name: "TACTIV CB"},
	{Code: "652", Name: "CREDICAPITAL"},
	{Code: "653", Name: "KUSPIT"},
	{Code: "656", Name: "UNAGRA"},
	{Code: "659", Name: "ASP INTEGRA OPC"},
	{Code: "661", Name: "ALTERNATIVOS"}, // nolint:misspell
	{Code: "670", Name: "LIBERTAD"},
	{Code: "677", Name: "CAJA POP MEXICA"},
	{Code: "680", Name: "CRISTOBAL COLON"},
	{Code: "683", Name: "CAJA TELEFONIST"},
	{Code: "684", Name: "OPM"},
	{Code: "685", Name: "FONDO (FIRA)"},
	{Code: "686", Name: "INVERCAP"},
	{Code: "688", Name: "CREDICLUB"},
	{Code: "689", Name: "FOMPED"},
	{Code: "703", Name: "TESORED"},
	{Code: "706", Name: "ARCUS"},
	{Code: "710", Name: "NVIO"},
}

// BankMap is a map of all banks in Mexico. The key is the bank code and the value is the bank.
var BankMap = func() map[string]Bank {
	m := make(map[string]Bank, len(Banks))
	for _, bank := range Banks {
		m[bank.Code] = bank
	}

	return m
}()

// GetBank returns a bank by its code.
func GetBank(code string) (Bank, bool) {
	val, ok := BankMap[code]

	return val, ok
}
