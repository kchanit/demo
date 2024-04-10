package errorresponse

import (
	"errors"
	"net/http"

	"github.com/centraldigital/cfw-core-lib/pkg/model/errormodel"
)

var (
	ErrDuplicateCustomerEmail = errormodel.CustomError{
		Status:      http.StatusConflict,
		Code:        "5001",
		Description: "Customer email is already existed",
	}

	ErrDuplicateCampaignId = errormodel.CustomError{
		Status:      http.StatusConflict,
		Code:        "5001",
		Description: "Campaign id is already existed",
	}

	ErrDuplicatedBannerRanking = errormodel.CustomError{
		Status:      http.StatusConflict,
		Code:        "5001",
		Description: "Ranking number is already existed",
	}

	ErrDuplicatedMiniBannerRanking = errormodel.CustomError{
		Status:      http.StatusConflict,
		Code:        "5001",
		Description: "Ranking number is already existed",
	}
	ErrDuplicatedWidgetName = errormodel.CustomError{
		Status:      http.StatusConflict,
		Code:        "5001",
		Description: "Widget name is already existed",
	}

	ErrUnableToLogout = errormodel.TechnicalError{
		Code:        "9203",
		Description: "unable to logout",
	}

	ErrGetCustomer = errormodel.BusinessError{
		Code:        "9202",
		Description: "unable to get customer detail",
	}

	ErrUpdateCoupon = errormodel.BusinessError{
		Code:        "9202",
		Description: "unable to update coupon detail",
	}

	ErrUpdateCouponQuota = errormodel.BusinessError{
		Code:        "9202",
		Description: "unable to update coupon quota",
	}

	ErrGetCouponSortConfig = errormodel.BusinessError{
		Code:        "9202",
		Description: "unable to get coupon sort config",
	}

	ErrUpdateCouponSortConfig = errormodel.BusinessError{
		Code:        "9202",
		Description: "unable to update coupon sort config",
	}

	ErrSendDateOutsidePeriod = errormodel.BusinessError{
		Code:        "9202",
		Description: "The send date falls outside the allowed period. Please input it again.",
	}

	ErrGetCouponSortingField = errormodel.BusinessError{
		Code:        "9202",
		Description: "unable to get coupon sort option [sorting_field]",
	}

	ErrGetCouponSortingSubTypeList = errormodel.BusinessError{
		Code:        "9202",
		Description: "unable to get coupon sort option [coupon_sub_type]",
	}

	ErrDeleteCampaign = errormodel.BusinessError{
		Code:        "9202",
		Description: "unalbe to delete campaign [landing page]",
	}

	ErrGetCampaignFromAdapter = errormodel.TechnicalError{
		Code:        "9203",
		Description: "unable to get campaing data",
	}

	ErrGetCampaignBannerFromAdapter = errormodel.TechnicalError{
		Code:        "9203",
		Description: "unable to get campaing-banner",
	}
	ErrGetCouponStatusList = errormodel.TechnicalError{
		Code:        "9204",
		Description: "unable to get coupon status list",
	}
	ErrGetCampaignStatusList = errormodel.TechnicalError{
		Code:        "9204",
		Description: "unable to get campaign status list",
	}

	ErrGetIconList = &errormodel.TechnicalError{
		Code:        "9204",
		Description: "unable to get icon wedget list",
	}

	ErrGetIconById = &errormodel.TechnicalError{
		Code:        "9204",
		Description: "unable to get icon wedget by id",
	}

	ErrUpdateIcon = &errormodel.TechnicalError{
		Code:        "9204",
		Description: "unable to update icon data",
	}

	ErrDeleteIcon = &errormodel.TechnicalError{
		Code:        "9204",
		Description: "unable to delete icon data",
	}

	ErrUploadImageRequireBase64Url = errormodel.BadRequestError(errors.New("require: base64 image url"))
	ErrUploadIconImage             = &errormodel.TechnicalError{
		Code:        "9204",
		Description: "unable to upload icon image",
	}
)
