package main
import("context";"fmt";"os";"net/url";scout "github.com/jsonscraper/uptimescout-api-go")
func main(){client:=scout.New(os.Getenv("UPTIMESCOUT_LICENSE_KEY"));data,err:=client.Request(context.Background(),"status",url.Values{"service":{"github"}});if err!=nil{panic(err)};fmt.Println(data)}
