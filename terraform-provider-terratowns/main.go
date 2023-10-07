// Package main is a special package that defines a standalone executable program.:
// This comment explains that the main package is a special package in Go,
// and it's used to define a standalone executable program.
package main

//fmt is short format, it contains functions for formatted I/O
import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
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

type Config struct {
	Endpoint string
	Token    string
	UserUuid string
}

// in golang, a titlecase function will get exported.
func Provider() *schema.Provider {
	var p *schema.Provider
	p = &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"terratowns_home": Resource(),
		},
		DataSourcesMap: map[string]*schema.Resource{},
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

// validateUUID is a ValidateFunc for checking if a given string is a valid UUID.
func validateUUID(v interface{}, k string) (ws []string, errors []error) {
	log.Print("validateUUID:start")
	value := v.(string) // Assert that the attribute value is a string.
	if _, err := uuid.Parse(value); err != nil {
		errors = append(errors, fmt.Errorf("invalid UUID format"))
	}
	log.Print("vaildateUUID:end")
	return
}

func providerConfigure(p *schema.Provider) schema.ConfigureContextFunc {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		log.Print("providerConfig:start")
		config := Config{
			Endpoint: d.Get("endpoint").(string),
			Token:    d.Get("token").(string),
			UserUuid: d.Get("user_uuid").(string),
		}
		log.Print("providerConfig:end")
		return &config, nil
	}
}

func Resource() *schema.Resource {
	log.Print("Resource:start")
	resource := &schema.Resource{
		CreateContext: resourceHouseCreate,
		ReadContext:   resourceHouseRead,
		UpdateContext: resourceHouseUpdate,
		DeleteContext: resourceHouseDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of home",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Description of home",
			},
			"domain_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Domain of home eg. *.cloudfront.net",
			},
			"town": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The town of which the home will belong to",
			},
			"content_version": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The content version of home",
			},
		},
	}
	log.Print("Resource:end")
	return resource
}

func resourceHouseCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Print("resourceHouseCreate:start")
	var diags diag.Diagnostics

	config := m.(*Config)

	payload := map[string]interface{}{
		"name":            d.Get("name").(string),
		"description":     d.Get("description").(string),
		"domain_name":     d.Get("domain_name").(string),
		"town":            d.Get("town").(string),
		"content_version": d.Get("content_version").(int),
	}
	payloadBytes, err := json.Marshal(payload)

	if err != nil {
		return diag.FromErr(err)
	}
	url := config.Endpoint + "/u/" + config.UserUuid + "/homes"
	log.Print("URL: " + url)
	//Construct the HTTP
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return diag.FromErr(err)
	}

	//Set Headers
	req.Header.Set("Authorization", "Bearer "+config.Token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return diag.FromErr(err)
	}

	defer resp.Body.Close()

	//parse response JSON
	var responseData map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		return diag.FromErr(err)
	}

	//StatusOk = 200 HTTP Response OK
	if resp.StatusCode != http.StatusOK {
		return diag.FromErr(fmt.Errorf("failed to create home resource, status_code: %d, status: %s, body: %s", resp.StatusCode, resp.Status, responseData))
	}

	//handle response status

	log.Print("resourceHouseCreate:end")

	homeUUID := responseData["uuid"].(string)
	d.SetId(homeUUID)
	return diags
}

func resourceHouseRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Print("resourceHouseRead:start")
	var diags diag.Diagnostics

	config := m.(*Config)

	homeUUID := d.Id()

	//Construct the HTTP
	url := config.Endpoint + "/u/" + config.UserUuid + "/homes/" + homeUUID
	log.Print("URL: " + url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	//Set Headers
	req.Header.Set("Authorization", "Bearer "+config.Token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return diag.FromErr(err)
	}

	defer resp.Body.Close()

	var responseData map[string]interface{}
	if resp.StatusCode == http.StatusOK {
		//parse response JSON
		if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
			return diag.FromErr(err)
		}
		d.Set("name", responseData["name"].(string))
		d.Set("description", responseData["description"].(string))
		d.Set("domain_name", responseData["domain_name"].(string))
		d.Set("town", responseData["town"].(string))
		d.Set("content_version", responseData["content_version"].(float64))
	} else if resp.StatusCode != http.StatusNotFound {
		d.SetId("")
	} else if resp.StatusCode != http.StatusOK {
		return diag.FromErr(fmt.Errorf("failed to read home resource, status_code: %d, status: %s, body: %s", resp.StatusCode, resp.Status, responseData))
	}

	log.Print("resourceHouseRead:end")
	return diags
}

func resourceHouseUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Print("resourceHouseUpdate:start")
	var diags diag.Diagnostics

	config := m.(*Config)

	homeUUID := d.Id()

	payload := map[string]interface{}{
		"name":            d.Get("name").(string),
		"description":     d.Get("description").(string),
		"content_version": d.Get("content_version").(int),
	}
	payloadBytes, err := json.Marshal(payload)

	if err != nil {
		return diag.FromErr(err)
	}

	url := config.Endpoint + "/u/" + config.UserUuid + "/homes/" + homeUUID
	log.Print("URL: " + url)
	//Construct the HTTP
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return diag.FromErr(err)
	}

	//Set Headers
	req.Header.Set("Authorization", "Bearer "+config.Token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return diag.FromErr(err)
	}

	defer resp.Body.Close()

	//StatusOk = 200 HTTP Response OK
	if resp.StatusCode != http.StatusOK {
		return diag.FromErr(fmt.Errorf("failed to update home resource, status_code: %d, status: %s", resp.StatusCode, resp.Status))
	}

	log.Print("resourceHouseUpdate:end")
	d.Set("name", payload["name"])
	d.Set("description", payload["description"])
	d.Set("content_version", payload["content_version"])
	return diags
}

func resourceHouseDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Print("resourceHouseDelete:start")
	var diags diag.Diagnostics

	config := m.(*Config)

	homeUUID := d.Id()

	url := config.Endpoint + "/u/" + config.UserUuid + "/homes/" + homeUUID
	log.Print("URL: " + url)
	//Construct the HTTP
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	//Set Headers
	req.Header.Set("Authorization", "Bearer "+config.Token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return diag.FromErr(err)
	}

	defer resp.Body.Close()

	//StatusOk = 200 HTTP Response OK
	if resp.StatusCode != http.StatusOK {
		return diag.FromErr(fmt.Errorf("failed to delete home resource, status_code: %d, status: %s", resp.StatusCode, resp.Status))
	}

	d.SetId("")

	log.Print("resourceHouseDelete:end")
	return diags
}
