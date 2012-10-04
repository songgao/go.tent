// Steve Phillips / elimisteve
// 2012.09.26

package tent

type Following struct {
	RemoteId    string      `json:"remoteid,omitempty"`
	Entity      Entity      `json:"entity,omitempty"`
	Permissions Permissions `json:"permissions,omitempty"`
	Id          Id          `json:"id,omitempty"`
	CreatedAt   int64       `json:"createdat,omitempty"`
	UpdatedAt   int64       `json:"updatedat,omitempty"`
	Groups      []Group     `json:"groups,omitempty"`
	Profile     Profile     `json:"profile,omitempty"`
	Licenses    []URL       `json:"licenses,omitempty"`
}
