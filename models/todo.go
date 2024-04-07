package models

type (
	Todo struct {
		Id   int    `json:"id" db:"id"`
		Note string `json:"note" db:"note"`
	}

	TodoCreateInput struct {
		Note string `json:"note" binding:"required"`
	}

	TodoReadInput struct {
		Id int
	}

	TodoReadOutput struct {
		Todo Todo `json:"todo"`
	}

	TodoUpdateInput struct {
		Id   int    `json:"id" binding:"required"`
		Note string `json:"note" binding:"required"`
	}

	TodoDeleteInput struct {
		Id int
	}
)
