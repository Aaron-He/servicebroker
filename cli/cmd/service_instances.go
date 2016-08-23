package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/cncf/servicebroker/k8s/service_controller/model"
	"github.com/spf13/cobra"
)

const (
	SERVICE_INSTANCES_URL = "/v2/service_instances"
)

var (
	service         string
	plan            string
	servicePlanGUID string
	parameters      string
)

func init() {
	RootCmd.AddCommand(serviceInstancesCmd)
	serviceInstancesCmd.AddCommand(createServiceInstancesCmd)
	createServiceInstancesCmd.Flags().StringVarP(&spaceGUID, "space_guid", "s", "default", "Space GUID on which to instantiate the service to")
	createServiceInstancesCmd.Flags().StringVarP(&parameters, "parameters", "p", "", "Parameters to pass to the service broker for creation (must be JSON object)")

}

var serviceInstancesCmd = &cobra.Command{
	Use:   "service-instances",
	Short: "Manage service instances",
	Long:  "Manage service instances",
}

var createServiceInstancesCmd = &cobra.Command{
	Use:   "create <NAME> <SERVICE> <PLAN>",
	Short: "Create a service instance",
	Long:  "Create a service instance",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 3 {
			return fmt.Errorf("need NAME SERVICE PLAN")
		}
		name := args[0]
		service := args[1]
		plan := args[2]

		servicePlanGUID, err := fetchServicePlanGUID(service, plan)
		if err != nil {
			return err
		}
		req := model.CreateServiceInstanceRequest{
			Name:            name,
			ServicePlanGUID: servicePlanGUID,
			SpaceID:         spaceGUID,
		}
		if len(parameters) > 0 {
			var m interface{}
			err := json.Unmarshal([]byte(parameters), &m)
			if err != nil {
				return err
			}
			req.Parameters = m.(map[string]interface{})
		}
		body, err := json.Marshal(req)
		if err != nil {
			return err
		}
		fmt.Printf("Sending body: %s\n\n", string(body))
		return callService(SERVICE_INSTANCES_URL, "POST", "create service instance", ioutil.NopCloser(bytes.NewReader(body)))
	},
}
