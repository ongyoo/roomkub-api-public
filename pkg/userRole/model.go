package userrole

import (
	mongo "github.com/ongyoo/roomkub-api/pkg/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRole struct {
	ID             primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Name           string               `json:"name" bson:"name"`
	PermissionID   []primitive.ObjectID `json:"permission_id" bson:"permissionId"`
	Slug           mongo.Encrypted      `json:"slug" bson:"slug"`
	SlugIdentityID string               `json:"slug_identity_id" bson:"slugIdentityId"`
	IsActive       bool                 `json:"is_active" bson:"isActive"`
}

// Request
type UserRoleRequest struct {
	UserRole `bson:",inline"`
}

type GetUserRoleListRequest struct {
	Page  int64  `uri:"page"`
	Limit int64  `uri:"limit"`
	Slug  string `uri:"slug"`
}

//Occupation         Occupation                          `json:"occupation" bson:"occupation" example:"COMPANY_OFFICER"  binding:"required,occupation" enums:"GOVERNMENT_OFFICER,STATE_ENTERPRISE_OFFICER,COMPANY_OFFICER,BUSINESS_OWNER,HOUSEWIFE,STUDENT,SELF_EMPLOYEE,FARMER,TEMPORARY_OFFICER,MONK,RETIREE"`
/*
type LogResult struct {
	ID   string `bson:"_id"`
	*Log `bson:",inline"`
}
*/

/*
type Occupation mongo.Encrypted

const (
	OccupationGovernmentOfficer      Occupation = "GOVERNMENT_OFFICER"
	OccupationStateEnterpriseOfficer Occupation = "STATE_ENTERPRISE_OFFICER"
	OccupationCompanyOfficer         Occupation = "COMPANY_OFFICER"
	OccupationBusinessOwner          Occupation = "BUSINESS_OWNER"
	OccupationHousewife              Occupation = "HOUSEWIFE"
	OccupationStudent                Occupation = "STUDENT"
	OccupationSelfEmployee           Occupation = "SELF_EMPLOYEE"
	OccupationFarmer                 Occupation = "FARMER"
	OccupationTemporaryOfficer       Occupation = "TEMPORARY_OFFICER"
	OccupationMonk                   Occupation = "MONK"
	OccupationRetiree                Occupation = "RETIREE"
)

func (o Occupation) GetKbankOccuCd() string {
	switch o {
	case OccupationGovernmentOfficer:
		return "01"
	case OccupationStateEnterpriseOfficer:
		return "02"
	case OccupationCompanyOfficer:
		return "03"
	case OccupationBusinessOwner:
		return "04"
	case OccupationHousewife:
		return "05"
	case OccupationStudent:
		return "06"
	case OccupationSelfEmployee:
		return "07"
	case OccupationFarmer:
		return "08"
	case OccupationTemporaryOfficer:
		return "09"
	case OccupationMonk:
		return "10"
	case OccupationRetiree:
		return "11"
	}
	return ""
}
*/
