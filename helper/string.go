package helper

import (
	"fmt"
	"strings"

	"github.com/LalatinaHub/LatinaApi/common/account/converter"
	"github.com/LalatinaHub/LatinaSub-go/db"
)

func MakeVPNMessage(account db.DBScheme) string {
	message := []string{}
	message = append(message, fmt.Sprintf("<code>PROTOCOL    </code> : <code>%s</code>", account.VPN))
	message = append(message, fmt.Sprintf("<code>REMARKS     </code> : <code>%s</code>", account.Remark))
	message = append(message, fmt.Sprintf("<code>CC          </code> : <code>%s</code>", account.CountryCode))
	message = append(message, fmt.Sprintf("<code>REGION      </code> : <code>%s</code>", account.Region))
	message = append(message, fmt.Sprintf("<code>SERVER      </code> : <code>%s</code>", account.Server))
	if account.Host != "" {
		message = append(message, fmt.Sprintf("<code>HOST        </code> : <code>%s</code>", account.Host))
	}
	if account.SNI != "" {
		message = append(message, fmt.Sprintf("<code>SNI         </code> : <code>%s</code>", account.SNI))
	}
	message = append(message, fmt.Sprintf("<code>PORT        </code> : <code>%d</code>", account.ServerPort))
	if account.UUID != "" {
		message = append(message, fmt.Sprintf("<code>UUID        </code> : <code>%s</code>", account.UUID))
	}
	if account.Password != "" {
		message = append(message, fmt.Sprintf("<code>PASSWORD    </code> : <code>%s</code>", account.Password))
	}
	message = append(message, fmt.Sprintf("<code>ISP         </code> : <code>%s</code>", account.Org))
	message = append(message, fmt.Sprintf("<code>MODE        </code> : <code>%s</code>", account.ConnMode))
	message = append(message, fmt.Sprintf("<code>TLS         </code> : <code>%t</code>", account.TLS))
	message = append(message, fmt.Sprintf("<code>NETWORK     </code> : <code>%s</code>", account.Transport))
	if account.Path != "" {
		message = append(message, fmt.Sprintf("<code>PATH        </code> : <code>%s</code>", account.Path))
	}
	if account.ServiceName != "" {
		message = append(message, fmt.Sprintf("<code>SERVICE NAME</code> : <code>%s</code>", account.ServiceName))
	}
	message = append(message, fmt.Sprintf("%s", "---------------------------------------------------"))
	message = append(message, fmt.Sprintf("<code>%s</code>", converter.ToRaw([]db.DBScheme{account})))
	message = append(message, fmt.Sprintf("%s", "---------------------------------------------------"))

	return strings.Join(message, "\n")
}
