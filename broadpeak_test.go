package broadpeakio

import (
	"testing"
	"time"
)

const apiKey = "TOKEN VALUE"

func TestAdInsertion(t *testing.T) {
	broadpeak := MakeClient(apiKey)

	var liveOptions = LiveInput{Name: "Test Ad", Url: "https://d1uz2gd9rhh5xg.cloudfront.net/out/v1/6e0f649095ca4131b16bd0f877048629/index.mpd?testgo"}

	createLiveResponse, err := broadpeak.CreateLive(liveOptions)
	if err != nil {
		t.Errorf("CreateLive failed with error:%v", err)
	}

	liveId := createLiveResponse.Id

	var adServerOptions = AdServerInput{Name: "Test", Url: "https://bitmovin-a.akamaihd.net/content/MI201109210084_1/mpds/f08e80da-bf1d-4e3d-8899-f0f6155f6efa.mpd"}

	createAdServerResponse, err := broadpeak.CreateAdServer(adServerOptions)
	if err != nil {
		t.Errorf("CreateAdServer failed with error:%v", err)
	}

	adServerId := createAdServerResponse.Id

	var options = CreateAdInsertionInput{Name: "Test", Source: &Identifiable{liveId}, LiveAdReplacement: &LiveAdReplacement{AdServer: &Identifiable{adServerId}}}

	createAdInsertionResponse, err := broadpeak.CreateAdInsertion(options)
	if err != nil {
		t.Errorf("CreateAdInsertion failed with error:%v", err)
	}

	id := createAdInsertionResponse.Id
	_, err = broadpeak.GetAdInsertion(uint(id))
	if err != nil {
		t.Errorf("GetAdInsertion failed with error:%v", err)
	}

	updateOptions := UpdateAdInsertionInput{Name: "Test1", LiveAdReplacement: &LiveAdReplacement{AdServer: &Identifiable{adServerId}}}
	_, err = broadpeak.UpdateAdInsertion(uint(id), updateOptions)
	if err != nil {
		t.Errorf("UpdateAdInsertion failed with error:%v", err)
	}

	_, err = broadpeak.DeleteAdInsertion(uint(id))
	if err != nil {
		t.Errorf("DeleteAdInsertion failed with error:%v", err)
	}

	_, err = broadpeak.DeleteAdServer(uint(adServerId))
	if err != nil {
		t.Errorf("DeleteAdServer failed with error:%v", err)
	}

	_, err = broadpeak.DeleteLive(uint(liveId))
	if err != nil {
		t.Errorf("DeleteLive failed with error:%v", err)
	}
}

func TestAdServer(t *testing.T) {
	broadpeak := MakeClient(apiKey)

	var options = AdServerInput{Name: "Test", Url: "https://bitmovin-a.akamaihd.net/content/MI201109210084_1/mpds/f08e80da-bf1d-4e3d-8899-f0f6155f6efa.mpd"}

	createAdServerResponse, err := broadpeak.CreateAdServer(options)
	if err != nil {
		t.Errorf("CreateAdServer failed with error:%v", err)
	}

	id := createAdServerResponse.Id
	_, err = broadpeak.GetAdServer(uint(id))
	if err != nil {
		t.Errorf("GetAdServer failed with error:%v", err)
	}

	updateOptions := AdServerInput{Name: "Test1", Url: "https://bitmovin-a.akamaihd.net/content/MI201109210084_1/mpds/f08e80da-bf1d-4e3d-8899-f0f6155f6efa.mpd"}
	_, err = broadpeak.UpdateAdServer(uint(id), updateOptions)
	if err != nil {
		t.Errorf("UpdateAdServer failed with error:%v", err)
	}

	_, err = broadpeak.DeleteAdServer(uint(id))
	if err != nil {
		t.Errorf("DeleteAdServer failed with error:%v", err)
	}
}

func TestApiKey(t *testing.T) {
	broadpeak := MakeClient(apiKey)

	timeVal := time.Now().UTC()

	_, err := broadpeak.CreateApiKey("TestApiKeySdk", timeVal.Format("2006-01-02T15:04:05Z"))
	if err != nil {
		t.Errorf("CreateApiKey failed with error:%v", err)
	}

	_, err = broadpeak.GetAllApiKeys(0, 0)
	if err != nil {
		t.Errorf("GetAllApiKeys failed with error:%v", err)
	}

	_, err = broadpeak.DeleteApiKey("TestApiKeySdk")
	if err != nil {
		t.Errorf("DeleteApiKey failed with error:%v", err)
	}
}

func TestApplicationStatus(t *testing.T) {
	broadpeak := MakeClient(apiKey)

	_, err := broadpeak.CheckApplicationStatus()
	if err != nil {
		t.Errorf("CheckApplicationStatus failed with error:%v", err)
	}
}

func TestAssetCatalog(t *testing.T) {
	broadpeak := MakeClient(apiKey)

	var options = AssetCatalogInput{Name: "TestAC", Url: "http://ftp.itec.aau.at/datasets/DASHDataset2014/TearsOfSteel/2sec/", AssetSample: "TearsOfSteel_2s_simple_2014_05_09.mpd"}

	createAssetCatalogResponse, err := broadpeak.CreateAssetCatalog(options)

	if err != nil {
		t.Errorf("CreateAssetCatalog failed with error:%v", err)
	}

	id := createAssetCatalogResponse.Id

	_, err = broadpeak.GetAssetCatalog(uint(id))
	if err != nil {
		t.Errorf("GetAssetCatalog failed with error:%v", err)
	}

	updateOptions := AssetCatalogInput{Name: "Test1", Url: "http://ftp.itec.aau.at/datasets/DASHDataset2014/TearsOfSteel/2sec/", AssetSample: "TearsOfSteel_2s_simple_2014_05_09.mpd"}
	_, err = broadpeak.UpdateAssetCatalog(uint(id), updateOptions)
	if err != nil {
		t.Errorf("UpdateAssetCatalog failed with error:%v", err)
	}

	_, err = broadpeak.DeleteAssetCatalog(uint(id))
	if err != nil {
		t.Errorf("DeleteAssetCatalog failed with error:%v", err)
	}
}

func TestAssets(t *testing.T) {
	broadpeak := MakeClient(apiKey)

	var options = AssetInput{Name: "Test", Url: "https://dash.akamaized.net/dash264/TestCasesIOP33/adapatationSetSwitching/5/manifest.mpd"}

	createAssetResponse, err := broadpeak.CreateAsset(options)

	if err != nil {
		t.Errorf("CreateAsset failed with error:%v", err)
	}

	id := createAssetResponse.Id

	_, err = broadpeak.GetAsset(uint(id))
	if err != nil {
		t.Errorf("GetAsset failed with error:%v", err)
	}

	updateOptions := AssetInput{Name: "Test1", Url: "https://dash.akamaized.net/dash264/TestCasesIOP33/adapatationSetSwitching/5/manifest.mpd"}
	_, err = broadpeak.UpdateAsset(uint(id), updateOptions)
	if err != nil {
		t.Errorf("UpdateAsset failed with error:%v", err)
	}

	_, err = broadpeak.DeleteAsset(uint(id))
	if err != nil {
		t.Errorf("DeleteAsset failed with error:%v", err)
	}
}

func TestCategories(t *testing.T) {
	broadpeak := MakeClient(apiKey)

	var subcategories = []Subcategory{{Key: "xyz", Value: "abc"}}

	createCategoryResponse, err := broadpeak.CreateCategory("Test", subcategories)
	if err != nil {
		t.Errorf("CreateCategory failed with error:%v", err)
	}

	id := createCategoryResponse.Id

	_, err = broadpeak.GetAllCategories(0, 0)
	if err != nil {
		t.Errorf("GetAllCategories failed with error:%v", err)
	}

	_, err = broadpeak.GetCategory(uint(id))
	if err != nil {
		t.Errorf("GetCategory failed with error:%v", err)
	}

	var updateSubcategories = []Subcategory{{Key: "newKey", Value: "newVal"}}
	_, err = broadpeak.UpdateCategory(uint(id), "Test", updateSubcategories)
	if err != nil {
		t.Errorf("UpdateCategory failed with error:%v", err)
	}

	_, err = broadpeak.DeleteCategory(uint(id))
	if err != nil {
		t.Errorf("DeleteCategory failed with error:%v", err)
	}
}

func TestContentReplacement(t *testing.T) {
	broadpeak := MakeClient(apiKey)

	var optionsLive = LiveInput{Name: "TestLive", Url: "https://livesim.dashif.org/livesim/chunkdur_1/ato_7/testpic4_8s/Manifest.mpd"}

	createLiveResponse, err := broadpeak.CreateLive(optionsLive)
	if err != nil {
		t.Errorf("CreateLive failed with error:%v", err)
	}

	idLive := createLiveResponse.Id

	var optionsAsset = AssetInput{Name: "TestAsset", Url: "https://dash.akamaized.net/dash264/TestCasesIOP33/adapatationSetSwitching/5/manifest.mpd"}

	createAssetResponse, err := broadpeak.CreateAsset(optionsAsset)

	if err != nil {
		t.Errorf("CreateAsset failed with error:%v", err)
	}

	idAsset := createAssetResponse.Id

	var options = CreateContentReplacementInput{Name: "Test", Source: &Identifiable{idLive}, Replacement: &Identifiable{idAsset}}

	createContentReplacement, err := broadpeak.CreateContentReplacement(options)

	if err != nil {
		t.Errorf("CreateContentReplacement failed with error:%v", err)
	}

	id := createContentReplacement.Id
	_, err = broadpeak.GetContentReplacement(uint(id))
	if err != nil {
		t.Errorf("GetContentReplacement failed with error:%v", err)
	}

	updateOptions := UpdateContentReplacementInput{Name: "Test1", Replacement: &Identifiable{idAsset}}
	_, err = broadpeak.UpdateContentReplacement(uint(id), updateOptions)
	if err != nil {
		t.Errorf("UpdateContentReplacement failed with error:%v", err)
	}

	_, err = broadpeak.DeleteContentReplacement(uint(id))
	if err != nil {
		t.Errorf("DeleteContentReplacement failed with error:%v", err)
	}

	_, err = broadpeak.DeleteLive(uint(idLive))
	if err != nil {
		t.Errorf("DeleteLive failed with error:%v", err)
	}

	_, err = broadpeak.DeleteAsset(uint(idAsset))
	if err != nil {
		t.Errorf("DeleteAsset failed with error:%v", err)
	}
}

func TestSlotsContentReplacement(t *testing.T) {
	broadpeak := MakeClient(apiKey)

	var optionsLive = LiveInput{Name: "TestLive", Url: "https://livesim.dashif.org/livesim/chunkdur_1/ato_7/testpic4_8s/Manifest.mpd"}

	createLiveResponse, err := broadpeak.CreateLive(optionsLive)
	if err != nil {
		t.Errorf("CreateLive failed with error:%v", err)
	}

	idLive := createLiveResponse.Id

	var optionsAsset = AssetInput{Name: "TestAsset", Url: "https://dash.akamaized.net/dash264/TestCasesIOP33/adapatationSetSwitching/5/manifest.mpd"}

	createAssetResponse, err := broadpeak.CreateAsset(optionsAsset)

	if err != nil {
		t.Errorf("CreateAsset failed with error:%v", err)
	}

	idAsset := createAssetResponse.Id

	var optionsCreateContentReplacement = CreateContentReplacementInput{Name: "Test", Source: &Identifiable{idLive}, Replacement: &Identifiable{idAsset}}

	createContentReplacement, err := broadpeak.CreateContentReplacement(optionsCreateContentReplacement)

	if err != nil {
		t.Errorf("CreateContentReplacement failed with error:%v", err)
	}

	idCreateContentReplacement := createContentReplacement.Id

	layout := "2006-01-02T15:04:05Z"
	currentTime := time.Now().UTC()
	formattedCurrentTime := currentTime.Format(layout)

	fiveDaysAfter := currentTime.AddDate(0, 0, 5)
	formattedFiveDaysAfter := fiveDaysAfter.Format(layout)

	sixDaysAfter := currentTime.AddDate(0, 0, 6)
	formattedSixDaysAfter := sixDaysAfter.Format(layout)

	var options = CreateContentReplacementSlotInput{StartTime: formattedFiveDaysAfter, Duration: 1200}

	createContentReplacementSlot, err := broadpeak.CreateContentReplacementSlot(uint(idCreateContentReplacement), options)

	if err != nil {
		t.Errorf("CreateContentReplacementSlot failed with error:%v", err)
	}

	id := createContentReplacementSlot.Id

	_, err = broadpeak.GetContentReplacementSlot(uint(idCreateContentReplacement), uint(id))
	if err != nil {
		t.Errorf("GetContentReplacementSlot failed with error:%v", err)
	}

	getAllSlotsOptions := GetAllSlotsInput{Offset: 0, Limit: 10, From: formattedCurrentTime, To: formattedSixDaysAfter}
	_, err = broadpeak.GetAllContentReplacementSlots(uint(idCreateContentReplacement), getAllSlotsOptions)
	if err != nil {
		t.Errorf("GetAllContentReplacementSlots failed with error:%v", err)
	}

	updateOptions := UpdateContentReplacementSlotInput{StartTime: formattedFiveDaysAfter, Duration: 1100}
	_, err = broadpeak.UpdateContentReplacementSlot(uint(idCreateContentReplacement), uint(id), updateOptions)
	if err != nil {
		t.Errorf("UpdateContentReplacementSlot failed with error:%v", err)
	}

	_, err = broadpeak.DeleteContentReplacementSlot(uint(idCreateContentReplacement), uint(id))
	if err != nil {
		t.Errorf("DeleteContentReplacementSlot failed with error:%v", err)
	}

	_, err = broadpeak.DeleteContentReplacement(uint(idCreateContentReplacement))
	if err != nil {
		t.Errorf("DeleteContentReplacement failed with error:%v", err)
	}

	_, err = broadpeak.DeleteLive(uint(idLive))
	if err != nil {
		t.Errorf("DeleteLive failed with error:%v", err)
	}

	_, err = broadpeak.DeleteAsset(uint(idAsset))
	if err != nil {
		t.Errorf("DeleteAsset failed with error:%v", err)
	}
}

func TestLive(t *testing.T) {
	broadpeak := MakeClient(apiKey)

	var options = LiveInput{Name: "Test", Url: "https://livesim.dashif.org/livesim/chunkdur_1/ato_7/testpic4_8s/Manifest.mpd"}

	createLiveResponse, err := broadpeak.CreateLive(options)
	if err != nil {
		t.Errorf("CreateLive failed with error:%v", err)
	}

	id := createLiveResponse.Id
	_, err = broadpeak.GetLive(uint(id))
	if err != nil {
		t.Errorf("GetLive failed with error:%v", err)
	}

	updateOptions := LiveInput{Name: "Test1", Url: "https://livesim.dashif.org/livesim/chunkdur_1/ato_7/testpic4_8s/Manifest.mpd"}
	_, err = broadpeak.UpdateLive(uint(id), updateOptions)
	if err != nil {
		t.Errorf("UpdateLive failed with error:%v", err)
	}

	_, err = broadpeak.DeleteLive(uint(id))
	if err != nil {
		t.Errorf("DeleteLive failed with error:%v", err)
	}
}

func TestServices(t *testing.T) {
	broadpeak := MakeClient(apiKey)

	_, err := broadpeak.GetAllServices(0, 5)
	if err != nil {
		t.Errorf("GetAllServices failed with error:%v", err)
	}
}

func TestSlate(t *testing.T) {
	broadpeak := MakeClient(apiKey)

	var options = SlateInput{Name: "Test", Url: "https://rachittestbroadpeak.s3.amazonaws.com/localisejs+spanish+conversion.jpg"}

	createSlateResponse, err := broadpeak.CreateSlate(options)

	if err != nil {
		t.Errorf("CreateSlate failed with error:%v", err)
	}

	id := createSlateResponse.Id
	_, err = broadpeak.GetSlate(uint(id))
	if err != nil {
		t.Errorf("GetSlate failed with error:%v", err)
	}

	updateOptions := SlateInput{Name: "Test1", Url: "https://rachittestbroadpeak.s3.amazonaws.com/localisejs+spanish+conversion.jpg"}
	_, err = broadpeak.UpdateSlate(uint(id), updateOptions)
	if err != nil {
		t.Errorf("UpdateSlate failed with error:%v", err)
	}

	_, err = broadpeak.DeleteSlate(uint(id))
	if err != nil {
		t.Errorf("DeleteSlate failed with error:%v", err)
	}
}

func TestSources(t *testing.T) {
	broadpeak := MakeClient(apiKey)

	_, err := broadpeak.GetAllSources(0, 0)
	if err != nil {
		t.Errorf("GetAllSources failed with error:%v", err)
	}
	_, err = broadpeak.CheckSourceStatus("live", "https://origin.broadpeak.io/bpk-tv/cycling/default/index.mpd")
	if err != nil {
		t.Errorf("CheckSourceStatus failed with error:%v", err)
	}
}

func TestTranscodingProfiles(t *testing.T) {
	broadpeak := MakeClient(apiKey)

	listProfiles, err := broadpeak.GetAllTranscodingProfiles(0, 1)
	if err != nil {
		t.Errorf("GetAllTranscodingProfiles failed with error:%v", err)
	}

	id := listProfiles[0].Id
	_, err = broadpeak.GetTranscodingProfile(uint(id))
	if err != nil {
		t.Errorf("GetTranscodingProfile failed with error:%v", err)
	}
}

/*func TestUsers(t *testing.T) {
	broadpeak := MakeClient(apiKey)

	var options = UserInput{FirstName: "Test", LastName: "Name", Email: "test1@test.com"}

	createUserResponse, err := broadpeak.CreateUser(options)

	if err != nil {
		t.Errorf("CreateUser failed with error:%v", err)
	}

	id := createUserResponse.Id

	_, err = broadpeak.GetAllUsers(0, 0, true)
	if err != nil {
		t.Errorf("GetAllUsers failed with error:%v", err)
	}

	userDetails, err := broadpeak.GetUser(uint(id))
	if err != nil {
		t.Errorf("GetUser failed with error:%v", err)
	}

	updateOptions := UserInput{FirstName: userDetails.FirstName, LastName: "NewLastName", Email: "test@test.com"}
	_, err = broadpeak.UpdateUser(uint(id), updateOptions)
	if err != nil {
		t.Errorf("UpdateUser failed with error:%v", err)
	}

	_, err = broadpeak.DeleteUser(uint(id))
	if err != nil {
		t.Errorf("DeleteUser failed with error:%v", err)
	}
}*/

func TestVirtualChannel(t *testing.T) {
	broadpeak := MakeClient(apiKey)

	var optionsLive = LiveInput{Name: "Test", Url: "https://livesim.dashif.org/livesim/chunkdur_1/ato_7/testpic4_8s/Manifest.mpd"}

	createLiveResponse, err := broadpeak.CreateLive(optionsLive)
	if err != nil {
		t.Errorf("CreateLive failed with error:%v", err)
	}

	idLive := createLiveResponse.Id

	var options = CreateVirtualChannelInput{Name: "Test", BaseLive: &Identifiable{idLive}}

	createVirtualChannel, err := broadpeak.CreateVirtualChannel(options)

	if err != nil {
		t.Errorf("CreateVirtualChannel failed with error:%v", err)
	}

	id := createVirtualChannel.Id
	_, err = broadpeak.GetVirtualChannel(uint(id))
	if err != nil {
		t.Errorf("GetVirtualChannel failed with error:%v", err)
	}

	updateOptions := UpdateVirtualChannelInput{Name: "Test1"}
	_, err = broadpeak.UpdateVirtualChannel(uint(id), updateOptions)
	if err != nil {
		t.Errorf("UpdateVirtualChannel failed with error:%v", err)
	}

	_, err = broadpeak.DeleteVirtualChannel(uint(id))
	if err != nil {
		t.Errorf("DeleteVirtualChannel failed with error:%v", err)
	}

	_, err = broadpeak.DeleteLive(uint(idLive))
	if err != nil {
		t.Errorf("DeleteLive failed with error:%v", err)
	}
}

func TestSlotsVirtualChannel(t *testing.T) {
	broadpeak := MakeClient(apiKey)

	var optionsLive = LiveInput{Name: "Test", Url: "https://livesim.dashif.org/livesim/chunkdur_1/ato_7/testpic4_8s/Manifest.mpd"}

	createLiveResponse, err := broadpeak.CreateLive(optionsLive)
	if err != nil {
		t.Errorf("CreateLive failed with error:%v", err)
	}

	idLive := createLiveResponse.Id

	var optionsCreateVirtualChannel = CreateVirtualChannelInput{Name: "Test", BaseLive: &Identifiable{idLive}}

	createVirtualChannel, err := broadpeak.CreateVirtualChannel(optionsCreateVirtualChannel)

	if err != nil {
		t.Errorf("CreateVirtualChannel failed with error:%v", err)
	}

	idVirtualChannel := createVirtualChannel.Id

	var optionsAsset = AssetInput{Name: "TestAsset", Url: "https://dash.akamaized.net/dash264/TestCasesIOP33/adapatationSetSwitching/5/manifest.mpd"}

	createAssetResponse, err := broadpeak.CreateAsset(optionsAsset)

	if err != nil {
		t.Errorf("CreateAsset failed with error:%v", err)
	}

	idAsset := createAssetResponse.Id

	layout := "2006-01-02T15:04:05Z"
	currentTime := time.Now().UTC()
	formattedCurrentTime := currentTime.Format(layout)

	fiveDaysAfter := currentTime.AddDate(0, 0, 5)
	formattedFiveDaysAfter := fiveDaysAfter.Format(layout)

	sixDaysAfter := currentTime.AddDate(0, 0, 6)
	formattedSixDaysAfter := sixDaysAfter.Format(layout)

	var options = CreateVirtualChannelSlotInput{StartTime: formattedFiveDaysAfter, Duration: 1200, Replacement: &Identifiable{idAsset}}

	createVirtualChannelSlot, err := broadpeak.CreateVirtualChannelSlot(uint(idVirtualChannel), options)

	if err != nil {
		t.Errorf("CreateVirtualChannelSlot failed with error:%v", err)
	}

	id := createVirtualChannelSlot.Id
	_, err = broadpeak.GetVirtualChannelSlot(uint(idVirtualChannel), uint(id))
	if err != nil {
		t.Errorf("GetVirtualChannelSlot failed with error:%v", err)
	}

	getAllSlotsOptions := GetAllSlotsInput{Offset: 0, Limit: 10, From: formattedCurrentTime, To: formattedSixDaysAfter}
	_, err = broadpeak.GetAllVirtualChannelSlots(uint(idVirtualChannel), getAllSlotsOptions)
	if err != nil {
		t.Errorf("GetAllVirtualChannelSlots failed with error:%v", err)
	}

	updateOptions := UpdateVirtualChannelSlotInput{StartTime: formattedFiveDaysAfter, Duration: 1100, Replacement: &Identifiable{idAsset}}
	_, err = broadpeak.UpdateVirtualChannelSlot(uint(idVirtualChannel), uint(id), updateOptions)
	if err != nil {
		t.Errorf("UpdateVirtualChannelSlot failed with error:%v", err)
	}

	_, err = broadpeak.DeleteVirtualChannelSlot(uint(idVirtualChannel), uint(id))
	if err != nil {
		t.Errorf("DeleteVirtualChannelSlot failed with error:%v", err)
	}

	_, err = broadpeak.DeleteVirtualChannel(uint(idVirtualChannel))
	if err != nil {
		t.Errorf("DeleteVirtualChannel failed with error:%v", err)
	}

	_, err = broadpeak.DeleteLive(uint(idLive))
	if err != nil {
		t.Errorf("DeleteLive failed with error:%v", err)
	}

	_, err = broadpeak.DeleteAsset(uint(idAsset))
	if err != nil {
		t.Errorf("DeleteAsset failed with error:%v", err)
	}
}

func TestConsumption(t *testing.T) {
	broadpeak := MakeClient(apiKey)

	layout := "2006-01-02T15:04:05Z"
	currentTime := time.Now().UTC()
	formattedCurrentTime := currentTime.Format(layout)

	fiveDaysAgo := currentTime.AddDate(0, 0, -5)
	formattedFiveDaysAgo := fiveDaysAgo.Format(layout)

	var options = AggregateConsumptionInput{From: formattedFiveDaysAgo, To: formattedCurrentTime}

	_, err := broadpeak.GetAggregateConsumption(options)
	if err != nil {
		t.Errorf("GetAggregateConsumption failed with error:%v", err)
	}
}

func TestSamples(t *testing.T) {
	broadpeak := MakeClient(apiKey)

	_, err := broadpeak.CreateSamples()
	if err != nil {
		t.Errorf("CreateSamples failed with error:%v", err)
	}

}
