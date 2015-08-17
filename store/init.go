package store

import (
	"fmt"
	"strings"

	"github.com/lessos/bigtree/btapi"

	"github.com/lessos/lessgo/types"
	"github.com/lessos/lessgo/utils"
	"github.com/lessos/lessgo/utilx"

	"github.com/lessos/lessids/idsapi"
)

type InitNew struct {
}

func (i InitNew) Init() {

	//
	role := idsapi.UserRole{
		Meta: types.ObjectMeta{
			ID:      utils.StringEncode16("1", 8),
			Name:    "Administrator",
			UserID:  utils.StringEncode16("sysadmin", 8),
			Created: utilx.TimeNow("atom"),
			Updated: utilx.TimeNow("atom"),
		},
		IdxID:  1,
		Desc:   "Root System Administrator",
		Status: 1,
	}
	rolejs, _ := utils.JsonEncode(role)
	BtAgent.ObjectSet(btapi.ObjectProposal{
		Meta: btapi.ObjectMeta{
			Path: fmt.Sprintf("/role/%s", role.Meta.ID),
		},
		Data: rolejs,
	})

	//
	role.Meta.ID = utils.StringEncode16("100", 8)
	role.Meta.Name = "Member"
	role.IdxID = 100
	role.Desc = "Universal Member"
	rolejs, _ = utils.JsonEncode(role)
	BtAgent.ObjectSet(btapi.ObjectProposal{
		Meta: btapi.ObjectMeta{
			Path: fmt.Sprintf("/role/%s", role.Meta.ID),
		},
		Data: rolejs,
	})

	//
	role.Meta.ID = utils.StringEncode16("1000", 8)
	role.Meta.Name = "Anonymous"
	role.IdxID = 1000
	role.Desc = "Anonymous Member"
	rolejs, _ = utils.JsonEncode(role)
	BtAgent.ObjectSet(btapi.ObjectProposal{
		Meta: btapi.ObjectMeta{
			Path: fmt.Sprintf("/role/%s", role.Meta.ID),
		},
		Data: rolejs,
	})

	//
	ps := []idsapi.AppPrivilege{
		{
			Privilege: "sys.admin",
			Roles:     []uint32{1},
			Desc:      "System Management",
		},
		{
			Privilege: "user.admin",
			Roles:     []uint32{1},
			Desc:      "User Management",
		},
	}

	inst := idsapi.AppInstance{
		Meta: types.ObjectMeta{
			ID:      utils.StringEncode16("lessids", 12),
			Created: utilx.TimeNow("atom"),
			Updated: utilx.TimeNow("atom"),
		},
		AppID:      "lessids",
		AppTitle:   "lessOS ID Service",
		Status:     1,
		Url:        "",
		Privileges: ps,
	}
	instjs, _ := utils.JsonEncode(inst)

	BtAgent.ObjectSet(btapi.ObjectProposal{
		Meta: btapi.ObjectMeta{
			Path: "/app-instance/" + inst.Meta.ID,
		},
		Data: instjs,
	})

	// privilege
	rps := map[uint32][]string{}
	for _, v := range ps {

		for _, rid := range v.Roles {

			if _, ok := rps[rid]; !ok {
				rps[rid] = []string{}
			}

			rps[rid] = append(rps[rid], v.Privilege)
		}
	}

	for rid, v := range rps {

		BtAgent.ObjectSet(btapi.ObjectProposal{
			Meta: btapi.ObjectMeta{
				Path: fmt.Sprintf("/role-privilege/%d/%s", rid, inst.Meta.ID),
			},
			Data: strings.Join(v, ","),
		})
	}
}