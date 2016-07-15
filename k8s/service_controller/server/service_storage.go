package server

import (
	"github.com/cncf/servicebroker/k8s/service_controller/model"
)

type ServiceBroker struct {
	// json info from create docs
	// http://apidocs.cloudfoundry.org/239/service_brokers/create_a_service_broker.html
	// CF uses name
	Name string `json:"name"`
	// CF uses broker_url, and I assume it has the port included
	// if it is a non standard port
	Hostname string `json:hostname""`
	Port     string `json:port""`
	// CF uses auth_username
	User string `json:user""`
	// CF uses auth_password
	Password string `json:password""`
}

type ServiceStorage interface {
	ListBrokers() ([]*ServiceBroker, error)
	GetBroker(string) (*ServiceBroker, error)
	GetInventory(string) (*model.Catalog, error)
	AddBroker(*ServiceBroker, *model.Catalog) error
	DeleteBroker(string) error

	ListServices(string) ([]*model.ServiceInstance, error)
	GetService(string, string) (*model.ServiceInstance, error)
	ServiceExists(string, string) bool
	AddService(string, *model.ServiceInstance) error
	DeleteService(string, string) error

	ListServiceBindings(string, string) ([]*model.ServiceBinding, error)
	GetServiceBinding(string, string, string) (*model.Credential, error)
	AddServiceBinding(string, *model.ServiceBinding, *model.Credential) error
	DeleteServiceBinding(string, string, string) error
}
