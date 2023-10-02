# Broadpeak.io SDK

## Package

To use the sdk in your application, you need to include the package: "github.com/broadpeak-io/go-sdk"

## Initializing the Broadpeak object

The main object of the SDK is broadpeakio. The Broadpeak object can be initialized using the below code:

	broadpeak := broadpeakio.MakeClient(apiKey)

API key: You can get the API key from account settings > API keys.

## How to use the SDK

Once you add the broadpeakio package and instantiate the broadpeak object you can call the methods of your choice by referencing those methods using the broadpeak object. For example:

    response, error := broadpeak.GetAllSources(0,0)

Following are the methods defined in the SDK with their required parameters and the expected response from the API call.

## Sources

### GetAllSources

#### Method Call

    response, err := broadpeak.GetAllSources(offset, limit)

#### Parameters

| Parameters | Type | Definition |
| -------------- | -------------- | -------------- |
| Offset | uint | Offset of the first result to return |
| Limit | uint | Maximum number of results to return |

#### Response

The response will be of the type `[]SourceOutput`

Type `SourceOutput` is defined as:

    {
        Id     uint   `json:"id"`
        Name   string `json:"name"`
        Url    string `json:"url"`
        Type   string `json:"type"`
        Format string `json:"format"`
    }

### GetSource

#### Method Call
	response, err = broadpeak.CheckSourceStatus(sourceType, sourceUrl)

#### Parameters

| Parameters | Type | Definition |
| -------------- | -------------- | -------------- |
| sourceType | string | Type of the source to be checked (live, asset...) |
| sourceUrl | string | URL of the source to be checked |

#### Response

The response will be of the type `SourceStatusOutput`

Type `SourceStatusOutput` is defined as:

    {
        MessageText   string `json:"messageText"`
        SeverityLEvel string `json:"severityLevel"`
    }

## Sources/Live

### CreateLiveSource

#### Method Call
	response, err := broadpeak.CreateLiveSource(options)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> options </td> <td> LiveSource </td>
<td>

```
{
    Name        string `json:"name,omitempty"` //required
    Url         string `json:"url,omitempty"`  //required
    Description string `json:"description,omitempty"`
    BackupIp    string `json:"backupIp,omitempty"`
}
```

</td>
</tr>
</table>

#### Response

The response will be of the type `LiveOutput`

Type `LiveOutput` is defined as:

    {
    	Name        string `json:"name"`
    	Description string `json:"description"`
    	Url         string `json:"url"`
    	BackupIp    string `json:"backupIp"`
    	Format      string `json:"format"`
    	Id          uint   `json:"id"`
    }

### GetLiveSource

#### Method Call

	response, err = broadpeak.GetLiveSource(Id)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> id</td> <td> uint</td>
<td>
Id of the live source assigned by broadpeak.io
</td>
</tr>
</table>

#### Response

The response will be of the type `LiveOutput`

Type `LiveOutput` is defined as:

    {
    	Name        string `json:"name"`
    	Description string `json:"description"`
    	Url         string `json:"url"`
    	BackupIp    string `json:"backupIp"`
    	Format      string `json:"format"`
    	Id          uint   `json:"id"`
    }

### UpdateLiveSource

#### Method Call

	response, err = broadpeak.UpdateLiveSource(Id, updateOptions)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> id</td> <td> uint</td>
<td>
Id of the live source assigned by broadpeak.io
</td>
</tr>
<tr>
<td> updateOptions</td> <td> LiveSource</td>
<td>

```
{
	Name        string `json:"name,omitempty"` //required
	Url         string `json:"url,omitempty"`  //required
	Description string `json:"description,omitempty"`
	BackupIp    string `json:"backupIp,omitempty"`
}
```

</td>
</tr>
</table>

#### Response

The response will be of the type `LiveOutput`

Type `LiveOutput` is defined as:

    {
    	Name        string `json:"name"`
    	Description string `json:"description"`
    	Url         string `json:"url"`
    	BackupIp    string `json:"backupIp"`
    	Format      string `json:"format"`
    	Id          uint   `json:"id"`
    }

### DeleteLiveSource

#### Method Call

	response, err = broadpeak.DeleteLiveSource(Id)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> id</td> <td> uint</td>
<td>
Id of the live source assigned by broadpeak.io
</td>
</tr>
</table>

#### Response
The response will be of the type `LiveOutput`

Type `LiveOutput` is defined as:

    {
        Name        string `json:"name"`
        Description string `json:"description"`
        Url         string `json:"url"`
        BackupIp    string `json:"backupIp"`
        Format      string `json:"format"`
        Id          uint   `json:"id"`
    }

## Sources/Asset

### CreateAsset

#### Method Call

response, err := broadpeak.CreateAsset(options)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> options</td> <td> Asset</td>
<td>

```
{
	Name        string `json:"name,omitempty"` //required
	Url         string `json:"url,omitempty"`  //required
	Description string `json:"description,omitempty"`
	BackupIp    string `json:"backupIp,omitempty"`
}
```

</td>
</tr>
</table>

#### Response

The response will be of the type `AssetOutput`

Type `AssetOutput` is defined as:

	{
    	Name        string `json:"name"`
    	Description string `json:"description"`
    	Url         string `json:"url"`
    	BackupIp    string `json:"backupIp"`
    	Format      string `json:"format"`
    	Id          uint   `json:"id"`
    }

### GetAsset

#### Method Call

	response, err = broadpeak.GetAsset(Id)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> id</td> <td> uint</td>
<td>
Id of the asset source assigned by broadpeak.io
</td>
</tr>
</table>

#### Response

The response will be of the type `AssetOutput`

Type `AssetOutput` is defined as:

	{
    	Name        string `json:"name"`
    	Description string `json:"description"`
    	Url         string `json:"url"`
    	BackupIp    string `json:"backupIp"`
    	Format      string `json:"format"`
    	Id          uint   `json:"id"`
    }

### UpdateAsset

#### Method Call

	response, err = broadpeak.UpdateAsset(Id, updateOptions)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> id</td> <td> uint</td>
<td>
Id of the asset source assigned by broadpeak.io
</td>
</tr>
<tr>
<td> updateOptions</td> <td> Asset</td>
<td>

```
{
	Name        string `json:"name,omitempty"` //required
	Url         string `json:"url,omitempty"`  //required
	Description string `json:"description,omitempty"`
	BackupIp    string `json:"backupIp,omitempty"`
}
```

</td>
</tr>
</table>

#### Response

The response will be of the type `AssetOutput`

Type `AssetOutput` is defined as:

	{
    	Name        string `json:"name"`
    	Description string `json:"description"`
    	Url         string `json:"url"`
    	BackupIp    string `json:"backupIp"`
    	Format      string `json:"format"`
    	Id          uint   `json:"id"`
    }

### DeleteAsset

#### Method Call

	response, err = broadpeak.DeleteAsset(Id)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> id</td> <td> uint</td>
<td>
Id of the asset source assigned by broadpeak.io
</td>
</tr>
</table>

### Response

The response will be of the type `AssetOutput`

Type `AssetOutput` is defined as:

	{
    	Name        string `json:"name"`
    	Description string `json:"description"`
    	Url         string `json:"url"`
    	BackupIp    string `json:"backupIp"`
    	Format      string `json:"format"`
    	Id          uint   `json:"id"`
    }

## Sources/AssetCatalog

### CreateAssetCatalog

#### Method Call

	response, err := broadpeak.CreateAssetCatalog(options)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> options</td> <td> AssetCatalog</td>
<td>

```
{
	Name        string `json:"name,omitempty"` //required
	Url         string `json:"url,omitempty"`  //required
	Description string `json:"description,omitempty"`
	AssetSample string `json:"assetSample,omitempty"` //required
	BackupIp    string `json:"backupIp,omitempty"`
}
```

</td>
</tr>
</table>

#### Response

The response will be of the type `AssetCatalogOutput`

Type `AssetCatalogOutput` is defined as:

	{
    	Name        string `json:"name"`
    	Description string `json:"description"`
    	Url         string `json:"url"`
    	BackupIp    string `json:"backupIp"`
    	AssetSample string `json:"assetSample"`
    	Id          uint   `json:"id"`
    }

### GetAssetCatalog

#### Method Call

	response, err = broadpeak.GetAssetCatalog(Id)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> id</td> <td> uint</td>
<td>
Id of the asset catalog source assigned by broadpeak.io
</td>
</tr>
</table>

#### Response

The response will be of the type `AssetCatalogOutput`

Type `AssetCatalogOutput` is defined as:

	{
    	Name        string `json:"name"`
    	Description string `json:"description"`
    	Url         string `json:"url"`
    	BackupIp    string `json:"backupIp"`
    	AssetSample string `json:"assetSample"`
    	Id          uint   `json:"id"`
    }

### UpdateAssetCatalog

#### Method Call

	response, err = broadpeak.UpdateAssetCatalog(Id, updateOptions)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> id</td> <td> uint</td>
<td>
Id of the asset catalog source assigned by broadpeak.io
</td>
</tr>
<tr>
<td> updateOptions</td> <td> AssetCatalog</td>
<td>

```
{
	Name        string `json:"name,omitempty"` //required
	Url         string `json:"url,omitempty"`  //required
	Description string `json:"description,omitempty"`
	BackupIp    string `json:"backupIp,omitempty"`
}
```

</td>
</tr>
</table>

#### Response

The response will be of the type `AssetCatalogOutput`

Type `AssetCatalogOutput` is defined as:

	{
    	Name        string `json:"name"`
    	Description string `json:"description"`
    	Url         string `json:"url"`
    	BackupIp    string `json:"backupIp"`
    	AssetSample string `json:"assetSample"`
    	Id          uint   `json:"id"`
    }

### DeleteAssetCatalog

#### Method Call

	response, err = broadpeak.DeleteAssetCatalog(Id)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> id</td> <td> uint</td>
<td>
Id of the asset catalog source assigned by broadpeak.io
</td>
</tr>
</table>

#### Response

The response will be of the type `AssetCatalogOutput`

Type `AssetCatalogOutput` is defined as:

    {
    	Name        string `json:"name"`
    	Description string `json:"description"`
    	Url         string `json:"url"`
    	BackupIp    string `json:"backupIp"`
    	AssetSample string `json:"assetSample"`
    	Id          uint   `json:"id"`
    }

## Sources/Slate

### CreateSlate

#### Method Call

	response, err := broadpeak.CreateSlate(options)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> options</td> <td> Slate</td>
<td>

```
{
	Name        string `json:"name"` //required
	Url         string `json:"url"`  //required
	Description string `json:"description"`
}
```

</td>
</tr>
</table>

#### Response

The response will be of the type `SlateOutput`

Type `SlateOutput` is defined as:

	{
    	Id          uint   `json:"id"`
    	Name        string `json:"name"`
    	Url         string `json:"url"`
    	Description string `json:"description"`
    	Format      string `json:"format"`
    }

### GetSlate

#### Method Call

	response, err = broadpeak.GetSlate(Id)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> id</td> <td> uint</td>
<td>
Id of the slate source assigned by broadpeak.io
</td>
</tr>
</table>

#### Response

The response will be of the type `SlateOutput`

Type `SlateOutput` is defined as:

	{
    	Id          uint   `json:"id"`
    	Name        string `json:"name"`
    	Url         string `json:"url"`
    	Description string `json:"description"`
    	Format      string `json:"format"`
    }

### UpdateSlate

#### Method Call

	response, err = broadpeak.UpdateSlate(Id, updateOptions)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> id</td> <td> uint</td>
<td>
Id of the slate source assigned by broadpeak.io
</td>
</tr>
<tr>
<td> updateOptions</td> <td> Slate</td>
<td>

```
{
	Name        string `json:"name"` //required
	Url         string `json:"url"`  //required
	Description string `json:"description"`
}
```

</td>
</tr>
</table>

#### Response

The response will be of the type `SlateOutput`

Type `SlateOutput` is defined as:

	{
    	Id          uint   `json:"id"`
    	Name        string `json:"name"`
    	Url         string `json:"url"`
    	Description string `json:"description"`
    	Format      string `json:"format"`
    }

### DeleteSlate

#### Method Call

	response, err = broadpeak.DeleteSlate(Id)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> id</td> <td> uint</td>
<td>
Id of the slate source assigned by broadpeak.io
</td>
</tr>
</table>

#### Response

The response will be of the type `SlateOutput`

Type `SlateOutput` is defined as:

	{
    	Id          uint   `json:"id"`
    	Name        string `json:"name"`
    	Url         string `json:"url"`
    	Description string `json:"description"`
    	Format      string `json:"format"`
    }

## Sources/AdServer

### CreateAdServer

#### Method Call

	response, err := broadpeak.CreateAdServer(options)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> options</td> <td> AdServer</td>
<td>

```
{
	Name        string `json:"name,omitempty"` //required
	Url         string `json:"url,omitempty"`  //required
	Description string `json:"description,omitempty"`
	Queries     string `json:"queries,omitempty"`
	Template    string `json:"template,omitempty"` //required
}
```

</td>
</tr>
</table>

#### Response

The response will be of the type `AdServerOutput`

Type `AdServerOutput` is defined as:

	{
    	Name        string `json:"name"`
    	Description string `json:"description"`
    	Url         string `json:"url"`
    	Queries     string `json:"queries"`
    	Template    string `json:"template"`
    	Id          uint   `json:"id"`
    }

### GetAdServer

#### Method Call

	response, err = broadpeak.GetAdServer(Id)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> id</td> <td> uint</td>
<td>
Id of the ad server source assigned by broadpeak.io
</td>
</tr>
</table>

#### Response

The response will be of the type `AdServerOutput`

Type `AdServerOutput` is defined as:

	{
    	Name        string `json:"name"`
    	Description string `json:"description"`
    	Url         string `json:"url"`
    	Queries     string `json:"queries"`
    	Template    string `json:"template"`
    	Id          uint   `json:"id"`
    }

### UpdateAdServer

#### Method Call

	response, err = broadpeak.UpdateAdServer(Id, updateOptions)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> id</td> <td> uint</td>
<td>
Id of the ad server source assigned by broadpeak.io
</td>
</tr>
<tr>
<td> updateOptions</td> <td> AdServer</td>
<td>

```
{
	Name        string `json:"name,omitempty"` //required
	Url         string `json:"url,omitempty"`  //required
	Description string `json:"description,omitempty"`
	Queries     string `json:"queries,omitempty"`
	Template    string `json:"template,omitempty"` //required
}
```

</td>
</tr>
</table>

#### Response

The response will be of the type `AdServerOutput`

Type `AdServerOutput` is defined as:

	{
    	Name        string `json:"name"`
    	Description string `json:"description"`
    	Url         string `json:"url"`
    	Queries     string `json:"queries"`
    	Template    string `json:"template"`
    	Id          uint   `json:"id"`
    }

### DeleteAdServer

#### Method Call

	response, err = broadpeak.DeleteAdServer(Id)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> id</td> <td> uint</td>
<td>
Id of the ad server source assigned by broadpeak.io
</td>
</tr>
</table>

#### Response

The response will be of the type `AdServerOutput`

Type `AdServerOutput` is defined as:

	{
    	Name        string `json:"name"`
    	Description string `json:"description"`
    	Url         string `json:"url"`
    	Queries     string `json:"queries"`
    	Template    string `json:"template"`
    	Id          uint   `json:"id"`
    }

## Services

### GetAllServices

#### Method Call

	response, err := broadpeak.GetAllServices(offset, limit)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> Offset</td> <td> uint</td>
<td>
Offset of the first result to return.
</td>
</tr>
<tr>
<td> Limit</td> <td> uint</td>
<td>
Maximum number of results to return
</td>
</tr>
</table>

#### Response

The response will be of the type `[]ServiceOutput`

Type `ServiceOutput` is defined as:

	{
        Name            string   `json:"name"`
        EnvironmentTags []string `json:"environmentTags"`
        Id              uint     `json:"id"`
        Type            []string `json:"type"`
        Url             string   `json:"url"`
        UpdateDate      string   `json:"updateDate"`
        CreationDate    string   `json:"creationDate"`
    }

## Sources/ContentReplacement

### CreateContentReplacement

#### Method Call

	response, err := broadpeak.CreateContentReplacement(options)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> options</td> <td> CreateContentReplacement</td>
<td>

```
{
	Name        string   `json:"name,omitempty"`        //required
	Replacement Identifiable `json:"replacement,omitempty"` //required
	Source      Identifiable `json:"source,omitempty"`      //required
	EnvTags     []string `json:"environmentTags,omitempty"`
}
```

</td>
</tr>
</table>

#### Response

The response will be of the type `ContentReplacementOutput`

Type `ContentReplacementOutput` is defined as:

	{
    	Name            string        `json:"name"`
    	EnvironmentTags []string      `json:"environmentTags"`
    	Replacement     IdNameUrlType `json:"replacement"`
    	Source          IdNameUrlType `json:"source"`
    	Id              uint          `json:"id"`
    	Url             string        `json:"url"`
    	UpdateDate      string        `json:"updateDate"`
    	CreationDate    string        `json:"creationDate"`
    }

### GetContentReplacement

#### Method Call

	response, err = broadpeak.GetContentReplacement(Id)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> id</td> <td> uint</td>
<td>
Id of the service
</td>
</tr>
</table>

#### Response

The response will be of the type `ContentReplacementOutput`

Type `ContentReplacementOutput` is defined as:

	{
    	Name            string        `json:"name"`
    	EnvironmentTags []string      `json:"environmentTags"`
    	Replacement     IdNameUrlType `json:"replacement"`
    	Source          IdNameUrlType `json:"source"`
    	Id              uint          `json:"id"`
    	Url             string        `json:"url"`
    	UpdateDate      string        `json:"updateDate"`
    	CreationDate    string        `json:"creationDate"`
    }

### UpdateContentReplacement

#### Method Call

	response, err = broadpeak.UpdateContentReplacement(Id, updateOptions)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> id</td> <td> uint</td>
<td>
Id of the service
</td>
</tr>
<tr>
<td> updateOptions</td> <td> UpdateContentReplacement</td>
<td>

```
{
	Name        string   `json:"name,omitempty"`        //required
	Replacement Identifiable `json:"replacement,omitempty"` //required
	EnvTags     []string `json:"environmentTags,omitempty"`
}
```

</td>
</tr>
</table>

#### Response

The response will be of the type `ContentReplacementOutput`

Type `ContentReplacementOutput` is defined as:

	{
    	Name            string        `json:"name"`
    	EnvironmentTags []string      `json:"environmentTags"`
    	Replacement     IdNameUrlType `json:"replacement"`
    	Source          IdNameUrlType `json:"source"`
    	Id              uint          `json:"id"`
    	Url             string        `json:"url"`
    	UpdateDate      string        `json:"updateDate"`
    	CreationDate    string        `json:"creationDate"`
    }

### DeleteContentReplacement

#### Method Call

	response, err = broadpeak.DeleteContentReplacement(Id)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> id</td> <td> uint</td>
<td>
Id of the service
</td>
</tr>
</table>

#### Response

The response will be of the type `ContentReplacementOutput`

Type `ContentReplacementOutput` is defined as:

	{
    	Name            string        `json:"name"`
    	EnvironmentTags []string      `json:"environmentTags"`
    	Replacement     IdNameUrlType `json:"replacement"`
    	Source          IdNameUrlType `json:"source"`
    	Id              uint          `json:"id"`
    	Url             string        `json:"url"`
    	UpdateDate      string        `json:"updateDate"`
    	CreationDate    string        `json:"creationDate"`
    }

## Sources/ContentReplacementSlot

### CreateContentReplacementSlot

#### Method Call

	response, err := broadpeak.CreateContentReplacementSlot(serviceId, options)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> serviceId</td> <td> uint</td>
<td>
Id of the service
</td>
</tr>
<tr>
<td> options</td> <td> CreateContentReplacementSlot</td>
<td>

```
{
	Name        string   `json:"name,omitempty"`
	StartTime   string   `json:"startTime,omitempty"` //required
	EndTime     string   `json:"endTime,omitempty"`
	Duration    uint     `json:"duration,omitempty"`
	Replacement Identifiable `json:"replacement,omitempty"`
}
```

</td>
</tr>
</table>

#### Response

The response will be of the type `ContentReplacementSlotOutput`

Type `ContentReplacementSlotOutput` is defined as:

	{
    	Name        string        `json:"name"`
    	StartTime   string        `json:"startTime"`
    	EndTime     string        `json:"endTime"`
    	Duration    uint          `json:"duration"`
    	Replacement IdNameUrlType `json:"replacement"`
    	Category    CategoryOutput     `json:"category"`
    	Id          uint          `json:"id"`
    }

### GetContentReplacementSlot

#### Method Call

	response, err = broadpeak.GetContentReplacementSlot(serviceId, Id)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> serviceId</td> <td> uint</td>
<td>
Id of the service
</td>
</tr>
<tr>
<td> id</td> <td> uint</td>
<td>
Id of the content replacement slot
</td>
</tr>
</table>

#### Response

The response will be of the type `ContentReplacementSlotOutput`

Type `ContentReplacementSlotOutput` is defined as:

	{
    	Name        string        `json:"name"`
    	StartTime   string        `json:"startTime"`
    	EndTime     string        `json:"endTime"`
    	Duration    uint          `json:"duration"`
    	Replacement IdNameUrlType `json:"replacement"`
    	Category    CategoryOutput     `json:"category"`
    	Id          uint          `json:"id"`
    }

### GetAllContentReplacementSlots

#### Method Call

	response, err = broadpeak.GetAllContentReplacementSlots(serviceId, getAllSlotsOptions)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> serviceId</td> <td> uint</td>
<td>
Id of the service
</td>
</tr>
<tr>
<td> getAllSlotsOptions</td> <td> GetAllSlots</td>
<td>

```
{
	Offset     uint
	Limit      uint
	Categories []uint
	From       uint
	To         uint
}
```

</td>
</tr>
</table>

#### Response

The response will be of the type `[]ContentReplacementSlotOutput`

Type `ContentReplacementSlotOutput` is defined as:

	{
    	Name        string        `json:"name"`
    	StartTime   string        `json:"startTime"`
    	EndTime     string        `json:"endTime"`
    	Duration    uint          `json:"duration"`
    	Replacement IdNameUrlType `json:"replacement"`
    	Category    CategoryOutput     `json:"category"`
    	Id          uint          `json:"id"`
    }

### UpdateContentReplacementSlot

#### Method Call

	response, err = broadpeak.UpdateContentReplacementSlot(serviceId, Id, updateOptions)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> serviceId</td> <td> uint</td>
<td>
Id of the service
</td>
</tr>
<tr>
<td> id</td> <td> uint</td>
<td>
Id of the content replacement slot
</td>
</tr>
<tr>
<td> updateOptions</td> <td> UpdateContentReplacementSlot</td>
<td>

```
{
	Name        string   `json:"name,omitempty"`
	StartTime   string   `json:"startTime,omitempty"` //required
	EndTime     string   `json:"endTime,omitempty"`
	Duration    uint     `json:"duration,omitempty"`
	Replacement Identifiable `json:"replacement,omitempty"`
	Category    Identifiable `json:"category,omitempty"`
}
```

</td>
</tr>
</table>

#### Response

The response will be of the type `ContentReplacementSlotOutput`

Type `ContentReplacementSlotOutput` is defined as:

	{
    	Name        string        `json:"name"`
    	StartTime   string        `json:"startTime"`
    	EndTime     string        `json:"endTime"`
    	Duration    uint          `json:"duration"`
    	Replacement IdNameUrlType `json:"replacement"`
    	Category    CategoryOutput     `json:"category"`
    	Id          uint          `json:"id"`
    }

### DeleteContentReplacementSlot

#### Method Call

	response, err = broadpeak.DeleteContentReplacementSlot(serviceId, Id)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> serviceId</td> <td> uint</td>
<td>
Id of the service
</td>
</tr>
<tr>
<td> id</td> <td> uint</td>
<td>
Id of the content replacement slot
</td>
</tr>
</table>

#### Response

The response will be of the type `ContentReplacementSlotOutput`

Type `ContentReplacementSlotOutput` is defined as:

	{
    	Name        string        `json:"name"`
    	StartTime   string        `json:"startTime"`
    	EndTime     string        `json:"endTime"`
    	Duration    uint          `json:"duration"`
    	Replacement IdNameUrlType `json:"replacement"`
    	Category    CategoryOutput     `json:"category"`
    	Id          uint          `json:"id"`
    }

## Sources/VirtualChannel

### CreateVirtualChannel

#### Method Call

	response, err := broadpeak.CreateVirtualChannel(options)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> options</td> <td> CreateVirtualChannel</td>
<td>

```
{
	Name                 string               `json:"name,omitempty"` //required
	EnvTags              []string             `json:"environmentTags,omitempty"`
	AdBreakInsertion     AdBreakInsertion     `json:"adBreakInsertion,omitempty"`
	TranscodingProfile   Identifiable             `json:"transcodingProfile,omitempty"`
	ServerSideAdTracking ServerSideAdTracking `json:"serverSideAdTracking,omitempty"`
	EnableAdTranscoding  bool                 `json:"enableAdTranscoding,omitempty"`
	BaseLive             Identifiable             `json:"baseLive,omitempty"` //required
}
```

</td>
</tr>
</table>

#### Response

The response will be of the type `VirtualChannelOutput`

Type `VirtualChannelOutput` is defined as:

	{
    	Name                 string               `json:"name"`
    	EnvironmentTags      []string             `json:"environmentTags"`
    	AdBreakInsertion     AdBreakInsertionOutput    `json:"adBreakInsertion"`
    	TranscodingProfiles  TranscodingProfileOutput `json:"transcodingProfile"`
    	ServerSideAdTracking ServerSideAdTracking `json:"serverSideAdTracking"`
    	EnableAdTranscoding  bool                 `json:"enableAdTranscoding"`
    	BaseLive             IdNameUrlType        `json:"baseLive"`
    	Id                   uint                 `json:"id"`
    	Url                  string               `json:"url"`
    	UpdateDate           string               `json:"updateDate"`
    	CreationDate         string               `json:"creationDate"`
    }

### GetVirtualChannel

#### Method Call

	response, err = broadpeak.GetVirtualChannel(Id)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> id</td> <td> uint</td>
<td>
Id of the virtual channel
</td>
</tr>
</table>

#### Response

The response will be of the type `VirtualChannelOutput`

Type `VirtualChannelOutput` is defined as:

	{
    	Name                 string               `json:"name"`
    	EnvironmentTags      []string             `json:"environmentTags"`
    	AdBreakInsertion     AdBreakInsertionOutput    `json:"adBreakInsertion"`
    	TranscodingProfiles  TranscodingProfileOutput `json:"transcodingProfile"`
    	ServerSideAdTracking ServerSideAdTracking `json:"serverSideAdTracking"`
    	EnableAdTranscoding  bool                 `json:"enableAdTranscoding"`
    	BaseLive             IdNameUrlType        `json:"baseLive"`
    	Id                   uint                 `json:"id"`
    	Url                  string               `json:"url"`
    	UpdateDate           string               `json:"updateDate"`
    	CreationDate         string               `json:"creationDate"`
    }

### UpdateVirtualChannel

#### Method Call

	response, err = broadpeak.UpdateVirtualChannel(Id, updateOptions)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> id</td> <td> uint</td>
<td>
Id of the virtual channel
</td>
</tr>
<tr>
<td> updateOptions</td> <td> UpdateVirtualChannel</td>
<td>

```
{
	Name                 string               `json:"name,omitempty"` //required
	EnvTags              []string             `json:"environmentTags,omitempty"`
	AdBreakInsertion     AdBreakInsertion     `json:"adBreakInsertion,omitempty"`
	TranscodingProfile   Identifiable             `json:"transcodingProfile,omitempty"`
	ServerSideAdTracking ServerSideAdTracking `json:"serverSideAdTracking,omitempty"`
	EnableAdTranscoding  bool                 `json:"enableAdTranscoding,omitempty"`
}
```

</td>
</tr>
</table>

#### Response

The response will be of the type `VirtualChannelOutput`

Type `VirtualChannelOutput` is defined as:

	{
    	Name                 string               `json:"name"`
    	EnvironmentTags      []string             `json:"environmentTags"`
    	AdBreakInsertion     AdBreakInsertionOutput    `json:"adBreakInsertion"`
    	TranscodingProfiles  TranscodingProfileOutput `json:"transcodingProfile"`
    	ServerSideAdTracking ServerSideAdTracking `json:"serverSideAdTracking"`
    	EnableAdTranscoding  bool                 `json:"enableAdTranscoding"`
    	BaseLive             IdNameUrlType        `json:"baseLive"`
    	Id                   uint                 `json:"id"`
    	Url                  string               `json:"url"`
    	UpdateDate           string               `json:"updateDate"`
    	CreationDate         string               `json:"creationDate"`
    }

### DeleteVirtualChannel

#### Method Call

	response, err = broadpeak.DeleteVirtualChannel(Id)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> id</td> <td> uint</td>
<td>
Id of the virtual channel
</td>
</tr>
</table>

#### Response

The response will be of the type `VirtualChannelOutput`

Type `VirtualChannelOutput` is defined as:

	{
    	Name                 string               `json:"name"`
    	EnvironmentTags      []string             `json:"environmentTags"`
    	AdBreakInsertion     AdBreakInsertionOutput    `json:"adBreakInsertion"`
    	TranscodingProfiles  TranscodingProfileOutput `json:"transcodingProfile"`
    	ServerSideAdTracking ServerSideAdTracking `json:"serverSideAdTracking"`
    	EnableAdTranscoding  bool                 `json:"enableAdTranscoding"`
    	BaseLive             IdNameUrlType        `json:"baseLive"`
    	Id                   uint                 `json:"id"`
    	Url                  string               `json:"url"`
    	UpdateDate           string               `json:"updateDate"`
    	CreationDate         string               `json:"creationDate"`
    }

## Sources/VirtualChannelSlot

### CreateVirtualChannelSlot

#### Method Call

	response, err := broadpeak.CreateVirtualChannelSlot(serviceId, options)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> serviceId</td> <td> uint</td>
<td>
Id of the service
</td>
</tr>
<tr>
<td> options</td> <td> CreateVirtualChannelSlot</td>
<td>

```
{
	Name        string   `json:"name,omitempty"`
	StartTime   string   `json:"startTime,omitempty"` //required
	EndTime     string   `json:"endTime,omitempty"`
	Duration    uint     `json:"duration,omitempty"`
	Replacement Identifiable `json:"replacement,omitempty"`
	Category    Identifiable `json:"category,omitempty"`
	Type        string   `json:"type,omitempty"`
}
```

</td>
</tr>
</table>

#### Response

The response will be of the type `VirtualChannelSlotOutput`

Type `VirtualChannelSlotOutput` is defined as:

	{
    	Name        string        `json:"name"`
    	StartTime   string        `json:"startTime"`
    	EndTime     string        `json:"endTime"`
    	Duration    uint          `json:"duration"`
    	Replacement IdNameUrlType `json:"replacement"`
    	Category    Identifiable      `json:"category"`
    	Type        string        `json:"type"`
    }

### GetVirtualChannelSlot

#### Method Call

	response, err = broadpeak.GetVirtualChannelSlot(serviceId, Id)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> serviceId</td> <td> uint</td>
<td>
Id of the service
</td>
</tr>
<tr>
<td> id</td> <td> uint</td>
<td>
Id of the virtual channel slot
</td>
</tr>
</table>

#### Response

The response will be of the type `VirtualChannelSlotOutput`

Type `VirtualChannelSlotOutput` is defined as:

	{
    	Name        string        `json:"name"`
    	StartTime   string        `json:"startTime"`
    	EndTime     string        `json:"endTime"`
    	Duration    uint          `json:"duration"`
    	Replacement IdNameUrlType `json:"replacement"`
    	Category    Identifiable      `json:"category"`
    	Type        string        `json:"type"`
    }

### GetAllVirtualChannelSlots

#### Method Call

	response, err = broadpeak.GetAllVirtualChannelSlots(serviceId, getAllSlotsOptions)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> serviceId</td> <td> uint</td>
<td>
Id of the service
</td>
</tr>
<tr>
<td> getAllSlotsOptions</td> <td> GetAllSlots</td>
<td>

```
{
	Offset     uint
	Limit      uint
	Categories []uint
	From       uint
	To         uint
}
```

</td>
</tr>
</table>

#### Response

The response will be of the type `[]VirtualChannelSlotOutput`

Type `VirtualChannelSlotOutput` is defined as:

	{
    	Name        string        `json:"name"`
    	StartTime   string        `json:"startTime"`
    	EndTime     string        `json:"endTime"`
    	Duration    uint          `json:"duration"`
    	Replacement IdNameUrlType `json:"replacement"`
    	Category    Identifiable      `json:"category"`
    	Type        string        `json:"type"`
    }

### UpdateVirtualChannelSlot

#### Method Call

	response, err = broadpeak.UpdateVirtualChannelSlot(serviceId, Id, updateOptions)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> serviceId</td> <td> uint</td>
<td>
Id of the service
</td>
</tr>
<tr>
<td> id</td> <td> uint</td>
<td>
Id of the virtual channel slot
</td>
</tr>
<tr>
<td> updateOptions</td> <td> UpdateVirtualChannelSlot</td>
<td>

```
{
	Name        string   `json:"name,omitempty"`
	StartTime   string   `json:"startTime,omitempty"` //required
	EndTime     string   `json:"endTime,omitempty"`
	Duration    uint     `json:"duration,omitempty"`
	Replacement Identifiable `json:"replacement,omitempty"`
	Category    Identifiable `json:"category,omitempty"`
	Type        string   `json:"type,omitempty"`
}
```

</td>
</tr>
</table>

#### Response

The response will be of the type `VirtualChannelSlotOutput`

Type `VirtualChannelSlotOutput` is defined as:

	{
    	Name        string        `json:"name"`
    	StartTime   string        `json:"startTime"`
    	EndTime     string        `json:"endTime"`
    	Duration    uint          `json:"duration"`
    	Replacement IdNameUrlType `json:"replacement"`
    	Category    Identifiable      `json:"category"`
    	Type        string        `json:"type"`
    }

### DeleteVirtualChannelSlot

#### Method Call

	response, err = broadpeak.DeleteVirtualChannelSlot(serviceId, Id)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> serviceId</td> <td> uint</td>
<td>
Id of the service
</td>
</tr>
<tr>
<td> id</td> <td> uint</td>
<td>
Id of the virtual channel slot
</td>
</tr>
</table>

#### Response

The response will be of the type `VirtualChannelSlotOutput`

Type `VirtualChannelSlotOutput` is defined as:

	{
    	Name        string        `json:"name"`
    	StartTime   string        `json:"startTime"`
    	EndTime     string        `json:"endTime"`
    	Duration    uint          `json:"duration"`
    	Replacement IdNameUrlType `json:"replacement"`
    	Category    Identifiable      `json:"category"`
    	Type        string        `json:"type"`
    }

## Sources/AdInsertion

### CreateAdInsertion

#### Method Call

	response, err := broadpeak.CreateAdInsertion(options)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> options</td> <td> CreateAdInsertion</td>
<td>

```
{
	Name                 string               `json:"name,omitempty"` //required
	EnvTags              []string             `json:"environmentTags,omitempty"`
	LiveAdPreRoll        LiveAdPreRoll        `json:"liveAdPreRoll,omitempty"`
	LiveAdReplacement    LiveAdReplacement    `json:"liveAdReplacement,omitempty"`
	VodAdInsertion       VodAdInsertion       `json:"vodAdInsertion,omitempty"`
	TranscodingProfile   Identifiable             `json:"transcodingProfile,omitempty"`
	EnableAdTranscoding  bool                 `json:"enableAdTranscoding,omitempty"`
	ServerSideAdTracking ServerSideAdTracking `json:"serverSideAdTracking,omitempty"`
	Source               Identifiable             `json:"source,omitempty"` //required
}
```

</td>
</tr>
</table>

#### Response

The response will be of the type `AdInsertionOutput`

Type `AdInsertionOutput` is defined as:

	{
    	Name                 string               `json:"name,omitempty"`
    	EnvironmentTags      []string             `json:"environmentTags,omitempty"`
    	LiveAdPreRoll        LiveAdPreRollOutput       `json:"liveAdPreRoll"`
    	LiveAdReplacement    LiveAdReplacementOutput   `json:"liveAdReplacement"`
    	VodAdInsertion       VodAdInsertionOutput      `json:"vodAdInsertion"`
    	TranscodingProfile   NameContentId        `json:"transcodingProfile"`
    	EnableAdTranscoding  bool                 `json:"enableAdTranscoding"`
    	ServerSideAdTracking ServerSideAdTracking `json:"serverSideAdTracking"`
    	Source               IdNameUrlQueries     `json:"source"`
    	Id                   uint                 `json:"id"`
    	CreationDate         string               `json:"creationDate"`
    	UpdateDate           string               `json:"updateDate"`
    	Url                  string               `json:"url"`
    }

### GetAdInsertion

#### Method Call

	response, err = broadpeak.GetAdInsertion(Id)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> id</td> <td> uint</td>
<td>
Id of the ad-insertion service
</td>
</tr>
</table>

#### Response

The response will be of the type `AdInsertionOutput`

Type `AdInsertionOutput` is defined as:

	{
    	Name                 string               `json:"name,omitempty"`
    	EnvironmentTags      []string             `json:"environmentTags,omitempty"`
    	LiveAdPreRoll        LiveAdPreRollOutput       `json:"liveAdPreRoll"`
    	LiveAdReplacement    LiveAdReplacementOutput   `json:"liveAdReplacement"`
    	VodAdInsertion       VodAdInsertionOutput      `json:"vodAdInsertion"`
    	TranscodingProfile   NameContentId        `json:"transcodingProfile"`
    	EnableAdTranscoding  bool                 `json:"enableAdTranscoding"`
    	ServerSideAdTracking ServerSideAdTracking `json:"serverSideAdTracking"`
    	Source               IdNameUrlQueries     `json:"source"`
    	Id                   uint                 `json:"id"`
    	CreationDate         string               `json:"creationDate"`
    	UpdateDate           string               `json:"updateDate"`
    	Url                  string               `json:"url"`
    }

### UpdateAdInsertion

#### Method Call

	response, err = broadpeak.UpdateAdInsertion(Id, updateOptions)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> id</td> <td> uint</td>
<td>
Id of the ad-insertion service
</td>
</tr>
<tr>
<td> updateOptions</td> <td> updateAdInsertion</td>
<td>

```
{
	Name                 string               `json:"name,omitempty"` //required
	EnvTags              []string             `json:"environmentTags,omitempty"`
	LiveAdPreRoll        LiveAdPreRoll        `json:"liveAdPreRoll,omitempty"`
	LiveAdReplacement    LiveAdReplacement    `json:"liveAdReplacement,omitempty"`
	VodAdInsertion       VodAdInsertion       `json:"vodAdInsertion,omitempty"`
	TranscodingProfile   Identifiable             `json:"transcodingProfile,omitempty"`
	EnableAdTranscoding  bool                 `json:"enableAdTranscoding,omitempty"`
	ServerSideAdTracking ServerSideAdTracking `json:"serverSideAdTracking,omitempty"`
}
```

</td>
</tr>
</table>

#### Response

The response will be of the type `AdInsertionOutput`

Type `AdInsertionOutput` is defined as:

	{
    	Name                 string               `json:"name,omitempty"`
    	EnvironmentTags      []string             `json:"environmentTags,omitempty"`
    	LiveAdPreRoll        LiveAdPreRollOutput       `json:"liveAdPreRoll"`
    	LiveAdReplacement    LiveAdReplacementOutput   `json:"liveAdReplacement"`
    	VodAdInsertion       VodAdInsertionOutput      `json:"vodAdInsertion"`
    	TranscodingProfile   NameContentId        `json:"transcodingProfile"`
    	EnableAdTranscoding  bool                 `json:"enableAdTranscoding"`
    	ServerSideAdTracking ServerSideAdTracking `json:"serverSideAdTracking"`
    	Source               IdNameUrlQueries     `json:"source"`
    	Id                   uint                 `json:"id"`
    	CreationDate         string               `json:"creationDate"`
    	UpdateDate           string               `json:"updateDate"`
    	Url                  string               `json:"url"`
    }

### DeleteAdInsertion

#### Method Call

	response, err = broadpeak.DeleteAdInsertion(Id)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> id</td> <td> uint</td>
<td>
Id of the ad-insertion service
</td>
</tr>
</table>

#### Response

The response will be of the type `AdInsertionOutput`

Type `AdInsertionOutput` is defined as:

	{
    	Name                 string               `json:"name,omitempty"`
    	EnvironmentTags      []string             `json:"environmentTags,omitempty"`
    	LiveAdPreRoll        LiveAdPreRollOutput       `json:"liveAdPreRoll"`
    	LiveAdReplacement    LiveAdReplacementOutput   `json:"liveAdReplacement"`
    	VodAdInsertion       VodAdInsertionOutput      `json:"vodAdInsertion"`
    	TranscodingProfile   NameContentId        `json:"transcodingProfile"`
    	EnableAdTranscoding  bool                 `json:"enableAdTranscoding"`
    	ServerSideAdTracking ServerSideAdTracking `json:"serverSideAdTracking"`
    	Source               IdNameUrlQueries     `json:"source"`
    	Id                   uint                 `json:"id"`
    	CreationDate         string               `json:"creationDate"`
    	UpdateDate           string               `json:"updateDate"`
    	Url                  string               `json:"url"`
    }

## TranscodingProfiles

### GetAllTranscodingProfiles

#### Method Call

	response, err := broadpeak.GetAllTranscodingProfiles(offset, limit)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td>Offset</td> <td> uint</td>
<td>
Offset of the first result to return.
</td>
</tr>
<tr>
<td> Limit</td> <td> uint</td>
<td>
Maximum number of results to return
</td>
</tr>
</table>

#### Response

The response will be of the type `[]TranscodingProfileOutput`

Type `TranscodingProfileOutput` is defined as:

	{
    	Name    string `json:"name"`
    	Content string `json:"content"`
    	Id      uint   `json:"id"`
    }

### GetTranscodingProfile

#### Method Call

	response, err = broadpeak.GetTranscodingProfile(Id)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> id</td> <td> uint</td>
<td>
Id of the transcoding profile
</td>
</tr>
</table>

#### Response

The response will be of the type `TranscodingProfileOutput`

Type `TranscodingProfileOutput` is defined as:

		{
    	Name    string `json:"name"`
    	Content string `json:"content"`
    	Id      uint   `json:"id"`
    }

## ESNI

### CreateUpdateMediaPoint

#### Method Call

	response, err := broadpeak.CreateUpdateMediaPoint(id)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> id</td> <td> string</td>
<td>
Media point id
</td>
</tr>
</table>

#### Response

The response will be of the type `string`

### GetMediaPoint

#### Method Call

	response, err = broadpeak.GetMediaPoint(Id)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> id</td> <td> string</td>
<td>
id of an existing MediaPoint
</td>
</tr>
</table>

## Categories

### CreateCategory

#### Method Call

	response, err := broadpeak.CreateCategory(name, subcategories)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> name</td> <td> string</td>
<td>
Name of the Category
</td>
</tr>
<tr>
<td> subcategories</td> <td> []Subcategory</td>
<td>
Array of type SubCategory:

```
{
	Key   string `json:"key"`
	Value string `json:"value"`
}
```

</td>
</tr>
</table>

#### Response

The response will be of the type `[]CategoryOutput`

Type `CategoryOutput` is defined as:

	{
    	Name        string        `json:"name"`
    	Subcategory []Subcategory `json:"subcategories"`
    	Id          uint          `json:"id"`
    }

### GetCategory

#### Method Call

	response, err = broadpeak.GetCategory(Id)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> id</td> <td> uint</td>
<td>
Id of the category
</td>
</tr>
</table>

#### Response

The response will be of the type `CategoryOutput`

Type `CategoryOutput` is defined as:

	{
    	Name        string        `json:"name"`
    	Subcategory []Subcategory `json:"subcategories"`
    	Id          uint          `json:"id"`
    }

### GetAllCategories

#### Method Call

response, err := broadpeak.GetAllCategories(offset, limit)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> Offset</td> <td> uint</td>
<td>
Offset of the first result to return.
</td>
</tr>
<tr>
<td> Limit</td> <td> uint</td>
<td>
Maximum number of results to return
</td>
</tr>
</table>

#### Response

The response will be of the type `[]CategoryOutput`

Type `CategoryOutput` is defined as:

	{
    	Name        string        `json:"name"`
    	Subcategory []Subcategory `json:"subcategories"`
    	Id          uint          `json:"id"`
    }

### UpdateCategory

#### Method Call

	response, err = broadpeak.UpdateCategory(id, name, subcategories)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> id</td> <td> uint</td>
<td>
Category id
</td>
</tr>
<tr>
<td> name</td> <td> string</td>
<td>
Category Name
</td>
</tr>
<tr>
<td> SubCategories</td> <td> []Subcategory</td>
<td>
Array of type SubCategory:

```
{
	Key   string `json:"key"`
	Value string `json:"value"`
}
```

</td>
</tr>
</table>

#### Response

The response will be of the type `CategoryOutput`

Type `CategoryOutput` is defined as:

	{
    	Name        string        `json:"name"`
    	Subcategory []Subcategory `json:"subcategories"`
    	Id          uint          `json:"id"`
    }

### DeleteCategory

#### Method Call

	response, err = broadpeak.DeleteCategory(Id)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> id</td> <td> uint</td>
<td>
Category id
</td>
</tr>
</table>

#### Response

The response will be of the type `string`

## Users

### GetUser

#### Method Call

	response, err = broadpeak.GetUser(Id)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> id</td> <td> uint</td>
<td>
Id of the user
</td>
</tr>
</table>

#### Response

The response will be of the type `UserOutput`

Type `UserOutput` is defined as:

	{
    	FirstName string `json:"firstName"`
    	LastName  string `json:"lastName"`
    	Email     string `json:"email"`
    	TenantId  string `json:"tenantId"`
    }

### GetAllUsers

#### Method Call

	response, err := broadpeak.GetAllUsers(offset, limit, withEmailStatus)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> Offset</td> <td> uint</td>
<td>
Offset of the first result to return.
</td>
</tr>
<tr>
<td> Limit</td> <td> uint</td>
<td>
Maximum number of results to return
</td>
</tr>
<tr>
<td> withEmailStatus</td> <td> bool</td>
<td>
Get email status flag
</td>
</tr>
</table>

#### Response

The response will be of the type `[]UserOutput`

Type `UserOutput` is defined as:

	{
    	FirstName string `json:"firstName"`
    	LastName  string `json:"lastName"`
    	Email     string `json:"email"`
    	TenantId  string `json:"tenantId"`
    }

### UpdateUser

#### Method Call

	response, err = broadpeak.UpdateUser(Id, updateOptions)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> id</td> <td> uint</td>
<td>
User id
</td>
</tr>
<tr>
<td> updateOptions</td> <td> User</td>
<td>

```
{
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email,omitempty"`
}
```

</td>
</tr>
</table>

#### Response

The response will be of the type `UserOutput`

Type `UserOutput` is defined as:

	{
    	FirstName string `json:"firstName"`
    	LastName  string `json:"lastName"`
    	Email     string `json:"email"`
    	TenantId  string `json:"tenantId"`
    }

### DeleteUser

#### Method Call

	response, err = broadpeak.DeleteUser(Id)

#### Parameters

<table>
<tr>
<td> Parameters </td> <td> Type </td> <td> Definition </td>
</tr>
<tr>
<td> id</td> <td> uint</td>
<td>
User id
</td>
</tr>
</table>

#### Response

The response will be of the type `string`

## ApplicationStatus

### CheckApplicationStatus

#### Method Call

	response, err = broadpeak.ApplicationStatus()

#### Parameters

The method call requires no parameters

#### Response

The response will be of the type `string`






