package deck

import (
	"fmt"
	"github.com/gofrs/uuid"
	"myl/common"
	"myl/utils"
	"myl/preset"
)

type Card struct {
	UUID uuid.UUID
	Code common.Code
}

func (c *Card) String() string {
	return fmt.Sprintf("<Card Code:%s UUID:%s>", c.Code, c.UUID)
}

type Cards []Card

// CardDisplay stores information exclusively related to visual display, like the text 
// of the abilities, the card description, etc.
type CardDisplay struct {}

// CardMetadata stores information compiled from abilities execution
type CardMetadata struct {}

func ItemToCard(item preset.Item) Card {
	return Card{
		UUID: utils.NewUUID4(),
		Code: item.Code,
	}
}
