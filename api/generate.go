//go:build generate
// +build generate

//go:generate bash -evc "go run sigs.k8s.io/controller-tools/cmd/controller-gen rbac:roleName=manager-role crd webhook paths=./... output:crd:artifacts:config=../config/crd/bases"
//go:generate bash -evc "DATE=$DOLLAR(date +'v%d%m%y-%H%M'); find ../config/crd/bases -type f | while read CRD; do go run github.com/kcp-dev/kcp/cli/cmd/kubectl-kcp crd snapshot -f $DOLLAR{CRD} --prefix $DOLLAR{DATE} > ../config/kcp/$DOLLAR(basename $DOLLAR{CRD}); done && sed -i '' 's/v[0-9-]*.widgets.data.my.domain/'$DOLLAR{DATE}'.widgets.data.my.domain/' ../config/kcp/apiexport.yaml"

package api

import (
	_ "github.com/kcp-dev/kcp/cli/cmd/kubectl-kcp"      //nolint:typecheck
	_ "sigs.k8s.io/controller-tools/cmd/controller-gen" //nolint:typecheck
)
