package xerror

var (
	Success = NewError(0, "成功")

	// 鉴权错误码
	UnauthorizedAuthNotExist  = NewError(100003, "鉴权失败，找不到对应的 AppKey 和 AppSecret")
	UnauthorizedTokenError    = NewError(100004, "鉴权失败，Token 错误")
	UnauthorizedTokenTimeout  = NewError(100001, "鉴权失败，Token 过期")
	UnauthorizedTokenGenerate = NewError(100006, "鉴权失败，Token 生成失败")
	// accesskey 错误
	AccessKeyError = NewError(100005, "accesskey 错误")
	// 签名错误
	SignatureError = NewError(100007, "签名错误")

	// 权限错误码
	PermissionNotAllow = NewError(100010, "权限不足")

	// 服务级错误码
	ServerError        = NewError(100101, "Internal Server Error")
	TooManyRequests    = NewError(100102, "Too Many Requests")
	ParamBindError     = NewError(100103, "参数信息有误")
	AuthorizationError = NewError(100104, "签名信息有误")
	CallHTTPError      = NewError(100105, "调用第三方 HTTP 接口失败")
	InvalidRequest     = NewError(100106, "Invalid request")
	DataNotFound       = NewError(100107, "")
	DataExists         = NewError(100108, "")

	CreateError        = NewError(100109, "创建错误")
	UpdateError        = NewError(100110, "更新错误")
	GetError           = NewError(100111, "查询错误")
	DeleteError        = NewError(100112, "删除错误")
	FromError          = NewError(100113, "表单错误")
	TransformError     = NewError(100114, "数据转换错误")
	ListError          = NewError(100115, "获取列表错误")
	DeviceInfoGetError = NewError(100116, "没有权限获取设备信息")
	// 缺少appkey
	AppKeyEmpty = NewError(100121, "缺少appkey")
	// 缺少 appsecret
	AppSecretEmpty = NewError(100122, "缺少 appsecret")
	// 缺失address
	AddressEmpty = NewError(100123, "缺少address")
	// 缺少endpoint
	EndpointEmpty = NewError(100124, "缺少endpoint")
	// 缺少sign
	SignEmpty = NewError(100125, "缺少sign")
	// 缺失sender
	SenderEmpty = NewError(100126, "缺少sender")
	// 缺失region
	RegionEmpty = NewError(100127, "缺少region")

	// 模块级错误码 - 用户模块
	IllegalUserName    = NewError(200101, "非法用户名")
	IllegalUserNameLen = NewError(200102, "Your name should be between 3 to 20 characters")
	UserCreateError    = NewError(200103, "创建用户失败")
	UserUpdateError    = NewError(200104, "更新用户失败")
	UserSearchError    = NewError(200105, "查询用户失败")
	UserNotLogin       = NewError(200106, "user not login")
	UserNotFound       = NewError(200107, "user not found")
	UserLoginError     = NewError(200108, "")
	EmailError         = NewError(200109, "Please enter a valid email")
	PhoneError         = NewError(200110, "")
	PhoneEmpty         = NewError(200111, "")
	PhoneLenghtError   = NewError(200112, "")
	PasswordError      = NewError(200113, "Your password should be between 3 and 50 characters")
	SMSCodeError       = NewError(200114, "sms code error")
	SMSHasSendFailed   = NewError(200115, "短信验证码已经发送")
	// 邮箱验证码已发送
	EmailCaptchaIsSendError = NewError(200116, "邮箱验证码已发送,请检查邮箱")
	// 邮箱验证码发送失败
	EmailCaptchaSendFailed = NewError(200117, "邮箱验证码发送失败")
	// 邮箱验证码不能为空
	EmailCaptchaEmpty = NewError(200118, "邮箱验证码不能为空")
	// 手机验证码不能为空
	PhoneCaptchaEmpty    = NewError(200119, "手机验证码不能为空")
	EmailCaptchaError    = NewError(200120, "邮箱验证码错误")
	PhoneCaptchaError    = NewError(200121, "手机验证码错误")
	TicketTimeoutError   = NewError(200122, "凭证过期")
	TickerinvalidError   = NewError(200123, "凭证无效")
	TicketNotverifyError = NewError(200124, "未扫码")
	PhoneSameError       = NewError(200125, "不能使用重复手机号")

	ConnectError     = NewError(101017, "Connect Error")
	DisConnectError  = NewError(101018, "DisConnect Error")
	ConnectOverError = NewError(101019, "The number of connections exceeded the limit")
	ConnectHasSend   = NewError(101020, "The connection request has been sent, please wait")
)
