package clicksend

import "errors"

// SMS is exactly what it sounds like
type SMS struct {
	// To - REQUIRED Recipient phone number in E164 format
	To string `json:"to"`
	// To - REQUIRED Text message
	Body string `json:"body"`
}

type SMSResponse struct {
	HTTPCode     int                   `json:"http_code"`
	ResponseCode ClickSendResponseCode `json:"response_code"`
	ResponseMsg  string                `json:"response_msg"`
	Data         struct {
		TotalPrice  float64 `json:"total_price"`
		TotalCount  int     `json:"total_count"`
		QueuedCount int     `json:"queued_count"`
		Messages    []struct {
			Direction string `json:"direction"`
			Date      int64  `json:"date"`
			To        string `json:"to"`
			Body      string `json:"body"`
			From      string `json:"from"`
			// Schedule     int64       `json:"schedule"`
			MessageID    string      `json:"message_id"`
			MessageParts int         `json:"message_parts"`
			MessagePrice string      `json:"message_price"`
			FromEmail    interface{} `json:"from_email"`
			ListID       interface{} `json:"list_id"`
			CustomString string      `json:"custom_string"`
			ContactID    interface{} `json:"contact_id"`
			UserID       int         `json:"user_id"`
			SubaccountID int         `json:"subaccount_id"`
			Country      string      `json:"country"`
			Carrier      string      `json:"carrier"`
			Status       string      `json:"status"`
		} `json:"messages"`
		Currency struct {
			CurrencyNameShort string `json:"currency_name_short"`
			CurrencyPrefixD   string `json:"currency_prefix_d"`
			CurrencyPrefixC   string `json:"currency_prefix_c"`
			CurrencyNameLong  string `json:"currency_name_long"`
		} `json:"_currency"`
	} `json:"data"`
}

type ClickSendResponseCode string

const (
	ClickSendResponseCodeSuccess                     ClickSendResponseCode = "SUCCESS"
	ClickSendResponseCodeMissingCredentials                                = "MISSING_CREDENTIALS"
	ClickSendResponseCodeAccountNotActivated                               = "ACCOUNT_NOT_ACTIVATED"
	ClickSendResponseCodeInvalidRecipient                                  = "INVALID_RECIPIENT"
	ClickSendResponseCodeThrottled                                         = "THROTTLED"
	ClickSendResponseCodeInvalidSenderID                                   = "INVALID_SENDER_ID"
	ClickSendResponseCodeInsufficientCredit                                = "INSUFFICIENT_CREDIT"
	ClickSendResponseCodeInvalidCredentials                                = "INVALID_CREDENTIALS"
	ClickSendResponseCodeAlreadyExists                                     = "ALREADY_EXISTS"
	ClickSendResponseCodeEmptyMessage                                      = "EMPTY_MESSAGE"
	ClickSendResponseCodeTooManyRecipients                                 = "TOO_MANY_RECIPIENTS"
	ClickSendResponseCodeMissingRequiredFields                             = "MISSING_REQUIRED_FIELDS"
	ClickSendResponseCodeInvalidSchedule                                   = "INVALID_SCHEDULE"
	ClickSendResponseCodeNotEnoughPermissionToListID                       = "NOT_ENOUGH_PERMISSION_TO_LIST_ID"
	ClickSendResponseCodeInternalError                                     = "INTERNAL_ERROR"
	ClickSendResponseCodeInvalidLang                                       = "INVALID_LANG"
	ClickSendResponseCodeInvalidVoice                                      = "INVALID_VOICE"
	ClickSendResponseCodeSubjectRequired                                   = "SUBJECT_REQUIRED"
	ClickSendResponseCodeInvalidMediaFile                                  = "INVALID_MEDIA_FILE"
	ClickSendResponseCodeSomethingIsWrong                                  = "SOMETHING_IS_WRONG"
)

func (c *Client) SendSMS(s *SMS) (*SMSResponse, error) {
	res := &SMSResponse{}

	if s == nil {
		return res, errors.New("The sms object is not set")
	}

	err := c.doRequest(parameters{
		Method:  "POST",
		Path:    "email/batch",
		Payload: s,
	}, &res)

	return res, err
}
