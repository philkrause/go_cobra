package cmd



import(
	"time"
	"github.com/spf13/cobra"
)

func printTime() *cobra.Command{
	return &cobra.Command{
		Use: "curtime",
		RunE: func(cmd *cobra.Command, args []string) error{
			now := time.Now()
			prettyTime := now.Format(time.RubyDate)
			cmd.Println("The current time is time is", prettyTime)
			return nil
		}
	}
}