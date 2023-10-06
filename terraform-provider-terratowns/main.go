// Package main is a special package that defines a standalone executable program.:
// This comment explains that the main package is a special package in Go,
// and it's used to define a standalone executable program.
package main

//fmt is short format, it contains functions for formatted I/O
import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"	
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"log"
	"context"

)

// func main() : Defines the main function, the entry point of the app.
// When you run() the program, it starts executing from this function.
// https://developer.hashicorp.com/terraform/tutorials/providers/provider-setup
func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: Provider,
	})

	//Formate PrintLine
	//Prints to standard output

	fmt.Println("Hello, World!")
}

type Config struct{
	Endpoint string
	Token string
	UserUuid string
}

// in golang, a titlecase function will get exported.
func Provider() *schema.Provider {
	var p *schema.Provider
	p = &schema.Provider{
		ResourcesMap:   map[string]*schema.Resource{
			"terratowns_home": Resource(),
		},
		DataSourcesMap: map[string]*schema.Resource{

		},
		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The endpoint for the external service",
			},
			"token": {
				Type:        schema.TypeString,
				Sensitive:   true, // make the token sensitive to hide it the logs
				Required:    true,
				Description: "Bearer token for authorization",
			},
			"user_uuid": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "UUID for configuration",
				ValidateFunc: validateUUID,
			},
		},
	}

	p.ConfigureContextFunc = providerConfigure(p)
	return p
}

//validateUUID is a ValidateFunc for checking if a given string is a valid UUID.
func validateUUID(v interface{}, k string) (ws []string, errors []error) {
	log.Print("validateUUID:start")
    value := v.(string) // Assert that the attribute value is a string.

    _, err := uuid.Parse(value)
    if err != nil {
        errors = append(errors, fmt.Errorf("invalid UUID format"))
    }
	log.Print("vaildateUUID:end")
	return 
}

func providerConfigure (p *schema.Provider) schema.ConfigureContextFunc{
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		log.Print("providerConfig:start")
		config := Config{
			Endpoint : d.Get("endpoint").(string),
			Token: d.Get("token").(string),
			UserUuid: d.Get("user_uuid").(string),
		}
		log.Print("providerConfig:end")
		return &config, nil
	}
}

func Resource() *schema.Resource{
	log.Print("Resource:start")
	resource := &schema.Resource{
		CreateContext: resourceHouseCreate,
		ReadContext: resourceHouseRead,
		UpdateContext: resourceHouseUpdate,
		DeleteContext: resourceHouseDelete,
	}
	log.Print("Resource:end")
	return resource
}

func resourceHouseCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
    var diags diag.Diagnostics
    return diags	
}


func resourceHouseRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
    var diags diag.Diagnostics
    return diags	
}


func resourceHouseUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
    var diags diag.Diagnostics
    return diags	
}

func resourceHouseDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
    var diags diag.Diagnostics
    return diags	
}
