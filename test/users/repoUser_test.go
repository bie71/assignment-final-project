package users_test

import (
	"fmt"
	"log"
	"testing"
)

func TestFindUserByUsername(t *testing.T) {
	user, err := repoUser.FindUserByUsername(ctx, "bi7")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(user)
}
