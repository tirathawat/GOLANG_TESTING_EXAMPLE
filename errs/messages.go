package errs

const (
	INTERNAL_SERVER_ERROR_TH = "เกิดข้อผิดพลาด"
	INTERNAL_SERVER_ERROR_EN = "Internal server error"

	SERVICE_UNAVAILABLE_ERROR_TH = "บริการไม่พร้อมใช้งาน"
	SERVICE_UNAVAILABLE_ERROR_EN = "Service unavailable"

	BAD_REQUEST_ERROR_TH = "คำร้องขอไม่ถูกต้อง"
	BAD_REQUEST_ERROR_EN = "Bad request"

	NOT_FOUND_ERROR_TH = "ไม่พบข้อมูล"
	NOT_FOUND_ERROR_EN = "Not found"
)

const (
	INSERT_ERROR_TH = "เกิดข้อผิดพลาดในการบันทึกข้อมูล"
	INSERT_ERROR_EN = "Saving data error"

	GET_ERROR_TH = "เกิดข้อผิดพลาดในการโหลดข้อมูล"
	GET_ERROR_EN = "Get data error"

	UPDATE_ERROR_TH = "เกิดข้อผิดพลาดในการอัพเดตข้อมูล"
	UPDATE_ERROR_EN = "Update data error"

	DELETE_ERROR_TH = "เกิดข้อผิดพลาดในการลบข้อมูล"
	DELETE_ERROR_EN = "Delete data error"
)

const (
	PURCHASE_ZERO_AMOUNT_TH = "ค่าใช้จ่ายไม่สามารถน้อยกว่าหรือเท่ากับศูนย์ได้"
	PURCHASE_ZERO_AMOUNT_EN = "Purchase amount could not be zero"

	PROMO_EXPIRED_TH = "โปรโมชันหมดอายุ"
	PROMO_EXPIRED_EN = "promotion has expired"

	PROMO_ID_NOT_FOUND_TH = "ไม่พบรหัสโปรชันในคำร้องขอ"
	PROMO_ID_NOT_FOUND_EN = "Promotion id not found"

	PURCHASE_NOT_FOUND_TH = "ไม่พบค่าใช้จ่ายในคำร้องขอ"
	PURCHASE_NOT_FOUND_EN = "Purchase amount not found"
)

var (
	ErrInternalServerError = NewInternalServerError(
		INTERNAL_SERVER_ERROR_TH,
		INTERNAL_SERVER_ERROR_EN,
	)
	ErrServiceUnavailableError = NewServiceUnavailableError(
		SERVICE_UNAVAILABLE_ERROR_TH,
		SERVICE_UNAVAILABLE_ERROR_EN,
	)
	ErrBadRequestError = NewBadRequestError(
		BAD_REQUEST_ERROR_TH,
		BAD_REQUEST_ERROR_EN,
	)
	ErrNotFoundError = NewNotFoundError(
		NOT_FOUND_ERROR_TH,
		NOT_FOUND_ERROR_EN,
	)
)

var (
	ErrInsertError = NewInternalServerError(INSERT_ERROR_TH, INSERT_ERROR_EN)
	ErrGetError    = NewInternalServerError(GET_ERROR_TH, GET_ERROR_EN)
	ErrUpdateError = NewInternalServerError(UPDATE_ERROR_TH, UPDATE_ERROR_EN)
	ErrDeleteError = NewInternalServerError(DELETE_ERROR_TH, DELETE_ERROR_EN)
)

var (
	ErrPurchaseZero     = NewBadRequestError(PURCHASE_ZERO_AMOUNT_TH, PURCHASE_ZERO_AMOUNT_EN)
	ErrPromoExpired     = NewBadRequestError(PROMO_EXPIRED_TH, PROMO_EXPIRED_EN)
	ErrPromoIDNotFound  = NewBadRequestError(PROMO_ID_NOT_FOUND_TH, PROMO_ID_NOT_FOUND_EN)
	ErrPurchaseNotFound = NewBadRequestError(PURCHASE_NOT_FOUND_TH, PURCHASE_NOT_FOUND_EN)
)
