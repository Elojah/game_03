package dto

import (
	"encoding/base64"
	"encoding/json"

	"github.com/elojah/game_03/pkg/rtc"
	"github.com/pion/webrtc/v3"
)

func (s SDP) MarshalRTC() (rtc.SDP, error) {
	var result webrtc.SessionDescription

	b, err := base64.StdEncoding.DecodeString(s.Encoded)
	if err != nil {
		return rtc.SDP{}, err
	}

	if err := json.Unmarshal(b, &result); err != nil {
		return rtc.SDP{}, err
	}

	return rtc.SDP(result), nil
}

func (s *SDP) UnmarshalRTC(sdp rtc.SDP) error {
	raw, err := json.Marshal(sdp)
	if err != nil {
		return err
	}

	s.Encoded = base64.StdEncoding.EncodeToString(raw)

	return nil
}

func (ic ICECandidate) MarshalRTC() (rtc.ICECandidate, error) {
	var result webrtc.ICECandidate

	b, err := base64.StdEncoding.DecodeString(ic.Encoded)
	if err != nil {
		return rtc.ICECandidate{}, err
	}

	if err := json.Unmarshal(b, &result); err != nil {
		return rtc.ICECandidate{}, err
	}

	return rtc.ICECandidate(result), nil
}

func (ic *ICECandidate) UnmarshalRTC(candidate rtc.ICECandidate) error {
	raw, err := json.Marshal(candidate)
	if err != nil {
		return err
	}

	ic.Encoded = base64.StdEncoding.EncodeToString(raw)

	return nil
}
