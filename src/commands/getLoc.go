package commands

import (
	"fmt"

	"github.com/imroc/req/v3"
	"github.com/spf13/cobra"
)

func Get() *cobra.Command {
	return &cobra.Command{
		Use:   "get [cep]",
		Short: "return loc data with CEP",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			viacepUrl := "https://viacep.com.br/ws/" + args[0] + "/json/"

			res, err := req.Get(viacepUrl)

			if err != nil {
				fmt.Println("err:", err)
				if res.Dump() != "" {
					fmt.Println("raw content:")
					fmt.Println(res.Dump())
				}
				return
			}

			if !res.IsSuccess() {
				fmt.Println(err)
			}

			fmt.Printf("CEP: %s\n", res)
		},
	}
}
