// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package atproto

// schema: com.atproto.admin.defs

import (
	"encoding/json"
	"fmt"

	"github.com/bluesky-social/indigo/lex/util"
)

// AdminDefs_AccountView is a "accountView" in the com.atproto.admin.defs schema.
type AdminDefs_AccountView struct {
	Did              string                     `json:"did" cborgen:"did"`
	Email            *string                    `json:"email,omitempty" cborgen:"email,omitempty"`
	EmailConfirmedAt *string                    `json:"emailConfirmedAt,omitempty" cborgen:"emailConfirmedAt,omitempty"`
	Handle           string                     `json:"handle" cborgen:"handle"`
	IndexedAt        string                     `json:"indexedAt" cborgen:"indexedAt"`
	InviteNote       *string                    `json:"inviteNote,omitempty" cborgen:"inviteNote,omitempty"`
	InvitedBy        *ServerDefs_InviteCode     `json:"invitedBy,omitempty" cborgen:"invitedBy,omitempty"`
	Invites          []*ServerDefs_InviteCode   `json:"invites,omitempty" cborgen:"invites,omitempty"`
	InvitesDisabled  *bool                      `json:"invitesDisabled,omitempty" cborgen:"invitesDisabled,omitempty"`
	RelatedRecords   []*util.LexiconTypeDecoder `json:"relatedRecords,omitempty" cborgen:"relatedRecords,omitempty"`
}

// AdminDefs_BlobView is a "blobView" in the com.atproto.admin.defs schema.
type AdminDefs_BlobView struct {
	Cid        string                      `json:"cid" cborgen:"cid"`
	CreatedAt  string                      `json:"createdAt" cborgen:"createdAt"`
	Details    *AdminDefs_BlobView_Details `json:"details,omitempty" cborgen:"details,omitempty"`
	MimeType   string                      `json:"mimeType" cborgen:"mimeType"`
	Moderation *AdminDefs_Moderation       `json:"moderation,omitempty" cborgen:"moderation,omitempty"`
	Size       int64                       `json:"size" cborgen:"size"`
}

type AdminDefs_BlobView_Details struct {
	AdminDefs_ImageDetails *AdminDefs_ImageDetails
	AdminDefs_VideoDetails *AdminDefs_VideoDetails
}

func (t *AdminDefs_BlobView_Details) MarshalJSON() ([]byte, error) {
	if t.AdminDefs_ImageDetails != nil {
		t.AdminDefs_ImageDetails.LexiconTypeID = "com.atproto.admin.defs#imageDetails"
		return json.Marshal(t.AdminDefs_ImageDetails)
	}
	if t.AdminDefs_VideoDetails != nil {
		t.AdminDefs_VideoDetails.LexiconTypeID = "com.atproto.admin.defs#videoDetails"
		return json.Marshal(t.AdminDefs_VideoDetails)
	}
	return nil, fmt.Errorf("cannot marshal empty enum")
}
func (t *AdminDefs_BlobView_Details) UnmarshalJSON(b []byte) error {
	typ, err := util.TypeExtract(b)
	if err != nil {
		return err
	}

	switch typ {
	case "com.atproto.admin.defs#imageDetails":
		t.AdminDefs_ImageDetails = new(AdminDefs_ImageDetails)
		return json.Unmarshal(b, t.AdminDefs_ImageDetails)
	case "com.atproto.admin.defs#videoDetails":
		t.AdminDefs_VideoDetails = new(AdminDefs_VideoDetails)
		return json.Unmarshal(b, t.AdminDefs_VideoDetails)

	default:
		return nil
	}
}

// AdminDefs_CommunicationTemplateView is a "communicationTemplateView" in the com.atproto.admin.defs schema.
type AdminDefs_CommunicationTemplateView struct {
	// contentMarkdown: Subject of the message, used in emails.
	ContentMarkdown string `json:"contentMarkdown" cborgen:"contentMarkdown"`
	CreatedAt       string `json:"createdAt" cborgen:"createdAt"`
	Disabled        bool   `json:"disabled" cborgen:"disabled"`
	Id              string `json:"id" cborgen:"id"`
	// lastUpdatedBy: DID of the user who last updated the template.
	LastUpdatedBy string `json:"lastUpdatedBy" cborgen:"lastUpdatedBy"`
	// name: Name of the template.
	Name string `json:"name" cborgen:"name"`
	// subject: Content of the template, can contain markdown and variable placeholders.
	Subject   *string `json:"subject,omitempty" cborgen:"subject,omitempty"`
	UpdatedAt string  `json:"updatedAt" cborgen:"updatedAt"`
}

// AdminDefs_ImageDetails is a "imageDetails" in the com.atproto.admin.defs schema.
//
// RECORDTYPE: AdminDefs_ImageDetails
type AdminDefs_ImageDetails struct {
	LexiconTypeID string `json:"$type,const=com.atproto.admin.defs#imageDetails" cborgen:"$type,const=com.atproto.admin.defs#imageDetails"`
	Height        int64  `json:"height" cborgen:"height"`
	Width         int64  `json:"width" cborgen:"width"`
}

// AdminDefs_ModEventAcknowledge is a "modEventAcknowledge" in the com.atproto.admin.defs schema.
//
// RECORDTYPE: AdminDefs_ModEventAcknowledge
type AdminDefs_ModEventAcknowledge struct {
	LexiconTypeID string  `json:"$type,const=com.atproto.admin.defs#modEventAcknowledge" cborgen:"$type,const=com.atproto.admin.defs#modEventAcknowledge"`
	Comment       *string `json:"comment,omitempty" cborgen:"comment,omitempty"`
}

// AdminDefs_ModEventComment is a "modEventComment" in the com.atproto.admin.defs schema.
//
// # Add a comment to a subject
//
// RECORDTYPE: AdminDefs_ModEventComment
type AdminDefs_ModEventComment struct {
	LexiconTypeID string `json:"$type,const=com.atproto.admin.defs#modEventComment" cborgen:"$type,const=com.atproto.admin.defs#modEventComment"`
	Comment       string `json:"comment" cborgen:"comment"`
	// sticky: Make the comment persistent on the subject
	Sticky *bool `json:"sticky,omitempty" cborgen:"sticky,omitempty"`
}

// AdminDefs_ModEventEmail is a "modEventEmail" in the com.atproto.admin.defs schema.
//
// # Keep a log of outgoing email to a user
//
// RECORDTYPE: AdminDefs_ModEventEmail
type AdminDefs_ModEventEmail struct {
	LexiconTypeID string `json:"$type,const=com.atproto.admin.defs#modEventEmail" cborgen:"$type,const=com.atproto.admin.defs#modEventEmail"`
	// comment: Additional comment about the outgoing comm.
	Comment *string `json:"comment,omitempty" cborgen:"comment,omitempty"`
	// subjectLine: The subject line of the email sent to the user.
	SubjectLine string `json:"subjectLine" cborgen:"subjectLine"`
}

// AdminDefs_ModEventEscalate is a "modEventEscalate" in the com.atproto.admin.defs schema.
//
// RECORDTYPE: AdminDefs_ModEventEscalate
type AdminDefs_ModEventEscalate struct {
	LexiconTypeID string  `json:"$type,const=com.atproto.admin.defs#modEventEscalate" cborgen:"$type,const=com.atproto.admin.defs#modEventEscalate"`
	Comment       *string `json:"comment,omitempty" cborgen:"comment,omitempty"`
}

// AdminDefs_ModEventLabel is a "modEventLabel" in the com.atproto.admin.defs schema.
//
// Apply/Negate labels on a subject
//
// RECORDTYPE: AdminDefs_ModEventLabel
type AdminDefs_ModEventLabel struct {
	LexiconTypeID   string   `json:"$type,const=com.atproto.admin.defs#modEventLabel" cborgen:"$type,const=com.atproto.admin.defs#modEventLabel"`
	Comment         *string  `json:"comment,omitempty" cborgen:"comment,omitempty"`
	CreateLabelVals []string `json:"createLabelVals" cborgen:"createLabelVals"`
	NegateLabelVals []string `json:"negateLabelVals" cborgen:"negateLabelVals"`
}

// AdminDefs_ModEventMute is a "modEventMute" in the com.atproto.admin.defs schema.
//
// # Mute incoming reports on a subject
//
// RECORDTYPE: AdminDefs_ModEventMute
type AdminDefs_ModEventMute struct {
	LexiconTypeID string  `json:"$type,const=com.atproto.admin.defs#modEventMute" cborgen:"$type,const=com.atproto.admin.defs#modEventMute"`
	Comment       *string `json:"comment,omitempty" cborgen:"comment,omitempty"`
	// durationInHours: Indicates how long the subject should remain muted.
	DurationInHours int64 `json:"durationInHours" cborgen:"durationInHours"`
}

// AdminDefs_ModEventReport is a "modEventReport" in the com.atproto.admin.defs schema.
//
// # Report a subject
//
// RECORDTYPE: AdminDefs_ModEventReport
type AdminDefs_ModEventReport struct {
	LexiconTypeID string  `json:"$type,const=com.atproto.admin.defs#modEventReport" cborgen:"$type,const=com.atproto.admin.defs#modEventReport"`
	Comment       *string `json:"comment,omitempty" cborgen:"comment,omitempty"`
	ReportType    *string `json:"reportType" cborgen:"reportType"`
}

// AdminDefs_ModEventResolveAppeal is a "modEventResolveAppeal" in the com.atproto.admin.defs schema.
//
// # Resolve appeal on a subject
//
// RECORDTYPE: AdminDefs_ModEventResolveAppeal
type AdminDefs_ModEventResolveAppeal struct {
	LexiconTypeID string `json:"$type,const=com.atproto.admin.defs#modEventResolveAppeal" cborgen:"$type,const=com.atproto.admin.defs#modEventResolveAppeal"`
	// comment: Describe resolution.
	Comment *string `json:"comment,omitempty" cborgen:"comment,omitempty"`
}

// AdminDefs_ModEventReverseTakedown is a "modEventReverseTakedown" in the com.atproto.admin.defs schema.
//
// # Revert take down action on a subject
//
// RECORDTYPE: AdminDefs_ModEventReverseTakedown
type AdminDefs_ModEventReverseTakedown struct {
	LexiconTypeID string `json:"$type,const=com.atproto.admin.defs#modEventReverseTakedown" cborgen:"$type,const=com.atproto.admin.defs#modEventReverseTakedown"`
	// comment: Describe reasoning behind the reversal.
	Comment *string `json:"comment,omitempty" cborgen:"comment,omitempty"`
}

// AdminDefs_ModEventTag is a "modEventTag" in the com.atproto.admin.defs schema.
//
// Add/Remove a tag on a subject
//
// RECORDTYPE: AdminDefs_ModEventTag
type AdminDefs_ModEventTag struct {
	LexiconTypeID string `json:"$type,const=com.atproto.admin.defs#modEventTag" cborgen:"$type,const=com.atproto.admin.defs#modEventTag"`
	// add: Tags to be added to the subject. If already exists, won't be duplicated.
	Add []string `json:"add" cborgen:"add"`
	// comment: Additional comment about added/removed tags.
	Comment *string `json:"comment,omitempty" cborgen:"comment,omitempty"`
	// remove: Tags to be removed to the subject. Ignores a tag If it doesn't exist, won't be duplicated.
	Remove []string `json:"remove" cborgen:"remove"`
}

// AdminDefs_ModEventTakedown is a "modEventTakedown" in the com.atproto.admin.defs schema.
//
// # Take down a subject permanently or temporarily
//
// RECORDTYPE: AdminDefs_ModEventTakedown
type AdminDefs_ModEventTakedown struct {
	LexiconTypeID string  `json:"$type,const=com.atproto.admin.defs#modEventTakedown" cborgen:"$type,const=com.atproto.admin.defs#modEventTakedown"`
	Comment       *string `json:"comment,omitempty" cborgen:"comment,omitempty"`
	// durationInHours: Indicates how long the takedown should be in effect before automatically expiring.
	DurationInHours *int64 `json:"durationInHours,omitempty" cborgen:"durationInHours,omitempty"`
}

// AdminDefs_ModEventUnmute is a "modEventUnmute" in the com.atproto.admin.defs schema.
//
// # Unmute action on a subject
//
// RECORDTYPE: AdminDefs_ModEventUnmute
type AdminDefs_ModEventUnmute struct {
	LexiconTypeID string `json:"$type,const=com.atproto.admin.defs#modEventUnmute" cborgen:"$type,const=com.atproto.admin.defs#modEventUnmute"`
	// comment: Describe reasoning behind the reversal.
	Comment *string `json:"comment,omitempty" cborgen:"comment,omitempty"`
}

// AdminDefs_ModEventView is a "modEventView" in the com.atproto.admin.defs schema.
type AdminDefs_ModEventView struct {
	CreatedAt       string                          `json:"createdAt" cborgen:"createdAt"`
	CreatedBy       string                          `json:"createdBy" cborgen:"createdBy"`
	CreatorHandle   *string                         `json:"creatorHandle,omitempty" cborgen:"creatorHandle,omitempty"`
	Event           *AdminDefs_ModEventView_Event   `json:"event" cborgen:"event"`
	Id              int64                           `json:"id" cborgen:"id"`
	Subject         *AdminDefs_ModEventView_Subject `json:"subject" cborgen:"subject"`
	SubjectBlobCids []string                        `json:"subjectBlobCids" cborgen:"subjectBlobCids"`
	SubjectHandle   *string                         `json:"subjectHandle,omitempty" cborgen:"subjectHandle,omitempty"`
}

// AdminDefs_ModEventViewDetail is a "modEventViewDetail" in the com.atproto.admin.defs schema.
type AdminDefs_ModEventViewDetail struct {
	CreatedAt    string                                `json:"createdAt" cborgen:"createdAt"`
	CreatedBy    string                                `json:"createdBy" cborgen:"createdBy"`
	Event        *AdminDefs_ModEventViewDetail_Event   `json:"event" cborgen:"event"`
	Id           int64                                 `json:"id" cborgen:"id"`
	Subject      *AdminDefs_ModEventViewDetail_Subject `json:"subject" cborgen:"subject"`
	SubjectBlobs []*AdminDefs_BlobView                 `json:"subjectBlobs" cborgen:"subjectBlobs"`
}

type AdminDefs_ModEventViewDetail_Event struct {
	AdminDefs_ModEventTakedown        *AdminDefs_ModEventTakedown
	AdminDefs_ModEventReverseTakedown *AdminDefs_ModEventReverseTakedown
	AdminDefs_ModEventComment         *AdminDefs_ModEventComment
	AdminDefs_ModEventReport          *AdminDefs_ModEventReport
	AdminDefs_ModEventLabel           *AdminDefs_ModEventLabel
	AdminDefs_ModEventAcknowledge     *AdminDefs_ModEventAcknowledge
	AdminDefs_ModEventEscalate        *AdminDefs_ModEventEscalate
	AdminDefs_ModEventMute            *AdminDefs_ModEventMute
	AdminDefs_ModEventEmail           *AdminDefs_ModEventEmail
	AdminDefs_ModEventResolveAppeal   *AdminDefs_ModEventResolveAppeal
}

func (t *AdminDefs_ModEventViewDetail_Event) MarshalJSON() ([]byte, error) {
	if t.AdminDefs_ModEventTakedown != nil {
		t.AdminDefs_ModEventTakedown.LexiconTypeID = "com.atproto.admin.defs#modEventTakedown"
		return json.Marshal(t.AdminDefs_ModEventTakedown)
	}
	if t.AdminDefs_ModEventReverseTakedown != nil {
		t.AdminDefs_ModEventReverseTakedown.LexiconTypeID = "com.atproto.admin.defs#modEventReverseTakedown"
		return json.Marshal(t.AdminDefs_ModEventReverseTakedown)
	}
	if t.AdminDefs_ModEventComment != nil {
		t.AdminDefs_ModEventComment.LexiconTypeID = "com.atproto.admin.defs#modEventComment"
		return json.Marshal(t.AdminDefs_ModEventComment)
	}
	if t.AdminDefs_ModEventReport != nil {
		t.AdminDefs_ModEventReport.LexiconTypeID = "com.atproto.admin.defs#modEventReport"
		return json.Marshal(t.AdminDefs_ModEventReport)
	}
	if t.AdminDefs_ModEventLabel != nil {
		t.AdminDefs_ModEventLabel.LexiconTypeID = "com.atproto.admin.defs#modEventLabel"
		return json.Marshal(t.AdminDefs_ModEventLabel)
	}
	if t.AdminDefs_ModEventAcknowledge != nil {
		t.AdminDefs_ModEventAcknowledge.LexiconTypeID = "com.atproto.admin.defs#modEventAcknowledge"
		return json.Marshal(t.AdminDefs_ModEventAcknowledge)
	}
	if t.AdminDefs_ModEventEscalate != nil {
		t.AdminDefs_ModEventEscalate.LexiconTypeID = "com.atproto.admin.defs#modEventEscalate"
		return json.Marshal(t.AdminDefs_ModEventEscalate)
	}
	if t.AdminDefs_ModEventMute != nil {
		t.AdminDefs_ModEventMute.LexiconTypeID = "com.atproto.admin.defs#modEventMute"
		return json.Marshal(t.AdminDefs_ModEventMute)
	}
	if t.AdminDefs_ModEventEmail != nil {
		t.AdminDefs_ModEventEmail.LexiconTypeID = "com.atproto.admin.defs#modEventEmail"
		return json.Marshal(t.AdminDefs_ModEventEmail)
	}
	if t.AdminDefs_ModEventResolveAppeal != nil {
		t.AdminDefs_ModEventResolveAppeal.LexiconTypeID = "com.atproto.admin.defs#modEventResolveAppeal"
		return json.Marshal(t.AdminDefs_ModEventResolveAppeal)
	}
	return nil, fmt.Errorf("cannot marshal empty enum")
}
func (t *AdminDefs_ModEventViewDetail_Event) UnmarshalJSON(b []byte) error {
	typ, err := util.TypeExtract(b)
	if err != nil {
		return err
	}

	switch typ {
	case "com.atproto.admin.defs#modEventTakedown":
		t.AdminDefs_ModEventTakedown = new(AdminDefs_ModEventTakedown)
		return json.Unmarshal(b, t.AdminDefs_ModEventTakedown)
	case "com.atproto.admin.defs#modEventReverseTakedown":
		t.AdminDefs_ModEventReverseTakedown = new(AdminDefs_ModEventReverseTakedown)
		return json.Unmarshal(b, t.AdminDefs_ModEventReverseTakedown)
	case "com.atproto.admin.defs#modEventComment":
		t.AdminDefs_ModEventComment = new(AdminDefs_ModEventComment)
		return json.Unmarshal(b, t.AdminDefs_ModEventComment)
	case "com.atproto.admin.defs#modEventReport":
		t.AdminDefs_ModEventReport = new(AdminDefs_ModEventReport)
		return json.Unmarshal(b, t.AdminDefs_ModEventReport)
	case "com.atproto.admin.defs#modEventLabel":
		t.AdminDefs_ModEventLabel = new(AdminDefs_ModEventLabel)
		return json.Unmarshal(b, t.AdminDefs_ModEventLabel)
	case "com.atproto.admin.defs#modEventAcknowledge":
		t.AdminDefs_ModEventAcknowledge = new(AdminDefs_ModEventAcknowledge)
		return json.Unmarshal(b, t.AdminDefs_ModEventAcknowledge)
	case "com.atproto.admin.defs#modEventEscalate":
		t.AdminDefs_ModEventEscalate = new(AdminDefs_ModEventEscalate)
		return json.Unmarshal(b, t.AdminDefs_ModEventEscalate)
	case "com.atproto.admin.defs#modEventMute":
		t.AdminDefs_ModEventMute = new(AdminDefs_ModEventMute)
		return json.Unmarshal(b, t.AdminDefs_ModEventMute)
	case "com.atproto.admin.defs#modEventEmail":
		t.AdminDefs_ModEventEmail = new(AdminDefs_ModEventEmail)
		return json.Unmarshal(b, t.AdminDefs_ModEventEmail)
	case "com.atproto.admin.defs#modEventResolveAppeal":
		t.AdminDefs_ModEventResolveAppeal = new(AdminDefs_ModEventResolveAppeal)
		return json.Unmarshal(b, t.AdminDefs_ModEventResolveAppeal)

	default:
		return nil
	}
}

type AdminDefs_ModEventViewDetail_Subject struct {
	AdminDefs_RepoView           *AdminDefs_RepoView
	AdminDefs_RepoViewNotFound   *AdminDefs_RepoViewNotFound
	AdminDefs_RecordView         *AdminDefs_RecordView
	AdminDefs_RecordViewNotFound *AdminDefs_RecordViewNotFound
}

func (t *AdminDefs_ModEventViewDetail_Subject) MarshalJSON() ([]byte, error) {
	if t.AdminDefs_RepoView != nil {
		t.AdminDefs_RepoView.LexiconTypeID = "com.atproto.admin.defs#repoView"
		return json.Marshal(t.AdminDefs_RepoView)
	}
	if t.AdminDefs_RepoViewNotFound != nil {
		t.AdminDefs_RepoViewNotFound.LexiconTypeID = "com.atproto.admin.defs#repoViewNotFound"
		return json.Marshal(t.AdminDefs_RepoViewNotFound)
	}
	if t.AdminDefs_RecordView != nil {
		t.AdminDefs_RecordView.LexiconTypeID = "com.atproto.admin.defs#recordView"
		return json.Marshal(t.AdminDefs_RecordView)
	}
	if t.AdminDefs_RecordViewNotFound != nil {
		t.AdminDefs_RecordViewNotFound.LexiconTypeID = "com.atproto.admin.defs#recordViewNotFound"
		return json.Marshal(t.AdminDefs_RecordViewNotFound)
	}
	return nil, fmt.Errorf("cannot marshal empty enum")
}
func (t *AdminDefs_ModEventViewDetail_Subject) UnmarshalJSON(b []byte) error {
	typ, err := util.TypeExtract(b)
	if err != nil {
		return err
	}

	switch typ {
	case "com.atproto.admin.defs#repoView":
		t.AdminDefs_RepoView = new(AdminDefs_RepoView)
		return json.Unmarshal(b, t.AdminDefs_RepoView)
	case "com.atproto.admin.defs#repoViewNotFound":
		t.AdminDefs_RepoViewNotFound = new(AdminDefs_RepoViewNotFound)
		return json.Unmarshal(b, t.AdminDefs_RepoViewNotFound)
	case "com.atproto.admin.defs#recordView":
		t.AdminDefs_RecordView = new(AdminDefs_RecordView)
		return json.Unmarshal(b, t.AdminDefs_RecordView)
	case "com.atproto.admin.defs#recordViewNotFound":
		t.AdminDefs_RecordViewNotFound = new(AdminDefs_RecordViewNotFound)
		return json.Unmarshal(b, t.AdminDefs_RecordViewNotFound)

	default:
		return nil
	}
}

type AdminDefs_ModEventView_Event struct {
	AdminDefs_ModEventTakedown        *AdminDefs_ModEventTakedown
	AdminDefs_ModEventReverseTakedown *AdminDefs_ModEventReverseTakedown
	AdminDefs_ModEventComment         *AdminDefs_ModEventComment
	AdminDefs_ModEventReport          *AdminDefs_ModEventReport
	AdminDefs_ModEventLabel           *AdminDefs_ModEventLabel
	AdminDefs_ModEventAcknowledge     *AdminDefs_ModEventAcknowledge
	AdminDefs_ModEventEscalate        *AdminDefs_ModEventEscalate
	AdminDefs_ModEventMute            *AdminDefs_ModEventMute
	AdminDefs_ModEventEmail           *AdminDefs_ModEventEmail
	AdminDefs_ModEventResolveAppeal   *AdminDefs_ModEventResolveAppeal
}

func (t *AdminDefs_ModEventView_Event) MarshalJSON() ([]byte, error) {
	if t.AdminDefs_ModEventTakedown != nil {
		t.AdminDefs_ModEventTakedown.LexiconTypeID = "com.atproto.admin.defs#modEventTakedown"
		return json.Marshal(t.AdminDefs_ModEventTakedown)
	}
	if t.AdminDefs_ModEventReverseTakedown != nil {
		t.AdminDefs_ModEventReverseTakedown.LexiconTypeID = "com.atproto.admin.defs#modEventReverseTakedown"
		return json.Marshal(t.AdminDefs_ModEventReverseTakedown)
	}
	if t.AdminDefs_ModEventComment != nil {
		t.AdminDefs_ModEventComment.LexiconTypeID = "com.atproto.admin.defs#modEventComment"
		return json.Marshal(t.AdminDefs_ModEventComment)
	}
	if t.AdminDefs_ModEventReport != nil {
		t.AdminDefs_ModEventReport.LexiconTypeID = "com.atproto.admin.defs#modEventReport"
		return json.Marshal(t.AdminDefs_ModEventReport)
	}
	if t.AdminDefs_ModEventLabel != nil {
		t.AdminDefs_ModEventLabel.LexiconTypeID = "com.atproto.admin.defs#modEventLabel"
		return json.Marshal(t.AdminDefs_ModEventLabel)
	}
	if t.AdminDefs_ModEventAcknowledge != nil {
		t.AdminDefs_ModEventAcknowledge.LexiconTypeID = "com.atproto.admin.defs#modEventAcknowledge"
		return json.Marshal(t.AdminDefs_ModEventAcknowledge)
	}
	if t.AdminDefs_ModEventEscalate != nil {
		t.AdminDefs_ModEventEscalate.LexiconTypeID = "com.atproto.admin.defs#modEventEscalate"
		return json.Marshal(t.AdminDefs_ModEventEscalate)
	}
	if t.AdminDefs_ModEventMute != nil {
		t.AdminDefs_ModEventMute.LexiconTypeID = "com.atproto.admin.defs#modEventMute"
		return json.Marshal(t.AdminDefs_ModEventMute)
	}
	if t.AdminDefs_ModEventEmail != nil {
		t.AdminDefs_ModEventEmail.LexiconTypeID = "com.atproto.admin.defs#modEventEmail"
		return json.Marshal(t.AdminDefs_ModEventEmail)
	}
	if t.AdminDefs_ModEventResolveAppeal != nil {
		t.AdminDefs_ModEventResolveAppeal.LexiconTypeID = "com.atproto.admin.defs#modEventResolveAppeal"
		return json.Marshal(t.AdminDefs_ModEventResolveAppeal)
	}
	return nil, fmt.Errorf("cannot marshal empty enum")
}
func (t *AdminDefs_ModEventView_Event) UnmarshalJSON(b []byte) error {
	typ, err := util.TypeExtract(b)
	if err != nil {
		return err
	}

	switch typ {
	case "com.atproto.admin.defs#modEventTakedown":
		t.AdminDefs_ModEventTakedown = new(AdminDefs_ModEventTakedown)
		return json.Unmarshal(b, t.AdminDefs_ModEventTakedown)
	case "com.atproto.admin.defs#modEventReverseTakedown":
		t.AdminDefs_ModEventReverseTakedown = new(AdminDefs_ModEventReverseTakedown)
		return json.Unmarshal(b, t.AdminDefs_ModEventReverseTakedown)
	case "com.atproto.admin.defs#modEventComment":
		t.AdminDefs_ModEventComment = new(AdminDefs_ModEventComment)
		return json.Unmarshal(b, t.AdminDefs_ModEventComment)
	case "com.atproto.admin.defs#modEventReport":
		t.AdminDefs_ModEventReport = new(AdminDefs_ModEventReport)
		return json.Unmarshal(b, t.AdminDefs_ModEventReport)
	case "com.atproto.admin.defs#modEventLabel":
		t.AdminDefs_ModEventLabel = new(AdminDefs_ModEventLabel)
		return json.Unmarshal(b, t.AdminDefs_ModEventLabel)
	case "com.atproto.admin.defs#modEventAcknowledge":
		t.AdminDefs_ModEventAcknowledge = new(AdminDefs_ModEventAcknowledge)
		return json.Unmarshal(b, t.AdminDefs_ModEventAcknowledge)
	case "com.atproto.admin.defs#modEventEscalate":
		t.AdminDefs_ModEventEscalate = new(AdminDefs_ModEventEscalate)
		return json.Unmarshal(b, t.AdminDefs_ModEventEscalate)
	case "com.atproto.admin.defs#modEventMute":
		t.AdminDefs_ModEventMute = new(AdminDefs_ModEventMute)
		return json.Unmarshal(b, t.AdminDefs_ModEventMute)
	case "com.atproto.admin.defs#modEventEmail":
		t.AdminDefs_ModEventEmail = new(AdminDefs_ModEventEmail)
		return json.Unmarshal(b, t.AdminDefs_ModEventEmail)
	case "com.atproto.admin.defs#modEventResolveAppeal":
		t.AdminDefs_ModEventResolveAppeal = new(AdminDefs_ModEventResolveAppeal)
		return json.Unmarshal(b, t.AdminDefs_ModEventResolveAppeal)

	default:
		return nil
	}
}

type AdminDefs_ModEventView_Subject struct {
	AdminDefs_RepoRef *AdminDefs_RepoRef
	RepoStrongRef     *RepoStrongRef
}

func (t *AdminDefs_ModEventView_Subject) MarshalJSON() ([]byte, error) {
	if t.AdminDefs_RepoRef != nil {
		t.AdminDefs_RepoRef.LexiconTypeID = "com.atproto.admin.defs#repoRef"
		return json.Marshal(t.AdminDefs_RepoRef)
	}
	if t.RepoStrongRef != nil {
		t.RepoStrongRef.LexiconTypeID = "com.atproto.repo.strongRef"
		return json.Marshal(t.RepoStrongRef)
	}
	return nil, fmt.Errorf("cannot marshal empty enum")
}
func (t *AdminDefs_ModEventView_Subject) UnmarshalJSON(b []byte) error {
	typ, err := util.TypeExtract(b)
	if err != nil {
		return err
	}

	switch typ {
	case "com.atproto.admin.defs#repoRef":
		t.AdminDefs_RepoRef = new(AdminDefs_RepoRef)
		return json.Unmarshal(b, t.AdminDefs_RepoRef)
	case "com.atproto.repo.strongRef":
		t.RepoStrongRef = new(RepoStrongRef)
		return json.Unmarshal(b, t.RepoStrongRef)

	default:
		return nil
	}
}

// AdminDefs_Moderation is a "moderation" in the com.atproto.admin.defs schema.
type AdminDefs_Moderation struct {
	SubjectStatus *AdminDefs_SubjectStatusView `json:"subjectStatus,omitempty" cborgen:"subjectStatus,omitempty"`
}

// AdminDefs_ModerationDetail is a "moderationDetail" in the com.atproto.admin.defs schema.
type AdminDefs_ModerationDetail struct {
	SubjectStatus *AdminDefs_SubjectStatusView `json:"subjectStatus,omitempty" cborgen:"subjectStatus,omitempty"`
}

// AdminDefs_RecordView is a "recordView" in the com.atproto.admin.defs schema.
//
// RECORDTYPE: AdminDefs_RecordView
type AdminDefs_RecordView struct {
	LexiconTypeID string                   `json:"$type,const=com.atproto.admin.defs#recordView" cborgen:"$type,const=com.atproto.admin.defs#recordView"`
	BlobCids      []string                 `json:"blobCids" cborgen:"blobCids"`
	Cid           string                   `json:"cid" cborgen:"cid"`
	IndexedAt     string                   `json:"indexedAt" cborgen:"indexedAt"`
	Moderation    *AdminDefs_Moderation    `json:"moderation" cborgen:"moderation"`
	Repo          *AdminDefs_RepoView      `json:"repo" cborgen:"repo"`
	Uri           string                   `json:"uri" cborgen:"uri"`
	Value         *util.LexiconTypeDecoder `json:"value" cborgen:"value"`
}

// AdminDefs_RecordViewDetail is a "recordViewDetail" in the com.atproto.admin.defs schema.
type AdminDefs_RecordViewDetail struct {
	Blobs      []*AdminDefs_BlobView       `json:"blobs" cborgen:"blobs"`
	Cid        string                      `json:"cid" cborgen:"cid"`
	IndexedAt  string                      `json:"indexedAt" cborgen:"indexedAt"`
	Labels     []*LabelDefs_Label          `json:"labels,omitempty" cborgen:"labels,omitempty"`
	Moderation *AdminDefs_ModerationDetail `json:"moderation" cborgen:"moderation"`
	Repo       *AdminDefs_RepoView         `json:"repo" cborgen:"repo"`
	Uri        string                      `json:"uri" cborgen:"uri"`
	Value      *util.LexiconTypeDecoder    `json:"value" cborgen:"value"`
}

// AdminDefs_RecordViewNotFound is a "recordViewNotFound" in the com.atproto.admin.defs schema.
//
// RECORDTYPE: AdminDefs_RecordViewNotFound
type AdminDefs_RecordViewNotFound struct {
	LexiconTypeID string `json:"$type,const=com.atproto.admin.defs#recordViewNotFound" cborgen:"$type,const=com.atproto.admin.defs#recordViewNotFound"`
	Uri           string `json:"uri" cborgen:"uri"`
}

// AdminDefs_RepoBlobRef is a "repoBlobRef" in the com.atproto.admin.defs schema.
//
// RECORDTYPE: AdminDefs_RepoBlobRef
type AdminDefs_RepoBlobRef struct {
	LexiconTypeID string  `json:"$type,const=com.atproto.admin.defs#repoBlobRef" cborgen:"$type,const=com.atproto.admin.defs#repoBlobRef"`
	Cid           string  `json:"cid" cborgen:"cid"`
	Did           string  `json:"did" cborgen:"did"`
	RecordUri     *string `json:"recordUri,omitempty" cborgen:"recordUri,omitempty"`
}

// AdminDefs_RepoRef is a "repoRef" in the com.atproto.admin.defs schema.
//
// RECORDTYPE: AdminDefs_RepoRef
type AdminDefs_RepoRef struct {
	LexiconTypeID string `json:"$type,const=com.atproto.admin.defs#repoRef" cborgen:"$type,const=com.atproto.admin.defs#repoRef"`
	Did           string `json:"did" cborgen:"did"`
}

// AdminDefs_RepoView is a "repoView" in the com.atproto.admin.defs schema.
//
// RECORDTYPE: AdminDefs_RepoView
type AdminDefs_RepoView struct {
	LexiconTypeID   string                     `json:"$type,const=com.atproto.admin.defs#repoView" cborgen:"$type,const=com.atproto.admin.defs#repoView"`
	Did             string                     `json:"did" cborgen:"did"`
	Email           *string                    `json:"email,omitempty" cborgen:"email,omitempty"`
	Handle          string                     `json:"handle" cborgen:"handle"`
	IndexedAt       string                     `json:"indexedAt" cborgen:"indexedAt"`
	InviteNote      *string                    `json:"inviteNote,omitempty" cborgen:"inviteNote,omitempty"`
	InvitedBy       *ServerDefs_InviteCode     `json:"invitedBy,omitempty" cborgen:"invitedBy,omitempty"`
	InvitesDisabled *bool                      `json:"invitesDisabled,omitempty" cborgen:"invitesDisabled,omitempty"`
	Moderation      *AdminDefs_Moderation      `json:"moderation" cborgen:"moderation"`
	RelatedRecords  []*util.LexiconTypeDecoder `json:"relatedRecords" cborgen:"relatedRecords"`
}

// AdminDefs_RepoViewDetail is a "repoViewDetail" in the com.atproto.admin.defs schema.
type AdminDefs_RepoViewDetail struct {
	Did              string                      `json:"did" cborgen:"did"`
	Email            *string                     `json:"email,omitempty" cborgen:"email,omitempty"`
	EmailConfirmedAt *string                     `json:"emailConfirmedAt,omitempty" cborgen:"emailConfirmedAt,omitempty"`
	Handle           string                      `json:"handle" cborgen:"handle"`
	IndexedAt        string                      `json:"indexedAt" cborgen:"indexedAt"`
	InviteNote       *string                     `json:"inviteNote,omitempty" cborgen:"inviteNote,omitempty"`
	InvitedBy        *ServerDefs_InviteCode      `json:"invitedBy,omitempty" cborgen:"invitedBy,omitempty"`
	Invites          []*ServerDefs_InviteCode    `json:"invites,omitempty" cborgen:"invites,omitempty"`
	InvitesDisabled  *bool                       `json:"invitesDisabled,omitempty" cborgen:"invitesDisabled,omitempty"`
	Labels           []*LabelDefs_Label          `json:"labels,omitempty" cborgen:"labels,omitempty"`
	Moderation       *AdminDefs_ModerationDetail `json:"moderation" cborgen:"moderation"`
	RelatedRecords   []*util.LexiconTypeDecoder  `json:"relatedRecords" cborgen:"relatedRecords"`
}

// AdminDefs_RepoViewNotFound is a "repoViewNotFound" in the com.atproto.admin.defs schema.
//
// RECORDTYPE: AdminDefs_RepoViewNotFound
type AdminDefs_RepoViewNotFound struct {
	LexiconTypeID string `json:"$type,const=com.atproto.admin.defs#repoViewNotFound" cborgen:"$type,const=com.atproto.admin.defs#repoViewNotFound"`
	Did           string `json:"did" cborgen:"did"`
}

// AdminDefs_ReportView is a "reportView" in the com.atproto.admin.defs schema.
type AdminDefs_ReportView struct {
	Comment             *string                       `json:"comment,omitempty" cborgen:"comment,omitempty"`
	CreatedAt           string                        `json:"createdAt" cborgen:"createdAt"`
	Id                  int64                         `json:"id" cborgen:"id"`
	ReasonType          *string                       `json:"reasonType" cborgen:"reasonType"`
	ReportedBy          string                        `json:"reportedBy" cborgen:"reportedBy"`
	ResolvedByActionIds []int64                       `json:"resolvedByActionIds" cborgen:"resolvedByActionIds"`
	Subject             *AdminDefs_ReportView_Subject `json:"subject" cborgen:"subject"`
	SubjectRepoHandle   *string                       `json:"subjectRepoHandle,omitempty" cborgen:"subjectRepoHandle,omitempty"`
}

// AdminDefs_ReportViewDetail is a "reportViewDetail" in the com.atproto.admin.defs schema.
type AdminDefs_ReportViewDetail struct {
	Comment           *string                             `json:"comment,omitempty" cborgen:"comment,omitempty"`
	CreatedAt         string                              `json:"createdAt" cborgen:"createdAt"`
	Id                int64                               `json:"id" cborgen:"id"`
	ReasonType        *string                             `json:"reasonType" cborgen:"reasonType"`
	ReportedBy        string                              `json:"reportedBy" cborgen:"reportedBy"`
	ResolvedByActions []*AdminDefs_ModEventView           `json:"resolvedByActions" cborgen:"resolvedByActions"`
	Subject           *AdminDefs_ReportViewDetail_Subject `json:"subject" cborgen:"subject"`
	SubjectStatus     *AdminDefs_SubjectStatusView        `json:"subjectStatus,omitempty" cborgen:"subjectStatus,omitempty"`
}

type AdminDefs_ReportViewDetail_Subject struct {
	AdminDefs_RepoView           *AdminDefs_RepoView
	AdminDefs_RepoViewNotFound   *AdminDefs_RepoViewNotFound
	AdminDefs_RecordView         *AdminDefs_RecordView
	AdminDefs_RecordViewNotFound *AdminDefs_RecordViewNotFound
}

func (t *AdminDefs_ReportViewDetail_Subject) MarshalJSON() ([]byte, error) {
	if t.AdminDefs_RepoView != nil {
		t.AdminDefs_RepoView.LexiconTypeID = "com.atproto.admin.defs#repoView"
		return json.Marshal(t.AdminDefs_RepoView)
	}
	if t.AdminDefs_RepoViewNotFound != nil {
		t.AdminDefs_RepoViewNotFound.LexiconTypeID = "com.atproto.admin.defs#repoViewNotFound"
		return json.Marshal(t.AdminDefs_RepoViewNotFound)
	}
	if t.AdminDefs_RecordView != nil {
		t.AdminDefs_RecordView.LexiconTypeID = "com.atproto.admin.defs#recordView"
		return json.Marshal(t.AdminDefs_RecordView)
	}
	if t.AdminDefs_RecordViewNotFound != nil {
		t.AdminDefs_RecordViewNotFound.LexiconTypeID = "com.atproto.admin.defs#recordViewNotFound"
		return json.Marshal(t.AdminDefs_RecordViewNotFound)
	}
	return nil, fmt.Errorf("cannot marshal empty enum")
}
func (t *AdminDefs_ReportViewDetail_Subject) UnmarshalJSON(b []byte) error {
	typ, err := util.TypeExtract(b)
	if err != nil {
		return err
	}

	switch typ {
	case "com.atproto.admin.defs#repoView":
		t.AdminDefs_RepoView = new(AdminDefs_RepoView)
		return json.Unmarshal(b, t.AdminDefs_RepoView)
	case "com.atproto.admin.defs#repoViewNotFound":
		t.AdminDefs_RepoViewNotFound = new(AdminDefs_RepoViewNotFound)
		return json.Unmarshal(b, t.AdminDefs_RepoViewNotFound)
	case "com.atproto.admin.defs#recordView":
		t.AdminDefs_RecordView = new(AdminDefs_RecordView)
		return json.Unmarshal(b, t.AdminDefs_RecordView)
	case "com.atproto.admin.defs#recordViewNotFound":
		t.AdminDefs_RecordViewNotFound = new(AdminDefs_RecordViewNotFound)
		return json.Unmarshal(b, t.AdminDefs_RecordViewNotFound)

	default:
		return nil
	}
}

type AdminDefs_ReportView_Subject struct {
	AdminDefs_RepoRef *AdminDefs_RepoRef
	RepoStrongRef     *RepoStrongRef
}

func (t *AdminDefs_ReportView_Subject) MarshalJSON() ([]byte, error) {
	if t.AdminDefs_RepoRef != nil {
		t.AdminDefs_RepoRef.LexiconTypeID = "com.atproto.admin.defs#repoRef"
		return json.Marshal(t.AdminDefs_RepoRef)
	}
	if t.RepoStrongRef != nil {
		t.RepoStrongRef.LexiconTypeID = "com.atproto.repo.strongRef"
		return json.Marshal(t.RepoStrongRef)
	}
	return nil, fmt.Errorf("cannot marshal empty enum")
}
func (t *AdminDefs_ReportView_Subject) UnmarshalJSON(b []byte) error {
	typ, err := util.TypeExtract(b)
	if err != nil {
		return err
	}

	switch typ {
	case "com.atproto.admin.defs#repoRef":
		t.AdminDefs_RepoRef = new(AdminDefs_RepoRef)
		return json.Unmarshal(b, t.AdminDefs_RepoRef)
	case "com.atproto.repo.strongRef":
		t.RepoStrongRef = new(RepoStrongRef)
		return json.Unmarshal(b, t.RepoStrongRef)

	default:
		return nil
	}
}

// AdminDefs_StatusAttr is a "statusAttr" in the com.atproto.admin.defs schema.
type AdminDefs_StatusAttr struct {
	Applied bool    `json:"applied" cborgen:"applied"`
	Ref     *string `json:"ref,omitempty" cborgen:"ref,omitempty"`
}

// AdminDefs_SubjectStatusView is a "subjectStatusView" in the com.atproto.admin.defs schema.
type AdminDefs_SubjectStatusView struct {
	// appealed: True indicates that the a previously taken moderator action was appealed against, by the author of the content. False indicates last appeal was resolved by moderators.
	Appealed *bool `json:"appealed,omitempty" cborgen:"appealed,omitempty"`
	// comment: Sticky comment on the subject.
	Comment *string `json:"comment,omitempty" cborgen:"comment,omitempty"`
	// createdAt: Timestamp referencing the first moderation status impacting event was emitted on the subject
	CreatedAt string `json:"createdAt" cborgen:"createdAt"`
	Id        int64  `json:"id" cborgen:"id"`
	// lastAppealedAt: Timestamp referencing when the author of the subject appealed a moderation action
	LastAppealedAt    *string                              `json:"lastAppealedAt,omitempty" cborgen:"lastAppealedAt,omitempty"`
	LastReportedAt    *string                              `json:"lastReportedAt,omitempty" cborgen:"lastReportedAt,omitempty"`
	LastReviewedAt    *string                              `json:"lastReviewedAt,omitempty" cborgen:"lastReviewedAt,omitempty"`
	LastReviewedBy    *string                              `json:"lastReviewedBy,omitempty" cborgen:"lastReviewedBy,omitempty"`
	MuteUntil         *string                              `json:"muteUntil,omitempty" cborgen:"muteUntil,omitempty"`
	ReviewState       *string                              `json:"reviewState" cborgen:"reviewState"`
	Subject           *AdminDefs_SubjectStatusView_Subject `json:"subject" cborgen:"subject"`
	SubjectBlobCids   []string                             `json:"subjectBlobCids,omitempty" cborgen:"subjectBlobCids,omitempty"`
	SubjectRepoHandle *string                              `json:"subjectRepoHandle,omitempty" cborgen:"subjectRepoHandle,omitempty"`
	SuspendUntil      *string                              `json:"suspendUntil,omitempty" cborgen:"suspendUntil,omitempty"`
	Tags              []string                             `json:"tags,omitempty" cborgen:"tags,omitempty"`
	Takendown         *bool                                `json:"takendown,omitempty" cborgen:"takendown,omitempty"`
	// updatedAt: Timestamp referencing when the last update was made to the moderation status of the subject
	UpdatedAt string `json:"updatedAt" cborgen:"updatedAt"`
}

type AdminDefs_SubjectStatusView_Subject struct {
	AdminDefs_RepoRef *AdminDefs_RepoRef
	RepoStrongRef     *RepoStrongRef
}

func (t *AdminDefs_SubjectStatusView_Subject) MarshalJSON() ([]byte, error) {
	if t.AdminDefs_RepoRef != nil {
		t.AdminDefs_RepoRef.LexiconTypeID = "com.atproto.admin.defs#repoRef"
		return json.Marshal(t.AdminDefs_RepoRef)
	}
	if t.RepoStrongRef != nil {
		t.RepoStrongRef.LexiconTypeID = "com.atproto.repo.strongRef"
		return json.Marshal(t.RepoStrongRef)
	}
	return nil, fmt.Errorf("cannot marshal empty enum")
}
func (t *AdminDefs_SubjectStatusView_Subject) UnmarshalJSON(b []byte) error {
	typ, err := util.TypeExtract(b)
	if err != nil {
		return err
	}

	switch typ {
	case "com.atproto.admin.defs#repoRef":
		t.AdminDefs_RepoRef = new(AdminDefs_RepoRef)
		return json.Unmarshal(b, t.AdminDefs_RepoRef)
	case "com.atproto.repo.strongRef":
		t.RepoStrongRef = new(RepoStrongRef)
		return json.Unmarshal(b, t.RepoStrongRef)

	default:
		return nil
	}
}

// AdminDefs_VideoDetails is a "videoDetails" in the com.atproto.admin.defs schema.
//
// RECORDTYPE: AdminDefs_VideoDetails
type AdminDefs_VideoDetails struct {
	LexiconTypeID string `json:"$type,const=com.atproto.admin.defs#videoDetails" cborgen:"$type,const=com.atproto.admin.defs#videoDetails"`
	Height        int64  `json:"height" cborgen:"height"`
	Length        int64  `json:"length" cborgen:"length"`
	Width         int64  `json:"width" cborgen:"width"`
}
