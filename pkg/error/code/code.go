package code

type BadRequest string

const (
	BadRequestPipelineOrMerchantNotFound BadRequest = "PIPELINE_OR_MERCHANT_NOT_FOUND"
)

type Internal string

type Conflict string

const (
	ConflictOngoingPipelineOfShopIdExists Conflict = "ON_GOING_PIPELINE_OF_SHOP_ID_EXIST"
	ConflictOngoingPipelineOfNidExists             = "ON_GOING_PIPELINE_OF_NID_EXIST"
	ConflictOngoingPipelineOfTaxIdExists           = "ON_GOING_PIPELINE_OF_TAX_ID_EXIST"
)

type Default string

const (
	DefaultInternalServerError Default = "6991"
)
