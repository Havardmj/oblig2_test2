package database



type MgoDB struct {
	MngoDBURL string
	MngoName string
	CurrencyCollection string
	WebHookCol string
}


type Currency struct {
	Base string `json:"base"`
	Date string `json:"date"`
	Rate map[string]float64 `json:"rates"`
}
type Webhookers struct {
	HId bson.ObjectId `bson:"_id"`
	HUrl string `json:"webhookURL"`
	Base string `json:"baseCurrency"`
	TargetCurrency string `json:"targetCurrency"`
	MinTriggerValue float32 `json:"minTriggerValue"`
	MaxTriggerValue float32 `json:"maxTriggerValue"`
}

func(db *currencyMongoDB) Init() {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()
}

func(db *currencyMongoDB) Addcurrency(cu Currency) {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	err = session.DB(db.MngoName).C(db.CurrencyCollection).Insert(cu)
	if err != nil {
		fmt.Printf("Something went wrong with adding data to mongoDB %v", err)
	}

}
func (db *MgoDB) AddWebHook(h Webhookers)(string, error) {
	session, err := mgo.Dial(db.MngoDBURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	h.HId = bson.NewObjectId()
	id := h.HId

	err = session.DB(db.MngoName).C(db.WebHookCol).Insert(h)
	if err != nil {
		fmt.Printf("something went wrong in Adding webhook %v", err.Error())
		return "", err
	}
	return id, nil

}
func (db *MgoDB) getAverage(target string) float32 {
	session, err := mgo.Dial(db.MngoDBURL)
	if err != nil {
		fmt.Printf("something went wrong with getAverage func %v", err)
	}
	defer session.Close()

	var res []Currency
	t := time.Now().AddDate(0, 0, -3).Format("2015-01-02")
	s := "rates." + target
	err = session.DB(db.MngoName).C(db.CurrencyCollection).Find(bson.M{sort.SearchString: target,"date": bson.M{"$gt": t}}).All(&res)
	if err != nil {
		fmt.Printf("did not retreive average %v", err)
		return nil
	}
	return res

}
func (db * MgoDB) getWebHook(obj string) (Webhookers, error) {
	tmp := Webhookers{}
	session,err := mgo.Dial(db.MngoDBURL)
	if err != nil {
		fmt.Printf("something went wrong with getWebHook %v", err)
		return tmp, err
	}
	defer session.Close()
	tmpObj := bson.ObjectId(obj)
	err = session.DB(db.MngoName).C(db.WebHookCol).Find(bson.M{"_id":tempID}).One(&tmp)
	return tmp, nil
}