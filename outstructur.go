package rosreestrcn

type ObjectInfo struct {
	ObjectID      string `json:"objectId"`
	Type          string `json:"type"`
	RegionKey     int    `json:"regionKey"`
	Source        int    `json:"source"`
	FirActualDate string `json:"firActualDate"`

	ObjectData struct {
		ID            string `json:"id"`
		RegionKey     int    `json:"regionKey"`
		SrcObject     int    `json:"srcObject"`
		ObjectType    string `json:"objectType"`
		ObjectName    string `json:"objectName"`
		Removed       int    `json:"removed"`
		DateLoad      string `json:"dateLoad"`
		AddressNote   string `json:"addressNote"`
		ObjectCn      string `json:"objectCn"`
		ObjectCon     string `json:"objectCon"`
		ObjectInv     string `json:"objectInv"`
		ObjectUn      string `json:"objectUn"`
		RsCode        string `json:"rsCode"`
		ActualDate    string `json:"actualDate"`
		BrkStatus     int    `json:"brkStatus"`
		BrkDate       string `json:"brkDate"`
		FormRights    string `json:"formRights"`
		ObjectAddress struct {
			ID            string `json:"id"`
			RegionKey     int    `json:"regionKey"`
			Okato         string `json:"okato"`
			Kladr         string `json:"kladr"`
			Region        string `json:"region"`
			District      string `json:"district"`
			DistrictType  string `json:"districtType"`
			Place         string `json:"place"`
			PlaceType     string `json:"placeType"`
			Locality      string `json:"locality"`
			LocalityType  string `json:"localityType"`
			Street        string `json:"street"`
			StreetType    string `json:"streetType"`
			House         string `json:"house"`
			Building      string `json:"building"`
			Structure     string `json:"structure"`
			Apartment     string `json:"apartment"`
			AddressNotes  string `json:"addressNotes"`
			MergedAddress string `json:"mergedAddress"`
		} `json:"objectAddress"`
	} `json:"objectData"`

	ParcelData struct {
		ID                      string  `json:"id"`
		RegionKey               int     `json:"regionKey"`
		ParcelCn                string  `json:"parcelCn"`
		ParcelStatus            string  `json:"parcelStatus"`
		DateCreate              string  `json:"dateCreate"`
		DateRemove              string  `json:"dateRemove"`
		CategoryType            string  `json:"categoryType"`
		AreaValue               float64 `json:"areaValue"`
		AreaType                string  `json:"areaType"`
		AreaUnit                string  `json:"areaUnit"`
		AreaTypeValue           string  `json:"areaTypeValue"`
		AreaUnitValue           string  `json:"areaUnitValue"`
		CategoryTypeValue       string  `json:"categoryTypeValue"`
		RightsReg               bool    `json:"rightsReg"`
		CadCost                 float64 `json:"cadCost"`
		CadUnit                 string  `json:"cadUnit"`
		DateCost                string  `json:"dateCost"`
		OksFlag                 int     `json:"oksFlag"`
		OksType                 string  `json:"oksType"`
		OksFloors               string  `json:"oksFloors"`
		OksUFloors              string  `json:"oksUFloors"`
		OksElementsConstruct    string  `json:"oksElementsConstruct"`
		OksYearUsed             string  `json:"oksYearUsed"`
		OksInventoryCost        float64 `json:"oksInventoryCost"`
		OksInn                  string  `json:"oksInn"`
		OksExecutor             string  `json:"oksExecutor"`
		OksYearBuilt            string  `json:"oksYearBuilt"`
		OksCostDate             string  `json:"oksCostDate"`
		RcType                  string  `json:"rcType"`
		RcDate                  string  `json:"rcDate"`
		GUIDUl                  string  `json:"guidUl"`
		GUIDFl                  string  `json:"guidFl"`
		CiSurname               string  `json:"ciSurname"`
		CiFirst                 string  `json:"ciFirst"`
		CiPatronymic            string  `json:"ciPatronymic"`
		CiNCertificate          string  `json:"ciNCertificate"`
		CiPhone                 string  `json:"ciPhone"`
		CiEmail                 string  `json:"ciEmail"`
		CiAddress               string  `json:"ciAddress"`
		CoName                  string  `json:"coName"`
		CoInn                   string  `json:"coInn"`
		UtilCode                string  `json:"utilCode"`
		UtilByDoc               string  `json:"utilByDoc"`
		CadastralBlockID        string  `json:"cadastralBlockId"`
		ParcelStatusStr         string  `json:"parcelStatusStr"`
		OksElementsConstructStr string  `json:"oksElementsConstructStr"`
		UtilCodeDesc            string  `json:"utilCodeDesc"`
	} `json:"parcelData"`
	RealtyData              string   `json:"realtyData"`
	PremisesData            string   `json:"premisesData"`
	RightEncumbranceObjects string   `json:"rightEncumbranceObjects"`
	OldNumbers              []string `json:"oldNumbers"`
}
