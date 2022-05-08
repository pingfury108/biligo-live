package live

import (
	"encoding/binary"
	"encoding/json"
)

// TODO msg注释移到struct上

type Transport struct {
	Msg   Msg
	Error error
}

type Msg interface {
	Cmd() string
	Raw() []byte
}
type base struct {
	raw []byte
}

func getData(raw []byte) json.RawMessage {
	var d struct {
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal(raw, &d); err != nil {
		return []byte{}
	}
	return d.Data
}

//

type MsgGeneral struct {
	base
}

func (m *MsgGeneral) Cmd() string {
	var cmd struct {
		CMD string `json:"cmd"`
	}
	if err := json.Unmarshal(m.raw, &cmd); err != nil {
		return ""
	}
	return cmd.CMD
}
func (m *MsgGeneral) Raw() []byte {
	return m.raw
}

//

type MsgHeartbeatReply struct {
	base
}

func (m *MsgHeartbeatReply) Cmd() string {
	return "HEARTBEAT_REPLY"
}
func (m *MsgHeartbeatReply) Raw() []byte {
	return m.raw
}
func (m *MsgHeartbeatReply) GetHot() int {
	return int(binary.BigEndian.Uint32(m.raw))
}

//

// MsgDanmaku 弹幕消息
type MsgDanmaku struct {
	base
}

func (m *MsgDanmaku) Cmd() string {
	return cmdDanmaku
}
func (m *MsgDanmaku) Raw() []byte {
	return m.raw
}

type Danmaku struct {
	SendMode     int    `json:"send_mode"`
	SendFontSize int    `json:"send_font_size"`
	DanmakuColor int64  `json:"danmaku_color"`
	Time         int64  `json:"time"`
	DMID         int64  `json:"dmid"`
	MsgType      int    `json:"msg_type"`
	Bubble       string `json:"bubble"`
	Content      string `json:"content"`
	MID          int64  `json:"mid"`
	Uname        string `json:"uname"`
	RoomAdmin    int    `json:"room_admin"`
	Vip          int    `json:"vip"`
	SVip         int    `json:"svip"`
	Rank         int    `json:"rank"`
	MobileVerify int    `json:"mobile_verify"`
	UnameColor   string `json:"uname_color"`
	MedalName    string `json:"medal_name"`
	UpName       string `json:"up_name"`
	MedalLevel   int    `json:"medal_level"`
	UserLevel    int    `json:"user_level"`
}

func (m *MsgDanmaku) Parse() (*Danmaku, error) {
	var t map[string]interface{}
	if err := json.Unmarshal(m.raw, &t); err != nil {
		return nil, err
	}
	info := t["info"]
	var dm = &Danmaku{}
	l := len(info.([]interface{}))
	if l >= 1 {
		h := info.([]interface{})[0].([]interface{})
		dm.SendMode = int(h[1].(float64))
		dm.SendFontSize = int(h[2].(float64))
		dm.DanmakuColor = int64(h[3].(float64))
		dm.Time = int64(h[4].(float64))
		dm.DMID = int64(h[5].(float64))
		dm.MsgType = int(h[10].(float64))
		dm.Bubble = h[11].(string)
	}
	if l >= 2 {
		dm.Content = info.([]interface{})[1].(string)
	}
	if l >= 3 {
		h := info.([]interface{})[2].([]interface{})
		dm.MID = int64(h[0].(float64))
		dm.Uname = h[1].(string)
		dm.RoomAdmin = int(h[2].(float64))
		dm.Vip = int(h[3].(float64))
		dm.SVip = int(h[4].(float64))
		dm.Rank = int(h[5].(float64))
		dm.MobileVerify = int(h[6].(float64))
		dm.UnameColor = h[7].(string)
	}
	if l >= 4 {
		h := info.([]interface{})[3].([]interface{})
		l2 := len(h)
		if l2 >= 1 {
			dm.MedalLevel = int(h[0].(float64))
		}
		if l2 >= 2 {
			dm.MedalName = h[1].(string)
		}
		if l2 >= 3 {
			dm.UpName = h[2].(string)
		}
	}
	if l >= 5 {
		dm.UserLevel = int(info.([]interface{})[4].([]interface{})[0].(float64))
	}
	return dm, nil
}

//

// MsgSendGift 投喂礼物
type MsgSendGift struct {
	base
}

func (m *MsgSendGift) Cmd() string {
	return cmdSendGift
}
func (m *MsgSendGift) Raw() []byte {
	return m.raw
}

type SendGift struct {
	Action         string `json:"action"`
	BatchComboID   string `json:"batch_combo_id"`
	BatchComboSend struct {
		Action        string      `json:"action"`
		BatchComboID  string      `json:"batch_combo_id"`
		BatchComboNum int         `json:"batch_combo_num"`
		BlindGift     interface{} `json:"blind_gift"`
		GiftID        int64       `json:"gift_id"`
		GiftName      string      `json:"gift_name"`
		GiftNum       int         `json:"gift_num"`
		SendMaster    interface{} `json:"send_master"`
		Uid           int         `json:"uid"`
		Uname         string      `json:"uname"`
	} `json:"batch_combo_send"`
	BeatID           string      `json:"beatId"`
	BizSource        string      `json:"biz_source"`
	BlindGift        interface{} `json:"blind_gift"`
	BroadcastID      int64       `json:"broadcast_id"`
	CoinType         string      `json:"coin_type"`
	ComboResourcesID int64       `json:"combo_resources_id"`
	ComboSend        struct {
		Action     string      `json:"action"`
		ComboID    string      `json:"combo_id"`
		ComboNum   int         `json:"combo_num"`
		GiftID     int64       `json:"gift_id"`
		GiftName   string      `json:"gift_name"`
		GiftNum    int         `json:"gift_num"`
		SendMaster interface{} `json:"send_master"`
		UID        int64       `json:"uid"`
		Uname      string      `json:"uname"`
	} `json:"combo_send"`
	ComboStayTime     int64   `json:"combo_stay_time"`
	ComboTotalCoin    int     `json:"combo_total_coin"`
	CritProb          int     `json:"crit_prob"`
	Demarcation       int     `json:"demarcation"`
	DiscountPrice     int     `json:"discount_price"`
	Dmscore           int     `json:"dmscore"`
	Draw              int     `json:"draw"`
	Effect            int     `json:"effect"`
	EffectBlock       int     `json:"effect_block"`
	Face              string  `json:"face"`
	FloatScResourceID int64   `json:"float_sc_resource_id"`
	GiftID            int64   `json:"giftId"`
	GiftName          string  `json:"giftName"`
	GiftType          int     `json:"giftType"`
	Gold              int     `json:"gold"`
	GuardLevel        int     `json:"guard_level"`
	IsFirst           bool    `json:"is_first"`
	IsSpecialBatch    int     `json:"is_special_batch"`
	Magnification     float64 `json:"magnification"`
	MedalInfo         struct {
		AnchorRoomid     int    `json:"anchor_roomid"`
		AnchorUname      string `json:"anchor_uname"`
		GuardLevel       int    `json:"guard_level"`
		IconID           int64  `json:"icon_id"`
		IsLighted        int    `json:"is_lighted"`
		MedalColor       int    `json:"medal_color"`
		MedalColorBorder int64  `json:"medal_color_border"`
		MedalColorEnd    int64  `json:"medal_color_end"`
		MedalColorStart  int64  `json:"medal_color_start"`
		MedalLevel       int    `json:"medal_level"`
		MedalName        string `json:"medal_name"`
		Special          string `json:"special"`
		TargetID         int    `json:"target_id"`
	} `json:"medal_info"`
	NameColor         string      `json:"name_color"`
	Num               int         `json:"num"`
	OriginalGiftName  string      `json:"original_gift_name"`
	Price             int         `json:"price"`
	Rcost             int         `json:"rcost"`
	Remain            int         `json:"remain"`
	Rnd               string      `json:"rnd"`
	SendMaster        interface{} `json:"send_master"`
	Silver            int         `json:"silver"`
	Super             int         `json:"super"`
	SuperBatchGiftNum int         `json:"super_batch_gift_num"`
	SuperGiftNum      int         `json:"super_gift_num"`
	SvgaBlock         int         `json:"svga_block"`
	TagImage          string      `json:"tag_image"`
	TID               string      `json:"tid"`
	Timestamp         int64       `json:"timestamp"`
	TopList           interface{} `json:"top_list"`
	TotalCoin         int         `json:"total_coin"`
	UID               int64       `json:"uid"`
	Uname             string      `json:"uname"`
}

func (m *MsgSendGift) Parse() (*SendGift, error) {
	var r = &SendGift{}
	if err := json.Unmarshal(getData(m.raw), &r); err != nil {
		return nil, err
	}
	return r, nil
}

//

// MsgComboSend 连击礼物
type MsgComboSend struct {
	base
}

func (m *MsgComboSend) Cmd() string {
	return cmdComboSend
}
func (m *MsgComboSend) Raw() []byte {
	return m.raw
}

//

// MsgFansUpdate 粉丝数量改变
type MsgFansUpdate struct {
	base
}

func (m *MsgFansUpdate) Cmd() string {
	return cmdRoomRealTimeMessageUpdate
}
func (m *MsgFansUpdate) Raw() []byte {
	return m.raw
}

type FansUpdate struct {
	// {
	// 	"fans_club": 49182,
	// 	"roomid": 545068,
	// 	"fans": 1384297,
	// 	"red_notice": -1
	// }
	FansClub  int   `json:"fans_club"`
	RoomID    int64 `json:"roomid"`
	Fans      int   `json:"fans"`
	RedNotice int   `json:"red_notice"`
}

func (m *MsgFansUpdate) Parse() (*FansUpdate, error) {
	var f = &FansUpdate{}
	if err := json.Unmarshal(getData(m.raw), &f); err != nil {
		return nil, err
	}
	return f, nil
}

//

// MsgOnlineRankCount 高能榜数量更新
type MsgOnlineRankCount struct {
	base
}

func (m *MsgOnlineRankCount) Cmd() string {
	return cmdOnlineRankCount
}
func (m *MsgOnlineRankCount) Raw() []byte {
	return m.raw
}
func (m *MsgOnlineRankCount) GetCount() int {
	var c struct {
		Count int `json:"count"`
	}
	if err := json.Unmarshal(getData(m.raw), &c); err != nil {
		return -1
	}
	return c.Count
}

//

// MsgSuperChatMessage 醒目留言
type MsgSuperChatMessage struct {
	base
}

func (m *MsgSuperChatMessage) Cmd() string {
	return cmdSuperChatMessage
}
func (m *MsgSuperChatMessage) Raw() []byte {
	return m.raw
}

type SuperChatMessage struct {
	BackgroundBottomColor string `json:"background_bottom_color"`
	Token                 string `json:"token"`
	BackgroundColorEnd    string `json:"background_color_end"`
	BackgroundImage       string `json:"background_image"`
	BackgroundIcon        string `json:"background_icon"`
	BackgroundPriceColor  string `json:"background_price_color"`
	DmScore               int    `json:"dmscore"`
	ID                    int64  `json:"id"`
	UserInfo              struct {
		UserLevel  int    `json:"user_level"`
		FaceFrame  string `json:"face_frame"`
		GuardLevel int    `json:"guard_level"`
		LevelColor string `json:"level_color"`
		Manager    int    `json:"manager"`
		Uname      string `json:"uname"`
		Title      string `json:"title"`
		Face       string `json:"face"`
		IsMainVip  int    `json:"is_main_vip"`
		IsSvip     int    `json:"is_svip"`
		IsVip      int    `json:"is_vip"`
		NameColor  string `json:"name_color"`
	} `json:"user_info"`
	IsSendAudit     int     `json:"is_send_audit"`
	Price           int     `json:"price"`
	BackgroundColor string  `json:"background_color"`
	ColorPoint      float64 `json:"color_point"`
	Gift            struct {
		GiftID   int64  `json:"gift_id"`
		GiftName string `json:"gift_name"`
		Num      int    `json:"num"`
	} `json:"gift"`
	MedalInfo struct {
		TargetID         int64  `json:"target_id"`
		AnchorRoomid     int    `json:"anchor_roomid"`
		AnchorUname      string `json:"anchor_uname"`
		GuardLevel       int    `json:"guard_level"`
		MedalColor       string `json:"medal_color"`
		MedalColorEnd    int    `json:"medal_color_end"`
		MedalLevel       int    `json:"medal_level"`
		Special          string `json:"special"`
		IconID           int64  `json:"icon_id"`
		IsLighted        int    `json:"is_lighted"`
		MedalColorBorder int    `json:"medal_color_border"`
		MedalColorStart  int    `json:"medal_color_start"`
		MedalName        string `json:"medal_name"`
	} `json:"medal_info"`
	TransMark            int    `json:"trans_mark"`
	Ts                   int    `json:"ts"`
	BackgroundColorStart string `json:"background_color_start"`
	EndTime              int64  `json:"end_time"`
	MessageFontColor     string `json:"message_font_color"`
	Rate                 int    `json:"rate"`
	MessageTrans         string `json:"message_trans"`
	StartTime            int64  `json:"start_time"`
	IsRanked             int    `json:"is_ranked"`
	Message              string `json:"message"`
	Time                 int64  `json:"time"`
	UID                  int64  `json:"uid"`
}

func (m *MsgSuperChatMessage) Parse() (*SuperChatMessage, error) {
	var r = &SuperChatMessage{}
	if err := json.Unmarshal(getData(m.raw), &r); err != nil {
		return nil, err
	}
	return r, nil
}

//

// MsgHotRankSettlement 荣登热门榜topX
type MsgHotRankSettlement struct {
	base
}

func (m *MsgHotRankSettlement) Cmd() string {
	return cmdHotRankSettlement
}
func (m *MsgHotRankSettlement) Raw() []byte {
	return m.raw
}

type HotRankSettlement struct {
	DmMsg     string `json:"dm_msg"`
	DmScore   int    `json:"dmscore"`
	Timestamp int64  `json:"timestamp"`
	Uname     string `json:"uname"`
	Url       string `json:"url"`
	AreaName  string `json:"area_name"`
	CacheKey  string `json:"cache_key"`
	Rank      int    `json:"rank"`
	Face      string `json:"face"`
	Icon      string `json:"icon"`
}

func (m *MsgHotRankSettlement) Parse() (*HotRankSettlement, error) {
	var r = &HotRankSettlement{}
	if err := json.Unmarshal(getData(m.raw), &r); err != nil {
		return nil, err
	}
	return r, nil
}

//

// MsgOnlineRankTop3 高能榜TOP3改变
type MsgOnlineRankTop3 struct {
	base
}

func (m *MsgOnlineRankTop3) Cmd() string {
	return cmdOnlineRankTop3
}
func (m *MsgOnlineRankTop3) Raw() []byte {
	return m.raw
}

type OnlineRankTop3 struct {
	DmScore int `json:"dmscore"`
	List    []struct {
		Msg  string `json:"msg"`
		Rank int    `json:"rank"`
	} `json:"list"`
}

func (m *MsgOnlineRankTop3) Parse() (*OnlineRankTop3, error) {
	var r = &OnlineRankTop3{}
	if err := json.Unmarshal(getData(m.raw), &r); err != nil {
		return nil, err
	}
	return r, nil
}

//

// MsgRoomBlockMsg 用户被禁言
type MsgRoomBlockMsg struct {
	base
}

func (m *MsgRoomBlockMsg) Cmd() string {
	return cmdRoomBlockMsg
}
func (m *MsgRoomBlockMsg) Raw() []byte {
	return m.raw
}

type RoomBlockMsg struct {
	Uname    string `json:"uname"`
	DmScore  int    `json:"dmscore"`
	Operator int    `json:"operator"`
	UID      int    `json:"uid"`
}

func (m *MsgRoomBlockMsg) Parse() (*RoomBlockMsg, error) {
	var r = &RoomBlockMsg{}
	if err := json.Unmarshal(getData(m.raw), &r); err != nil {
		return nil, err
	}
	return r, nil
}

//

// MsgStopLiveRoomList 刚刚停止了直播的直播间
type MsgStopLiveRoomList struct {
	base
}

func (m *MsgStopLiveRoomList) Cmd() string {
	return cmdStopLiveRoomList
}
func (m *MsgStopLiveRoomList) Raw() []byte {
	return m.raw
}

// GetList 返回停播直播间号数组
func (m *MsgStopLiveRoomList) GetList() ([]int64, error) {
	var r struct {
		RoomIDList []int64 `json:"room_id_list"`
	}
	if err := json.Unmarshal(getData(m.raw), &r); err != nil {
		return nil, err
	}
	return r.RoomIDList, nil
}

//

// MsgOnlineRankV2 高能榜数据
type MsgOnlineRankV2 struct {
	base
}

func (m *MsgOnlineRankV2) Cmd() string {
	return cmdOnlineRankV2
}
func (m *MsgOnlineRankV2) Raw() []byte {
	return m.raw
}

type OnlineRankV2 struct {
	List []struct {
		GuardLevel int    `json:"guard_level"` // 3:舰长 2:提督 1:总督?
		UID        int64  `json:"uid"`
		Face       string `json:"face"`
		Score      string `json:"score"`
		Uname      string `json:"uname"`
		Rank       int    `json:"rank"`
	} `json:"list"`
	RankType string `json:"rank_type"`
}

func (m *MsgOnlineRankV2) Parse() (*OnlineRankV2, error) {
	var r = &OnlineRankV2{}
	if err := json.Unmarshal(getData(m.raw), &r); err != nil {
		return nil, err
	}
	return r, nil
}

//

// MsgNoticeMsg 广播消息(别的直播间投递高价礼物对所有直播间发起的广播)
type MsgNoticeMsg struct {
	base
}

func (m *MsgNoticeMsg) Cmd() string {
	return cmdNoticeMsg
}
func (m *MsgNoticeMsg) Raw() []byte {
	return m.raw
}

type NoticeMsg struct {
	BusinessID string `json:"business_id"`
	Full       struct {
		HeadIcon    string `json:"head_icon"`
		TailIcon    string `json:"tail_icon"`
		HeadIconFa  string `json:"head_icon_fa"`
		TailIconFa  string `json:"tail_icon_fa"`
		Background  string `json:"background"`
		Highlight   string `json:"highlight"`
		HeadIconFan int    `json:"head_icon_fan"`
		TailIconFan int    `json:"tail_icon_fan"`
		Color       string `json:"color"`
		Time        int64  `json:"time"`
	} `json:"full"`
	Half struct {
		Time       int64  `json:"time"`
		HeadIcon   string `json:"head_icon"`
		TailIcon   string `json:"tail_icon"`
		Background string `json:"background"`
		Color      string `json:"color"`
		Highlight  string `json:"highlight"`
	} `json:"half"`
	ID         int64  `json:"id"`
	LinkUrl    string `json:"link_url"`
	MsgCommon  string `json:"msg_common"`
	MsgSelf    string `json:"msg_self"`
	MsgType    int    `json:"msg_type"`
	Name       string `json:"name"`
	RealRoomID int64  `json:"real_roomid"`
	RoomID     int64  `json:"roomid"`
	Scatter    struct {
		Min int `json:"min"`
		Max int `json:"max"`
	} `json:"scatter"`
	ShieldUID int64 `json:"shield_uid"`
	Side      struct {
		HeadIcon   string `json:"head_icon"`
		Background string `json:"background"`
		Color      string `json:"color"`
		Highlight  string `json:"highlight"`
		Border     string `json:"border"`
	} `json:"side"`
}

func (m *MsgNoticeMsg) Parse() (*NoticeMsg, error) {
	var r = &NoticeMsg{}
	if err := json.Unmarshal(m.raw, &r); err != nil {
		return nil, err
	}
	return r, nil
}

//

// MsgHotRankChanged 热门榜改变
type MsgHotRankChanged struct {
	base
}

func (m *MsgHotRankChanged) Cmd() string {
	return cmdHotRankChanged
}
func (m *MsgHotRankChanged) Raw() []byte {
	return m.raw
}

type HotRankChanged struct {
	Rank        int    `json:"rank"`
	Timestamp   int64  `json:"timestamp"`
	WebUrl      string `json:"web_url"`
	LiveUrl     string `json:"live_url"`
	LiveLinkUrl string `json:"live_link_url"`
	AreaName    string `json:"area_name"`
	Trend       int    `json:"trend"`
	Countdown   int    `json:"countdown"`
	BlinkUrl    string `json:"blink_url"`
	PCLinkUrl   string `json:"pc_link_url"`
	Icon        string `json:"icon"`
}

func (m *MsgHotRankChanged) Parse() (*HotRankChanged, error) {
	var r = &HotRankChanged{}
	if err := json.Unmarshal(getData(m.raw), &r); err != nil {
		return nil, err
	}
	return r, nil
}

//

// MsgGuardBuy 用户上舰长
type MsgGuardBuy struct {
	base
}

func (m *MsgGuardBuy) Cmd() string {
	return cmdGuardBuy
}
func (m *MsgGuardBuy) Raw() []byte {
	return m.raw
}

type GuardBuy struct {
	GuardLevel int    `json:"guard_level"`
	Price      int    `json:"price"`
	UID        int64  `json:"uid"`
	Num        int    `json:"num"`
	GiftID     int64  `json:"gift_id"`
	GiftName   string `json:"gift_name"`
	StartTime  int64  `json:"start_time"`
	EndTime    int64  `json:"end_time"`
	Username   string `json:"username"`
}

func (m *MsgGuardBuy) Parse() (*GuardBuy, error) {
	var r = &GuardBuy{}
	if err := json.Unmarshal(getData(m.raw), &r); err != nil {
		return nil, err
	}
	return r, nil
}

//

// MsgSuperChatMessageJPN 醒目留言日文翻译?
type MsgSuperChatMessageJPN struct {
	base
}

func (m *MsgSuperChatMessageJPN) Cmd() string {
	return cmdSuperChatMessageJPN
}
func (m *MsgSuperChatMessageJPN) Raw() []byte {
	return m.raw
}

type SuperChatMessageJPN struct {
	UID       string `json:"uid"`
	IsRanked  int    `json:"is_ranked"`
	MedalInfo struct {
		MedalColor   string `json:"medal_color"`
		IconID       int64  `json:"icon_id"`
		TargetID     int64  `json:"target_id"`
		Special      string `json:"special"`
		AnchorUname  string `json:"anchor_uname"`
		AnchorRoomid int    `json:"anchor_roomid"`
		MedalLevel   int    `json:"medal_level"`
		MedalName    string `json:"medal_name"`
	} `json:"medal_info"`
	UserInfo struct {
		UserLevel  int    `json:"user_level"`
		LevelColor string `json:"level_color"`
		IsVip      int    `json:"is_vip"`
		IsSvip     int    `json:"is_svip"`
		IsMainVip  int    `json:"is_main_vip"`
		Title      string `json:"title"`
		Uname      string `json:"uname"`
		Face       string `json:"face"`
		Manager    int    `json:"manager"`
		FaceFrame  string `json:"face_frame"`
		GuardLevel int    `json:"guard_level"`
	} `json:"user_info"`
	ID                   string `json:"id"`
	MessageJpn           string `json:"message_jpn"`
	Time                 int64  `json:"time"`
	Rate                 int    `json:"rate"`
	BackgroundImage      string `json:"background_image"`
	BackgroundIcon       string `json:"background_icon"`
	BackgroundPriceColor string `json:"background_price_color"`
	Token                string `json:"token"`
	Gift                 struct {
		Num      int    `json:"num"`
		GiftID   int64  `json:"gift_id"`
		GiftName string `json:"gift_name"`
	} `json:"gift"`
	Price                 int    `json:"price"`
	Message               string `json:"message"`
	BackgroundColor       string `json:"background_color"`
	BackgroundBottomColor string `json:"background_bottom_color"`
	TS                    int64  `json:"ts"`
	StartTime             int64  `json:"start_time"`
	EndTime               int64  `json:"end_time"`
}

func (m *MsgSuperChatMessageJPN) Parse() (*SuperChatMessageJPN, error) {
	var r = &SuperChatMessageJPN{}
	if err := json.Unmarshal(getData(m.raw), &r); err != nil {
		return nil, err
	}
	return r, nil
}

//

// MsgUserToastMsg 上船附带的通知
type MsgUserToastMsg struct {
	base
}

func (m *MsgUserToastMsg) Cmd() string {
	return cmdUserToastMsg
}
func (m *MsgUserToastMsg) Raw() []byte {
	return m.raw
}

type UserToastMsg struct {
	GuardLevel       int    `json:"guard_level"`
	OpType           int    `json:"op_type"`
	PayflowID        string `json:"payflow_id"`
	Unit             string `json:"unit"`
	IsShow           int    `json:"is_show"`
	Num              int    `json:"num"`
	Price            int64  `json:"price"`
	StartTime        int64  `json:"start_time"`
	SvgaBlock        int    `json:"svga_block"`
	UserShow         bool   `json:"user_show"`
	Color            string `json:"color"`
	EndTime          int64  `json:"end_time"`
	RoleName         string `json:"role_name"`
	ToastMsg         string `json:"toast_msg"`
	UID              int64  `json:"uid"`
	AnchorShow       bool   `json:"anchor_show"`
	DmScore          int    `json:"dmscore"`
	TargetGuardCount int    `json:"target_guard_count"`
	Username         string `json:"username"`
}

func (m *MsgUserToastMsg) Parse() (*UserToastMsg, error) {
	var r = &UserToastMsg{}
	if err := json.Unmarshal(getData(m.raw), &r); err != nil {
		return nil, err
	}
	return r, nil
}

//

// MsgSuperChatMessageDelete 删除醒目留言 (似乎有时候并不会发,同时结束时间在 CmdSuperChatMessage 可以获取)
type MsgSuperChatMessageDelete struct {
	base
}

func (m *MsgSuperChatMessageDelete) Cmd() string {
	return cmdSuperChatMessageDelete
}
func (m *MsgSuperChatMessageDelete) Raw() []byte {
	return m.raw
}

// GetList 返回id数组
func (m *MsgSuperChatMessageDelete) GetList() ([]int64, error) {
	var r struct {
		IDS []int64 `json:"ids"`
	}
	if err := json.Unmarshal(getData(m.raw), &r); err != nil {
		return nil, err
	}
	return r.IDS, nil
}

//

// MsgAnchorLotStart 天选之人开始完整信息
type MsgAnchorLotStart struct {
	base
}

func (m *MsgAnchorLotStart) Cmd() string {
	return cmdAnchorLotStart
}
func (m *MsgAnchorLotStart) Raw() []byte {
	return m.raw
}

type AnchorLotStart struct {
	MaxTime        int    `json:"max_time"`
	Danmu          string `json:"danmu"`
	GiftNum        int    `json:"gift_num"`
	JoinType       int    `json:"join_type"`
	AwardImage     string `json:"award_image"`
	GiftPrice      int    `json:"gift_price"`
	GiftID         int64  `json:"gift_id"`
	GiftName       string `json:"gift_name"`
	GoodsID        int64  `json:"goods_id"`
	RoomID         int64  `json:"room_id"`
	Time           int64  `json:"time"`
	Url            string `json:"url"`
	CurGiftNum     int    `json:"cur_gift_num"`
	CurrentTime    int64  `json:"current_time"`
	LotStatus      int    `json:"lot_status"`
	RequireType    int    `json:"require_type"`
	WebUrl         string `json:"web_url"`
	GoawayTime     int    `json:"goaway_time"`
	IsBroadcast    int    `json:"is_broadcast"`
	RequireValue   int    `json:"require_value"`
	ShowPanel      int    `json:"show_panel"`
	Status         int    `json:"status"`
	ID             int64  `json:"id"`
	RequireText    string `json:"require_text"`
	AwardNum       int    `json:"award_num"`
	AssetIcon      string `json:"asset_icon"`
	AwardName      string `json:"award_name"`
	SendGiftEnsure int    `json:"send_gift_ensure"`
}

func (m *MsgAnchorLotStart) Parse() (*AnchorLotStart, error) {
	var r = &AnchorLotStart{}
	if err := json.Unmarshal(getData(m.raw), &r); err != nil {
		return nil, err
	}
	return r, nil
}

//

// MsgAnchorLotCheckStatus 天选时刻前的审核
type MsgAnchorLotCheckStatus struct {
	base
}

func (m *MsgAnchorLotCheckStatus) Cmd() string {
	return cmdAnchorLotCheckStatus
}
func (m *MsgAnchorLotCheckStatus) Raw() []byte {
	return m.raw
}

/*
	{
	  "id": 1890708,
	  "reject_reason": "",
	  "status": 4,
	  "uid": 2920960
	}
*/
type AnchorLotCheckStatus struct {
	ID           int64  `json:"id"`
	RejectReason string `json:"reject_reason"`
	Status       int    `json:"status"`
	Uid          int64  `json:"uid"`
}

func (m *MsgAnchorLotCheckStatus) Parse() (*AnchorLotCheckStatus, error) {
	var r = &AnchorLotCheckStatus{}
	if err := json.Unmarshal(getData(m.raw), &r); err != nil {
		return nil, err
	}
	return r, nil
}

//

// MsgAnchorLotAward 天选结果推送
type MsgAnchorLotAward struct {
	base
}

func (m *MsgAnchorLotAward) Cmd() string {
	return cmdAnchorLotAward
}
func (m *MsgAnchorLotAward) Raw() []byte {
	return m.raw
}

type AnchorLotAward struct {
	LotStatus  int    `json:"lot_status"`
	Url        string `json:"url"`
	WebUrl     string `json:"web_url"`
	AwardImage string `json:"award_image"`
	AwardName  string `json:"award_name"`
	AwardNum   int    `json:"award_num"`
	AwardUsers []struct {
		Uname string `json:"uname"`
		Face  string `json:"face"`
		Level int    `json:"level"`
		Color int64  `json:"color"`
		UID   int64  `json:"uid"`
	} `json:"award_users"`
	ID int64 `json:"id"`
}

func (m *MsgAnchorLotAward) Parse() (*AnchorLotAward, error) {
	var r = &AnchorLotAward{}
	if err := json.Unmarshal(getData(m.raw), &r); err != nil {
		return nil, err
	}
	return r, nil
}

// MsgAnchorLotEnd 天选之人获奖id
type MsgAnchorLotEnd struct {
	base
}

func (m *MsgAnchorLotEnd) Cmd() string {
	return cmdAnchorLotEnd
}
func (m *MsgAnchorLotEnd) Raw() []byte {
	return m.raw
}
func (m *MsgAnchorLotEnd) GetID() int64 {
	var r struct {
		ID int64 `json:"id"`
	}
	if err := json.Unmarshal(getData(m.raw), &r); err != nil {
		return -1
	}
	return r.ID
}

//

// MsgRoomChange 房间信息改变
type MsgRoomChange struct {
	base
}

func (m *MsgRoomChange) Cmd() string {
	return cmdRoomChange
}
func (m *MsgRoomChange) Raw() []byte {
	return m.raw
}

type RoomChange struct {
	ParentAreaID   int    `json:"parent_area_id"`
	AreaName       string `json:"area_name"`
	ParentAreaName string `json:"parent_area_name"`
	LiveKey        string `json:"live_key"`
	SubSessionKey  string `json:"sub_session_key"`
	Title          string `json:"title"`
	AreaID         int    `json:"area_id"`
}

func (m *MsgRoomChange) Parse() (*RoomChange, error) {
	var r = &RoomChange{}
	if err := json.Unmarshal(getData(m.raw), &r); err != nil {
		return nil, err
	}
	return r, nil
}

//

// MsgVoiceJoinList 连麦申请、取消连麦申请
type MsgVoiceJoinList struct {
	base
}

func (m *MsgVoiceJoinList) Cmd() string {
	return cmdVoiceJoinList
}
func (m *MsgVoiceJoinList) Raw() []byte {
	return m.raw
}

type VoiceJoinList struct {
	RoomID     int64 `json:"room_id"`
	Category   int   `json:"category"`
	ApplyCount int   `json:"apply_count"`
	RedPoint   int   `json:"red_point"`
	Refresh    int   `json:"refresh"`
}

func (m *MsgVoiceJoinList) Parse() (*VoiceJoinList, error) {
	var r = &VoiceJoinList{}
	if err := json.Unmarshal(getData(m.raw), &r); err != nil {
		return nil, err
	}
	return r, nil
}

//

// MsgVoiceJoinRoomCountInfo 申请连麦队列变化
type MsgVoiceJoinRoomCountInfo struct {
	base
}

func (m *MsgVoiceJoinRoomCountInfo) Cmd() string {
	return cmdVoiceJoinRoomCountInfo
}
func (m *MsgVoiceJoinRoomCountInfo) Raw() []byte {
	return m.raw
}

type VoiceJoinRoomCountInfo struct {
	ApplyCount  int   `json:"apply_count"`
	NotifyCount int   `json:"notify_count"`
	RedPoint    int   `json:"red_point"`
	RoomID      int64 `json:"room_id"`
	RootStatus  int   `json:"root_status"`
	RoomStatus  int   `json:"room_status"`
}

func (m *MsgVoiceJoinRoomCountInfo) Parse() (*VoiceJoinRoomCountInfo, error) {
	var r = &VoiceJoinRoomCountInfo{}
	if err := json.Unmarshal(getData(m.raw), &r); err != nil {
		return nil, err
	}
	return r, nil
}

//

// MsgAttention 用户关注
type MsgAttention struct {
	base
}

func (m *MsgAttention) Cmd() string {
	return cmdAttention
}
func (m *MsgAttention) Raw() []byte {
	return m.raw
}

//

// MsgShare 用户分享直播间
type MsgShare struct {
	base
}

func (m *MsgShare) Cmd() string {
	return cmdShare
}
func (m *MsgShare) Raw() []byte {
	return m.raw
}

//

// MsgSpecialAttention 特别关注直播间
type MsgSpecialAttention struct {
	base
}

func (m *MsgSpecialAttention) Cmd() string {
	return cmdSpecialAttention
}
func (m *MsgSpecialAttention) Raw() []byte {
	return m.raw
}

//

type MsgSysMsg struct {
	base
}

func (m *MsgSysMsg) Cmd() string {
	return cmdSysMsg
}
func (m *MsgSysMsg) Raw() []byte {
	return m.raw
}

//

// MsgPreparing 下播
type MsgPreparing struct {
	base
}

func (m *MsgPreparing) Cmd() string {
	return cmdPreparing
}
func (m *MsgPreparing) Raw() []byte {
	return m.raw
}

//

// MsgLive 开播
type MsgLive struct {
	base
}

func (m *MsgLive) Cmd() string {
	return cmdLive
}
func (m *MsgLive) Raw() []byte {
	return m.raw
}

//

// MsgRoomRank 排名改变
type MsgRoomRank struct {
	base
}

func (m *MsgRoomRank) Cmd() string {
	return cmdRoomRank
}
func (m *MsgRoomRank) Raw() []byte {
	return m.raw
}

//

type MsgRoomLimit struct {
	base
}

func (m *MsgRoomLimit) Cmd() string {
	return cmdRoomLimit
}
func (m *MsgRoomLimit) Raw() []byte {
	return m.raw
}

//

type MsgBlock struct {
	base
}

func (m *MsgBlock) Cmd() string {
	return cmdBlock
}
func (m *MsgBlock) Raw() []byte {
	return m.raw
}

//

type MsgPkPre struct {
	base
}

func (m *MsgPkPre) Cmd() string {
	return cmdPkPre
}
func (m *MsgPkPre) Raw() []byte {
	return m.raw
}

//

// MsgPkEnd PK判断胜负
type MsgPkEnd struct {
	base
}

func (m *MsgPkEnd) Cmd() string {
	return cmdPkEnd
}
func (m *MsgPkEnd) Raw() []byte {
	return m.raw
}

//

type MsgPkSettle struct {
	base
}

func (m *MsgPkSettle) Cmd() string {
	return cmdPkSettle
}
func (m *MsgPkSettle) Raw() []byte {
	return m.raw
}

//

type MsgSysGift struct {
	base
}

func (m *MsgSysGift) Cmd() string {
	return cmdSysGift
}
func (m *MsgSysGift) Raw() []byte {
	return m.raw
}

//

// MsgHotRank 热门榜xx榜topX
type MsgHotRank struct {
	base
}

func (m *MsgHotRank) Cmd() string {
	return cmdHotRank
}
func (m *MsgHotRank) Raw() []byte {
	return m.raw
}

//

type MsgActivityRedPacket struct {
	base
}

func (m *MsgActivityRedPacket) Cmd() string {
	return cmdActivityRedPacket
}
func (m *MsgActivityRedPacket) Raw() []byte {
	return m.raw
}

//

type MsgPkMicEnd struct {
	base
}

func (m *MsgPkMicEnd) Cmd() string {
	return cmdPkMicEnd
}
func (m *MsgPkMicEnd) Raw() []byte {
	return m.raw
}

//

type MsgPlayTag struct {
	base
}

func (m *MsgPlayTag) Cmd() string {
	return cmdPlayTag
}
func (m *MsgPlayTag) Raw() []byte {
	return m.raw
}

//

// MsgGuardMsg 舰长消息
type MsgGuardMsg struct {
	base
}

func (m *MsgGuardMsg) Cmd() string {
	return cmdGuardMsg
}
func (m *MsgGuardMsg) Raw() []byte {
	return m.raw
}

//

type MsgPlayProgressBar struct {
	base
}

func (m *MsgPlayProgressBar) Cmd() string {
	return cmdPlayProgressBar
}
func (m *MsgPlayProgressBar) Raw() []byte {
	return m.raw
}

//

type MsgHotRoomNotify struct {
	base
}

func (m *MsgHotRoomNotify) Cmd() string {
	return cmdHotRoomNotify
}
func (m *MsgHotRoomNotify) Raw() []byte {
	return m.raw
}

//

type MsgRefresh struct {
	base
}

func (m *MsgRefresh) Cmd() string {
	return cmdRefresh
}
func (m *MsgRefresh) Raw() []byte {
	return m.raw
}

//

type MsgRound struct {
	base
}

func (m *MsgRound) Cmd() string {
	return cmdRound
}
func (m *MsgRound) Raw() []byte {
	return m.raw
}

//

type MsgWelcomeGuard struct {
	base
}

func (m *MsgWelcomeGuard) Cmd() string {
	return cmdWelcomeGuard
}
func (m *MsgWelcomeGuard) Raw() []byte {
	return m.raw
}

//

// MsgEntryEffect 舰长、高能榜、老爷进入
type MsgEntryEffect struct {
	base
}

func (m *MsgEntryEffect) Cmd() string {
	return cmdEntryEffect
}
func (m *MsgEntryEffect) Raw() []byte {
	return m.raw
}

//

// MsgWelcome 欢迎进入房间(似乎已废弃)
type MsgWelcome struct {
	base
}

func (m *MsgWelcome) Cmd() string {
	return cmdWelcome
}
func (m *MsgWelcome) Raw() []byte {
	return m.raw
}

//

type MsgLiveInteractiveGame struct {
	base
}

func (m *MsgLiveInteractiveGame) Cmd() string {
	return cmdLiveInteractiveGame
}
func (m *MsgLiveInteractiveGame) Raw() []byte {
	return m.raw
}

//

// MsgVoiceJoinStatus 开始连麦、结束连麦
type MsgVoiceJoinStatus struct {
	base
}

func (m *MsgVoiceJoinStatus) Cmd() string {
	return cmdVoiceJoinStatus
}
func (m *MsgVoiceJoinStatus) Raw() []byte {
	return m.raw
}

//

// MsgCutOff 被超管切断
type MsgCutOff struct {
	base
}

func (m *MsgCutOff) Cmd() string {
	return cmdCutOff
}
func (m *MsgCutOff) Raw() []byte {
	return m.raw
}

//

// MsgSpecialGift 节奏风暴
type MsgSpecialGift struct {
	base
}

func (m *MsgSpecialGift) Cmd() string {
	return cmdSpecialGift
}
func (m *MsgSpecialGift) Raw() []byte {
	return m.raw
}

//

// MsgNewGuardCount 船员数量改变事件
type MsgNewGuardCount struct {
	base
}

func (m *MsgNewGuardCount) Cmd() string {
	return cmdNewGuardCount
}
func (m *MsgNewGuardCount) Raw() []byte {
	return m.raw
}

//

// MsgRoomAdmins 房管数量改变
type MsgRoomAdmins struct {
	base
}

func (m *MsgRoomAdmins) Cmd() string {
	return cmdRoomAdmins
}
func (m *MsgRoomAdmins) Raw() []byte {
	return m.raw
}

//

type MsgActivityBannerUpdateV2 struct {
	base
}

func (m *MsgActivityBannerUpdateV2) Cmd() string {
	return cmdActivityBannerUpdateV2
}
func (m *MsgActivityBannerUpdateV2) Raw() []byte {
	return m.raw
}

//

// MsgInteractWord 用户进入直播间
type MsgInteractWord struct {
	base
}

func (m *MsgInteractWord) Cmd() string {
	return cmdInteractWord
}
func (m *MsgInteractWord) Raw() []byte {
	return m.raw
}

type InteractWord struct {
	TailIcon     int    `json:"tail_icon"`
	UID          int64  `json:"uid"`
	Uname        string `json:"uname"`
	UnameColor   string `json:"uname_color"`
	Dmscore      int    `json:"dmscore"`
	Score        int64  `json:"score"`
	SpreadDesc   string `json:"spread_desc"`
	Timestamp    int64  `json:"timestamp"`
	Identities   []int  `json:"identities"`
	IsSpread     int    `json:"is_spread"`
	Roomid       int    `json:"roomid"`
	TriggerTime  int64  `json:"trigger_time"`
	Contribution struct {
		Grade int `json:"grade"`
	} `json:"contribution"`
	FansMedal struct {
		MedalColor       int64  `json:"medal_color"`
		MedalColorStart  int64  `json:"medal_color_start"`
		MedalLevel       int    `json:"medal_level"`
		Score            int    `json:"score"`
		TargetId         int    `json:"target_id"`
		GuardLevel       int    `json:"guard_level"`
		IconId           int    `json:"icon_id"`
		IsLighted        int    `json:"is_lighted"`
		MedalName        string `json:"medal_name"`
		Special          string `json:"special"`
		AnchorRoomID     int64  `json:"anchor_roomid"`
		MedalColorBorder int64  `json:"medal_color_border"`
		MedalColorEnd    int64  `json:"medal_color_end"`
	} `json:"fans_medal"`
	MsgType    int    `json:"msg_type"`
	SpreadInfo string `json:"spread_info"`
}

func (m *MsgInteractWord) Parse() (*InteractWord, error) {
	var r = &InteractWord{}
	if err := json.Unmarshal(getData(m.raw), &r); err != nil {
		return nil, err
	}
	return r, nil
}

//

// MsgPkBattlePre 大乱斗准备，10秒后开始
type MsgPkBattlePre struct {
	base
}

func (m *MsgPkBattlePre) Cmd() string {
	return cmdPkBattlePre
}
func (m *MsgPkBattlePre) Raw() []byte {
	return m.raw
}

//

type MsgPkBattleSettle struct {
	base
}

func (m *MsgPkBattleSettle) Cmd() string {
	return cmdPkBattleSettle
}
func (m *MsgPkBattleSettle) Raw() []byte {
	return m.raw
}

//

// MsgPkBattleStart 大乱斗开始
type MsgPkBattleStart struct {
	base
}

func (m *MsgPkBattleStart) Cmd() string {
	return cmdPkBattleStart
}
func (m *MsgPkBattleStart) Raw() []byte {
	return m.raw
}

//

// MsgPkBattleProcess 大乱斗双方送礼
type MsgPkBattleProcess struct {
	base
}

func (m *MsgPkBattleProcess) Cmd() string {
	return cmdPkBattleProcess
}
func (m *MsgPkBattleProcess) Raw() []byte {
	return m.raw
}

//

// MsgPkEnding 大乱斗尾声，最后几秒
type MsgPkEnding struct {
	base
}

func (m *MsgPkEnding) Cmd() string {
	return cmdPkEnding
}
func (m *MsgPkEnding) Raw() []byte {
	return m.raw
}

//

// MsgPkBattleEnd 大乱斗结束
type MsgPkBattleEnd struct {
	base
}

func (m *MsgPkBattleEnd) Cmd() string {
	return cmdPkBattleEnd
}
func (m *MsgPkBattleEnd) Raw() []byte {
	return m.raw
}

//

type MsgPkBattleSettleUser struct {
	base
}

func (m *MsgPkBattleSettleUser) Cmd() string {
	return cmdPkBattleSettleUser
}
func (m *MsgPkBattleSettleUser) Raw() []byte {
	return m.raw
}

//

type MsgPkBattleSettleV2 struct {
	base
}

func (m *MsgPkBattleSettleV2) Cmd() string {
	return cmdPkBattleSettleV2
}
func (m *MsgPkBattleSettleV2) Raw() []byte {
	return m.raw
}

//

// MsgPkLotteryStart 大乱斗胜利后的抽奖
type MsgPkLotteryStart struct {
	base
}

func (m *MsgPkLotteryStart) Cmd() string {
	return cmdPkLotteryStart
}
func (m *MsgPkLotteryStart) Raw() []byte {
	return m.raw
}

//

// MsgPkBestUname PK最佳助攻
type MsgPkBestUname struct {
	base
}

func (m *MsgPkBestUname) Cmd() string {
	return cmdPkBestUname
}
func (m *MsgPkBestUname) Raw() []byte {
	return m.raw
}

//

// MsgCallOnOpposite 本直播间的观众跑去对面串门
type MsgCallOnOpposite struct {
	base
}

func (m *MsgCallOnOpposite) Cmd() string {
	return cmdCallOnOpposite
}
func (m *MsgCallOnOpposite) Raw() []byte {
	return m.raw
}

//

// MsgAttentionOpposite 本直播间观众关注了对面主播
type MsgAttentionOpposite struct {
	base
}

func (m *MsgAttentionOpposite) Cmd() string {
	return cmdAttentionOpposite
}
func (m *MsgAttentionOpposite) Raw() []byte {
	return m.raw
}

//

// MsgShareOpposite 本直播间观众分享了对面直播间
type MsgShareOpposite struct {
	base
}

func (m *MsgShareOpposite) Cmd() string {
	return cmdShareOpposite
}
func (m *MsgShareOpposite) Raw() []byte {
	return m.raw
}

//

// MsgAttentionOnOpposite 对面观众关注了本直播间
type MsgAttentionOnOpposite struct {
	base
}

func (m *MsgAttentionOnOpposite) Cmd() string {
	return cmdAttentionOnOpposite
}
func (m *MsgAttentionOnOpposite) Raw() []byte {
	return m.raw
}

//

// MsgPkMatchInfo 获取对面直播间信息
type MsgPkMatchInfo struct {
	base
}

func (m *MsgPkMatchInfo) Cmd() string {
	return cmdPkMatchInfo
}
func (m *MsgPkMatchInfo) Raw() []byte {
	return m.raw
}

//

// MsgPkMatchOnlineGuard 获取对面直播间舰长在线人数
type MsgPkMatchOnlineGuard struct {
	base
}

func (m *MsgPkMatchOnlineGuard) Cmd() string {
	return cmdPkMatchOnlineGuard
}
func (m *MsgPkMatchOnlineGuard) Raw() []byte {
	return m.raw
}

//

// MsgPkWinningStreak 大乱斗连胜事件
type MsgPkWinningStreak struct {
	base
}

func (m *MsgPkWinningStreak) Cmd() string {
	return cmdPkWinningStreak
}
func (m *MsgPkWinningStreak) Raw() []byte {
	return m.raw
}

//

// MsgPkDanmuMsg 对面的弹幕消息
type MsgPkDanmuMsg struct {
	base
}

func (m *MsgPkDanmuMsg) Cmd() string {
	return cmdPkDanmuMsg
}
func (m *MsgPkDanmuMsg) Raw() []byte {
	return m.raw
}

//

// MsgPkSendGift 对面的礼物消息
type MsgPkSendGift struct {
	base
}

func (m *MsgPkSendGift) Cmd() string {
	return cmdPkSendGift
}
func (m *MsgPkSendGift) Raw() []byte {
	return m.raw
}

//

// MsgPkInteractWord 对面的用户进入
type MsgPkInteractWord struct {
	base
}

func (m *MsgPkInteractWord) Cmd() string {
	return cmdPkInteractWord
}
func (m *MsgPkInteractWord) Raw() []byte {
	return m.raw
}

//

// MsgPkAttention 对面新增关注
type MsgPkAttention struct {
	base
}

func (m *MsgPkAttention) Cmd() string {
	return cmdPkAttention
}
func (m *MsgPkAttention) Raw() []byte {
	return m.raw
}

//

// MsgPkShare 对面有人分享直播间
type MsgPkShare struct {
	base
}

func (m *MsgPkShare) Cmd() string {
	return cmdPkShare
}
func (m *MsgPkShare) Raw() []byte {
	return m.raw
}

type WatChed struct {
	Num       int    `json:"num"`        //144450
	TextLarge string `json:"text_large"` //14.4万人看过
	TextSmall string `json:"text_small"` //14.4万
}

// MsgWatChed 直播间看过人数变化
type MsgWatChed struct {
	base
}

func (m *MsgWatChed) Cmd() string {
	return cmdWatChedChange
}
func (m *MsgWatChed) Raw() []byte {
	return m.raw
}

func (m *MsgWatChed) Parse() (*WatChed, error) {
	var watChed WatChed
	err := json.Unmarshal(getData(m.raw), &watChed)
	if err != nil {
		return nil, err
	}
	return &watChed, nil
}
